package alertroute

type AlertRouteEscalationConfig struct {
	// Should we auto cancel escalations when all alerts are resolved?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#auto_cancel_escalations AlertRoute#auto_cancel_escalations}
	AutoCancelEscalations any `field:"required" json:"autoCancelEscalations" yaml:"autoCancelEscalations"`
	// Targets for escalation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#escalation_targets AlertRoute#escalation_targets}
	EscalationTargets any `field:"required" json:"escalationTargets" yaml:"escalationTargets"`
}
