package agent


type AgentGrafanaLokiConfig struct {
	// API URL endpoint to the Grafana Loki instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#url Agent#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

