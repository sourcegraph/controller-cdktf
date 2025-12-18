package monitorv2


type MonitorV2ActionsActionWebhookHeaders struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#header MonitorV2#header}.
	Header *string `field:"required" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#value MonitorV2#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

