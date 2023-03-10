package googlecomputeinstancetemplate


type GoogleComputeInstanceTemplateDisk struct {
	// Whether or not the disk should be auto-deleted. This defaults to true.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#auto_delete GoogleComputeInstanceTemplate#auto_delete}
	AutoDelete interface{} `field:"optional" json:"autoDelete" yaml:"autoDelete"`
	// Indicates that this is a boot disk.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#boot GoogleComputeInstanceTemplate#boot}
	Boot interface{} `field:"optional" json:"boot" yaml:"boot"`
	// A unique device name that is reflected into the /dev/ tree of a Linux operating system running within the instance.
	//
	// If not specified, the server chooses a default device name to apply to this disk.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#device_name GoogleComputeInstanceTemplate#device_name}
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// disk_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#disk_encryption_key GoogleComputeInstanceTemplate#disk_encryption_key}
	DiskEncryptionKey *GoogleComputeInstanceTemplateDiskDiskEncryptionKey `field:"optional" json:"diskEncryptionKey" yaml:"diskEncryptionKey"`
	// Name of the disk. When not provided, this defaults to the name of the instance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#disk_name GoogleComputeInstanceTemplate#disk_name}
	DiskName *string `field:"optional" json:"diskName" yaml:"diskName"`
	// The size of the image in gigabytes.
	//
	// If not specified, it will inherit the size of its base image. For SCRATCH disks, the size must be exactly 375GB.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#disk_size_gb GoogleComputeInstanceTemplate#disk_size_gb}
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// The Google Compute Engine disk type. Such as "pd-ssd", "local-ssd", "pd-balanced" or "pd-standard".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#disk_type GoogleComputeInstanceTemplate#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
	// Specifies the disk interface to use for attaching this disk.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#interface GoogleComputeInstanceTemplate#interface}
	Interface *string `field:"optional" json:"interface" yaml:"interface"`
	// A set of key/value label pairs to assign to disks,.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#labels GoogleComputeInstanceTemplate#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The mode in which to attach this disk, either READ_WRITE or READ_ONLY.
	//
	// If you are attaching or creating a boot disk, this must read-write mode.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#mode GoogleComputeInstanceTemplate#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// A list (short name or id) of resource policies to attach to this disk.
	//
	// Currently a max of 1 resource policy is supported.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#resource_policies GoogleComputeInstanceTemplate#resource_policies}
	ResourcePolicies *[]*string `field:"optional" json:"resourcePolicies" yaml:"resourcePolicies"`
	// The name (not self_link) of the disk (such as those managed by google_compute_disk) to attach.
	//
	// ~> Note: Either source or source_image is required when creating a new instance except for when creating a local SSD.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#source GoogleComputeInstanceTemplate#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
	// The image from which to initialize this disk.
	//
	// This can be one of: the image's self_link, projects/{project}/global/images/{image}, projects/{project}/global/images/family/{family}, global/images/{image}, global/images/family/{family}, family/{family}, {project}/{family}, {project}/{image}, {family}, or {image}. ~> Note: Either source or source_image is required when creating a new instance except for when creating a local SSD.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#source_image GoogleComputeInstanceTemplate#source_image}
	SourceImage *string `field:"optional" json:"sourceImage" yaml:"sourceImage"`
	// The type of Google Compute Engine disk, can be either "SCRATCH" or "PERSISTENT".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#type GoogleComputeInstanceTemplate#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

