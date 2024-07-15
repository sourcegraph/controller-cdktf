package googlecontainerawsnodepool


type GoogleContainerAwsNodePoolConfigAutoscalingMetricsCollection struct {
	// The frequency at which EC2 Auto Scaling sends aggregated data to AWS CloudWatch. The only valid value is "1Minute".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_node_pool#granularity GoogleContainerAwsNodePool#granularity}
	Granularity *string `field:"required" json:"granularity" yaml:"granularity"`
	// The metrics to enable.
	//
	// For a list of valid metrics, see https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_EnableMetricsCollection.html. If you specify granularity and don't specify any metrics, all metrics are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_aws_node_pool#metrics GoogleContainerAwsNodePool#metrics}
	Metrics *[]*string `field:"optional" json:"metrics" yaml:"metrics"`
}

