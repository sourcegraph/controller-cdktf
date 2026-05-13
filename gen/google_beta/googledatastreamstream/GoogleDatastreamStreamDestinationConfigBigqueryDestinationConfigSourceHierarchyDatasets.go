package googledatastreamstream


type GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasets struct {
	// dataset_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#dataset_template GoogleDatastreamStream#dataset_template}
	DatasetTemplate *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasetsDatasetTemplate `field:"required" json:"datasetTemplate" yaml:"datasetTemplate"`
	// Optional.
	//
	// The project id of the BigQuery dataset. If not specified, the project will be inferred from the stream resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#project_id GoogleDatastreamStream#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

