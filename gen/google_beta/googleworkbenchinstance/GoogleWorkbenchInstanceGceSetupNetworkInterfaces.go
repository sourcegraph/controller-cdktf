package googleworkbenchinstance


type GoogleWorkbenchInstanceGceSetupNetworkInterfaces struct {
	// Optional. The name of the VPC that this VM instance is in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workbench_instance#network GoogleWorkbenchInstance#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// Optional.
	//
	// The type of vNIC to be used on this interface. This
	// may be gVNIC or VirtioNet. Possible values: ["VIRTIO_NET", "GVNIC"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workbench_instance#nic_type GoogleWorkbenchInstance#nic_type}
	NicType *string `field:"optional" json:"nicType" yaml:"nicType"`
	// Optional. The name of the subnet that this VM instance is in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workbench_instance#subnet GoogleWorkbenchInstance#subnet}
	Subnet *string `field:"optional" json:"subnet" yaml:"subnet"`
}

