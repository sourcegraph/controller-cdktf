package googleprivilegedaccessmanagersettings


type GooglePrivilegedAccessManagerSettingsEmailNotificationSettingsCustomNotificationBehaviorRequesterNotifications struct {
	// Notification mode for entitlement assigned. Possible values: ["NOTIFICATION_MODE_UNSPECIFIED", "ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_privileged_access_manager_settings#entitlement_assigned GooglePrivilegedAccessManagerSettings#entitlement_assigned}
	EntitlementAssigned *string `field:"optional" json:"entitlementAssigned" yaml:"entitlementAssigned"`
	// Notification mode for grant activated. Possible values: ["NOTIFICATION_MODE_UNSPECIFIED", "ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_privileged_access_manager_settings#grant_activated GooglePrivilegedAccessManagerSettings#grant_activated}
	GrantActivated *string `field:"optional" json:"grantActivated" yaml:"grantActivated"`
	// Notification mode for grant activation failed. Possible values: ["NOTIFICATION_MODE_UNSPECIFIED", "ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_privileged_access_manager_settings#grant_activation_failed GooglePrivilegedAccessManagerSettings#grant_activation_failed}
	GrantActivationFailed *string `field:"optional" json:"grantActivationFailed" yaml:"grantActivationFailed"`
	// Notification mode for grant denied. Possible values: ["NOTIFICATION_MODE_UNSPECIFIED", "ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_privileged_access_manager_settings#grant_denied GooglePrivilegedAccessManagerSettings#grant_denied}
	GrantDenied *string `field:"optional" json:"grantDenied" yaml:"grantDenied"`
	// Notification mode for grant ended. Possible values: ["NOTIFICATION_MODE_UNSPECIFIED", "ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_privileged_access_manager_settings#grant_ended GooglePrivilegedAccessManagerSettings#grant_ended}
	GrantEnded *string `field:"optional" json:"grantEnded" yaml:"grantEnded"`
	// Notification mode for grant expired. Possible values: ["NOTIFICATION_MODE_UNSPECIFIED", "ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_privileged_access_manager_settings#grant_expired GooglePrivilegedAccessManagerSettings#grant_expired}
	GrantExpired *string `field:"optional" json:"grantExpired" yaml:"grantExpired"`
	// Notification mode for grant externally modified. Possible values: ["NOTIFICATION_MODE_UNSPECIFIED", "ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_privileged_access_manager_settings#grant_externally_modified GooglePrivilegedAccessManagerSettings#grant_externally_modified}
	GrantExternallyModified *string `field:"optional" json:"grantExternallyModified" yaml:"grantExternallyModified"`
	// Notification mode for grant revoked. Possible values: ["NOTIFICATION_MODE_UNSPECIFIED", "ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_privileged_access_manager_settings#grant_revoked GooglePrivilegedAccessManagerSettings#grant_revoked}
	GrantRevoked *string `field:"optional" json:"grantRevoked" yaml:"grantRevoked"`
}

