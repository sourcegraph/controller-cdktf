package googlesecurityscannerscanconfig


type GoogleSecurityScannerScanConfigSchedule struct {
	// The duration of time between executions in days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_security_scanner_scan_config#interval_duration_days GoogleSecurityScannerScanConfig#interval_duration_days}
	IntervalDurationDays *float64 `field:"required" json:"intervalDurationDays" yaml:"intervalDurationDays"`
	// A timestamp indicates when the next run will be scheduled.
	//
	// The value is refreshed
	// by the server after each run. If unspecified, it will default to current server time,
	// which means the scan will be scheduled to start immediately.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_security_scanner_scan_config#schedule_time GoogleSecurityScannerScanConfig#schedule_time}
	ScheduleTime *string `field:"optional" json:"scheduleTime" yaml:"scheduleTime"`
}

