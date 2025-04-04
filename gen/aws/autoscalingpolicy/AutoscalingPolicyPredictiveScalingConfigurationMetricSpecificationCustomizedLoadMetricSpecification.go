package autoscalingpolicy


type AutoscalingPolicyPredictiveScalingConfigurationMetricSpecificationCustomizedLoadMetricSpecification struct {
	// metric_data_queries block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/autoscaling_policy#metric_data_queries AutoscalingPolicy#metric_data_queries}
	MetricDataQueries interface{} `field:"required" json:"metricDataQueries" yaml:"metricDataQueries"`
}

