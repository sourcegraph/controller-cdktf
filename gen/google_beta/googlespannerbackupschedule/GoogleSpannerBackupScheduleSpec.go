package googlespannerbackupschedule


type GoogleSpannerBackupScheduleSpec struct {
	// cron_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_spanner_backup_schedule#cron_spec GoogleSpannerBackupSchedule#cron_spec}
	CronSpec *GoogleSpannerBackupScheduleSpecCronSpec `field:"optional" json:"cronSpec" yaml:"cronSpec"`
}

