package googlecestoolset


type GoogleCesToolsetOpenApiToolset struct {
	// The OpenAPI schema of the toolset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#open_api_schema GoogleCesToolset#open_api_schema}
	OpenApiSchema *string `field:"required" json:"openApiSchema" yaml:"openApiSchema"`
	// api_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#api_authentication GoogleCesToolset#api_authentication}
	ApiAuthentication *GoogleCesToolsetOpenApiToolsetApiAuthentication `field:"optional" json:"apiAuthentication" yaml:"apiAuthentication"`
	// If true, the agent will ignore unknown fields in the API response for all operations defined in the OpenAPI schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#ignore_unknown_fields GoogleCesToolset#ignore_unknown_fields}
	IgnoreUnknownFields interface{} `field:"optional" json:"ignoreUnknownFields" yaml:"ignoreUnknownFields"`
	// service_directory_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#service_directory_config GoogleCesToolset#service_directory_config}
	ServiceDirectoryConfig *GoogleCesToolsetOpenApiToolsetServiceDirectoryConfig `field:"optional" json:"serviceDirectoryConfig" yaml:"serviceDirectoryConfig"`
	// tls_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#tls_config GoogleCesToolset#tls_config}
	TlsConfig *GoogleCesToolsetOpenApiToolsetTlsConfig `field:"optional" json:"tlsConfig" yaml:"tlsConfig"`
}

