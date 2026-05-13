package datastreamstream


type DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSpannerIdentifier struct {
	// The table name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#table DatastreamStream#table}
	Table *string `field:"required" json:"table" yaml:"table"`
	// The schema name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#schema DatastreamStream#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

