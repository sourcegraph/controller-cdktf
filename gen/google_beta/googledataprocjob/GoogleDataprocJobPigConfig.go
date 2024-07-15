package googledataprocjob


type GoogleDataprocJobPigConfig struct {
	// Whether to continue executing queries if a query fails.
	//
	// The default value is false. Setting to true can be useful when executing independent parallel queries. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#continue_on_failure GoogleDataprocJob#continue_on_failure}
	ContinueOnFailure interface{} `field:"optional" json:"continueOnFailure" yaml:"continueOnFailure"`
	// HCFS URIs of jar files to add to the CLASSPATH of the Pig Client and Hadoop MapReduce (MR) tasks.
	//
	// Can contain Pig UDFs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#jar_file_uris GoogleDataprocJob#jar_file_uris}
	JarFileUris *[]*string `field:"optional" json:"jarFileUris" yaml:"jarFileUris"`
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#logging_config GoogleDataprocJob#logging_config}
	LoggingConfig *GoogleDataprocJobPigConfigLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// A mapping of property names to values, used to configure Pig.
	//
	// Properties that conflict with values set by the Cloud Dataproc API may be overwritten. Can include properties set in /etc/hadoop/conf/*-site.xml, /etc/pig/conf/pig.properties, and classes in user code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#properties GoogleDataprocJob#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// HCFS URI of file containing Hive script to execute as the job. Conflicts with query_list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#query_file_uri GoogleDataprocJob#query_file_uri}
	QueryFileUri *string `field:"optional" json:"queryFileUri" yaml:"queryFileUri"`
	// The list of Hive queries or statements to execute as part of the job. Conflicts with query_file_uri.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#query_list GoogleDataprocJob#query_list}
	QueryList *[]*string `field:"optional" json:"queryList" yaml:"queryList"`
	// Mapping of query variable names to values (equivalent to the Pig command: name=[value]).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#script_variables GoogleDataprocJob#script_variables}
	ScriptVariables *map[string]*string `field:"optional" json:"scriptVariables" yaml:"scriptVariables"`
}

