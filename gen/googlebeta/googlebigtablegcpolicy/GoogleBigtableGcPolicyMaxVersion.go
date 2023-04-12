package googlebigtablegcpolicy


type GoogleBigtableGcPolicyMaxVersion struct {
	// Number of version before applying the GC policy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigtable_gc_policy#number GoogleBigtableGcPolicy#number}
	Number *float64 `field:"required" json:"number" yaml:"number"`
}

