package backupdrrestoreworkload


type BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesIpv6AccessConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#external_ip BackupDrRestoreWorkload#external_ip}.
	ExternalIp *string `field:"optional" json:"externalIp" yaml:"externalIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#external_ipv6 BackupDrRestoreWorkload#external_ipv6}.
	ExternalIpv6 *string `field:"optional" json:"externalIpv6" yaml:"externalIpv6"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#external_ipv6_prefix_length BackupDrRestoreWorkload#external_ipv6_prefix_length}.
	ExternalIpv6PrefixLength *float64 `field:"optional" json:"externalIpv6PrefixLength" yaml:"externalIpv6PrefixLength"`
	// Optional. The name of this access configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#name BackupDrRestoreWorkload#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Possible values: ["NETWORK_TIER_UNSPECIFIED", "PREMIUM", "STANDARD"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#network_tier BackupDrRestoreWorkload#network_tier}
	NetworkTier *string `field:"optional" json:"networkTier" yaml:"networkTier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#public_ptr_domain_name BackupDrRestoreWorkload#public_ptr_domain_name}.
	PublicPtrDomainName *string `field:"optional" json:"publicPtrDomainName" yaml:"publicPtrDomainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#set_public_ptr BackupDrRestoreWorkload#set_public_ptr}.
	SetPublicPtr interface{} `field:"optional" json:"setPublicPtr" yaml:"setPublicPtr"`
	// Optional. The type of configuration. Possible values: ["ACCESS_TYPE_UNSPECIFIED", "ONE_TO_ONE_NAT", "DIRECT_IPV6"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#type BackupDrRestoreWorkload#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

