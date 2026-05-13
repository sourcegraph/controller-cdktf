package googledialogflowcxtool


type GoogleDialogflowCxToolOpenApiSpec struct {
	// The OpenAPI schema specified as a text.
	//
	// This field is part of a union field 'schema': only one of 'textSchema' may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool#text_schema GoogleDialogflowCxTool#text_schema}
	TextSchema *string `field:"required" json:"textSchema" yaml:"textSchema"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool#authentication GoogleDialogflowCxTool#authentication}
	Authentication *GoogleDialogflowCxToolOpenApiSpecAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// service_directory_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool#service_directory_config GoogleDialogflowCxTool#service_directory_config}
	ServiceDirectoryConfig *GoogleDialogflowCxToolOpenApiSpecServiceDirectoryConfig `field:"optional" json:"serviceDirectoryConfig" yaml:"serviceDirectoryConfig"`
	// tls_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_tool#tls_config GoogleDialogflowCxTool#tls_config}
	TlsConfig *GoogleDialogflowCxToolOpenApiSpecTlsConfig `field:"optional" json:"tlsConfig" yaml:"tlsConfig"`
}

