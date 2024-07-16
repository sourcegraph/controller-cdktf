package googlenetappvolume


type GoogleNetappVolumeSnapshotPolicy struct {
	// daily_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#daily_schedule GoogleNetappVolume#daily_schedule}
	DailySchedule *GoogleNetappVolumeSnapshotPolicyDailySchedule `field:"optional" json:"dailySchedule" yaml:"dailySchedule"`
	// Enables automated snapshot creation according to defined schedule.
	//
	// Default is false.
	// To disable automatic snapshot creation you have to remove the whole snapshot_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#enabled GoogleNetappVolume#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// hourly_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#hourly_schedule GoogleNetappVolume#hourly_schedule}
	HourlySchedule *GoogleNetappVolumeSnapshotPolicyHourlySchedule `field:"optional" json:"hourlySchedule" yaml:"hourlySchedule"`
	// monthly_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#monthly_schedule GoogleNetappVolume#monthly_schedule}
	MonthlySchedule *GoogleNetappVolumeSnapshotPolicyMonthlySchedule `field:"optional" json:"monthlySchedule" yaml:"monthlySchedule"`
	// weekly_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_netapp_volume#weekly_schedule GoogleNetappVolume#weekly_schedule}
	WeeklySchedule *GoogleNetappVolumeSnapshotPolicyWeeklySchedule `field:"optional" json:"weeklySchedule" yaml:"weeklySchedule"`
}

