package googleeventarctrigger


type GoogleEventarcTriggerDestinationNetworkConfig struct {
	// Required. Name of the NetworkAttachment that allows access to the destination VPC. Format: `projects/{PROJECT_ID}/regions/{REGION}/networkAttachments/{NETWORK_ATTACHMENT_NAME}`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_eventarc_trigger#network_attachment GoogleEventarcTrigger#network_attachment}
	NetworkAttachment *string `field:"required" json:"networkAttachment" yaml:"networkAttachment"`
}

