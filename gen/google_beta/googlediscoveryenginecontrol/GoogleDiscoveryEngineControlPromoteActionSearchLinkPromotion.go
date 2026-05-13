package googlediscoveryenginecontrol


type GoogleDiscoveryEngineControlPromoteActionSearchLinkPromotion struct {
	// The title of the promoted link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_control#title GoogleDiscoveryEngineControl#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// The description of the promoted link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_control#description GoogleDiscoveryEngineControl#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The document to promote.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_control#document GoogleDiscoveryEngineControl#document}
	Document *string `field:"optional" json:"document" yaml:"document"`
	// Return promotions for basic site search.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_control#enabled GoogleDiscoveryEngineControl#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The image URI of the promoted link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_control#image_uri GoogleDiscoveryEngineControl#image_uri}
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
	// The URI to promote.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_control#uri GoogleDiscoveryEngineControl#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

