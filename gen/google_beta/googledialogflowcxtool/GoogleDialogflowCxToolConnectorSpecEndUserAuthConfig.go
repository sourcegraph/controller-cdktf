package googledialogflowcxtool


type GoogleDialogflowCxToolConnectorSpecEndUserAuthConfig struct {
	// oauth2_auth_code_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool#oauth2_auth_code_config GoogleDialogflowCxTool#oauth2_auth_code_config}
	Oauth2AuthCodeConfig *GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOauth2AuthCodeConfig `field:"optional" json:"oauth2AuthCodeConfig" yaml:"oauth2AuthCodeConfig"`
	// oauth2_jwt_bearer_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool#oauth2_jwt_bearer_config GoogleDialogflowCxTool#oauth2_jwt_bearer_config}
	Oauth2JwtBearerConfig *GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOauth2JwtBearerConfig `field:"optional" json:"oauth2JwtBearerConfig" yaml:"oauth2JwtBearerConfig"`
}

