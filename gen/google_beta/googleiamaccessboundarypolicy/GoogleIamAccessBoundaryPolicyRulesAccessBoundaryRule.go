package googleiamaccessboundarypolicy


type GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRule struct {
	// availability_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iam_access_boundary_policy#availability_condition GoogleIamAccessBoundaryPolicy#availability_condition}
	AvailabilityCondition *GoogleIamAccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityCondition `field:"optional" json:"availabilityCondition" yaml:"availabilityCondition"`
	// A list of permissions that may be allowed for use on the specified resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iam_access_boundary_policy#available_permissions GoogleIamAccessBoundaryPolicy#available_permissions}
	AvailablePermissions *[]*string `field:"optional" json:"availablePermissions" yaml:"availablePermissions"`
	// The full resource name of a Google Cloud resource entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_iam_access_boundary_policy#available_resource GoogleIamAccessBoundaryPolicy#available_resource}
	AvailableResource *string `field:"optional" json:"availableResource" yaml:"availableResource"`
}

