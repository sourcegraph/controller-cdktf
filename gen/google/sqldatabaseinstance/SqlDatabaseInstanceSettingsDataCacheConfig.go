package sqldatabaseinstance

type SqlDatabaseInstanceSettingsDataCacheConfig struct {
	// Whether data cache is enabled for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/sql_database_instance#data_cache_enabled SqlDatabaseInstance#data_cache_enabled}
	DataCacheEnabled any `field:"optional" json:"dataCacheEnabled" yaml:"dataCacheEnabled"`
}
