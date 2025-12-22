package monitorv2


type MonitorV2Rules struct {
	// The alarm level (Critical, Error, Informational, None, Warning).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#level MonitorV2#level}
	Level *string `field:"required" json:"level" yaml:"level"`
	// count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#count MonitorV2#count}
	Count *MonitorV2RulesCount `field:"optional" json:"count" yaml:"count"`
	// promote block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#promote MonitorV2#promote}
	Promote *MonitorV2RulesPromote `field:"optional" json:"promote" yaml:"promote"`
	// threshold block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#threshold MonitorV2#threshold}
	Threshold *MonitorV2RulesThreshold `field:"optional" json:"threshold" yaml:"threshold"`
}

