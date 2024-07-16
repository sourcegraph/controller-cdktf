package googleosconfigospolicyassignment


type GoogleOsConfigOsPolicyAssignmentInstanceFilter struct {
	// Target all VMs in the project. If true, no other criteria is permitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#all GoogleOsConfigOsPolicyAssignment#all}
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// exclusion_labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#exclusion_labels GoogleOsConfigOsPolicyAssignment#exclusion_labels}
	ExclusionLabels interface{} `field:"optional" json:"exclusionLabels" yaml:"exclusionLabels"`
	// inclusion_labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#inclusion_labels GoogleOsConfigOsPolicyAssignment#inclusion_labels}
	InclusionLabels interface{} `field:"optional" json:"inclusionLabels" yaml:"inclusionLabels"`
	// inventories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#inventories GoogleOsConfigOsPolicyAssignment#inventories}
	Inventories interface{} `field:"optional" json:"inventories" yaml:"inventories"`
}

