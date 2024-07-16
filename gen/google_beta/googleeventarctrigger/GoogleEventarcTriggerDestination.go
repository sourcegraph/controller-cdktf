package googleeventarctrigger


type GoogleEventarcTriggerDestination struct {
	// cloud_run_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_eventarc_trigger#cloud_run_service GoogleEventarcTrigger#cloud_run_service}
	CloudRunService *GoogleEventarcTriggerDestinationCloudRunService `field:"optional" json:"cloudRunService" yaml:"cloudRunService"`
	// gke block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_eventarc_trigger#gke GoogleEventarcTrigger#gke}
	Gke *GoogleEventarcTriggerDestinationGke `field:"optional" json:"gke" yaml:"gke"`
	// http_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_eventarc_trigger#http_endpoint GoogleEventarcTrigger#http_endpoint}
	HttpEndpoint *GoogleEventarcTriggerDestinationHttpEndpoint `field:"optional" json:"httpEndpoint" yaml:"httpEndpoint"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_eventarc_trigger#network_config GoogleEventarcTrigger#network_config}
	NetworkConfig *GoogleEventarcTriggerDestinationNetworkConfig `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// The resource name of the Workflow whose Executions are triggered by the events.
	//
	// The Workflow resource should be deployed in the same project as the trigger. Format: `projects/{project}/locations/{location}/workflows/{workflow}`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_eventarc_trigger#workflow GoogleEventarcTrigger#workflow}
	Workflow *string `field:"optional" json:"workflow" yaml:"workflow"`
}

