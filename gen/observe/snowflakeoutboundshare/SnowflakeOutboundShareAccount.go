package snowflakeoutboundshare


type SnowflakeOutboundShareAccount struct {
	// The name of the Snowflake account to share with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/snowflake_outbound_share#account SnowflakeOutboundShare#account}
	Account *string `field:"required" json:"account" yaml:"account"`
	// The name of the Snowflake organization to share with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/snowflake_outbound_share#organization SnowflakeOutboundShare#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
}

