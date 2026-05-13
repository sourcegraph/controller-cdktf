package googleprivilegedaccessmanagersettings


type GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehavior struct {
	// admin_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_privileged_access_manager_settings#admin_notifications GooglePrivilegedAccessManagerSettings#admin_notifications}
	AdminNotifications *GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorAdminNotifications `field:"optional" json:"adminNotifications" yaml:"adminNotifications"`
	// approver_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_privileged_access_manager_settings#approver_notifications GooglePrivilegedAccessManagerSettings#approver_notifications}
	ApproverNotifications *GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorApproverNotifications `field:"optional" json:"approverNotifications" yaml:"approverNotifications"`
	// requester_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_privileged_access_manager_settings#requester_notifications GooglePrivilegedAccessManagerSettings#requester_notifications}
	RequesterNotifications *GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotifications `field:"optional" json:"requesterNotifications" yaml:"requesterNotifications"`
}

