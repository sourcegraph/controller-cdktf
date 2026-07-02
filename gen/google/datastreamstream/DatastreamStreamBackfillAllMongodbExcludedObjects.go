package datastreamstream

type DatastreamStreamBackfillAllMongodbExcludedObjects struct {
	// databases block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/datastream_stream#databases DatastreamStream#databases}
	Databases any `field:"required" json:"databases" yaml:"databases"`
}
