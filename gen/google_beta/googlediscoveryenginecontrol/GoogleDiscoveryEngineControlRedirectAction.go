package googlediscoveryenginecontrol


type GoogleDiscoveryEngineControlRedirectAction struct {
	// The URI to redirect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_control#redirect_uri GoogleDiscoveryEngineControl#redirect_uri}
	RedirectUri *string `field:"required" json:"redirectUri" yaml:"redirectUri"`
}

