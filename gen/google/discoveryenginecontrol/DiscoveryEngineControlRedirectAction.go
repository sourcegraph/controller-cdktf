package discoveryenginecontrol


type DiscoveryEngineControlRedirectAction struct {
	// The URI to redirect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#redirect_uri DiscoveryEngineControl#redirect_uri}
	RedirectUri *string `field:"required" json:"redirectUri" yaml:"redirectUri"`
}

