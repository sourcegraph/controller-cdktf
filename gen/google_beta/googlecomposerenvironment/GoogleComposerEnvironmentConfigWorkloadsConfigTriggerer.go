package googlecomposerenvironment


type GoogleComposerEnvironmentConfigWorkloadsConfigTriggerer struct {
	// The number of triggerers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_environment#count GoogleComposerEnvironment#count}
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// CPU request and limit for a single Airflow triggerer replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_environment#cpu GoogleComposerEnvironment#cpu}
	Cpu *float64 `field:"required" json:"cpu" yaml:"cpu"`
	// Memory (GB) request and limit for a single Airflow triggerer replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_environment#memory_gb GoogleComposerEnvironment#memory_gb}
	MemoryGb *float64 `field:"required" json:"memoryGb" yaml:"memoryGb"`
}

