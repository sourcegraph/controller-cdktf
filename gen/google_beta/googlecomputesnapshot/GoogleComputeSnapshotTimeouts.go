package googlecomputesnapshot


type GoogleComputeSnapshotTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_snapshot#create GoogleComputeSnapshot#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_snapshot#delete GoogleComputeSnapshot#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_snapshot#update GoogleComputeSnapshot#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

