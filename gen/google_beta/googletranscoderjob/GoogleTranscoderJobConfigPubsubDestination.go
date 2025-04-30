package googletranscoderjob


type GoogleTranscoderJobConfigPubsubDestination struct {
	// The name of the Pub/Sub topic to publish job completion notification to. For example: projects/{project}/topics/{topic}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#topic GoogleTranscoderJob#topic}
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
}

