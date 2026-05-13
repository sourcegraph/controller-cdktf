package googlebackupdrrestoreworkload


type GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfaces struct {
	// access_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#access_configs GoogleBackupDrRestoreWorkload#access_configs}
	AccessConfigs interface{} `field:"optional" json:"accessConfigs" yaml:"accessConfigs"`
	// alias_ip_ranges block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#alias_ip_ranges GoogleBackupDrRestoreWorkload#alias_ip_ranges}
	AliasIpRanges interface{} `field:"optional" json:"aliasIpRanges" yaml:"aliasIpRanges"`
	// Optional. The prefix length of the primary internal IPv6 range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#internal_ipv6_prefix_length GoogleBackupDrRestoreWorkload#internal_ipv6_prefix_length}
	InternalIpv6PrefixLength *float64 `field:"optional" json:"internalIpv6PrefixLength" yaml:"internalIpv6PrefixLength"`
	// Optional. An IPv4 internal IP address to assign to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#ip_address GoogleBackupDrRestoreWorkload#ip_address}
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// ipv6_access_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#ipv6_access_configs GoogleBackupDrRestoreWorkload#ipv6_access_configs}
	Ipv6AccessConfigs interface{} `field:"optional" json:"ipv6AccessConfigs" yaml:"ipv6AccessConfigs"`
	// Possible values: ["UNSPECIFIED_IPV6_ACCESS_TYPE", "INTERNAL", "EXTERNAL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#ipv6_access_type GoogleBackupDrRestoreWorkload#ipv6_access_type}
	Ipv6AccessType *string `field:"optional" json:"ipv6AccessType" yaml:"ipv6AccessType"`
	// Optional. An IPv6 internal network address for this network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#ipv6_address GoogleBackupDrRestoreWorkload#ipv6_address}
	Ipv6Address *string `field:"optional" json:"ipv6Address" yaml:"ipv6Address"`
	// Optional. URL of the VPC network resource for this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#network GoogleBackupDrRestoreWorkload#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#network_attachment GoogleBackupDrRestoreWorkload#network_attachment}.
	NetworkAttachment *string `field:"optional" json:"networkAttachment" yaml:"networkAttachment"`
	// Possible values: ["NIC_TYPE_UNSPECIFIED", "VIRTIO_NET", "GVNIC"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#nic_type GoogleBackupDrRestoreWorkload#nic_type}
	NicType *string `field:"optional" json:"nicType" yaml:"nicType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#queue_count GoogleBackupDrRestoreWorkload#queue_count}.
	QueueCount *float64 `field:"optional" json:"queueCount" yaml:"queueCount"`
	// Possible values: ["STACK_TYPE_UNSPECIFIED", "IPV4_ONLY", "IPV4_IPV6"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#stack_type GoogleBackupDrRestoreWorkload#stack_type}
	StackType *string `field:"optional" json:"stackType" yaml:"stackType"`
	// Optional. The URL of the Subnetwork resource for this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#subnetwork GoogleBackupDrRestoreWorkload#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
}

