package googledialogflowcxtoolversion


type GoogleDialogflowCxToolVersionToolOpenApiSpecAuthenticationBearerTokenConfig struct {
	// Optional.
	//
	// The name of the SecretManager secret version resource storing the Bearer token. If this field is set, the 'token' field will be ignored.
	// Format: projects/{project}/secrets/{secret}/versions/{version}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#secret_version_for_token GoogleDialogflowCxToolVersion#secret_version_for_token}
	SecretVersionForToken *string `field:"optional" json:"secretVersionForToken" yaml:"secretVersionForToken"`
	// Optional.
	//
	// The text token appended to the text Bearer to the request Authorization header.
	// [Session parameters reference](https://cloud.google.com/dialogflow/cx/docs/concept/parameter#session-ref) can be used to pass the token dynamically, e.g. '$session.params.parameter-id'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#token GoogleDialogflowCxToolVersion#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
}

