package googlediscoveryengineaclconfig


type GoogleDiscoveryEngineAclConfigIdpConfig struct {
	// external_idp_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_acl_config#external_idp_config GoogleDiscoveryEngineAclConfig#external_idp_config}
	ExternalIdpConfig *GoogleDiscoveryEngineAclConfigIdpConfigExternalIdpConfig `field:"optional" json:"externalIdpConfig" yaml:"externalIdpConfig"`
	// Identity provider type. Possible values: ["GSUITE", "THIRD_PARTY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_acl_config#idp_type GoogleDiscoveryEngineAclConfig#idp_type}
	IdpType *string `field:"optional" json:"idpType" yaml:"idpType"`
}

