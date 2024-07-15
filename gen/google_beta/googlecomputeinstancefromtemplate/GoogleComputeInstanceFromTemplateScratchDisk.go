package googlecomputeinstancefromtemplate


type GoogleComputeInstanceFromTemplateScratchDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_template#device_name GoogleComputeInstanceFromTemplate#device_name}.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_template#interface GoogleComputeInstanceFromTemplate#interface}.
	Interface *string `field:"optional" json:"interface" yaml:"interface"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_from_template#size GoogleComputeInstanceFromTemplate#size}.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

