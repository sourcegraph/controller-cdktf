package accesspolicy


type AccessPolicyApprovalGroup struct {
	// Number of approvals needed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#approvals_needed AccessPolicy#approvals_needed}
	ApprovalsNeeded *float64 `field:"required" json:"approvalsNeeded" yaml:"approvalsNeeded"`
	// List of emails to request approval from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#email_addresses AccessPolicy#email_addresses}
	EmailAddresses *[]*string `field:"optional" json:"emailAddresses" yaml:"emailAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy#email_list_uuid AccessPolicy#email_list_uuid}.
	EmailListUuid *string `field:"optional" json:"emailListUuid" yaml:"emailListUuid"`
}

