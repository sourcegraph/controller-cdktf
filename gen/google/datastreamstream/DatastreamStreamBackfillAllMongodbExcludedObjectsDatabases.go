package datastreamstream


type DatastreamStreamBackfillAllMongodbExcludedObjectsDatabases struct {
	// Database name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#database DatastreamStream#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// collections block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#collections DatastreamStream#collections}
	Collections interface{} `field:"optional" json:"collections" yaml:"collections"`
}

