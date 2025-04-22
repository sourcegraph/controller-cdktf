package googledataprocgdcsparkapplication


type GoogleDataprocGdcSparkApplicationSparkSqlApplicationConfig struct {
	// HCFS URIs of jar files to be added to the Spark CLASSPATH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_gdc_spark_application#jar_file_uris GoogleDataprocGdcSparkApplication#jar_file_uris}
	JarFileUris *[]*string `field:"optional" json:"jarFileUris" yaml:"jarFileUris"`
	// The HCFS URI of the script that contains SQL queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_gdc_spark_application#query_file_uri GoogleDataprocGdcSparkApplication#query_file_uri}
	QueryFileUri *string `field:"optional" json:"queryFileUri" yaml:"queryFileUri"`
	// query_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_gdc_spark_application#query_list GoogleDataprocGdcSparkApplication#query_list}
	QueryList *GoogleDataprocGdcSparkApplicationSparkSqlApplicationConfigQueryListStruct `field:"optional" json:"queryList" yaml:"queryList"`
	// Mapping of query variable names to values (equivalent to the Spark SQL command: SET 'name="value";').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_gdc_spark_application#script_variables GoogleDataprocGdcSparkApplication#script_variables}
	ScriptVariables *map[string]*string `field:"optional" json:"scriptVariables" yaml:"scriptVariables"`
}

