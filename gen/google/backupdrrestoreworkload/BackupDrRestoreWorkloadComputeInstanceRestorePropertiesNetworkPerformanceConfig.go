package backupdrrestoreworkload


type BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkPerformanceConfig struct {
	// Possible values: ["TIER_UNSPECIFIED", "DEFAULT", "TIER_1"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/backup_dr_restore_workload#total_egress_bandwidth_tier BackupDrRestoreWorkload#total_egress_bandwidth_tier}
	TotalEgressBandwidthTier *string `field:"optional" json:"totalEgressBandwidthTier" yaml:"totalEgressBandwidthTier"`
}

