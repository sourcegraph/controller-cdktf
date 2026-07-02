package poller

type PollerCloudwatchMetricsQueryResourceFilter struct {
	// tag_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#tag_filter Poller#tag_filter}
	TagFilter any `field:"required" json:"tagFilter" yaml:"tagFilter"`
	// Metric dimension name for resource identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#dimension_name Poller#dimension_name}
	DimensionName *string `field:"optional" json:"dimensionName" yaml:"dimensionName"`
	// Regular expression for extracting identifier out of resource ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#pattern Poller#pattern}
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// Resource type to filter for as supported by `aws resourcegroupstaggingapi get-resources`, e.g. `ec2:instance`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#resource_type Poller#resource_type}
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
}
