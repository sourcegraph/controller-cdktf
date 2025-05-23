package spannerbackupschedule


type SpannerBackupScheduleSpec struct {
	// cron_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/spanner_backup_schedule#cron_spec SpannerBackupSchedule#cron_spec}
	CronSpec *SpannerBackupScheduleSpecCronSpec `field:"optional" json:"cronSpec" yaml:"cronSpec"`
}

