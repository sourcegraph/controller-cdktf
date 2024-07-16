package googledataprocmetastoreservice


type GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfig struct {
	// Defines whether autoscaling is enabled. The default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_metastore_service#autoscaling_enabled GoogleDataprocMetastoreService#autoscaling_enabled}
	AutoscalingEnabled interface{} `field:"optional" json:"autoscalingEnabled" yaml:"autoscalingEnabled"`
	// limit_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_metastore_service#limit_config GoogleDataprocMetastoreService#limit_config}
	LimitConfig *GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfig `field:"optional" json:"limitConfig" yaml:"limitConfig"`
}

