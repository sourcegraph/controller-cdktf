package alertroute


type AlertRouteEscalationConfigEscalationTargets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_route#escalation_paths AlertRoute#escalation_paths}.
	EscalationPaths *AlertRouteEscalationConfigEscalationTargetsEscalationPaths `field:"optional" json:"escalationPaths" yaml:"escalationPaths"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_route#users AlertRoute#users}.
	Users *AlertRouteEscalationConfigEscalationTargetsUsers `field:"optional" json:"users" yaml:"users"`
}

