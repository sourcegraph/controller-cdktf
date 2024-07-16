package googleloggingmetric


type GoogleLoggingMetricMetricDescriptorLabels struct {
	// The label key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_metric#key GoogleLoggingMetric#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// A human-readable description for the label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_metric#description GoogleLoggingMetric#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The type of data that can be assigned to the label. Default value: "STRING" Possible values: ["BOOL", "INT64", "STRING"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_metric#value_type GoogleLoggingMetric#value_type}
	ValueType *string `field:"optional" json:"valueType" yaml:"valueType"`
}

