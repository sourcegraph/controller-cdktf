package googlecloudfunctionsfunction


type GoogleCloudfunctionsFunctionEventTriggerFailurePolicy struct {
	// Whether the function should be retried on failure. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_cloudfunctions_function#retry GoogleCloudfunctionsFunction#retry}
	Retry interface{} `field:"required" json:"retry" yaml:"retry"`
}

