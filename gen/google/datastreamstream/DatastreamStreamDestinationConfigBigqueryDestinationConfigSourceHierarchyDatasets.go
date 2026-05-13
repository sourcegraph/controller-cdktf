package datastreamstream


type DatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasets struct {
	// dataset_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#dataset_template DatastreamStream#dataset_template}
	DatasetTemplate *DatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasetsDatasetTemplate `field:"required" json:"datasetTemplate" yaml:"datasetTemplate"`
	// Optional.
	//
	// The project id of the BigQuery dataset. If not specified, the project will be inferred from the stream resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#project_id DatastreamStream#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

