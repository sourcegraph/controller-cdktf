package monitorv2


type MonitorV2NoDataRules struct {
	// Allows for the user to specify how long they'd like the missing data alert to persist for before it resolves by itself.
	//
	// If not provided, the default expiration time will be set to 24 hours. The expiration must be identical across all rules.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#expiration MonitorV2#expiration}
	Expiration *string `field:"optional" json:"expiration" yaml:"expiration"`
	// threshold block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2#threshold MonitorV2#threshold}
	Threshold *MonitorV2NoDataRulesThreshold `field:"optional" json:"threshold" yaml:"threshold"`
}

