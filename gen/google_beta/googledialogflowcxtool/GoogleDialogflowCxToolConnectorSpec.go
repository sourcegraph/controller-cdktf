package googledialogflowcxtool


type GoogleDialogflowCxToolConnectorSpec struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool#actions GoogleDialogflowCxTool#actions}
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The full resource name of the referenced Integration Connectors Connection. Format: projects/* /locations/* /connections/*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool#name GoogleDialogflowCxTool#name}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	Name *string `field:"required" json:"name" yaml:"name"`
	// end_user_auth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool#end_user_auth_config GoogleDialogflowCxTool#end_user_auth_config}
	EndUserAuthConfig *GoogleDialogflowCxToolConnectorSpecEndUserAuthConfig `field:"optional" json:"endUserAuthConfig" yaml:"endUserAuthConfig"`
}

