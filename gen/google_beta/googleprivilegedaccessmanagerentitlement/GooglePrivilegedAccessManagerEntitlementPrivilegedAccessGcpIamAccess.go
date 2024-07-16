package googleprivilegedaccessmanagerentitlement


type GooglePrivilegedAccessManagerEntitlementPrivilegedAccessGcpIamAccess struct {
	// Name of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#resource GooglePrivilegedAccessManagerEntitlement#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// The type of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#resource_type GooglePrivilegedAccessManagerEntitlement#resource_type}
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// role_bindings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#role_bindings GooglePrivilegedAccessManagerEntitlement#role_bindings}
	RoleBindings interface{} `field:"required" json:"roleBindings" yaml:"roleBindings"`
}

