package googlecestoolset


type GoogleCesToolsetMcpToolsetApiAuthenticationApiKeyConfig struct {
	// The name of the SecretManager secret version resource storing the API key.
	//
	// Format: 'projects/{project}/secrets/{secret}/versions/{version}'
	// Note: You should grant 'roles/secretmanager.secretAccessor' role to the CES
	// service agent
	// 'service-@gcp-sa-ces.iam.gserviceaccount.com'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#api_key_secret_version GoogleCesToolset#api_key_secret_version}
	ApiKeySecretVersion *string `field:"required" json:"apiKeySecretVersion" yaml:"apiKeySecretVersion"`
	// The parameter name or the header name of the API key.
	//
	// E.g., If the API request is "https://example.com/act?X-Api-Key=", "X-Api-Key" would be the parameter name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#key_name GoogleCesToolset#key_name}
	KeyName *string `field:"required" json:"keyName" yaml:"keyName"`
	// Key location in the request.
	//
	// For API key auth on MCP toolsets,
	// the API key can only be sent in the request header.
	// Possible values:
	// HEADER
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_toolset#request_location GoogleCesToolset#request_location}
	RequestLocation *string `field:"required" json:"requestLocation" yaml:"requestLocation"`
}

