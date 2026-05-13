package googledialogflowcxtool


type GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOauth2JwtBearerConfig struct {
	// Client key value or parameter name to pass it through.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool#client_key GoogleDialogflowCxTool#client_key}
	ClientKey *string `field:"required" json:"clientKey" yaml:"clientKey"`
	// Issuer value or parameter name to pass it through.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool#issuer GoogleDialogflowCxTool#issuer}
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Subject value or parameter name to pass it through.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool#subject GoogleDialogflowCxTool#subject}
	Subject *string `field:"required" json:"subject" yaml:"subject"`
}

