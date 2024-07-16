package googlealloydbcluster


type GoogleAlloydbClusterAutomatedBackupPolicyWeeklySchedule struct {
	// start_times block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#start_times GoogleAlloydbCluster#start_times}
	StartTimes interface{} `field:"required" json:"startTimes" yaml:"startTimes"`
	// The days of the week to perform a backup.
	//
	// At least one day of the week must be provided. Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#days_of_week GoogleAlloydbCluster#days_of_week}
	DaysOfWeek *[]*string `field:"optional" json:"daysOfWeek" yaml:"daysOfWeek"`
}

