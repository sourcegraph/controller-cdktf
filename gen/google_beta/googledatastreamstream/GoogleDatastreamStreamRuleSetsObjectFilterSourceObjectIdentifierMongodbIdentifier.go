package googledatastreamstream


type GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMongodbIdentifier struct {
	// The MongoDB collection name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#collection GoogleDatastreamStream#collection}
	Collection *string `field:"required" json:"collection" yaml:"collection"`
	// The MongoDB database name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#database GoogleDatastreamStream#database}
	Database *string `field:"required" json:"database" yaml:"database"`
}

