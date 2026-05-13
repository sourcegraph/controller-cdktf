package googledatastreamstream


type GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSpannerIdentifier struct {
	// The table name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#table GoogleDatastreamStream#table}
	Table *string `field:"required" json:"table" yaml:"table"`
	// The schema name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#schema GoogleDatastreamStream#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

