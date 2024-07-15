package googledataprocjob


type GoogleDataprocJobSparkConfig struct {
	// HCFS URIs of archives to be extracted in the working directory of .jar, .tar, .tar.gz, .tgz, and .zip.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#archive_uris GoogleDataprocJob#archive_uris}
	ArchiveUris *[]*string `field:"optional" json:"archiveUris" yaml:"archiveUris"`
	// The arguments to pass to the driver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#args GoogleDataprocJob#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// HCFS URIs of files to be copied to the working directory of Spark drivers and distributed tasks.
	//
	// Useful for naively parallel tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#file_uris GoogleDataprocJob#file_uris}
	FileUris *[]*string `field:"optional" json:"fileUris" yaml:"fileUris"`
	// HCFS URIs of jar files to add to the CLASSPATHs of the Spark driver and tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#jar_file_uris GoogleDataprocJob#jar_file_uris}
	JarFileUris *[]*string `field:"optional" json:"jarFileUris" yaml:"jarFileUris"`
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#logging_config GoogleDataprocJob#logging_config}
	LoggingConfig *GoogleDataprocJobSparkConfigLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// The class containing the main method of the driver.
	//
	// Must be in a provided jar or jar that is already on the classpath. Conflicts with main_jar_file_uri
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#main_class GoogleDataprocJob#main_class}
	MainClass *string `field:"optional" json:"mainClass" yaml:"mainClass"`
	// The HCFS URI of jar file containing the driver jar. Conflicts with main_class.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#main_jar_file_uri GoogleDataprocJob#main_jar_file_uri}
	MainJarFileUri *string `field:"optional" json:"mainJarFileUri" yaml:"mainJarFileUri"`
	// A mapping of property names to values, used to configure Spark.
	//
	// Properties that conflict with values set by the Cloud Dataproc API may be overwritten. Can include properties set in /etc/spark/conf/spark-defaults.conf and classes in user code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_job#properties GoogleDataprocJob#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

