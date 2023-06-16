package autoscalingpolicy


type AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/autoscaling_policy#metric_name AutoscalingPolicy#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/autoscaling_policy#namespace AutoscalingPolicy#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/autoscaling_policy#statistic AutoscalingPolicy#statistic}.
	Statistic *string `field:"required" json:"statistic" yaml:"statistic"`
	// metric_dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/autoscaling_policy#metric_dimension AutoscalingPolicy#metric_dimension}
	MetricDimension interface{} `field:"optional" json:"metricDimension" yaml:"metricDimension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/autoscaling_policy#unit AutoscalingPolicy#unit}.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

