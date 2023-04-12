package googlecloudbuildtrigger


type GoogleCloudbuildTriggerWebhookConfig struct {
	// Resource name for the secret required as a URL parameter.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudbuild_trigger#secret GoogleCloudbuildTrigger#secret}
	Secret *string `field:"required" json:"secret" yaml:"secret"`
}

