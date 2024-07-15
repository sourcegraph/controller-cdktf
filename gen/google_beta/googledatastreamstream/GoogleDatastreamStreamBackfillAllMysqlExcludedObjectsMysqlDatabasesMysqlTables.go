package googledatastreamstream


type GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabasesMysqlTables struct {
	// Table name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#table GoogleDatastreamStream#table}
	Table *string `field:"required" json:"table" yaml:"table"`
	// mysql_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#mysql_columns GoogleDatastreamStream#mysql_columns}
	MysqlColumns interface{} `field:"optional" json:"mysqlColumns" yaml:"mysqlColumns"`
}

