package googledialogflowcxtoolversion


type GoogleDialogflowCxToolVersionToolConnectorSpecEndUserAuthConfigOauth2JwtBearerConfig struct {
	// Client key value or parameter name to pass it through.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#client_key GoogleDialogflowCxToolVersion#client_key}
	ClientKey *string `field:"required" json:"clientKey" yaml:"clientKey"`
	// Issuer value or parameter name to pass it through.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#issuer GoogleDialogflowCxToolVersion#issuer}
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Subject value or parameter name to pass it through.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#subject GoogleDialogflowCxToolVersion#subject}
	Subject *string `field:"required" json:"subject" yaml:"subject"`
}

