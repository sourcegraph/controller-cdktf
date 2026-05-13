package googledatastreamstream


type GoogleDatastreamStreamBackfillAllMongodbExcludedObjectsDatabasesCollections struct {
	// Collection name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#collection GoogleDatastreamStream#collection}
	Collection *string `field:"required" json:"collection" yaml:"collection"`
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#fields GoogleDatastreamStream#fields}
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
}

