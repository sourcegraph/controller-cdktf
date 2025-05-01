package googlecolabschedule


type GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource struct {
	// The Cloud Storage uri pointing to the ipynb file. Format: gs://bucket/notebook_file.ipynb.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_colab_schedule#uri GoogleColabSchedule#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// The version of the Cloud Storage object to read.
	//
	// If unset, the current version of the object is read. See https://cloud.google.com/storage/docs/metadata#generation-number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_colab_schedule#generation GoogleColabSchedule#generation}
	Generation *string `field:"optional" json:"generation" yaml:"generation"`
}

