package reportsystemhealthreview


type ReportSystemHealthReviewThresholds struct {
	// Min value for the Green status (e.g. healthy).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/report_system_health_review#green_gt ReportSystemHealthReview#green_gt}
	GreenGt *float64 `field:"required" json:"greenGt" yaml:"greenGt"`
	// Max value for the Red status (e.g. exhausted budget).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/report_system_health_review#red_lte ReportSystemHealthReview#red_lte}
	RedLte *float64 `field:"required" json:"redLte" yaml:"redLte"`
	// ShowNoData customizes the report to either show or hide rows with no data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/report_system_health_review#show_no_data ReportSystemHealthReview#show_no_data}
	ShowNoData interface{} `field:"optional" json:"showNoData" yaml:"showNoData"`
}

