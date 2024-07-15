package googleworkbenchinstance


type GoogleWorkbenchInstanceGceSetup struct {
	// accelerator_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workbench_instance#accelerator_configs GoogleWorkbenchInstance#accelerator_configs}
	AcceleratorConfigs interface{} `field:"optional" json:"acceleratorConfigs" yaml:"acceleratorConfigs"`
	// boot_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workbench_instance#boot_disk GoogleWorkbenchInstance#boot_disk}
	BootDisk *GoogleWorkbenchInstanceGceSetupBootDisk `field:"optional" json:"bootDisk" yaml:"bootDisk"`
	// container_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workbench_instance#container_image GoogleWorkbenchInstance#container_image}
	ContainerImage *GoogleWorkbenchInstanceGceSetupContainerImage `field:"optional" json:"containerImage" yaml:"containerImage"`
	// data_disks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workbench_instance#data_disks GoogleWorkbenchInstance#data_disks}
	DataDisks *GoogleWorkbenchInstanceGceSetupDataDisks `field:"optional" json:"dataDisks" yaml:"dataDisks"`
	// Optional. If true, no external IP will be assigned to this VM instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workbench_instance#disable_public_ip GoogleWorkbenchInstance#disable_public_ip}
	DisablePublicIp interface{} `field:"optional" json:"disablePublicIp" yaml:"disablePublicIp"`
	// Optional. Flag to enable ip forwarding or not, default false/off. https://cloud.google.com/vpc/docs/using-routes#canipforward.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workbench_instance#enable_ip_forwarding GoogleWorkbenchInstance#enable_ip_forwarding}
	EnableIpForwarding interface{} `field:"optional" json:"enableIpForwarding" yaml:"enableIpForwarding"`
	// Optional. The machine type of the VM instance. https://cloud.google.com/compute/docs/machine-resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workbench_instance#machine_type GoogleWorkbenchInstance#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// Optional. Custom metadata to apply to this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workbench_instance#metadata GoogleWorkbenchInstance#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// network_interfaces block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workbench_instance#network_interfaces GoogleWorkbenchInstance#network_interfaces}
	NetworkInterfaces interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// service_accounts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workbench_instance#service_accounts GoogleWorkbenchInstance#service_accounts}
	ServiceAccounts interface{} `field:"optional" json:"serviceAccounts" yaml:"serviceAccounts"`
	// shielded_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workbench_instance#shielded_instance_config GoogleWorkbenchInstance#shielded_instance_config}
	ShieldedInstanceConfig *GoogleWorkbenchInstanceGceSetupShieldedInstanceConfig `field:"optional" json:"shieldedInstanceConfig" yaml:"shieldedInstanceConfig"`
	// Optional. The Compute Engine tags to add to instance (see [Tagging instances](https://cloud.google.com/compute/docs/label-or-tag-resources#tags)).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workbench_instance#tags GoogleWorkbenchInstance#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// vm_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workbench_instance#vm_image GoogleWorkbenchInstance#vm_image}
	VmImage *GoogleWorkbenchInstanceGceSetupVmImage `field:"optional" json:"vmImage" yaml:"vmImage"`
}

