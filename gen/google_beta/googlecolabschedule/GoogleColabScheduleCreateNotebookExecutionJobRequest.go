package googlecolabschedule


type GoogleColabScheduleCreateNotebookExecutionJobRequest struct {
	// notebook_execution_job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_colab_schedule#notebook_execution_job GoogleColabSchedule#notebook_execution_job}
	NotebookExecutionJob *GoogleColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob `field:"required" json:"notebookExecutionJob" yaml:"notebookExecutionJob"`
}

