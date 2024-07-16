package googleprivilegedaccessmanagerentitlement


type GooglePrivilegedAccessManagerEntitlementPrivilegedAccess struct {
	// gcp_iam_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#gcp_iam_access GooglePrivilegedAccessManagerEntitlement#gcp_iam_access}
	GcpIamAccess *GooglePrivilegedAccessManagerEntitlementPrivilegedAccessGcpIamAccess `field:"required" json:"gcpIamAccess" yaml:"gcpIamAccess"`
}

