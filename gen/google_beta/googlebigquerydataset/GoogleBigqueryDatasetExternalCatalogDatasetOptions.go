package googlebigquerydataset


type GoogleBigqueryDatasetExternalCatalogDatasetOptions struct {
	// The storage location URI for all tables in the dataset.
	//
	// Equivalent to hive metastore's
	// database locationUri. Maximum length of 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_bigquery_dataset#default_storage_location_uri GoogleBigqueryDataset#default_storage_location_uri}
	DefaultStorageLocationUri *string `field:"optional" json:"defaultStorageLocationUri" yaml:"defaultStorageLocationUri"`
	// A map of key value pairs defining the parameters and properties of the open source schema. Maximum size of 2Mib.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_bigquery_dataset#parameters GoogleBigqueryDataset#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

