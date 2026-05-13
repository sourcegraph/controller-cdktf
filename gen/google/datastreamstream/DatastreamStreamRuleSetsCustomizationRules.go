package datastreamstream


type DatastreamStreamRuleSetsCustomizationRules struct {
	// bigquery_clustering block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#bigquery_clustering DatastreamStream#bigquery_clustering}
	BigqueryClustering *DatastreamStreamRuleSetsCustomizationRulesBigqueryClustering `field:"optional" json:"bigqueryClustering" yaml:"bigqueryClustering"`
	// bigquery_partitioning block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#bigquery_partitioning DatastreamStream#bigquery_partitioning}
	BigqueryPartitioning *DatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioning `field:"optional" json:"bigqueryPartitioning" yaml:"bigqueryPartitioning"`
}

