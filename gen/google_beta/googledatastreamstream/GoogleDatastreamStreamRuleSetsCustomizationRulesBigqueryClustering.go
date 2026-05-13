package googledatastreamstream


type GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryClustering struct {
	// Column names to set as clustering columns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#columns GoogleDatastreamStream#columns}
	Columns *[]*string `field:"required" json:"columns" yaml:"columns"`
}

