package datastreamstream


type DatastreamStreamBackfillAllMongodbExcludedObjects struct {
	// databases block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#databases DatastreamStream#databases}
	Databases interface{} `field:"required" json:"databases" yaml:"databases"`
}

