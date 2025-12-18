package poller


type PollerCloudwatchMetricsQueryResourceFilterTagFilter struct {
	// Tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#key Poller#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Set of acceptable tag values. Exact matches only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/poller#values Poller#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

