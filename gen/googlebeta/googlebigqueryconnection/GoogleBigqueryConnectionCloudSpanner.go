package googlebigqueryconnection


type GoogleBigqueryConnectionCloudSpanner struct {
	// Cloud Spanner database in the form 'project/instance/database'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_connection#database GoogleBigqueryConnection#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// If parallelism should be used when reading from Cloud Spanner.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_connection#use_parallelism GoogleBigqueryConnection#use_parallelism}
	UseParallelism interface{} `field:"optional" json:"useParallelism" yaml:"useParallelism"`
}

