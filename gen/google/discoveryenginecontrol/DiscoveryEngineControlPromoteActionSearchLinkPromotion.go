package discoveryenginecontrol


type DiscoveryEngineControlPromoteActionSearchLinkPromotion struct {
	// The title of the promoted link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#title DiscoveryEngineControl#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// The description of the promoted link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#description DiscoveryEngineControl#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The document to promote.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#document DiscoveryEngineControl#document}
	Document *string `field:"optional" json:"document" yaml:"document"`
	// Return promotions for basic site search.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#enabled DiscoveryEngineControl#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The image URI of the promoted link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#image_uri DiscoveryEngineControl#image_uri}
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
	// The URI to promote.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#uri DiscoveryEngineControl#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

