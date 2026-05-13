package googleeventarctrigger


type GoogleEventarcTriggerRetryPolicy struct {
	// The maximum number of delivery attempts for any message. The only valid value is 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_eventarc_trigger#max_attempts GoogleEventarcTrigger#max_attempts}
	MaxAttempts *float64 `field:"optional" json:"maxAttempts" yaml:"maxAttempts"`
}

