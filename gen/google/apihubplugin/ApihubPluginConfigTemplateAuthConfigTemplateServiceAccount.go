package apihubplugin


type ApihubPluginConfigTemplateAuthConfigTemplateServiceAccount struct {
	// The service account to be used for authenticating request.
	//
	// The 'iam.serviceAccounts.getAccessToken' permission should be granted on
	// this service account to the impersonator service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/apihub_plugin#service_account ApihubPlugin#service_account}
	ServiceAccount *string `field:"required" json:"serviceAccount" yaml:"serviceAccount"`
}

