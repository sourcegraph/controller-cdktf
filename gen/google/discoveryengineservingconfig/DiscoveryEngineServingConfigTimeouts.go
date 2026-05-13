package discoveryengineservingconfig


type DiscoveryEngineServingConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_serving_config#create DiscoveryEngineServingConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_serving_config#delete DiscoveryEngineServingConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_serving_config#update DiscoveryEngineServingConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

