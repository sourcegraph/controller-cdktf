package apigatewayv2integration


type Apigatewayv2IntegrationResponseParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/apigatewayv2_integration#mappings Apigatewayv2Integration#mappings}.
	Mappings *map[string]*string `field:"required" json:"mappings" yaml:"mappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/apigatewayv2_integration#status_code Apigatewayv2Integration#status_code}.
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
}

