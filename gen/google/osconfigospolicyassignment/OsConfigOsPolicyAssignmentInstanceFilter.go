package osconfigospolicyassignment

type OsConfigOsPolicyAssignmentInstanceFilter struct {
	// Target all VMs in the project. If true, no other criteria is permitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/os_config_os_policy_assignment#all OsConfigOsPolicyAssignment#all}
	All any `field:"optional" json:"all" yaml:"all"`
	// exclusion_labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/os_config_os_policy_assignment#exclusion_labels OsConfigOsPolicyAssignment#exclusion_labels}
	ExclusionLabels any `field:"optional" json:"exclusionLabels" yaml:"exclusionLabels"`
	// inclusion_labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/os_config_os_policy_assignment#inclusion_labels OsConfigOsPolicyAssignment#inclusion_labels}
	InclusionLabels any `field:"optional" json:"inclusionLabels" yaml:"inclusionLabels"`
	// inventories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/os_config_os_policy_assignment#inventories OsConfigOsPolicyAssignment#inventories}
	Inventories any `field:"optional" json:"inventories" yaml:"inventories"`
}
