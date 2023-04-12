package googlemonitoringslo


type GoogleMonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceDistributionCutRange struct {
	// max value for the range (inclusive). If not given, will be set to "infinity", defining an open range ">= range.min".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_monitoring_slo#max GoogleMonitoringSlo#max}
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// Min value for the range (inclusive). If not given, will be set to "-infinity", defining an open range "< range.max".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_monitoring_slo#min GoogleMonitoringSlo#min}
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

