package googlecomposerenvironment


type GoogleComposerEnvironmentConfigWorkloadsConfigScheduler struct {
	// The number of schedulers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_environment#count GoogleComposerEnvironment#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// CPU request and limit for a single Airflow scheduler replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_environment#cpu GoogleComposerEnvironment#cpu}
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// Memory (GB) request and limit for a single Airflow scheduler replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_environment#memory_gb GoogleComposerEnvironment#memory_gb}
	MemoryGb *float64 `field:"optional" json:"memoryGb" yaml:"memoryGb"`
	// Storage (GB) request and limit for a single Airflow scheduler replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_composer_environment#storage_gb GoogleComposerEnvironment#storage_gb}
	StorageGb *float64 `field:"optional" json:"storageGb" yaml:"storageGb"`
}

