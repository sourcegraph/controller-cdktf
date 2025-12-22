package dataobservemonitorv2


type DataObserveMonitorV2Rules struct {
	// count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#count DataObserveMonitorV2#count}
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// promote block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#promote DataObserveMonitorV2#promote}
	Promote interface{} `field:"optional" json:"promote" yaml:"promote"`
	// threshold block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#threshold DataObserveMonitorV2#threshold}
	Threshold interface{} `field:"optional" json:"threshold" yaml:"threshold"`
}

