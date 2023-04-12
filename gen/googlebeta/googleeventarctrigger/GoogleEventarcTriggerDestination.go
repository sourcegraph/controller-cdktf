package googleeventarctrigger


type GoogleEventarcTriggerDestination struct {
	// [WARNING] Configuring a Cloud Function in Trigger is not supported as of today.
	//
	// The Cloud Function resource name. Format: projects/{project}/locations/{location}/functions/{function}
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_eventarc_trigger#cloud_function GoogleEventarcTrigger#cloud_function}
	CloudFunction *string `field:"optional" json:"cloudFunction" yaml:"cloudFunction"`
	// cloud_run_service block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_eventarc_trigger#cloud_run_service GoogleEventarcTrigger#cloud_run_service}
	CloudRunService *GoogleEventarcTriggerDestinationCloudRunService `field:"optional" json:"cloudRunService" yaml:"cloudRunService"`
	// gke block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_eventarc_trigger#gke GoogleEventarcTrigger#gke}
	Gke *GoogleEventarcTriggerDestinationGke `field:"optional" json:"gke" yaml:"gke"`
	// The resource name of the Workflow whose Executions are triggered by the events.
	//
	// The Workflow resource should be deployed in the same project as the trigger. Format: `projects/{project}/locations/{location}/workflows/{workflow}`
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_eventarc_trigger#workflow GoogleEventarcTrigger#workflow}
	Workflow *string `field:"optional" json:"workflow" yaml:"workflow"`
}

