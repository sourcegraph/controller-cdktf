package googlemonitoringslo


type GoogleMonitoringSloBasicSliLatency struct {
	// A duration string, e.g. 10s. Good service is defined to be the count of requests made to this service that return in no more than threshold.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_monitoring_slo#threshold GoogleMonitoringSlo#threshold}
	Threshold *string `field:"required" json:"threshold" yaml:"threshold"`
}

