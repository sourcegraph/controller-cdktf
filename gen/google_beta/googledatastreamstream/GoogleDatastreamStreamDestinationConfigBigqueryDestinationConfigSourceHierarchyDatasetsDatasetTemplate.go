package googledatastreamstream


type GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasetsDatasetTemplate struct {
	// The geographic location where the dataset should reside. See https://cloud.google.com/bigquery/docs/locations for supported locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#location GoogleDatastreamStream#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// If supplied, every created dataset will have its name prefixed by the provided value.
	//
	// The prefix and name will be separated by an underscore. i.e. _.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#dataset_id_prefix GoogleDatastreamStream#dataset_id_prefix}
	DatasetIdPrefix *string `field:"optional" json:"datasetIdPrefix" yaml:"datasetIdPrefix"`
	// Describes the Cloud KMS encryption key that will be used to protect destination BigQuery table.
	//
	// The BigQuery Service Account associated with your project requires access to this
	// encryption key. i.e. projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{cryptoKey}.
	// See https://cloud.google.com/bigquery/docs/customer-managed-encryption for more information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_stream#kms_key_name GoogleDatastreamStream#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
}

