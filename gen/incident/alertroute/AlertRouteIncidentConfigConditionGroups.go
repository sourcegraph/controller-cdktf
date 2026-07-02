package alertroute

type AlertRouteIncidentConfigConditionGroups struct {
	// The prerequisite conditions that must all be satisfied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#conditions AlertRoute#conditions}
	Conditions any `field:"required" json:"conditions" yaml:"conditions"`
}
