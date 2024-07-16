package googleprivilegedaccessmanagerentitlement


type GooglePrivilegedAccessManagerEntitlementRequesterJustificationConfig struct {
	// not_mandatory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#not_mandatory GooglePrivilegedAccessManagerEntitlement#not_mandatory}
	NotMandatory *GooglePrivilegedAccessManagerEntitlementRequesterJustificationConfigNotMandatory `field:"optional" json:"notMandatory" yaml:"notMandatory"`
	// unstructured block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#unstructured GooglePrivilegedAccessManagerEntitlement#unstructured}
	Unstructured *GooglePrivilegedAccessManagerEntitlementRequesterJustificationConfigUnstructured `field:"optional" json:"unstructured" yaml:"unstructured"`
}

