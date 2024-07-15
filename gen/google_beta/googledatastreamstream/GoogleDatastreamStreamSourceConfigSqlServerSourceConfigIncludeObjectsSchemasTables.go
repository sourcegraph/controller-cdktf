package googledatastreamstream


type GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemasTables struct {
	// Table name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#table GoogleDatastreamStream#table}
	Table *string `field:"required" json:"table" yaml:"table"`
	// columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#columns GoogleDatastreamStream#columns}
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
}

