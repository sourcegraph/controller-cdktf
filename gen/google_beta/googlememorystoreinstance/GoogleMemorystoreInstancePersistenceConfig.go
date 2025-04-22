package googlememorystoreinstance


type GoogleMemorystoreInstancePersistenceConfig struct {
	// aof_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_memorystore_instance#aof_config GoogleMemorystoreInstance#aof_config}
	AofConfig *GoogleMemorystoreInstancePersistenceConfigAofConfig `field:"optional" json:"aofConfig" yaml:"aofConfig"`
	// Optional. Current persistence mode.   Possible values: DISABLED RDB AOF Possible values: ["DISABLED", "RDB", "AOF"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_memorystore_instance#mode GoogleMemorystoreInstance#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// rdb_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_memorystore_instance#rdb_config GoogleMemorystoreInstance#rdb_config}
	RdbConfig *GoogleMemorystoreInstancePersistenceConfigRdbConfig `field:"optional" json:"rdbConfig" yaml:"rdbConfig"`
}

