package googleeventarcpipeline


type GoogleEventarcPipelineRetryPolicy struct {
	// The maximum number of delivery attempts for any message.
	//
	// The value must
	// be between 1 and 100.
	// The default value for this field is 5.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_eventarc_pipeline#max_attempts GoogleEventarcPipeline#max_attempts}
	MaxAttempts *float64 `field:"optional" json:"maxAttempts" yaml:"maxAttempts"`
	// The maximum amount of seconds to wait between retry attempts.
	//
	// The value
	// must be between 1 and 600.
	// The default value for this field is 60.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_eventarc_pipeline#max_retry_delay GoogleEventarcPipeline#max_retry_delay}
	MaxRetryDelay *string `field:"optional" json:"maxRetryDelay" yaml:"maxRetryDelay"`
	// The minimum amount of seconds to wait between retry attempts.
	//
	// The value
	// must be between 1 and 600.
	// The default value for this field is 5.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_eventarc_pipeline#min_retry_delay GoogleEventarcPipeline#min_retry_delay}
	MinRetryDelay *string `field:"optional" json:"minRetryDelay" yaml:"minRetryDelay"`
}

