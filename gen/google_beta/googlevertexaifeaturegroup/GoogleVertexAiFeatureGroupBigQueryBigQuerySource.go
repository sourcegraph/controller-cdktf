package googlevertexaifeaturegroup


type GoogleVertexAiFeatureGroupBigQueryBigQuerySource struct {
	// BigQuery URI to a table, up to 2000 characters long. For example: 'bq://projectId.bqDatasetId.bqTableId.'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_group#input_uri GoogleVertexAiFeatureGroup#input_uri}
	InputUri *string `field:"required" json:"inputUri" yaml:"inputUri"`
}

