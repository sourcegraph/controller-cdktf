package googleprivilegedaccessmanagerentitlement


type GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsSteps struct {
	// approvers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#approvers GooglePrivilegedAccessManagerEntitlement#approvers}
	Approvers *GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsApprovers `field:"required" json:"approvers" yaml:"approvers"`
	// How many users from the above list need to approve.
	//
	// If there are not enough distinct users in the list above then the workflow
	// will indefinitely block. Should always be greater than 0. Currently 1 is the only
	// supported value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#approvals_needed GooglePrivilegedAccessManagerEntitlement#approvals_needed}
	ApprovalsNeeded *float64 `field:"optional" json:"approvalsNeeded" yaml:"approvalsNeeded"`
	// Optional. Additional email addresses to be notified when a grant is pending approval.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#approver_email_recipients GooglePrivilegedAccessManagerEntitlement#approver_email_recipients}
	ApproverEmailRecipients *[]*string `field:"optional" json:"approverEmailRecipients" yaml:"approverEmailRecipients"`
}

