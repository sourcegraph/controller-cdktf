package alertroute


type AlertRouteIncidentTemplateSeverity struct {
	// Strategy for merging severity when multiple alerts create/update the same incident. Possible values are: `first-wins`, `max`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_route#merge_strategy AlertRoute#merge_strategy}
	MergeStrategy *string `field:"required" json:"mergeStrategy" yaml:"mergeStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/alert_route#binding AlertRoute#binding}.
	Binding *AlertRouteIncidentTemplateSeverityBinding `field:"optional" json:"binding" yaml:"binding"`
}

