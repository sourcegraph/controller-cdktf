package googledataplextask


type GoogleDataplexTaskNotebook struct {
	// Path to input notebook.
	//
	// This can be the Cloud Storage URI of the notebook file or the path to a Notebook Content. The execution args are accessible as environment variables (TASK_key=value).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_task#notebook GoogleDataplexTask#notebook}
	Notebook *string `field:"required" json:"notebook" yaml:"notebook"`
	// Cloud Storage URIs of archives to be extracted into the working directory of each executor.
	//
	// Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_task#archive_uris GoogleDataplexTask#archive_uris}
	ArchiveUris *[]*string `field:"optional" json:"archiveUris" yaml:"archiveUris"`
	// Cloud Storage URIs of files to be placed in the working directory of each executor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_task#file_uris GoogleDataplexTask#file_uris}
	FileUris *[]*string `field:"optional" json:"fileUris" yaml:"fileUris"`
	// infrastructure_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataplex_task#infrastructure_spec GoogleDataplexTask#infrastructure_spec}
	InfrastructureSpec *GoogleDataplexTaskNotebookInfrastructureSpec `field:"optional" json:"infrastructureSpec" yaml:"infrastructureSpec"`
}

