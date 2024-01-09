package agent


type AgentDatadogConfig struct {
	// `com` or `eu`, Datadog SaaS instance, which corresponds to one of Datadog's two locations (https://www.datadoghq.com/ in the U.S. or https://datadoghq.eu/ in the European Union).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#site Agent#site}
	Site *string `field:"required" json:"site" yaml:"site"`
}

