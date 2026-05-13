package cesapp


type CesAppLoggingSettingsBigqueryExportSettings struct {
	// The BigQuery dataset to export the data to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#dataset CesApp#dataset}
	Dataset *string `field:"optional" json:"dataset" yaml:"dataset"`
	// Indicates whether the BigQuery export is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#enabled CesApp#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The project ID of the BigQuery dataset to export the data to.
	//
	// Note: If the BigQuery dataset is in a different project from the app, you should grant
	// roles/bigquery.admin role to the CES service agent service-<PROJECT-
	// NUMBER>@gcp-sa-ces.iam.gserviceaccount.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#project CesApp#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

