package alertroute


type AlertRouteIncidentConfigGroupingKeys struct {
	// The alert attribute ID to use as a grouping key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_route#reference AlertRoute#reference}
	Reference *string `field:"required" json:"reference" yaml:"reference"`
}

