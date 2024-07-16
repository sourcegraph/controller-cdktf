package googlebigqueryjob


type GoogleBigqueryJobQueryUserDefinedFunctionResources struct {
	// An inline resource that contains code for a user-defined function (UDF).
	//
	// Providing a inline code resource is equivalent to providing a URI for a file containing the same code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_job#inline_code GoogleBigqueryJob#inline_code}
	InlineCode *string `field:"optional" json:"inlineCode" yaml:"inlineCode"`
	// A code resource to load from a Google Cloud Storage URI (gs://bucket/path).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_job#resource_uri GoogleBigqueryJob#resource_uri}
	ResourceUri *string `field:"optional" json:"resourceUri" yaml:"resourceUri"`
}

