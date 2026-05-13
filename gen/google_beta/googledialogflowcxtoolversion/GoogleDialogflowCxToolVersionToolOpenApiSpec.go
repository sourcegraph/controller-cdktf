package googledialogflowcxtoolversion


type GoogleDialogflowCxToolVersionToolOpenApiSpec struct {
	// The OpenAPI schema specified as a text.
	//
	// This field is part of a union field 'schema': only one of 'textSchema' may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#text_schema GoogleDialogflowCxToolVersion#text_schema}
	TextSchema *string `field:"required" json:"textSchema" yaml:"textSchema"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#authentication GoogleDialogflowCxToolVersion#authentication}
	Authentication *GoogleDialogflowCxToolVersionToolOpenApiSpecAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// service_directory_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#service_directory_config GoogleDialogflowCxToolVersion#service_directory_config}
	ServiceDirectoryConfig *GoogleDialogflowCxToolVersionToolOpenApiSpecServiceDirectoryConfig `field:"optional" json:"serviceDirectoryConfig" yaml:"serviceDirectoryConfig"`
	// tls_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool_version#tls_config GoogleDialogflowCxToolVersion#tls_config}
	TlsConfig *GoogleDialogflowCxToolVersionToolOpenApiSpecTlsConfig `field:"optional" json:"tlsConfig" yaml:"tlsConfig"`
}

