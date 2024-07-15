package googlebigtablegcpolicy


type GoogleBigtableGcPolicyMaxAge struct {
	// Number of days before applying GC policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigtable_gc_policy#days GoogleBigtableGcPolicy#days}
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// Duration before applying GC policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigtable_gc_policy#duration GoogleBigtableGcPolicy#duration}
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
}

