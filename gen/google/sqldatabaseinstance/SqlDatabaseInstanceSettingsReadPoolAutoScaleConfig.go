package sqldatabaseinstance


type SqlDatabaseInstanceSettingsReadPoolAutoScaleConfig struct {
	// True if auto scale in is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/sql_database_instance#disable_scale_in SqlDatabaseInstance#disable_scale_in}
	DisableScaleIn interface{} `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// True if Read Pool Auto Scale is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/sql_database_instance#enabled SqlDatabaseInstance#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Maximum number of nodes in the read pool.
	//
	// If set to lower than current node count, node count will be updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/sql_database_instance#max_node_count SqlDatabaseInstance#max_node_count}
	MaxNodeCount *float64 `field:"optional" json:"maxNodeCount" yaml:"maxNodeCount"`
	// Minimum number of nodes in the read pool.
	//
	// If set to higher than current node count, node count will be updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/sql_database_instance#min_node_count SqlDatabaseInstance#min_node_count}
	MinNodeCount *float64 `field:"optional" json:"minNodeCount" yaml:"minNodeCount"`
	// The cooldown period for scale in operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/sql_database_instance#scale_in_cooldown_seconds SqlDatabaseInstance#scale_in_cooldown_seconds}
	ScaleInCooldownSeconds *float64 `field:"optional" json:"scaleInCooldownSeconds" yaml:"scaleInCooldownSeconds"`
	// The cooldown period for scale out operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/sql_database_instance#scale_out_cooldown_seconds SqlDatabaseInstance#scale_out_cooldown_seconds}
	ScaleOutCooldownSeconds *float64 `field:"optional" json:"scaleOutCooldownSeconds" yaml:"scaleOutCooldownSeconds"`
	// target_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/sql_database_instance#target_metrics SqlDatabaseInstance#target_metrics}
	TargetMetrics interface{} `field:"optional" json:"targetMetrics" yaml:"targetMetrics"`
}

