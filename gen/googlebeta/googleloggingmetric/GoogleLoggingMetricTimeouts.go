package googleloggingmetric


type GoogleLoggingMetricTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_logging_metric#create GoogleLoggingMetric#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_logging_metric#delete GoogleLoggingMetric#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_logging_metric#update GoogleLoggingMetric#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

