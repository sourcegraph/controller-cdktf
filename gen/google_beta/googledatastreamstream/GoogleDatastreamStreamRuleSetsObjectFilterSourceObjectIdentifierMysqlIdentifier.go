package googledatastreamstream


type GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMysqlIdentifier struct {
	// The database name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#database GoogleDatastreamStream#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// The table name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#table GoogleDatastreamStream#table}
	Table *string `field:"required" json:"table" yaml:"table"`
}

