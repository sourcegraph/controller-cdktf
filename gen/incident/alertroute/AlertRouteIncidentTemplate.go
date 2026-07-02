package alertroute

type AlertRouteIncidentTemplate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#name AlertRoute#name}.
	Name *AlertRouteIncidentTemplateName `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#summary AlertRoute#summary}.
	Summary *AlertRouteIncidentTemplateSummary `field:"required" json:"summary" yaml:"summary"`
	// Custom fields configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#custom_fields AlertRoute#custom_fields}
	CustomFields any `field:"optional" json:"customFields" yaml:"customFields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#incident_mode AlertRoute#incident_mode}.
	IncidentMode *AlertRouteIncidentTemplateIncidentMode `field:"optional" json:"incidentMode" yaml:"incidentMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#incident_type AlertRoute#incident_type}.
	IncidentType *AlertRouteIncidentTemplateIncidentType `field:"optional" json:"incidentType" yaml:"incidentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#severity AlertRoute#severity}.
	Severity *AlertRouteIncidentTemplateSeverity `field:"optional" json:"severity" yaml:"severity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#start_in_triage AlertRoute#start_in_triage}.
	StartInTriage *AlertRouteIncidentTemplateStartInTriage `field:"optional" json:"startInTriage" yaml:"startInTriage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#workspace AlertRoute#workspace}.
	Workspace *AlertRouteIncidentTemplateWorkspace `field:"optional" json:"workspace" yaml:"workspace"`
}
