package monitorv2action


type MonitorV2ActionWebhookHeaders struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2_action#header MonitorV2Action#header}.
	Header *string `field:"required" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2_action#value MonitorV2Action#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

