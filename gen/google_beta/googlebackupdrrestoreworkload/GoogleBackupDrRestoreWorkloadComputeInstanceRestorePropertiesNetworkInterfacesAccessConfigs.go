package googlebackupdrrestoreworkload


type GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAccessConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#external_ip GoogleBackupDrRestoreWorkload#external_ip}.
	ExternalIp *string `field:"optional" json:"externalIp" yaml:"externalIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#external_ipv6 GoogleBackupDrRestoreWorkload#external_ipv6}.
	ExternalIpv6 *string `field:"optional" json:"externalIpv6" yaml:"externalIpv6"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#external_ipv6_prefix_length GoogleBackupDrRestoreWorkload#external_ipv6_prefix_length}.
	ExternalIpv6PrefixLength *float64 `field:"optional" json:"externalIpv6PrefixLength" yaml:"externalIpv6PrefixLength"`
	// Optional. The name of this access configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#name GoogleBackupDrRestoreWorkload#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Possible values: ["NETWORK_TIER_UNSPECIFIED", "PREMIUM", "STANDARD"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#network_tier GoogleBackupDrRestoreWorkload#network_tier}
	NetworkTier *string `field:"optional" json:"networkTier" yaml:"networkTier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#public_ptr_domain_name GoogleBackupDrRestoreWorkload#public_ptr_domain_name}.
	PublicPtrDomainName *string `field:"optional" json:"publicPtrDomainName" yaml:"publicPtrDomainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#set_public_ptr GoogleBackupDrRestoreWorkload#set_public_ptr}.
	SetPublicPtr interface{} `field:"optional" json:"setPublicPtr" yaml:"setPublicPtr"`
	// Optional. The type of configuration. Possible values: ["ACCESS_TYPE_UNSPECIFIED", "ONE_TO_ONE_NAT", "DIRECT_IPV6"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#type GoogleBackupDrRestoreWorkload#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

