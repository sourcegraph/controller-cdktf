package googledataprocjob


type GoogleDataprocJobSparksqlConfig struct {
	// HCFS URIs of jar files to be added to the Spark CLASSPATH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#jar_file_uris GoogleDataprocJob#jar_file_uris}
	JarFileUris *[]*string `field:"optional" json:"jarFileUris" yaml:"jarFileUris"`
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#logging_config GoogleDataprocJob#logging_config}
	LoggingConfig *GoogleDataprocJobSparksqlConfigLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// A mapping of property names to values, used to configure Spark SQL's SparkConf.
	//
	// Properties that conflict with values set by the Cloud Dataproc API may be overwritten.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#properties GoogleDataprocJob#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// The HCFS URI of the script that contains SQL queries. Conflicts with query_list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#query_file_uri GoogleDataprocJob#query_file_uri}
	QueryFileUri *string `field:"optional" json:"queryFileUri" yaml:"queryFileUri"`
	// The list of SQL queries or statements to execute as part of the job. Conflicts with query_file_uri.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#query_list GoogleDataprocJob#query_list}
	QueryList *[]*string `field:"optional" json:"queryList" yaml:"queryList"`
	// Mapping of query variable names to values (equivalent to the Spark SQL command: SET name="value";).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#script_variables GoogleDataprocJob#script_variables}
	ScriptVariables *map[string]*string `field:"optional" json:"scriptVariables" yaml:"scriptVariables"`
}

