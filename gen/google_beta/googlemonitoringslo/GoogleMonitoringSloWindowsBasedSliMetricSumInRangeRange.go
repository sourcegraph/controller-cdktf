package googlemonitoringslo


type GoogleMonitoringSloWindowsBasedSliMetricSumInRangeRange struct {
	// max value for the range (inclusive). If not given, will be set to "infinity", defining an open range ">= range.min".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_slo#max GoogleMonitoringSlo#max}
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// Min value for the range (inclusive). If not given, will be set to "-infinity", defining an open range "< range.max".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_slo#min GoogleMonitoringSlo#min}
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

