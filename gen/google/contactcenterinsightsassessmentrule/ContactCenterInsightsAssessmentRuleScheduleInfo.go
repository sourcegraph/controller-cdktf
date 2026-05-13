package contactcenterinsightsassessmentrule


type ContactCenterInsightsAssessmentRuleScheduleInfo struct {
	// End time of the schedule.
	//
	// If not specified, will keep scheduling new
	// pipelines for execution until the schedule is no longer active or deleted.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and
	// up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_assessment_rule#end_time ContactCenterInsightsAssessmentRule#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The groc expression.
	//
	// Format: 'every number [synchronized]'
	// Cron syntax is not supported.
	// Time units can be: minutes, hours
	// Synchronized is optional and indicates that the schedule should be
	// synchronized to the start of the interval: every 5 minutes synchronized
	// means 00:00, 00:05 ...
	// Otherwise the start time is random within the interval.
	// Example: 'every 5 minutes'
	// could be  00:02, 00:07, 00:12, ...
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_assessment_rule#schedule ContactCenterInsightsAssessmentRule#schedule}
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// Start time of the schedule.
	//
	// If not specified, will start as soon as the
	// schedule is created.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and
	// up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_assessment_rule#start_time ContactCenterInsightsAssessmentRule#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// The timezone to use for the groc expression. If not specified, defaults to UTC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_assessment_rule#time_zone ContactCenterInsightsAssessmentRule#time_zone}
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

