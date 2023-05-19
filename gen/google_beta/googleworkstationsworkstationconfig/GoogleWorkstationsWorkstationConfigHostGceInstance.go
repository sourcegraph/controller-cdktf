package googleworkstationsworkstationconfig


type GoogleWorkstationsWorkstationConfigHostGceInstance struct {
	// Size of the boot disk in GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_workstations_workstation_config#boot_disk_size_gb GoogleWorkstationsWorkstationConfigA#boot_disk_size_gb}
	BootDiskSizeGb *float64 `field:"optional" json:"bootDiskSizeGb" yaml:"bootDiskSizeGb"`
	// confidential_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_workstations_workstation_config#confidential_instance_config GoogleWorkstationsWorkstationConfigA#confidential_instance_config}
	ConfidentialInstanceConfig *GoogleWorkstationsWorkstationConfigHostGceInstanceConfidentialInstanceConfig `field:"optional" json:"confidentialInstanceConfig" yaml:"confidentialInstanceConfig"`
	// Whether instances have no public IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_workstations_workstation_config#disable_public_ip_addresses GoogleWorkstationsWorkstationConfigA#disable_public_ip_addresses}
	DisablePublicIpAddresses interface{} `field:"optional" json:"disablePublicIpAddresses" yaml:"disablePublicIpAddresses"`
	// The name of a Compute Engine machine type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_workstations_workstation_config#machine_type GoogleWorkstationsWorkstationConfigA#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// Number of instances to pool for faster workstation startup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_workstations_workstation_config#pool_size GoogleWorkstationsWorkstationConfigA#pool_size}
	PoolSize *float64 `field:"optional" json:"poolSize" yaml:"poolSize"`
	// Email address of the service account that will be used on VM instances used to support this config.
	//
	// This service account must have permission to pull the specified container image. If not set, VMs will run without a service account, in which case the image must be publicly accessible.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_workstations_workstation_config#service_account GoogleWorkstationsWorkstationConfigA#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// shielded_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_workstations_workstation_config#shielded_instance_config GoogleWorkstationsWorkstationConfigA#shielded_instance_config}
	ShieldedInstanceConfig *GoogleWorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfig `field:"optional" json:"shieldedInstanceConfig" yaml:"shieldedInstanceConfig"`
	// Network tags to add to the Compute Engine machines backing the Workstations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_workstations_workstation_config#tags GoogleWorkstationsWorkstationConfigA#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

