package discoveryengineaclconfig


type DiscoveryEngineAclConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_acl_config#create DiscoveryEngineAclConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_acl_config#delete DiscoveryEngineAclConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_acl_config#update DiscoveryEngineAclConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

