package googlecomputeimage


type GoogleComputeImageTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_image#create GoogleComputeImage#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_image#delete GoogleComputeImage#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_image#update GoogleComputeImage#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

