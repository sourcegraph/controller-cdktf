package googleprivilegedaccessmanagerentitlement


type GooglePrivilegedAccessManagerEntitlementApprovalWorkflow struct {
	// manual_approvals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#manual_approvals GooglePrivilegedAccessManagerEntitlement#manual_approvals}
	ManualApprovals *GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovals `field:"required" json:"manualApprovals" yaml:"manualApprovals"`
}

