package cestoolset


type CesToolsetOpenApiToolsetApiAuthenticationApiKeyConfig struct {
	// The name of the SecretManager secret version resource storing the API key.
	//
	// Format: 'projects/{project}/secrets/{secret}/versions/{version}'
	// Note: You should grant 'roles/secretmanager.secretAccessor' role to the CES
	// service agent
	// 'service-@gcp-sa-ces.iam.gserviceaccount.com'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_toolset#api_key_secret_version CesToolset#api_key_secret_version}
	ApiKeySecretVersion *string `field:"required" json:"apiKeySecretVersion" yaml:"apiKeySecretVersion"`
	// The parameter name or the header name of the API key.
	//
	// E.g., If the API request is "https://example.com/act?X-Api-Key=", "X-Api-Key" would be the parameter name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_toolset#key_name CesToolset#key_name}
	KeyName *string `field:"required" json:"keyName" yaml:"keyName"`
	// Key location in the request. Possible values: HEADER QUERY_STRING.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_toolset#request_location CesToolset#request_location}
	RequestLocation *string `field:"required" json:"requestLocation" yaml:"requestLocation"`
}

