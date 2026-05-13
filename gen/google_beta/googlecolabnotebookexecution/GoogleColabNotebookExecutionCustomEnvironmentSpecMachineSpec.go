package googlecolabnotebookexecution


type GoogleColabNotebookExecutionCustomEnvironmentSpecMachineSpec struct {
	// The number of accelerators used by the runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_colab_notebook_execution#accelerator_count GoogleColabNotebookExecution#accelerator_count}
	AcceleratorCount *float64 `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// The type of hardware accelerator used by the runtime. If specified, acceleratorCount must also be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_colab_notebook_execution#accelerator_type GoogleColabNotebookExecution#accelerator_type}
	AcceleratorType *string `field:"optional" json:"acceleratorType" yaml:"acceleratorType"`
	// The Compute Engine machine type selected for the runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_colab_notebook_execution#machine_type GoogleColabNotebookExecution#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
}

