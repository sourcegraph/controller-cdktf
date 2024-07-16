package googleosconfigguestpolicies


type GoogleOsConfigGuestPoliciesAssignmentGroupLabels struct {
	// Google Compute Engine instance labels that must be present for an instance to be included in this assignment group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#labels GoogleOsConfigGuestPolicies#labels}
	Labels *map[string]*string `field:"required" json:"labels" yaml:"labels"`
}

