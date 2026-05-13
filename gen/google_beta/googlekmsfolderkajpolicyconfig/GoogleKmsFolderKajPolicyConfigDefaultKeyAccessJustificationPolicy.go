package googlekmsfolderkajpolicyconfig


type GoogleKmsFolderKajPolicyConfigDefaultKeyAccessJustificationPolicy struct {
	// A KeyAccessJustificationsPolicy specifies zero or more allowed AccessReason values for encrypt, decrypt, and sign operations on a CryptoKey.
	//
	// Possible values: ["CUSTOMER_INITIATED_SUPPORT", "GOOGLE_INITIATED_SERVICE", "THIRD_PARTY_DATA_REQUEST", "GOOGLE_INITIATED_REVIEW", "CUSTOMER_INITIATED_ACCESS", "GOOGLE_INITIATED_SYSTEM_OPERATION", "REASON_NOT_EXPECTED", "MODIFIED_CUSTOMER_INITIATED_ACCESS", "MODIFIED_GOOGLE_INITIATED_SYSTEM_OPERATION", "GOOGLE_RESPONSE_TO_PRODUCTION_ALERT", "CUSTOMER_AUTHORIZED_WORKFLOW_SERVICING"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_kms_folder_kaj_policy_config#allowed_access_reasons GoogleKmsFolderKajPolicyConfig#allowed_access_reasons}
	AllowedAccessReasons *[]*string `field:"optional" json:"allowedAccessReasons" yaml:"allowedAccessReasons"`
}

