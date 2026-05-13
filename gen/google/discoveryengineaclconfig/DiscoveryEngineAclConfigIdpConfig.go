package discoveryengineaclconfig


type DiscoveryEngineAclConfigIdpConfig struct {
	// external_idp_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_acl_config#external_idp_config DiscoveryEngineAclConfig#external_idp_config}
	ExternalIdpConfig *DiscoveryEngineAclConfigIdpConfigExternalIdpConfig `field:"optional" json:"externalIdpConfig" yaml:"externalIdpConfig"`
	// Identity provider type. Possible values: ["GSUITE", "THIRD_PARTY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_acl_config#idp_type DiscoveryEngineAclConfig#idp_type}
	IdpType *string `field:"optional" json:"idpType" yaml:"idpType"`
}

