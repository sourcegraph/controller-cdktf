package googlemonitoringslo


type GoogleMonitoringSloRequestBasedSliDistributionCut struct {
	// A TimeSeries [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters) aggregating values to quantify the good service provided.
	//
	// Must have ValueType = DISTRIBUTION and
	// MetricKind = DELTA or MetricKind = CUMULATIVE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_slo#distribution_filter GoogleMonitoringSlo#distribution_filter}
	DistributionFilter *string `field:"required" json:"distributionFilter" yaml:"distributionFilter"`
	// range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_slo#range GoogleMonitoringSlo#range}
	Range *GoogleMonitoringSloRequestBasedSliDistributionCutRange `field:"required" json:"range" yaml:"range"`
}

