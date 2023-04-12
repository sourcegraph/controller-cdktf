package googlecomputeinstancetemplate


type GoogleComputeInstanceTemplateServiceAccount struct {
	// A list of service scopes.
	//
	// Both OAuth2 URLs and gcloud short names are supported. To allow full access to all Cloud APIs, use the cloud-platform scope.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#scopes GoogleComputeInstanceTemplate#scopes}
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// The service account e-mail address. If not given, the default Google Compute Engine service account is used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_template#email GoogleComputeInstanceTemplate#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
}

