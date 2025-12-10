package alertroute


type AlertRouteIncidentConfig struct {
	// Should triage incidents be declined when alerts are resolved?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#auto_decline_enabled AlertRoute#auto_decline_enabled}
	AutoDeclineEnabled interface{} `field:"required" json:"autoDeclineEnabled" yaml:"autoDeclineEnabled"`
	// Groups of prerequisite conditions. All conditions in at least one group must be satisfied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#condition_groups AlertRoute#condition_groups}
	ConditionGroups interface{} `field:"required" json:"conditionGroups" yaml:"conditionGroups"`
	// How long should the escalation defer time be?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#defer_time_seconds AlertRoute#defer_time_seconds}
	DeferTimeSeconds *float64 `field:"required" json:"deferTimeSeconds" yaml:"deferTimeSeconds"`
	// Whether incident creation is enabled for this alert route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#enabled AlertRoute#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Which attributes should this alert route use to group alerts?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#grouping_keys AlertRoute#grouping_keys}
	GroupingKeys interface{} `field:"required" json:"groupingKeys" yaml:"groupingKeys"`
	// How large should the grouping window be?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#grouping_window_seconds AlertRoute#grouping_window_seconds}
	GroupingWindowSeconds *float64 `field:"required" json:"groupingWindowSeconds" yaml:"groupingWindowSeconds"`
	// Should grouped alerts automatically be related to active incidents without confirmation?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#auto_relate_grouped_alerts AlertRoute#auto_relate_grouped_alerts}
	AutoRelateGroupedAlerts interface{} `field:"optional" json:"autoRelateGroupedAlerts" yaml:"autoRelateGroupedAlerts"`
}

