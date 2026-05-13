package googlegkehubscoperbacrolebinding


type GoogleGkeHubScopeRbacRoleBindingRole struct {
	// CustomRole is the custom Kubernetes ClusterRole to be used.
	//
	// The custom role format must be allowlisted in the rbacrolebindingactuation feature and RFC 1123 compliant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gke_hub_scope_rbac_role_binding#custom_role GoogleGkeHubScopeRbacRoleBinding#custom_role}
	CustomRole *string `field:"optional" json:"customRole" yaml:"customRole"`
	// PredefinedRole is an ENUM representation of the default Kubernetes Roles Possible values: ["UNKNOWN", "ADMIN", "EDIT", "VIEW"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gke_hub_scope_rbac_role_binding#predefined_role GoogleGkeHubScopeRbacRoleBinding#predefined_role}
	PredefinedRole *string `field:"optional" json:"predefinedRole" yaml:"predefinedRole"`
}

