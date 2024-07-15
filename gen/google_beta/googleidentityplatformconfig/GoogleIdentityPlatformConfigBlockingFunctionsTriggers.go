package googleidentityplatformconfig


type GoogleIdentityPlatformConfigBlockingFunctionsTriggers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#event_type GoogleIdentityPlatformConfig#event_type}.
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// HTTP URI trigger for the Cloud Function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#function_uri GoogleIdentityPlatformConfig#function_uri}
	FunctionUri *string `field:"required" json:"functionUri" yaml:"functionUri"`
}

