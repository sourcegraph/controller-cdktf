package agent


type AgentAzureMonitorConfig struct {
	// Azure Tenant Id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#tenant_id Agent#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
}

