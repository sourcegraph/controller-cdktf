package reportsystemhealthreview


type ReportSystemHealthReviewFiltersSlo struct {
	// Unique name of the resource, must conform to the naming convention from [DNS RFC1123](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/report_system_health_review#name ReportSystemHealthReview#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of the Nobl9 project the resource sits in, must conform to the naming convention from [DNS RFC1123](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/report_system_health_review#project ReportSystemHealthReview#project}
	Project *string `field:"required" json:"project" yaml:"project"`
}

