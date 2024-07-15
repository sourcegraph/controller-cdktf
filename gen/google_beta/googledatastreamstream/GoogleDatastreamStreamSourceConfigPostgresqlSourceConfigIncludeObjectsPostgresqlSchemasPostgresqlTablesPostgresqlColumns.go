package googledatastreamstream


type GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumns struct {
	// Column name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#column GoogleDatastreamStream#column}
	Column *string `field:"optional" json:"column" yaml:"column"`
	// The PostgreSQL data type. Full data types list can be found here: https://www.postgresql.org/docs/current/datatype.html.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#data_type GoogleDatastreamStream#data_type}
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// Whether or not the column can accept a null value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#nullable GoogleDatastreamStream#nullable}
	Nullable interface{} `field:"optional" json:"nullable" yaml:"nullable"`
	// The ordinal position of the column in the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#ordinal_position GoogleDatastreamStream#ordinal_position}
	OrdinalPosition *float64 `field:"optional" json:"ordinalPosition" yaml:"ordinalPosition"`
	// Whether or not the column represents a primary key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#primary_key GoogleDatastreamStream#primary_key}
	PrimaryKey interface{} `field:"optional" json:"primaryKey" yaml:"primaryKey"`
}

