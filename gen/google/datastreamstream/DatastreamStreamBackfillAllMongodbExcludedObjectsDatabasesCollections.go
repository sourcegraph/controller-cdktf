package datastreamstream


type DatastreamStreamBackfillAllMongodbExcludedObjectsDatabasesCollections struct {
	// Collection name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#collection DatastreamStream#collection}
	Collection *string `field:"required" json:"collection" yaml:"collection"`
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#fields DatastreamStream#fields}
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
}

