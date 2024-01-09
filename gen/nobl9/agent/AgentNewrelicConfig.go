package agent


type AgentNewrelicConfig struct {
	// ID number assigned to the New Relic user account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/agent#account_id Agent#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
}

