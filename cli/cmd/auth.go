package cmd

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/keptn/go-utils/pkg/common/kubeutils"
	"github.com/keptn/keptn/cli/internal"

	"github.com/keptn/keptn/cli/internal/auth"
	"github.com/keptn/keptn/cli/pkg/credentialmanager"

	"github.com/spf13/cobra"
)

type authCmdParams struct {
	endPoint             *string
	apiToken             *string
	exportConfig         *bool
	acceptContext        bool
	secure               *bool
	skipNamespaceListing *bool
	oauth                *bool
	oauthLogout          *bool
	oauthDiscovery       *string
	oauthClientID        *string
	oauthScopes          *[]string
	oauthClientSecret    *string
}

type smartKeptnAuthParams struct {
	ingressName    string
	serviceName    string
	secretName     string
	insecurePrefix string
}

var smartKeptnAuth smartKeptnAuthParams

var authParams *authCmdParams

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth --endpoint=https://api.keptn.MY.DOMAIN.COM --api-token=SECRET_TOKEN",
	Args:  cobra.NoArgs,
	Short: "Authenticates the Keptn CLI against a Keptn installation",
	Long: `Authenticates the Keptn CLI against a Keptn installation using an endpoint and an API token. 
The endpoint and API token are automatically configured during the Keptn installation.
If the authentication is successful, the endpoint and the API token are stored in a password store of the underlying operating system.
More precisely, the Keptn CLI stores the endpoint and API token using *pass* in case of Linux, using *Keychain* in case of macOS, or *Wincred* in case of Windows.

**Note:** If you receive a warning *Using a file-based storage for the key because the password-store seems to be not set up.* this is because a password store could not be found in your environment. In this case, the credentials are stored in *~/.keptn/.keptn* in your home directory.
	`,
	Example: `keptn auth --endpoint=https://api.keptn.MY.DOMAIN.COM --api-token=abcd-0123-wxyz-7890
keptn auth				# Automatically fetch the endpoint & api-token from current kubernetes context
keptn auth --secure			# Authenticates against the https endpoint
keptn auth --skip-namespace-listing # To skip the listing of namespaces and use the namespace passed with "--namespace" flag (default namespace is 'keptn')
`,
	SilenceUsage: true,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		smartKeptnAuth = smartKeptnAuthParams{
			ingressName:    "api-keptn-ingress",
			serviceName:    "api-gateway-nginx",
			secretName:     "keptn-api-token",
			insecurePrefix: "http://",
		}
		return verifyAuthParams(authParams, smartKeptnAuth)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		authenticator := NewAuthenticator(namespace, credentialmanager.NewCredentialManager(authParams.acceptContext), auth.NewLocalFileOauthStore())
		if *authParams.exportConfig {
			endpoint, apiToken, err := authenticator.GetCredentials()
			if err != nil {
				return err
			}
			fmt.Println("Endpoint: ", endpoint.String())
			fmt.Println("API Token: ", apiToken)
			return nil
		}

		if *authParams.oauthLogout {
			store := auth.NewLocalFileOauthStore()
			err := store.Wipe()
			if err != nil {
				return err
			}
			fmt.Println("Successfully logged out")
			return nil
		}

		if *authParams.oauth {

			if *authParams.oauthDiscovery == "" {
				return fmt.Errorf("Unable to login: No OAuth Discovery URL provided")
			}
			if *authParams.oauthClientID == "" {
				return fmt.Errorf("Unable to login: No OAuth Client ID provided")
			}
			oauth := auth.NewOauthAuthenticator(auth.NewOauthDiscovery(&http.Client{}), auth.NewLocalFileOauthStore(), auth.NewBrowser(), &auth.ClosingRedirectHandler{})

			clientValues := auth.OauthClientValues{
				OauthDiscoveryURL: *authParams.oauthDiscovery,
				OauthClientID:     *authParams.oauthClientID,
				OauthClientSecret: *authParams.oauthClientSecret,
				OauthScopes:       *authParams.oauthScopes,
			}
			if err := oauth.Auth(clientValues); err != nil {
				return err
			}
		}

		return authenticator.Auth(AuthenticatorOptions{Endpoint: *authParams.endPoint, APIToken: *authParams.apiToken})
	},
}

func init() {

	rootCmd.AddCommand(authCmd)
	authParams = &authCmdParams{}
	authParams.endPoint = authCmd.Flags().StringP("endpoint", "e", "", "The endpoint exposed by the Keptn installation (e.g., api.keptn.127.0.0.1.xip.io)")
	authParams.apiToken = authCmd.Flags().StringP("api-token", "a", "", "The API token to communicate with the Keptn installation")
	authParams.exportConfig = authCmd.Flags().BoolP("export", "c", false, "To export the current cluster config i.e API token and Endpoint")
	authParams.oauth = authCmd.Flags().Bool("oauth", false, "Use OAuth Authorization Code Flow to log in")
	authParams.oauthLogout = authCmd.Flags().Bool("oauth-logout", false, "Disable Oauth access")
	authParams.oauthDiscovery = authCmd.Flags().String("oauth-discovery", "", "Well known discovery URL used for OAuth")
	authParams.oauthClientID = authCmd.Flags().String("oauth-client-id", "", "Oauth Client ID used for OAuth")
	authParams.oauthScopes = authCmd.Flags().StringSlice("oauth-scopes", []string{}, "Oauth scopes used for OAuth")
	authParams.oauthClientSecret = authCmd.Flags().String("oauth-client-secret", "", "Oauth Client Secret used for OAuth")
	authParams.secure = authCmd.Flags().BoolP("secure", "s", false, "To make http/https request to auto fetched endpoint while authentication")
	authParams.skipNamespaceListing = authCmd.Flags().BoolP("skip-namespace-listing", "i", false, "To skip the listing of namespaces and use the namespace passed with \"--namespace\" flag (default namespace is 'keptn')")
	authCmd.Flags().BoolVarP(&authParams.acceptContext, "yes", "y", false, "Automatically accept change of Kubernetes Context")
}

