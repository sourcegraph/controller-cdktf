package datastreamstream

type DatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabases struct {
	// Database name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/datastream_stream#database DatastreamStream#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// mysql_tables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/datastream_stream#mysql_tables DatastreamStream#mysql_tables}
	MysqlTables any `field:"optional" json:"mysqlTables" yaml:"mysqlTables"`
}
