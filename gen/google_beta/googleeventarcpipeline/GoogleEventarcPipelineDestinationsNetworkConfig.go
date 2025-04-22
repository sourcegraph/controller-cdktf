package googleeventarcpipeline


type GoogleEventarcPipelineDestinationsNetworkConfig struct {
	// Name of the NetworkAttachment that allows access to the consumer VPC. Format: 'projects/{PROJECT_ID}/regions/{REGION}/networkAttachments/{NETWORK_ATTACHMENT_NAME}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_eventarc_pipeline#network_attachment GoogleEventarcPipeline#network_attachment}
	NetworkAttachment *string `field:"required" json:"networkAttachment" yaml:"networkAttachment"`
}

