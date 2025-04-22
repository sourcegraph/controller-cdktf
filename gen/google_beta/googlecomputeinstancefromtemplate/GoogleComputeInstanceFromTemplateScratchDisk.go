package googlecomputeinstancefromtemplate


type GoogleComputeInstanceFromTemplateScratchDisk struct {
	// The disk interface used for attaching this disk. One of SCSI or NVME.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_instance_from_template#interface GoogleComputeInstanceFromTemplate#interface}
	Interface *string `field:"required" json:"interface" yaml:"interface"`
	// Name with which the attached disk is accessible under /dev/disk/by-id/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_instance_from_template#device_name GoogleComputeInstanceFromTemplate#device_name}
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// The size of the disk in gigabytes. One of 375 or 3000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_instance_from_template#size GoogleComputeInstanceFromTemplate#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

