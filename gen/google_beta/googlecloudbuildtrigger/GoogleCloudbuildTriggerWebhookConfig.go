package googlecloudbuildtrigger


type GoogleCloudbuildTriggerWebhookConfig struct {
	// Resource name for the secret required as a URL parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#secret GoogleCloudbuildTrigger#secret}
	Secret *string `field:"required" json:"secret" yaml:"secret"`
}

