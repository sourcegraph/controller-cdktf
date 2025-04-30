package googledataprocbatch


type GoogleDataprocBatchSparkSqlBatch struct {
	// HCFS URIs of jar files to be added to the Spark CLASSPATH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_batch#jar_file_uris GoogleDataprocBatch#jar_file_uris}
	JarFileUris *[]*string `field:"optional" json:"jarFileUris" yaml:"jarFileUris"`
	// The HCFS URI of the script that contains Spark SQL queries to execute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_batch#query_file_uri GoogleDataprocBatch#query_file_uri}
	QueryFileUri *string `field:"optional" json:"queryFileUri" yaml:"queryFileUri"`
	// Mapping of query variable names to values (equivalent to the Spark SQL command: SET name="value";).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_batch#query_variables GoogleDataprocBatch#query_variables}
	QueryVariables *map[string]*string `field:"optional" json:"queryVariables" yaml:"queryVariables"`
}

