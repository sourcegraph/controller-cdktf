package googleprivilegedaccessmanagerentitlement


type GooglePrivilegedAccessManagerEntitlementPrivilegedAccessGcpIamAccessRoleBindings struct {
	// IAM role to be granted. https://cloud.google.com/iam/docs/roles-overview.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#role GooglePrivilegedAccessManagerEntitlement#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// The expression field of the IAM condition to be associated with the role.
	//
	// If specified, a user with an active grant for this entitlement would be able to access the resource only if this condition evaluates to true for their request.
	// https://cloud.google.com/iam/docs/conditions-overview#attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#condition_expression GooglePrivilegedAccessManagerEntitlement#condition_expression}
	ConditionExpression *string `field:"optional" json:"conditionExpression" yaml:"conditionExpression"`
}

