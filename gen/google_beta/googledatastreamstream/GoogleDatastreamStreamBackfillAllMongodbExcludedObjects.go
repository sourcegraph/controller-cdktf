package googledatastreamstream


type GoogleDatastreamStreamBackfillAllMongodbExcludedObjects struct {
	// databases block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#databases GoogleDatastreamStream#databases}
	Databases interface{} `field:"required" json:"databases" yaml:"databases"`
}

