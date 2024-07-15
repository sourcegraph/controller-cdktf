package googledatastreamstream


type GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemasTablesColumns struct {
	// Column name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#column GoogleDatastreamStream#column}
	Column *string `field:"optional" json:"column" yaml:"column"`
	// The SQL Server data type. Full data types list can be found here: https://learn.microsoft.com/en-us/sql/t-sql/data-types/data-types-transact-sql?view=sql-server-ver16.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#data_type GoogleDatastreamStream#data_type}
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
}

