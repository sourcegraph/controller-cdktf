package dataobservequery


type DataObserveQueryPoll struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/query#interval DataObserveQuery#interval}.
	Interval *string `field:"optional" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/query#timeout DataObserveQuery#timeout}.
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
}