func verifyAuthParams(authParams *authCmdParams, smartKeptnAuth smartKeptnAuthParams) error {

	const parametersRequiredMessage = "keptn auth requires api-token and endpoint \n\n" +
		"For more information on how to obtain token and endpoint to go https://keptn.sh/docs/%s/reference/cli/#authenticate-keptn-cli \n\n" +
		"Alternatively, to quickly access Keptn, you can use a port-forward and then authenticate your Keptn CLI: \n" +
		"- kubectl -n %s port-forward service/api-gateway-nginx 8080:80 \n" +
		"- keptn auth --endpoint=http://localhost:8080/api --api-token=$(kubectl get secret keptn-api-token -n %s -ojsonpath={.data.keptn-api-token} | base64 --decode)"

	var err error
	if *authParams.exportConfig {
		return nil
	}

	if *authParams.oauth {
		return nil
	}

	if !*authParams.skipNamespaceListing && (authParams.endPoint == nil || *authParams.endPoint == "") && (authParams.apiToken == nil || *authParams.apiToken == "") {
		namespace, err = smartKeptnCLIAuth()
		if err != nil {
			return internal.OnAPIError(err)
		}
	}

	err = smartFetchKeptnAuthParameters(authParams, smartKeptnAuth)
	if err != nil {
		return fmt.Errorf(err.Error()+parametersRequiredMessage, getReleaseDocsURL(), namespace, namespace)
	}
	return nil
}

// Removes given prefix and adds the right one according to cmd paramters
func addCorrectHttpPrefix(authParams *authCmdParams) string {
	prefix := "http://"

	if *authParams.secure || strings.HasPrefix(*authParams.endPoint, "https://") {
		prefix = "https://"
	}

	url := strings.TrimPrefix(*authParams.endPoint, "https://")
	url = strings.TrimPrefix(url, "http://")

	return prefix + url
}

type resolveFunc func(string) ([]string, error)
type sleepFunc func(time.Duration)

func smartFetchKeptnAuthParameters(authParams *authCmdParams, smartKeptnAuth smartKeptnAuthParams) error {
	if authParams.endPoint == nil || *authParams.endPoint == "" {
		keptnEndpointProvider, err := kubeutils.NewKeptnEndpointProvider(false)
		if err != nil {
			return err
		}
		*authParams.endPoint, err = keptnEndpointProvider.GetKeptnEndpointFromIngress(context.TODO(), namespace, smartKeptnAuth.ingressName)
		if err != nil {
			*authParams.endPoint, err = keptnEndpointProvider.GetKeptnEndpointFromService(context.TODO(), namespace, smartKeptnAuth.serviceName)
			if err != nil {
				return fmt.Errorf("Cannot automatically fetch the endpoint\n" + err.Error() + "\n\n")
			}
		}
		*authParams.endPoint = *authParams.endPoint + "/api"
	}

	*authParams.endPoint = addCorrectHttpPrefix(authParams)

	if authParams.apiToken == nil || *authParams.apiToken == "" {
		keptnApiTokenProvider, err := kubeutils.NewAPITokenProvider(false)
		if err != nil {
			return err
		}
		*authParams.apiToken, err = keptnApiTokenProvider.GetKeptnAPITokenFromSecret(context.TODO(), namespace, smartKeptnAuth.secretName)
		if err != nil {
			return fmt.Errorf("Error in fetching the api-token\n" + err.Error() + "\nCLI is not authenticated")
		}
	}

	return nil
}

func smartKeptnCLIAuth() (string, error) {
	namespaceManager, err := kubeutils.NewNamespaceManager(false)
	if err != nil {
		return "", err
	}
	keptnInstallations, err := namespaceManager.GetKeptnManagedNamespace(context.TODO())
	if err != nil {
		return "", errors.New("Could not get current Kubernetes context from KUBECONFIG: " + err.Error())
	}

	if len(keptnInstallations) > 1 {
		fmt.Println("There are multiple Keptn installations on your Kubernetes cluster, please select the correct one from the list below to continue")
		for index, keptnInstallation := range keptnInstallations {
			fmt.Printf("\t%d - %s\n", index, keptnInstallation)
		}
		fmt.Println("Please select the correct keptn installation: ")
		reader := bufio.NewReader(os.Stdin)
		in, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		inp, err := strconv.Atoi(strings.TrimSuffix(in, "\n"))
		if err != nil {
			return "", err
		}
		if inp >= len(keptnInstallations) {
			return "", fmt.Errorf("Please select the correct keptn installation")
		}
		return keptnInstallations[inp], nil
	} else if len(keptnInstallations) == 0 {
		return "", errors.New("We haven't found any Keptn Installation, Please follow the upgrade guide and patch the namespace with the annotation & label 'keptn.sh/managed-by: keptn'")
	}
	return keptnInstallations[0], nil
}
