package sqldatabaseinstance

type SqlDatabaseInstanceSettingsConnectionPoolConfig struct {
	// Whether Managed Connection Pool is enabled for this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/sql_database_instance#connection_pooling_enabled SqlDatabaseInstance#connection_pooling_enabled}
	ConnectionPoolingEnabled any `field:"optional" json:"connectionPoolingEnabled" yaml:"connectionPoolingEnabled"`
	// flags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/sql_database_instance#flags SqlDatabaseInstance#flags}
	Flags any `field:"optional" json:"flags" yaml:"flags"`
}
