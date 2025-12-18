package poller


type PollerCloudwatchMetricsQueryDimension struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#name Poller#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#value Poller#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

