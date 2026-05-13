package googledialogflowcxwebhook


type GoogleDialogflowCxWebhookGenericWebServiceOauthConfig struct {
	// The client ID provided by the 3rd party platform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_webhook#client_id GoogleDialogflowCxWebhook#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The token endpoint provided by the 3rd party platform to exchange an access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_webhook#token_endpoint GoogleDialogflowCxWebhook#token_endpoint}
	TokenEndpoint *string `field:"required" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// The client secret provided by the 3rd party platform.  If the 'secret_version_for_client_secret' field is set, this field will be ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_webhook#client_secret GoogleDialogflowCxWebhook#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The OAuth scopes to grant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_webhook#scopes GoogleDialogflowCxWebhook#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	// The name of the SecretManager secret version resource storing the client secret.
	//
	// If this field is set, the 'client_secret' field will be
	// ignored.
	// Format: 'projects/{project}/secrets/{secret}/versions/{version}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_webhook#secret_version_for_client_secret GoogleDialogflowCxWebhook#secret_version_for_client_secret}
	SecretVersionForClientSecret *string `field:"optional" json:"secretVersionForClientSecret" yaml:"secretVersionForClientSecret"`
}

