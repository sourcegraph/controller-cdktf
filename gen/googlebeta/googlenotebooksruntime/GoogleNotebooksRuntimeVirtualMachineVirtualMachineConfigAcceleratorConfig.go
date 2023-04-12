package googlenotebooksruntime


type GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigAcceleratorConfig struct {
	// Count of cores of this accelerator.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_notebooks_runtime#core_count GoogleNotebooksRuntime#core_count}
	CoreCount *float64 `field:"optional" json:"coreCount" yaml:"coreCount"`
	// Accelerator model. For valid values, see 'https://cloud.google.com/vertex-ai/docs/workbench/reference/ rest/v1/projects.locations.runtimes#AcceleratorType'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_notebooks_runtime#type GoogleNotebooksRuntime#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

