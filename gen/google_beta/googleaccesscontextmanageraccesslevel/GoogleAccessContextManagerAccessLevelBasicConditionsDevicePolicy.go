package googleaccesscontextmanageraccesslevel


type GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicy struct {
	// A list of allowed device management levels. An empty list allows all management levels. Possible values: ["MANAGEMENT_UNSPECIFIED", "NONE", "BASIC", "COMPLETE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_access_level#allowed_device_management_levels GoogleAccessContextManagerAccessLevel#allowed_device_management_levels}
	AllowedDeviceManagementLevels *[]*string `field:"optional" json:"allowedDeviceManagementLevels" yaml:"allowedDeviceManagementLevels"`
	// A list of allowed encryptions statuses. An empty list allows all statuses. Possible values: ["ENCRYPTION_UNSPECIFIED", "ENCRYPTION_UNSUPPORTED", "UNENCRYPTED", "ENCRYPTED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_access_level#allowed_encryption_statuses GoogleAccessContextManagerAccessLevel#allowed_encryption_statuses}
	AllowedEncryptionStatuses *[]*string `field:"optional" json:"allowedEncryptionStatuses" yaml:"allowedEncryptionStatuses"`
	// os_constraints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_access_level#os_constraints GoogleAccessContextManagerAccessLevel#os_constraints}
	OsConstraints interface{} `field:"optional" json:"osConstraints" yaml:"osConstraints"`
	// Whether the device needs to be approved by the customer admin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_access_level#require_admin_approval GoogleAccessContextManagerAccessLevel#require_admin_approval}
	RequireAdminApproval interface{} `field:"optional" json:"requireAdminApproval" yaml:"requireAdminApproval"`
	// Whether the device needs to be corp owned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_access_level#require_corp_owned GoogleAccessContextManagerAccessLevel#require_corp_owned}
	RequireCorpOwned interface{} `field:"optional" json:"requireCorpOwned" yaml:"requireCorpOwned"`
	// Whether or not screenlock is required for the DevicePolicy to be true. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_access_level#require_screen_lock GoogleAccessContextManagerAccessLevel#require_screen_lock}
	RequireScreenLock interface{} `field:"optional" json:"requireScreenLock" yaml:"requireScreenLock"`
}

