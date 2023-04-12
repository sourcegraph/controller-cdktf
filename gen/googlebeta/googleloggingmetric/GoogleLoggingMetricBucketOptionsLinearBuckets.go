package googleloggingmetric


type GoogleLoggingMetricBucketOptionsLinearBuckets struct {
	// Must be greater than 0.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_logging_metric#num_finite_buckets GoogleLoggingMetric#num_finite_buckets}
	NumFiniteBuckets *float64 `field:"optional" json:"numFiniteBuckets" yaml:"numFiniteBuckets"`
	// Lower bound of the first bucket.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_logging_metric#offset GoogleLoggingMetric#offset}
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
	// Must be greater than 0.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_logging_metric#width GoogleLoggingMetric#width}
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

