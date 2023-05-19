package loadbalancer


type LoadBalancerRulesFixedResponse struct {
	// The value of the HTTP context-type header for this fixed response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#content_type LoadBalancer#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The value of the HTTP location header for this fixed response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#location LoadBalancer#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The text used as the html body for this fixed response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#message_body LoadBalancer#message_body}
	MessageBody *string `field:"optional" json:"messageBody" yaml:"messageBody"`
	// The HTTP status code used for this fixed response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#status_code LoadBalancer#status_code}
	StatusCode *float64 `field:"optional" json:"statusCode" yaml:"statusCode"`
}

