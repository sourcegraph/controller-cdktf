package googlecomputeresourcepolicy


type GoogleComputeResourcePolicySnapshotSchedulePolicySchedule struct {
	// daily_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_resource_policy#daily_schedule GoogleComputeResourcePolicy#daily_schedule}
	DailySchedule *GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleDailySchedule `field:"optional" json:"dailySchedule" yaml:"dailySchedule"`
	// hourly_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_resource_policy#hourly_schedule GoogleComputeResourcePolicy#hourly_schedule}
	HourlySchedule *GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule `field:"optional" json:"hourlySchedule" yaml:"hourlySchedule"`
	// weekly_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_resource_policy#weekly_schedule GoogleComputeResourcePolicy#weekly_schedule}
	WeeklySchedule *GoogleComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule `field:"optional" json:"weeklySchedule" yaml:"weeklySchedule"`
}

