package googleiapsettings


type GoogleIapSettingsAccessSettingsReauthSettings struct {
	// Reauth session lifetime, how long before a user has to reauthenticate again.
	//
	// A duration in seconds with up to nine fractional digits, ending with 's'.
	// Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_iap_settings#max_age GoogleIapSettings#max_age}
	MaxAge *string `field:"required" json:"maxAge" yaml:"maxAge"`
	// Reauth method requested. The possible values are:.
	//
	// * 'LOGIN': Prompts the user to log in again.
	// * 'SECURE_KEY': User must use their secure key 2nd factor device.
	// * 'ENROLLED_SECOND_FACTORS': User can use any enabled 2nd factor. Possible values: ["LOGIN", "SECURE_KEY", "ENROLLED_SECOND_FACTORS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_iap_settings#method GoogleIapSettings#method}
	Method *string `field:"required" json:"method" yaml:"method"`
	// How IAP determines the effective policy in cases of hierarchical policies.
	//
	// Policies are merged from higher in the hierarchy to lower in the hierarchy.
	// The possible values are:
	//
	// * 'MINIMUM': This policy acts as a minimum to other policies, lower in the hierarchy.
	// 		   Effective policy may only be the same or stricter.
	// * 'DEFAULT': This policy acts as a default if no other reauth policy is set. Possible values: ["MINIMUM", "DEFAULT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_iap_settings#policy_type GoogleIapSettings#policy_type}
	PolicyType *string `field:"required" json:"policyType" yaml:"policyType"`
}

