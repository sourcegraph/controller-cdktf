package apigatewayintegration

type ApiGatewayIntegrationTlsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/api_gateway_integration#insecure_skip_verification ApiGatewayIntegration#insecure_skip_verification}.
	InsecureSkipVerification any `field:"optional" json:"insecureSkipVerification" yaml:"insecureSkipVerification"`
}
