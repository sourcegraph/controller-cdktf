package agent


type AgentLogicMonitorConfig struct {
	// LogicMonitor Account name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/agent#account Agent#account}
	Account *string `field:"required" json:"account" yaml:"account"`
}

