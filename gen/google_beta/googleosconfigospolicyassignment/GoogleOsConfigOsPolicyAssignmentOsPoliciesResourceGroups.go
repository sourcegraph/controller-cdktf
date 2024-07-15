package googleosconfigospolicyassignment


type GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroups struct {
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#resources GoogleOsConfigOsPolicyAssignment#resources}
	Resources interface{} `field:"required" json:"resources" yaml:"resources"`
	// inventory_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#inventory_filters GoogleOsConfigOsPolicyAssignment#inventory_filters}
	InventoryFilters interface{} `field:"optional" json:"inventoryFilters" yaml:"inventoryFilters"`
}

