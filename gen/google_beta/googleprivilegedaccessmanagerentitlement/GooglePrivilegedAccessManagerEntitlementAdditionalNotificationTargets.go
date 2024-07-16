package googleprivilegedaccessmanagerentitlement


type GooglePrivilegedAccessManagerEntitlementAdditionalNotificationTargets struct {
	// Optional. Additional email addresses to be notified when a principal(requester) is granted access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#admin_email_recipients GooglePrivilegedAccessManagerEntitlement#admin_email_recipients}
	AdminEmailRecipients *[]*string `field:"optional" json:"adminEmailRecipients" yaml:"adminEmailRecipients"`
	// Optional. Additional email address to be notified about an eligible entitlement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privileged_access_manager_entitlement#requester_email_recipients GooglePrivilegedAccessManagerEntitlement#requester_email_recipients}
	RequesterEmailRecipients *[]*string `field:"optional" json:"requesterEmailRecipients" yaml:"requesterEmailRecipients"`
}

