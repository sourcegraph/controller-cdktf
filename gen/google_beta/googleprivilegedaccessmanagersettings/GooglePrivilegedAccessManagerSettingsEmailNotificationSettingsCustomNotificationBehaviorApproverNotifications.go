package googleprivilegedaccessmanagersettings


type GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorApproverNotifications struct {
	// Notification mode for pending approval. Possible values: ["NOTIFICATION_MODE_UNSPECIFIED", "ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_privileged_access_manager_settings#pending_approval GooglePrivilegedAccessManagerSettings#pending_approval}
	PendingApproval *string `field:"optional" json:"pendingApproval" yaml:"pendingApproval"`
}

