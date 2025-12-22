package monitorv2action


type MonitorV2ActionEmail struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2_action#subject MonitorV2Action#subject}.
	Subject *string `field:"required" json:"subject" yaml:"subject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2_action#addresses MonitorV2Action#addresses}.
	Addresses *[]*string `field:"optional" json:"addresses" yaml:"addresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2_action#body MonitorV2Action#body}.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2_action#fragments MonitorV2Action#fragments}.
	Fragments *string `field:"optional" json:"fragments" yaml:"fragments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2_action#users MonitorV2Action#users}.
	Users *[]*string `field:"optional" json:"users" yaml:"users"`
}

