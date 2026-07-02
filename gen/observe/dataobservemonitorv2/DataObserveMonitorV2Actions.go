package dataobservemonitorv2

type DataObserveMonitorV2Actions struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#action DataObserveMonitorV2#action}
	Action any `field:"optional" json:"action" yaml:"action"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#conditions DataObserveMonitorV2#conditions}
	Conditions any `field:"optional" json:"conditions" yaml:"conditions"`
}
