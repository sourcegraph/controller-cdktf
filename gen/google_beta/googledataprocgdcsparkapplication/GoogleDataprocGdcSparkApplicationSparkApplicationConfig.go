package googledataprocgdcsparkapplication


type GoogleDataprocGdcSparkApplicationSparkApplicationConfig struct {
	// HCFS URIs of archives to be extracted into the working directory of each executor.
	//
	// Supported file types: '.jar', '.tar', '.tar.gz', '.tgz', and '.zip'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_gdc_spark_application#archive_uris GoogleDataprocGdcSparkApplication#archive_uris}
	ArchiveUris *[]*string `field:"optional" json:"archiveUris" yaml:"archiveUris"`
	// The arguments to pass to the driver.
	//
	// Do not include arguments that can be set as application properties, such as '--conf', since a collision can occur that causes an incorrect application submission.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_gdc_spark_application#args GoogleDataprocGdcSparkApplication#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// HCFS URIs of files to be placed in the working directory of each executor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_gdc_spark_application#file_uris GoogleDataprocGdcSparkApplication#file_uris}
	FileUris *[]*string `field:"optional" json:"fileUris" yaml:"fileUris"`
	// HCFS URIs of jar files to add to the classpath of the Spark driver and tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_gdc_spark_application#jar_file_uris GoogleDataprocGdcSparkApplication#jar_file_uris}
	JarFileUris *[]*string `field:"optional" json:"jarFileUris" yaml:"jarFileUris"`
	// The name of the driver main class.
	//
	// The jar file that contains the class must be in the classpath or specified in 'jar_file_uris'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_gdc_spark_application#main_class GoogleDataprocGdcSparkApplication#main_class}
	MainClass *string `field:"optional" json:"mainClass" yaml:"mainClass"`
	// The HCFS URI of the jar file that contains the main class.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_gdc_spark_application#main_jar_file_uri GoogleDataprocGdcSparkApplication#main_jar_file_uri}
	MainJarFileUri *string `field:"optional" json:"mainJarFileUri" yaml:"mainJarFileUri"`
}

