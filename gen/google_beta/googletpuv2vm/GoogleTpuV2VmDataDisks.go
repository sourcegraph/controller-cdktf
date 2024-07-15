package googletpuv2vm


type GoogleTpuV2VmDataDisks struct {
	// Specifies the full path to an existing disk. For example: "projects/my-project/zones/us-central1-c/disks/my-disk".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#source_disk GoogleTpuV2Vm#source_disk}
	SourceDisk *string `field:"required" json:"sourceDisk" yaml:"sourceDisk"`
	// The mode in which to attach this disk.
	//
	// If not specified, the default is READ_WRITE
	// mode. Only applicable to dataDisks. Default value: "READ_WRITE" Possible values: ["READ_WRITE", "READ_ONLY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#mode GoogleTpuV2Vm#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

