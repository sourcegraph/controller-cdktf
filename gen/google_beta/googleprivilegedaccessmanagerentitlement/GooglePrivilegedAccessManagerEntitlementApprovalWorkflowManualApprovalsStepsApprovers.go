package googleprivilegedaccessmanagerentitlement


type GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsApprovers struct {
	// Users who are being allowed for the operation.
	//
	// Each entry should be a valid v1 IAM Principal Identifier. Format for these is documented at: https://cloud.google.com/iam/docs/principal-identifiers#v1
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#principals GooglePrivilegedAccessManagerEntitlement#principals}
	Principals *[]*string `field:"required" json:"principals" yaml:"principals"`
}

