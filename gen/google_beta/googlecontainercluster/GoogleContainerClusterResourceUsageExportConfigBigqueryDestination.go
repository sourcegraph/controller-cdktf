package googlecontainercluster


type GoogleContainerClusterResourceUsageExportConfigBigqueryDestination struct {
	// The ID of a BigQuery Dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#dataset_id GoogleContainerCluster#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
}

