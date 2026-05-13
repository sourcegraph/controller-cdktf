package googleapihubplugin


type GoogleApihubPluginConfigTemplateAuthConfigTemplateServiceAccount struct {
	// The service account to be used for authenticating request.
	//
	// The 'iam.serviceAccounts.getAccessToken' permission should be granted on
	// this service account to the impersonator service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apihub_plugin#service_account GoogleApihubPlugin#service_account}
	ServiceAccount *string `field:"required" json:"serviceAccount" yaml:"serviceAccount"`
}

