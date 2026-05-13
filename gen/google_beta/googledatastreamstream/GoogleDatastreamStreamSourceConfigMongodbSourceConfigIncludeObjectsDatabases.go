package googledatastreamstream


type GoogleDatastreamStreamSourceConfigMongodbSourceConfigIncludeObjectsDatabases struct {
	// collections block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#collections GoogleDatastreamStream#collections}
	Collections interface{} `field:"optional" json:"collections" yaml:"collections"`
	// Database name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#database GoogleDatastreamStream#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
}

