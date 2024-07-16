package googlegkehubscoperbacrolebinding


type GoogleGkeHubScopeRbacRoleBindingRole struct {
	// PredefinedRole is an ENUM representation of the default Kubernetes Roles Possible values: ["UNKNOWN", "ADMIN", "EDIT", "VIEW"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_scope_rbac_role_binding#predefined_role GoogleGkeHubScopeRbacRoleBinding#predefined_role}
	PredefinedRole *string `field:"optional" json:"predefinedRole" yaml:"predefinedRole"`
}

