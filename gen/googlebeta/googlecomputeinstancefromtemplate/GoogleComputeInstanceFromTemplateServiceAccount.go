package googlecomputeinstancefromtemplate


type GoogleComputeInstanceFromTemplateServiceAccount struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_from_template#email GoogleComputeInstanceFromTemplate#email}.
	Email *string `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_from_template#scopes GoogleComputeInstanceFromTemplate#scopes}.
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

