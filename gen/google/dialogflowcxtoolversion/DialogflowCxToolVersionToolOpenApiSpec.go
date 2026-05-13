package dialogflowcxtoolversion


type DialogflowCxToolVersionToolOpenApiSpec struct {
	// The OpenAPI schema specified as a text.
	//
	// This field is part of a union field 'schema': only one of 'textSchema' may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_tool_version#text_schema DialogflowCxToolVersion#text_schema}
	TextSchema *string `field:"required" json:"textSchema" yaml:"textSchema"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_tool_version#authentication DialogflowCxToolVersion#authentication}
	Authentication *DialogflowCxToolVersionToolOpenApiSpecAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// service_directory_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_tool_version#service_directory_config DialogflowCxToolVersion#service_directory_config}
	ServiceDirectoryConfig *DialogflowCxToolVersionToolOpenApiSpecServiceDirectoryConfig `field:"optional" json:"serviceDirectoryConfig" yaml:"serviceDirectoryConfig"`
	// tls_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dialogflow_cx_tool_version#tls_config DialogflowCxToolVersion#tls_config}
	TlsConfig *DialogflowCxToolVersionToolOpenApiSpecTlsConfig `field:"optional" json:"tlsConfig" yaml:"tlsConfig"`
}

