package googlecomputeinstancefromtemplate


type GoogleComputeInstanceFromTemplateBootDiskInitializeParams struct {
	// The image from which this disk was initialised.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_from_template#image GoogleComputeInstanceFromTemplate#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
	// A set of key/value label pairs assigned to the disk.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_from_template#labels GoogleComputeInstanceFromTemplate#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The size of the image in gigabytes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_from_template#size GoogleComputeInstanceFromTemplate#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// The Google Compute Engine disk type. Such as pd-standard, pd-ssd or pd-balanced.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_from_template#type GoogleComputeInstanceFromTemplate#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

