package apigatewayusageplan


type ApiGatewayUsagePlanApiStagesThrottle struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan#path ApiGatewayUsagePlan#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan#burst_limit ApiGatewayUsagePlan#burst_limit}.
	BurstLimit *float64 `field:"optional" json:"burstLimit" yaml:"burstLimit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_usage_plan#rate_limit ApiGatewayUsagePlan#rate_limit}.
	RateLimit *float64 `field:"optional" json:"rateLimit" yaml:"rateLimit"`
}
