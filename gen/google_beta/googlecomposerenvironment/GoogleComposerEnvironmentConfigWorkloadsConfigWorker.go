package googlecomposerenvironment


type GoogleComposerEnvironmentConfigWorkloadsConfigWorker struct {
	// CPU request and limit for a single Airflow worker replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_environment#cpu GoogleComposerEnvironment#cpu}
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// Maximum number of workers for autoscaling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_environment#max_count GoogleComposerEnvironment#max_count}
	MaxCount *float64 `field:"optional" json:"maxCount" yaml:"maxCount"`
	// Memory (GB) request and limit for a single Airflow worker replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_environment#memory_gb GoogleComposerEnvironment#memory_gb}
	MemoryGb *float64 `field:"optional" json:"memoryGb" yaml:"memoryGb"`
	// Minimum number of workers for autoscaling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_environment#min_count GoogleComposerEnvironment#min_count}
	MinCount *float64 `field:"optional" json:"minCount" yaml:"minCount"`
	// Storage (GB) request and limit for a single Airflow worker replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_environment#storage_gb GoogleComposerEnvironment#storage_gb}
	StorageGb *float64 `field:"optional" json:"storageGb" yaml:"storageGb"`
}

