package googleloggingmetric


type GoogleLoggingMetricBucketOptionsExplicitBuckets struct {
	// The values must be monotonically increasing.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_logging_metric#bounds GoogleLoggingMetric#bounds}
	Bounds *[]*float64 `field:"required" json:"bounds" yaml:"bounds"`
}

