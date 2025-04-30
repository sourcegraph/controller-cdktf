package googlerediscluster


type GoogleRedisClusterAutomatedBackupConfigFixedFrequencySchedule struct {
	// start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_redis_cluster#start_time GoogleRedisCluster#start_time}
	StartTime *GoogleRedisClusterAutomatedBackupConfigFixedFrequencyScheduleStartTime `field:"required" json:"startTime" yaml:"startTime"`
}

