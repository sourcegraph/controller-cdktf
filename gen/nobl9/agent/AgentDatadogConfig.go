package agent


type AgentDatadogConfig struct {
	// Datadog SaaS instance that corresponds to one of Datadog's available locations (see [Datadog docs](https://docs.datadoghq.com/getting_started/site/#access-the-datadog-site)for an up to date list):   - `datadoghq.com` (formerly referred to as `com`)   - `us3.datadoghq.com`   - `us5.datadoghq.com`   - `datadoghq.eu` (formerly referred to as `eu`)   - `ddog-gov.com`   - `ap1.datadoghq.com`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/agent#site Agent#site}
	Site *string `field:"required" json:"site" yaml:"site"`
}

