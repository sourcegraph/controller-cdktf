package googlebigquerydatasetaccess


type GoogleBigqueryDatasetAccessDatasetA struct {
	// dataset block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_dataset_access#dataset GoogleBigqueryDatasetAccessA#dataset}
	Dataset *GoogleBigqueryDatasetAccessDatasetDatasetA `field:"required" json:"dataset" yaml:"dataset"`
	// Which resources in the dataset this entry applies to.
	//
	// Currently, only views are supported,
	// but additional target types may be added in the future. Possible values: VIEWS
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_dataset_access#target_types GoogleBigqueryDatasetAccessA#target_types}
	TargetTypes *[]*string `field:"required" json:"targetTypes" yaml:"targetTypes"`
}

