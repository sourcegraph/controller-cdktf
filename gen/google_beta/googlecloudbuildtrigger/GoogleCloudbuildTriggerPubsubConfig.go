package googlecloudbuildtrigger


type GoogleCloudbuildTriggerPubsubConfig struct {
	// The name of the topic from which this subscription is receiving messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#topic GoogleCloudbuildTrigger#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// Service account that will make the push request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#service_account_email GoogleCloudbuildTrigger#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
}

