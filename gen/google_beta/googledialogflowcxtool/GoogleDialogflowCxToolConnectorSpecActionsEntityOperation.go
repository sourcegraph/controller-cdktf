package googledialogflowcxtool


type GoogleDialogflowCxToolConnectorSpecActionsEntityOperation struct {
	// ID of the entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool#entity_id GoogleDialogflowCxTool#entity_id}
	EntityId *string `field:"required" json:"entityId" yaml:"entityId"`
	// The operation to perform on the entity. Possible values: ["LIST", "CREATE", "UPDATE", "DELETE", "GET"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool#operation GoogleDialogflowCxTool#operation}
	Operation *string `field:"required" json:"operation" yaml:"operation"`
}

