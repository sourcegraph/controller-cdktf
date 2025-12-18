package monitorv2action


type MonitorV2ActionWebhook struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2_action#body MonitorV2Action#body}.
	Body *string `field:"required" json:"body" yaml:"body"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2_action#method MonitorV2Action#method}.
	Method *string `field:"required" json:"method" yaml:"method"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2_action#url MonitorV2Action#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2_action#fragments MonitorV2Action#fragments}.
	Fragments *string `field:"optional" json:"fragments" yaml:"fragments"`
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2_action#headers MonitorV2Action#headers}
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
}

