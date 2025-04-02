package reportsystemhealthreview


type ReportSystemHealthReviewTimeFrame struct {
	// snapshot block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/report_system_health_review#snapshot ReportSystemHealthReview#snapshot}
	Snapshot *ReportSystemHealthReviewTimeFrameSnapshot `field:"required" json:"snapshot" yaml:"snapshot"`
	// Timezone name in IANA Time Zone Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/report_system_health_review#time_zone ReportSystemHealthReview#time_zone}
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
}

