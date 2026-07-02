package dataobservemonitor

type DataObserveMonitorRule struct {
	// change block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor#change DataObserveMonitor#change}
	Change any `field:"optional" json:"change" yaml:"change"`
	// count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor#count DataObserveMonitor#count}
	Count any `field:"optional" json:"count" yaml:"count"`
	// facet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor#facet DataObserveMonitor#facet}
	Facet any `field:"optional" json:"facet" yaml:"facet"`
	// group_by_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor#group_by_group DataObserveMonitor#group_by_group}
	GroupByGroup any `field:"optional" json:"groupByGroup" yaml:"groupByGroup"`
	// log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor#log DataObserveMonitor#log}
	Log any `field:"optional" json:"log" yaml:"log"`
	// promote block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor#promote DataObserveMonitor#promote}
	Promote any `field:"optional" json:"promote" yaml:"promote"`
	// threshold block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor#threshold DataObserveMonitor#threshold}
	Threshold any `field:"optional" json:"threshold" yaml:"threshold"`
}
