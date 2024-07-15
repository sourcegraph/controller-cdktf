package googleloggingmetric


type GoogleLoggingMetricBucketOptions struct {
	// explicit_buckets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_metric#explicit_buckets GoogleLoggingMetric#explicit_buckets}
	ExplicitBuckets *GoogleLoggingMetricBucketOptionsExplicitBuckets `field:"optional" json:"explicitBuckets" yaml:"explicitBuckets"`
	// exponential_buckets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_metric#exponential_buckets GoogleLoggingMetric#exponential_buckets}
	ExponentialBuckets *GoogleLoggingMetricBucketOptionsExponentialBuckets `field:"optional" json:"exponentialBuckets" yaml:"exponentialBuckets"`
	// linear_buckets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_metric#linear_buckets GoogleLoggingMetric#linear_buckets}
	LinearBuckets *GoogleLoggingMetricBucketOptionsLinearBuckets `field:"optional" json:"linearBuckets" yaml:"linearBuckets"`
}

