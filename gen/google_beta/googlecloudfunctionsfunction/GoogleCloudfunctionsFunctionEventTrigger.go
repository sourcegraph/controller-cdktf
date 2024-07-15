package googlecloudfunctionsfunction


type GoogleCloudfunctionsFunctionEventTrigger struct {
	// The type of event to observe.
	//
	// For example: "google.storage.object.finalize". See the documentation on calling Cloud Functions for a full reference of accepted triggers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions_function#event_type GoogleCloudfunctionsFunction#event_type}
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// The name or partial URI of the resource from which to observe events. For example, "myBucket" or "projects/my-project/topics/my-topic".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions_function#resource GoogleCloudfunctionsFunction#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// failure_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions_function#failure_policy GoogleCloudfunctionsFunction#failure_policy}
	FailurePolicy *GoogleCloudfunctionsFunctionEventTriggerFailurePolicy `field:"optional" json:"failurePolicy" yaml:"failurePolicy"`
}

