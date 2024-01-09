package agent


type AgentAmazonPrometheusConfig struct {
	// AWS region e.g., eu-central-1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#region Agent#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// Base URL to Amazon Prometheus server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#url Agent#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

