package workstationsworkstationconfig


type WorkstationsWorkstationConfigHostGceInstance struct {
	// accelerators block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#accelerators WorkstationsWorkstationConfigA#accelerators}
	Accelerators interface{} `field:"optional" json:"accelerators" yaml:"accelerators"`
	// boost_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#boost_configs WorkstationsWorkstationConfigA#boost_configs}
	BoostConfigs interface{} `field:"optional" json:"boostConfigs" yaml:"boostConfigs"`
	// Size of the boot disk in GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#boot_disk_size_gb WorkstationsWorkstationConfigA#boot_disk_size_gb}
	BootDiskSizeGb *float64 `field:"optional" json:"bootDiskSizeGb" yaml:"bootDiskSizeGb"`
	// confidential_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#confidential_instance_config WorkstationsWorkstationConfigA#confidential_instance_config}
	ConfidentialInstanceConfig *WorkstationsWorkstationConfigHostGceInstanceConfidentialInstanceConfig `field:"optional" json:"confidentialInstanceConfig" yaml:"confidentialInstanceConfig"`
	// Whether instances have no public IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#disable_public_ip_addresses WorkstationsWorkstationConfigA#disable_public_ip_addresses}
	DisablePublicIpAddresses interface{} `field:"optional" json:"disablePublicIpAddresses" yaml:"disablePublicIpAddresses"`
	// Whether to disable SSH access to the VM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#disable_ssh WorkstationsWorkstationConfigA#disable_ssh}
	DisableSsh interface{} `field:"optional" json:"disableSsh" yaml:"disableSsh"`
	// Whether to enable nested virtualization on the Compute Engine VMs backing the Workstations.
	//
	// See https://cloud.google.com/workstations/docs/reference/rest/v1/projects.locations.workstationClusters.workstationConfigs#GceInstance.FIELDS.enable_nested_virtualization
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#enable_nested_virtualization WorkstationsWorkstationConfigA#enable_nested_virtualization}
	EnableNestedVirtualization interface{} `field:"optional" json:"enableNestedVirtualization" yaml:"enableNestedVirtualization"`
	// The name of a Compute Engine machine type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#machine_type WorkstationsWorkstationConfigA#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// Number of instances to pool for faster workstation startup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#pool_size WorkstationsWorkstationConfigA#pool_size}
	PoolSize *float64 `field:"optional" json:"poolSize" yaml:"poolSize"`
	// Email address of the service account that will be used on VM instances used to support this config.
	//
	// This service account must have permission to pull the specified container image. If not set, VMs will run without a service account, in which case the image must be publicly accessible.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#service_account WorkstationsWorkstationConfigA#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// Scopes to grant to the service_account.
	//
	// Various scopes are automatically added based on feature usage. When specified, users of workstations under this configuration must have 'iam.serviceAccounts.actAs' on the service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#service_account_scopes WorkstationsWorkstationConfigA#service_account_scopes}
	ServiceAccountScopes *[]*string `field:"optional" json:"serviceAccountScopes" yaml:"serviceAccountScopes"`
	// shielded_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#shielded_instance_config WorkstationsWorkstationConfigA#shielded_instance_config}
	ShieldedInstanceConfig *WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfig `field:"optional" json:"shieldedInstanceConfig" yaml:"shieldedInstanceConfig"`
	// Network tags to add to the Compute Engine machines backing the Workstations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#tags WorkstationsWorkstationConfigA#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Resource manager tags to be bound to the VM instances backing the Workstations.
	//
	// Tag keys and values have the same definition as
	// https://docs.cloud.google.com/resource-manager/docs/tags/tags-overview
	// Keys must be in the format 'tagKeys/{tag_key_id}', and
	// values are in the format 'tagValues/456'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_config#vm_tags WorkstationsWorkstationConfigA#vm_tags}
	VmTags *map[string]*string `field:"optional" json:"vmTags" yaml:"vmTags"`
}

