package googleclouddomainsregistration


type GoogleClouddomainsRegistrationDnsSettingsCustomDnsDsRecords struct {
	// The algorithm used to generate the referenced DNSKEY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddomains_registration#algorithm GoogleClouddomainsRegistration#algorithm}
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// The digest generated from the referenced DNSKEY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddomains_registration#digest GoogleClouddomainsRegistration#digest}
	Digest *string `field:"optional" json:"digest" yaml:"digest"`
	// The hash function used to generate the digest of the referenced DNSKEY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddomains_registration#digest_type GoogleClouddomainsRegistration#digest_type}
	DigestType *string `field:"optional" json:"digestType" yaml:"digestType"`
	// The key tag of the record. Must be set in range 0 -- 65535.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddomains_registration#key_tag GoogleClouddomainsRegistration#key_tag}
	KeyTag *float64 `field:"optional" json:"keyTag" yaml:"keyTag"`
}

