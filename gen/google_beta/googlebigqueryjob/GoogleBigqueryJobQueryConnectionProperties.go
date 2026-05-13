package googlebigqueryjob


type GoogleBigqueryJobQueryConnectionProperties struct {
	// The key of the property to set.
	//
	// Currently supported connection properties:
	// * 'dataset_project_id': represents the default project for datasets that are used in the query
	// * 'time_zone': represents the default timezone used to run the query
	// * 'session_id': associates the query with a given session
	// * 'query_label': associates the query with a given job label
	// * 'service_account': indicates the service account to use to run a continuous query
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_job#key GoogleBigqueryJob#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of the property to set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_job#value GoogleBigqueryJob#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

