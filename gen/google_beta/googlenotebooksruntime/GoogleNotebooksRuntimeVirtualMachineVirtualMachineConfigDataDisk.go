package googlenotebooksruntime


type GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDisk struct {
	// initialize_params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_runtime#initialize_params GoogleNotebooksRuntime#initialize_params}
	InitializeParams *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigDataDiskInitializeParams `field:"optional" json:"initializeParams" yaml:"initializeParams"`
	// "Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME.
	//
	// The default is SCSI. Persistent
	// disks must always use SCSI and the request will fail if you attempt
	// to attach a persistent disk in any other format than SCSI. Local SSDs
	// can use either NVME or SCSI. For performance characteristics of SCSI
	// over NVMe, see Local SSD performance. Valid values: * NVME * SCSI".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_runtime#interface GoogleNotebooksRuntime#interface}
	Interface *string `field:"optional" json:"interface" yaml:"interface"`
	// The mode in which to attach this disk, either READ_WRITE or READ_ONLY.
	//
	// If not specified, the default is to attach
	// the disk in READ_WRITE mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_runtime#mode GoogleNotebooksRuntime#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Specifies a valid partial or full URL to an existing Persistent Disk resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_runtime#source GoogleNotebooksRuntime#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Specifies the type of the disk, either SCRATCH or PERSISTENT. If not specified, the default is PERSISTENT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_runtime#type GoogleNotebooksRuntime#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

