package reportsystemhealthreview


type ReportSystemHealthReviewTimeFrameSnapshot struct {
	// The method of reporting time frame [past/latest].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/report_system_health_review#point ReportSystemHealthReview#point}
	Point *string `field:"required" json:"point" yaml:"point"`
	// Date and time of the past snapshot in RFC3339 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/report_system_health_review#date_time ReportSystemHealthReview#date_time}
	DateTime *string `field:"optional" json:"dateTime" yaml:"dateTime"`
	// The recurrence rule for the report past snapshot. The expected value is a string in RRULE format. Example: `FREQ=MONTHLY;BYMONTHDAY=1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/report_system_health_review#rrule ReportSystemHealthReview#rrule}
	Rrule *string `field:"optional" json:"rrule" yaml:"rrule"`
}

