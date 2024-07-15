package googlegkehubmembershiprbacrolebinding


type GoogleGkeHubMembershipRbacRoleBindingRole struct {
	// PredefinedRole is an ENUM representation of the default Kubernetes Roles Possible values: ["UNKNOWN", "ADMIN", "EDIT", "VIEW", "ANTHOS_SUPPORT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_membership_rbac_role_binding#predefined_role GoogleGkeHubMembershipRbacRoleBinding#predefined_role}
	PredefinedRole *string `field:"required" json:"predefinedRole" yaml:"predefinedRole"`
}

