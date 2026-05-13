package googledialogflowcxagent


type GoogleDialogflowCxAgentGenAppBuilderSettings struct {
	// The full name of the Gen App Builder engine related to this agent if there is one.
	//
	// Format: projects/{Project ID}/locations/{Location ID}/collections/{Collection ID}/engines/{Engine ID}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_agent#engine GoogleDialogflowCxAgent#engine}
	Engine *string `field:"required" json:"engine" yaml:"engine"`
}

