package googletpuv2vm


type GoogleTpuV2VmServiceAccount struct {
	// Email address of the service account. If empty, default Compute service account will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#email GoogleTpuV2Vm#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
	// The list of scopes to be made available for this service account.
	//
	// If empty, access to all
	// Cloud APIs will be allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#scope GoogleTpuV2Vm#scope}
	Scope *[]*string `field:"optional" json:"scope" yaml:"scope"`
}

