package alloydbinstance


type AlloydbInstanceConnectionPoolConfig struct {
	// Whether to enabled Managed Connection Pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/alloydb_instance#enabled AlloydbInstance#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Flags for configuring managed connection pooling when it is enabled.
	//
	// These flags will only be set if 'connection_pool_config.enabled' is
	// true.
	// Please see
	// https://cloud.google.com/alloydb/docs/configure-managed-connection-pooling#configuration-options
	// for a comprehensive list of flags that can be set. To specify the flags
	// in Terraform, please remove the "connection-pooling-" prefix and use
	// underscores instead of dashes in the name. For example,
	// "connection-pooling-pool-mode" would be "pool_mode".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/alloydb_instance#flags AlloydbInstance#flags}
	Flags *map[string]*string `field:"optional" json:"flags" yaml:"flags"`
}

