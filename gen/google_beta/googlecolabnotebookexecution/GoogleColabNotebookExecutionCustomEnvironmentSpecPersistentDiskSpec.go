package googlecolabnotebookexecution


type GoogleColabNotebookExecutionCustomEnvironmentSpecPersistentDiskSpec struct {
	// The disk size of the runtime in GB.
	//
	// If specified, the diskType must also be specified. The minimum size is 10GB and the maximum is 65536GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_colab_notebook_execution#disk_size_gb GoogleColabNotebookExecution#disk_size_gb}
	DiskSizeGb *string `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// The type of the persistent disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_colab_notebook_execution#disk_type GoogleColabNotebookExecution#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
}

