package cestoolset


type CesToolsetOpenApiToolset struct {
	// The OpenAPI schema of the toolset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_toolset#open_api_schema CesToolset#open_api_schema}
	OpenApiSchema *string `field:"required" json:"openApiSchema" yaml:"openApiSchema"`
	// api_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_toolset#api_authentication CesToolset#api_authentication}
	ApiAuthentication *CesToolsetOpenApiToolsetApiAuthentication `field:"optional" json:"apiAuthentication" yaml:"apiAuthentication"`
	// If true, the agent will ignore unknown fields in the API response for all operations defined in the OpenAPI schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_toolset#ignore_unknown_fields CesToolset#ignore_unknown_fields}
	IgnoreUnknownFields interface{} `field:"optional" json:"ignoreUnknownFields" yaml:"ignoreUnknownFields"`
	// service_directory_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_toolset#service_directory_config CesToolset#service_directory_config}
	ServiceDirectoryConfig *CesToolsetOpenApiToolsetServiceDirectoryConfig `field:"optional" json:"serviceDirectoryConfig" yaml:"serviceDirectoryConfig"`
	// tls_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_toolset#tls_config CesToolset#tls_config}
	TlsConfig *CesToolsetOpenApiToolsetTlsConfig `field:"optional" json:"tlsConfig" yaml:"tlsConfig"`
}

