package apihubplugininstance


type ApihubPluginInstanceAuthConfigGoogleServiceAccountConfig struct {
	// The service account to be used for authenticating request.
	//
	// The 'iam.serviceAccounts.getAccessToken' permission should be granted on
	// this service account to the impersonator service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/apihub_plugin_instance#service_account ApihubPluginInstance#service_account}
	ServiceAccount *string `field:"required" json:"serviceAccount" yaml:"serviceAccount"`
}

