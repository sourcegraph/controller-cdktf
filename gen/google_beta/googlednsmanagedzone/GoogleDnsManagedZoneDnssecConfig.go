package googlednsmanagedzone


type GoogleDnsManagedZoneDnssecConfig struct {
	// default_key_specs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#default_key_specs GoogleDnsManagedZone#default_key_specs}
	DefaultKeySpecs interface{} `field:"optional" json:"defaultKeySpecs" yaml:"defaultKeySpecs"`
	// Identifies what kind of resource this is.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#kind GoogleDnsManagedZone#kind}
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Specifies the mechanism used to provide authenticated denial-of-existence responses.
	//
	// non_existence can only be updated when the state is 'off'. Possible values: ["nsec", "nsec3"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#non_existence GoogleDnsManagedZone#non_existence}
	NonExistence *string `field:"optional" json:"nonExistence" yaml:"nonExistence"`
	// Specifies whether DNSSEC is enabled, and what mode it is in Possible values: ["off", "on", "transfer"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#state GoogleDnsManagedZone#state}
	State *string `field:"optional" json:"state" yaml:"state"`
}

