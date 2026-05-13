package backupdrrestoreworkload


type BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesAliasIpRanges struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#ip_cidr_range BackupDrRestoreWorkload#ip_cidr_range}.
	IpCidrRange *string `field:"optional" json:"ipCidrRange" yaml:"ipCidrRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#subnetwork_range_name BackupDrRestoreWorkload#subnetwork_range_name}.
	SubnetworkRangeName *string `field:"optional" json:"subnetworkRangeName" yaml:"subnetworkRangeName"`
}

