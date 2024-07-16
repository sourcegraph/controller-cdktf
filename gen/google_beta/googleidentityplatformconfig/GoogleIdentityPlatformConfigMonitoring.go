package googleidentityplatformconfig


type GoogleIdentityPlatformConfigMonitoring struct {
	// request_logging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#request_logging GoogleIdentityPlatformConfig#request_logging}
	RequestLogging *GoogleIdentityPlatformConfigMonitoringRequestLogging `field:"optional" json:"requestLogging" yaml:"requestLogging"`
}

