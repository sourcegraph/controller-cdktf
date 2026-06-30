package alertroute

type AlertRouteAlertSources struct {
	// The alert source ID that will match for the route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#alert_source_id AlertRoute#alert_source_id}
	AlertSourceId *string `field:"required" json:"alertSourceId" yaml:"alertSourceId"`
	// Groups of prerequisite conditions. All conditions in at least one group must be satisfied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#condition_groups AlertRoute#condition_groups}
	ConditionGroups any `field:"required" json:"conditionGroups" yaml:"conditionGroups"`
}
