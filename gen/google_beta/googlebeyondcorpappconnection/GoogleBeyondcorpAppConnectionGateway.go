package googlebeyondcorpappconnection


type GoogleBeyondcorpAppConnectionGateway struct {
	// AppGateway name in following format: projects/{project_id}/locations/{locationId}/appgateways/{gateway_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_beyondcorp_app_connection#app_gateway GoogleBeyondcorpAppConnection#app_gateway}
	AppGateway *string `field:"required" json:"appGateway" yaml:"appGateway"`
	// The type of hosting used by the gateway. Refer to https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#Type_1 for a list of possible values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_beyondcorp_app_connection#type GoogleBeyondcorpAppConnection#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

