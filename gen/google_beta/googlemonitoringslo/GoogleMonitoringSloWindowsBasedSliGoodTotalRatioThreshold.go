package googlemonitoringslo


type GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThreshold struct {
	// basic_sli_performance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_slo#basic_sli_performance GoogleMonitoringSlo#basic_sli_performance}
	BasicSliPerformance *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformance `field:"optional" json:"basicSliPerformance" yaml:"basicSliPerformance"`
	// performance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_slo#performance GoogleMonitoringSlo#performance}
	Performance *GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformance `field:"optional" json:"performance" yaml:"performance"`
	// If window performance >= threshold, the window is counted as good.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_slo#threshold GoogleMonitoringSlo#threshold}
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
}

