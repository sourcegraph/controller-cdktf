package dataobservemonitorv2


type DataObserveMonitorV2ActionsAction struct {
	// email block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#email DataObserveMonitorV2#email}
	Email interface{} `field:"optional" json:"email" yaml:"email"`
	// webhook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#webhook DataObserveMonitorV2#webhook}
	Webhook interface{} `field:"optional" json:"webhook" yaml:"webhook"`
}

