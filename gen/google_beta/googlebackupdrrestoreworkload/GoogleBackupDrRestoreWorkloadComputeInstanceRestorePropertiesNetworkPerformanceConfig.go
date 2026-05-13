package googlebackupdrrestoreworkload


type GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkPerformanceConfig struct {
	// Possible values: ["TIER_UNSPECIFIED", "DEFAULT", "TIER_1"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_backup_dr_restore_workload#total_egress_bandwidth_tier GoogleBackupDrRestoreWorkload#total_egress_bandwidth_tier}
	TotalEgressBandwidthTier *string `field:"optional" json:"totalEgressBandwidthTier" yaml:"totalEgressBandwidthTier"`
}

