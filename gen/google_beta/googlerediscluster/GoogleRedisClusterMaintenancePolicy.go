package googlerediscluster


type GoogleRedisClusterMaintenancePolicy struct {
	// weekly_maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_redis_cluster#weekly_maintenance_window GoogleRedisCluster#weekly_maintenance_window}
	WeeklyMaintenanceWindow interface{} `field:"optional" json:"weeklyMaintenanceWindow" yaml:"weeklyMaintenanceWindow"`
}

