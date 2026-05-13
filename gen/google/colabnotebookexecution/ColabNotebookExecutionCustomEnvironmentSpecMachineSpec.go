package colabnotebookexecution


type ColabNotebookExecutionCustomEnvironmentSpecMachineSpec struct {
	// The number of accelerators used by the runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/colab_notebook_execution#accelerator_count ColabNotebookExecution#accelerator_count}
	AcceleratorCount *float64 `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// The type of hardware accelerator used by the runtime. If specified, acceleratorCount must also be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/colab_notebook_execution#accelerator_type ColabNotebookExecution#accelerator_type}
	AcceleratorType *string `field:"optional" json:"acceleratorType" yaml:"acceleratorType"`
	// The Compute Engine machine type selected for the runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/colab_notebook_execution#machine_type ColabNotebookExecution#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
}

