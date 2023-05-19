package googlecomputeinstance


type GoogleComputeInstanceBootDiskInitializeParams struct {
	// The image from which this disk was initialised.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_compute_instance#image GoogleComputeInstance#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
	// A set of key/value label pairs assigned to the disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_compute_instance#labels GoogleComputeInstance#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The size of the image in gigabytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_compute_instance#size GoogleComputeInstance#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// The Google Compute Engine disk type. Such as pd-standard, pd-ssd or pd-balanced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_compute_instance#type GoogleComputeInstance#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

