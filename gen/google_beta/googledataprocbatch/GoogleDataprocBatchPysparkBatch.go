package googledataprocbatch


type GoogleDataprocBatchPysparkBatch struct {
	// HCFS URIs of archives to be extracted into the working directory of each executor.
	//
	// Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_batch#archive_uris GoogleDataprocBatch#archive_uris}
	ArchiveUris *[]*string `field:"optional" json:"archiveUris" yaml:"archiveUris"`
	// The arguments to pass to the driver.
	//
	// Do not include arguments that can be set as batch
	// properties, such as --conf, since a collision can occur that causes an incorrect batch submission.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_batch#args GoogleDataprocBatch#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// HCFS URIs of files to be placed in the working directory of each executor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_batch#file_uris GoogleDataprocBatch#file_uris}
	FileUris *[]*string `field:"optional" json:"fileUris" yaml:"fileUris"`
	// HCFS URIs of jar files to add to the classpath of the Spark driver and tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_batch#jar_file_uris GoogleDataprocBatch#jar_file_uris}
	JarFileUris *[]*string `field:"optional" json:"jarFileUris" yaml:"jarFileUris"`
	// The HCFS URI of the main Python file to use as the Spark driver. Must be a .py file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_batch#main_python_file_uri GoogleDataprocBatch#main_python_file_uri}
	MainPythonFileUri *string `field:"optional" json:"mainPythonFileUri" yaml:"mainPythonFileUri"`
	// HCFS file URIs of Python files to pass to the PySpark framework. Supported file types: .py, .egg, and .zip.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_batch#python_file_uris GoogleDataprocBatch#python_file_uris}
	PythonFileUris *[]*string `field:"optional" json:"pythonFileUris" yaml:"pythonFileUris"`
}

