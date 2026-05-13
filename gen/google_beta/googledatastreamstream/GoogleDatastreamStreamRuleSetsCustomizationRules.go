package googledatastreamstream


type GoogleDatastreamStreamRuleSetsCustomizationRules struct {
	// bigquery_clustering block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#bigquery_clustering GoogleDatastreamStream#bigquery_clustering}
	BigqueryClustering *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryClustering `field:"optional" json:"bigqueryClustering" yaml:"bigqueryClustering"`
	// bigquery_partitioning block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#bigquery_partitioning GoogleDatastreamStream#bigquery_partitioning}
	BigqueryPartitioning *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioning `field:"optional" json:"bigqueryPartitioning" yaml:"bigqueryPartitioning"`
}

