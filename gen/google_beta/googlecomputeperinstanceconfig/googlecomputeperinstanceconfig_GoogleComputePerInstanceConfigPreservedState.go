package googlecomputeperinstanceconfig


type GoogleComputePerInstanceConfigPreservedState struct {
	// disk block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_per_instance_config#disk GoogleComputePerInstanceConfig#disk}
	Disk interface{} `field:"optional" json:"disk" yaml:"disk"`
	// Preserved metadata defined for this instance. This is a list of key->value pairs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_per_instance_config#metadata GoogleComputePerInstanceConfig#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
}

