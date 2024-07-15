package googledataprocworkflowtemplate


type GoogleDataprocWorkflowTemplateJobsSparkSqlJob struct {
	// Optional. HCFS URIs of jar files to be added to the Spark CLASSPATH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#jar_file_uris GoogleDataprocWorkflowTemplate#jar_file_uris}
	JarFileUris *[]*string `field:"optional" json:"jarFileUris" yaml:"jarFileUris"`
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#logging_config GoogleDataprocWorkflowTemplate#logging_config}
	LoggingConfig *GoogleDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// Optional.
	//
	// A mapping of property names to values, used to configure Spark SQL's SparkConf. Properties that conflict with values set by the Dataproc API may be overwritten.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#properties GoogleDataprocWorkflowTemplate#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// The HCFS URI of the script that contains SQL queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#query_file_uri GoogleDataprocWorkflowTemplate#query_file_uri}
	QueryFileUri *string `field:"optional" json:"queryFileUri" yaml:"queryFileUri"`
	// query_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#query_list GoogleDataprocWorkflowTemplate#query_list}
	QueryList *GoogleDataprocWorkflowTemplateJobsSparkSqlJobQueryListStruct `field:"optional" json:"queryList" yaml:"queryList"`
	// Optional. Mapping of query variable names to values (equivalent to the Spark SQL command: SET `name="value";`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#script_variables GoogleDataprocWorkflowTemplate#script_variables}
	ScriptVariables *map[string]*string `field:"optional" json:"scriptVariables" yaml:"scriptVariables"`
}

