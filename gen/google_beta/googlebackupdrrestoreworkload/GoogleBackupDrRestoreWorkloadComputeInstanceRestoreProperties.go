package googlebackupdrrestoreworkload


type GoogleBackupDrRestoreWorkloadComputeInstanceRestoreProperties struct {
	// Required. Name of the compute instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#name GoogleBackupDrRestoreWorkload#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// advanced_machine_features block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#advanced_machine_features GoogleBackupDrRestoreWorkload#advanced_machine_features}
	AdvancedMachineFeatures *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeatures `field:"optional" json:"advancedMachineFeatures" yaml:"advancedMachineFeatures"`
	// allocation_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#allocation_affinity GoogleBackupDrRestoreWorkload#allocation_affinity}
	AllocationAffinity *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinity `field:"optional" json:"allocationAffinity" yaml:"allocationAffinity"`
	// Optional. Allows this instance to send and receive packets with non-matching destination or source IPs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#can_ip_forward GoogleBackupDrRestoreWorkload#can_ip_forward}
	CanIpForward interface{} `field:"optional" json:"canIpForward" yaml:"canIpForward"`
	// confidential_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#confidential_instance_config GoogleBackupDrRestoreWorkload#confidential_instance_config}
	ConfidentialInstanceConfig *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesConfidentialInstanceConfig `field:"optional" json:"confidentialInstanceConfig" yaml:"confidentialInstanceConfig"`
	// Optional. Whether the resource should be protected against deletion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#deletion_protection GoogleBackupDrRestoreWorkload#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Optional. An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#description GoogleBackupDrRestoreWorkload#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// disks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#disks GoogleBackupDrRestoreWorkload#disks}
	Disks interface{} `field:"optional" json:"disks" yaml:"disks"`
	// display_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#display_device GoogleBackupDrRestoreWorkload#display_device}
	DisplayDevice *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDevice `field:"optional" json:"displayDevice" yaml:"displayDevice"`
	// guest_accelerators block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#guest_accelerators GoogleBackupDrRestoreWorkload#guest_accelerators}
	GuestAccelerators interface{} `field:"optional" json:"guestAccelerators" yaml:"guestAccelerators"`
	// Optional. Specifies the hostname of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#hostname GoogleBackupDrRestoreWorkload#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// instance_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#instance_encryption_key GoogleBackupDrRestoreWorkload#instance_encryption_key}
	InstanceEncryptionKey *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesInstanceEncryptionKey `field:"optional" json:"instanceEncryptionKey" yaml:"instanceEncryptionKey"`
	// Optional. KeyRevocationActionType of the instance. Possible values: ["KEY_REVOCATION_ACTION_TYPE_UNSPECIFIED", "NONE", "STOP"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#key_revocation_action_type GoogleBackupDrRestoreWorkload#key_revocation_action_type}
	KeyRevocationActionType *string `field:"optional" json:"keyRevocationActionType" yaml:"keyRevocationActionType"`
	// labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#labels GoogleBackupDrRestoreWorkload#labels}
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// Optional. Full or partial URL of the machine type resource to use for this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#machine_type GoogleBackupDrRestoreWorkload#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#metadata GoogleBackupDrRestoreWorkload#metadata}
	Metadata *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Optional. Minimum CPU platform to use for this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#min_cpu_platform GoogleBackupDrRestoreWorkload#min_cpu_platform}
	MinCpuPlatform *string `field:"optional" json:"minCpuPlatform" yaml:"minCpuPlatform"`
	// network_interfaces block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#network_interfaces GoogleBackupDrRestoreWorkload#network_interfaces}
	NetworkInterfaces interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// network_performance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#network_performance_config GoogleBackupDrRestoreWorkload#network_performance_config}
	NetworkPerformanceConfig *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkPerformanceConfig `field:"optional" json:"networkPerformanceConfig" yaml:"networkPerformanceConfig"`
	// params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#params GoogleBackupDrRestoreWorkload#params}
	Params *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesParams `field:"optional" json:"params" yaml:"params"`
	// Optional. The private IPv6 google access type for the VM. Possible values: ["INSTANCE_PRIVATE_IPV6_GOOGLE_ACCESS_UNSPECIFIED", "INHERIT_FROM_SUBNETWORK", "ENABLE_OUTBOUND_VM_ACCESS_TO_GOOGLE", "ENABLE_BIDIRECTIONAL_ACCESS_TO_GOOGLE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#private_ipv6_google_access GoogleBackupDrRestoreWorkload#private_ipv6_google_access}
	PrivateIpv6GoogleAccess *string `field:"optional" json:"privateIpv6GoogleAccess" yaml:"privateIpv6GoogleAccess"`
	// Optional. Resource policies applied to this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#resource_policies GoogleBackupDrRestoreWorkload#resource_policies}
	ResourcePolicies *[]*string `field:"optional" json:"resourcePolicies" yaml:"resourcePolicies"`
	// scheduling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#scheduling GoogleBackupDrRestoreWorkload#scheduling}
	Scheduling *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesScheduling `field:"optional" json:"scheduling" yaml:"scheduling"`
	// service_accounts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#service_accounts GoogleBackupDrRestoreWorkload#service_accounts}
	ServiceAccounts interface{} `field:"optional" json:"serviceAccounts" yaml:"serviceAccounts"`
	// shielded_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#shielded_instance_config GoogleBackupDrRestoreWorkload#shielded_instance_config}
	ShieldedInstanceConfig *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfig `field:"optional" json:"shieldedInstanceConfig" yaml:"shieldedInstanceConfig"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#tags GoogleBackupDrRestoreWorkload#tags}
	Tags *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesTags `field:"optional" json:"tags" yaml:"tags"`
}

