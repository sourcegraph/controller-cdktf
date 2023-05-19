package googledatastreamstream


type GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabases struct {
	// Database name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_datastream_stream#database GoogleDatastreamStream#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// mysql_tables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_datastream_stream#mysql_tables GoogleDatastreamStream#mysql_tables}
	MysqlTables interface{} `field:"optional" json:"mysqlTables" yaml:"mysqlTables"`
}

