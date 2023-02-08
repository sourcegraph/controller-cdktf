package googlenotebooksruntime


type GoogleNotebooksRuntimeVirtualMachine struct {
	// virtual_machine_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_notebooks_runtime#virtual_machine_config GoogleNotebooksRuntime#virtual_machine_config}
	VirtualMachineConfig *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfig `field:"optional" json:"virtualMachineConfig" yaml:"virtualMachineConfig"`
}

