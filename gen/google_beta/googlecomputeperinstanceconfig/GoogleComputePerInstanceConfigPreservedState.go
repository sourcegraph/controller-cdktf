package googlecomputeperinstanceconfig


type GoogleComputePerInstanceConfigPreservedState struct {
	// disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_per_instance_config#disk GoogleComputePerInstanceConfig#disk}
	Disk interface{} `field:"optional" json:"disk" yaml:"disk"`
	// external_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_per_instance_config#external_ip GoogleComputePerInstanceConfig#external_ip}
	ExternalIp interface{} `field:"optional" json:"externalIp" yaml:"externalIp"`
	// internal_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_per_instance_config#internal_ip GoogleComputePerInstanceConfig#internal_ip}
	InternalIp interface{} `field:"optional" json:"internalIp" yaml:"internalIp"`
	// Preserved metadata defined for this instance. This is a list of key->value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_per_instance_config#metadata GoogleComputePerInstanceConfig#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
}

