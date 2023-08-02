package googlebigqueryconnection


type GoogleBigqueryConnectionCloudSpanner struct {
	// Cloud Spanner database in the form 'project/instance/database'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_bigquery_connection#database GoogleBigqueryConnection#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// If parallelism should be used when reading from Cloud Spanner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_bigquery_connection#use_parallelism GoogleBigqueryConnection#use_parallelism}
	UseParallelism interface{} `field:"optional" json:"useParallelism" yaml:"useParallelism"`
	// If the serverless analytics service should be used to read data from Cloud Spanner.
	//
	// useParallelism must be set when using serverless analytics
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_bigquery_connection#use_serverless_analytics GoogleBigqueryConnection#use_serverless_analytics}
	UseServerlessAnalytics interface{} `field:"optional" json:"useServerlessAnalytics" yaml:"useServerlessAnalytics"`
}

