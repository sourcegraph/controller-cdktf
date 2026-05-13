package datastreamstream


type DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMysqlIdentifier struct {
	// The database name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#database DatastreamStream#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// The table name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#table DatastreamStream#table}
	Table *string `field:"required" json:"table" yaml:"table"`
}

