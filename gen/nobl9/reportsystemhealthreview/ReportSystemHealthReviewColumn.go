package reportsystemhealthreview


type ReportSystemHealthReviewColumn struct {
	// Column display name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/report_system_health_review#display_name ReportSystemHealthReview#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// label block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/report_system_health_review#label ReportSystemHealthReview#label}
	Label interface{} `field:"required" json:"label" yaml:"label"`
}

