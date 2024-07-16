package googlebigtablegcpolicy


type GoogleBigtableGcPolicyMaxVersion struct {
	// Number of version before applying the GC policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigtable_gc_policy#number GoogleBigtableGcPolicy#number}
	Number *float64 `field:"required" json:"number" yaml:"number"`
}

