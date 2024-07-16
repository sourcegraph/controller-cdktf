package googlebigqueryroutine


type GoogleBigqueryRoutineRemoteFunctionOptions struct {
	// Fully qualified name of the user-provided connection object which holds the authentication information to send requests to the remote service.
	//
	// Format: "projects/{projectId}/locations/{locationId}/connections/{connectionId}"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_routine#connection GoogleBigqueryRoutine#connection}
	Connection *string `field:"optional" json:"connection" yaml:"connection"`
	// Endpoint of the user-provided remote service, e.g. 'https://us-east1-my_gcf_project.cloudfunctions.net/remote_add'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_routine#endpoint GoogleBigqueryRoutine#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Max number of rows in each batch sent to the remote service.
	//
	// If absent or if 0,
	// BigQuery dynamically decides the number of rows in a batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_routine#max_batching_rows GoogleBigqueryRoutine#max_batching_rows}
	MaxBatchingRows *string `field:"optional" json:"maxBatchingRows" yaml:"maxBatchingRows"`
	// User-defined context as a set of key/value pairs, which will be sent as function invocation context together with batched arguments in the requests to the remote service.
	//
	// The total number of bytes of keys and values must be less than 8KB.
	//
	// An object containing a list of "key": value pairs. Example:
	// '{ "name": "wrench", "mass": "1.3kg", "count": "3" }'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_routine#user_defined_context GoogleBigqueryRoutine#user_defined_context}
	UserDefinedContext *map[string]*string `field:"optional" json:"userDefinedContext" yaml:"userDefinedContext"`
}

