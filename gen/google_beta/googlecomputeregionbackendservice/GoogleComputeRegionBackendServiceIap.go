package googlecomputeregionbackendservice


type GoogleComputeRegionBackendServiceIap struct {
	// OAuth2 Client ID for IAP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_backend_service#oauth2_client_id GoogleComputeRegionBackendService#oauth2_client_id}
	Oauth2ClientId *string `field:"required" json:"oauth2ClientId" yaml:"oauth2ClientId"`
	// OAuth2 Client Secret for IAP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_backend_service#oauth2_client_secret GoogleComputeRegionBackendService#oauth2_client_secret}
	Oauth2ClientSecret *string `field:"required" json:"oauth2ClientSecret" yaml:"oauth2ClientSecret"`
}

