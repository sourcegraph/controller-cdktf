package googlednsmanagedzone


type GoogleDnsManagedZoneDnssecConfigDefaultKeySpecs struct {
	// String mnemonic specifying the DNSSEC algorithm of this key Possible values: ["ecdsap256sha256", "ecdsap384sha384", "rsasha1", "rsasha256", "rsasha512"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#algorithm GoogleDnsManagedZone#algorithm}
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// Length of the keys in bits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#key_length GoogleDnsManagedZone#key_length}
	KeyLength *float64 `field:"optional" json:"keyLength" yaml:"keyLength"`
	// Specifies whether this is a key signing key (KSK) or a zone signing key (ZSK).
	//
	// Key signing keys have the Secure Entry
	// Point flag set and, when active, will only be used to sign
	// resource record sets of type DNSKEY. Zone signing keys do
	// not have the Secure Entry Point flag set and will be used
	// to sign all other types of resource record sets. Possible values: ["keySigning", "zoneSigning"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#key_type GoogleDnsManagedZone#key_type}
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
	// Identifies what kind of resource this is.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dns_managed_zone#kind GoogleDnsManagedZone#kind}
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
}

