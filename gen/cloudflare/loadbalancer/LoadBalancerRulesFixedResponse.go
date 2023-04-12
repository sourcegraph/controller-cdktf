package loadbalancer


type LoadBalancerRulesFixedResponse struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#content_type LoadBalancer#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#location LoadBalancer#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#message_body LoadBalancer#message_body}.
	MessageBody *string `field:"optional" json:"messageBody" yaml:"messageBody"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#status_code LoadBalancer#status_code}.
	StatusCode *float64 `field:"optional" json:"statusCode" yaml:"statusCode"`
}

