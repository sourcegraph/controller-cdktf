package poller


type PollerCloudwatchMetricsQuery struct {
	// AWS Metric Namespace to query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#namespace Poller#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#dimension Poller#dimension}
	Dimension interface{} `field:"optional" json:"dimension" yaml:"dimension"`
	// Metric names to filter down to.
	//
	// If more than one metric name is provided, `ListMetrics` will be called with no filter on metric names.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#metric_names Poller#metric_names}
	MetricNames *[]*string `field:"optional" json:"metricNames" yaml:"metricNames"`
	// resource_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#resource_filter Poller#resource_filter}
	ResourceFilter interface{} `field:"optional" json:"resourceFilter" yaml:"resourceFilter"`
}

