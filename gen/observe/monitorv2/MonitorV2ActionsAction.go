package monitorv2


type MonitorV2ActionsAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#type MonitorV2#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#description MonitorV2#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// email block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#email MonitorV2#email}
	Email *MonitorV2ActionsActionEmail `field:"optional" json:"email" yaml:"email"`
	// webhook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#webhook MonitorV2#webhook}
	Webhook *MonitorV2ActionsActionWebhook `field:"optional" json:"webhook" yaml:"webhook"`
}

