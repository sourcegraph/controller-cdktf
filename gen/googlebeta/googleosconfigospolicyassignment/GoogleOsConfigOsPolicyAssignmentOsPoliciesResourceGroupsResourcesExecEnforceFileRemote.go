package googleosconfigospolicyassignment


type GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesExecEnforceFileRemote struct {
	// Required.
	//
	// URI from which to fetch the object. It should contain both the protocol and path following the format `{protocol}://{location}`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_os_config_os_policy_assignment#uri GoogleOsConfigOsPolicyAssignment#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// SHA256 checksum of the remote file.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_os_config_os_policy_assignment#sha256_checksum GoogleOsConfigOsPolicyAssignment#sha256_checksum}
	Sha256Checksum *string `field:"optional" json:"sha256Checksum" yaml:"sha256Checksum"`
}

