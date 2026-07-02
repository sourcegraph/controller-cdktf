package reportsystemhealthreview

type ReportSystemHealthReviewFilters struct {
	// label block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/report_system_health_review#label ReportSystemHealthReview#label}
	Label any `field:"optional" json:"label" yaml:"label"`
	// Projects to pull data for report from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/report_system_health_review#projects ReportSystemHealthReview#projects}
	Projects *[]*string `field:"optional" json:"projects" yaml:"projects"`
	// service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/report_system_health_review#service ReportSystemHealthReview#service}
	Service any `field:"optional" json:"service" yaml:"service"`
	// slo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/report_system_health_review#slo ReportSystemHealthReview#slo}
	Slo any `field:"optional" json:"slo" yaml:"slo"`
}
