package passwordpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PasswordPolicyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Policy Name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#name PasswordPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Authentication Provider: OKTA, ACTIVE_DIRECTORY or LDAP.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#auth_provider PasswordPolicy#auth_provider}
	AuthProvider *string `field:"optional" json:"authProvider" yaml:"authProvider"`
	// Enable or disable voice call recovery: ACTIVE or INACTIVE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#call_recovery PasswordPolicy#call_recovery}
	CallRecovery *string `field:"optional" json:"callRecovery" yaml:"callRecovery"`
	// Policy Description.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#description PasswordPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Enable or disable email password recovery: ACTIVE or INACTIVE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#email_recovery PasswordPolicy#email_recovery}
	EmailRecovery *string `field:"optional" json:"emailRecovery" yaml:"emailRecovery"`
	// List of Group IDs to Include.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#groups_included PasswordPolicy#groups_included}
	GroupsIncluded *[]*string `field:"optional" json:"groupsIncluded" yaml:"groupsIncluded"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#id PasswordPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Number of minutes before a locked account is unlocked: 0 = no limit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#password_auto_unlock_minutes PasswordPolicy#password_auto_unlock_minutes}
	PasswordAutoUnlockMinutes *float64 `field:"optional" json:"passwordAutoUnlockMinutes" yaml:"passwordAutoUnlockMinutes"`
	// Check Passwords Against Common Password Dictionary.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#password_dictionary_lookup PasswordPolicy#password_dictionary_lookup}
	PasswordDictionaryLookup interface{} `field:"optional" json:"passwordDictionaryLookup" yaml:"passwordDictionaryLookup"`
	// User firstName attribute must be excluded from the password.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#password_exclude_first_name PasswordPolicy#password_exclude_first_name}
	PasswordExcludeFirstName interface{} `field:"optional" json:"passwordExcludeFirstName" yaml:"passwordExcludeFirstName"`
	// User lastName attribute must be excluded from the password.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#password_exclude_last_name PasswordPolicy#password_exclude_last_name}
	PasswordExcludeLastName interface{} `field:"optional" json:"passwordExcludeLastName" yaml:"passwordExcludeLastName"`
	// If the user name must be excluded from the password.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#password_exclude_username PasswordPolicy#password_exclude_username}
	PasswordExcludeUsername interface{} `field:"optional" json:"passwordExcludeUsername" yaml:"passwordExcludeUsername"`
	// Length in days a user will be warned before password expiry: 0 = no warning.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#password_expire_warn_days PasswordPolicy#password_expire_warn_days}
	PasswordExpireWarnDays *float64 `field:"optional" json:"passwordExpireWarnDays" yaml:"passwordExpireWarnDays"`
	// Number of distinct passwords that must be created before they can be reused: 0 = none.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#password_history_count PasswordPolicy#password_history_count}
	PasswordHistoryCount *float64 `field:"optional" json:"passwordHistoryCount" yaml:"passwordHistoryCount"`
	// Notification channels to use to notify a user when their account has been locked.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#password_lockout_notification_channels PasswordPolicy#password_lockout_notification_channels}
	PasswordLockoutNotificationChannels *[]*string `field:"optional" json:"passwordLockoutNotificationChannels" yaml:"passwordLockoutNotificationChannels"`
	// Length in days a password is valid before expiry: 0 = no limit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#password_max_age_days PasswordPolicy#password_max_age_days}
	PasswordMaxAgeDays *float64 `field:"optional" json:"passwordMaxAgeDays" yaml:"passwordMaxAgeDays"`
	// Number of unsuccessful login attempts allowed before lockout: 0 = no limit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#password_max_lockout_attempts PasswordPolicy#password_max_lockout_attempts}
	PasswordMaxLockoutAttempts *float64 `field:"optional" json:"passwordMaxLockoutAttempts" yaml:"passwordMaxLockoutAttempts"`
	// Minimum time interval in minutes between password changes: 0 = no limit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#password_min_age_minutes PasswordPolicy#password_min_age_minutes}
	PasswordMinAgeMinutes *float64 `field:"optional" json:"passwordMinAgeMinutes" yaml:"passwordMinAgeMinutes"`
	// Minimum password length.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#password_min_length PasswordPolicy#password_min_length}
	PasswordMinLength *float64 `field:"optional" json:"passwordMinLength" yaml:"passwordMinLength"`
	// If a password must contain at least one lower case letter: 0 = no, 1 = yes.
	//
	// Default = 1
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#password_min_lowercase PasswordPolicy#password_min_lowercase}
	PasswordMinLowercase *float64 `field:"optional" json:"passwordMinLowercase" yaml:"passwordMinLowercase"`
	// If a password must contain at least one number: 0 = no, 1 = yes. Default = 1.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#password_min_number PasswordPolicy#password_min_number}
	PasswordMinNumber *float64 `field:"optional" json:"passwordMinNumber" yaml:"passwordMinNumber"`
	// If a password must contain at least one symbol (!@#$%^&*): 0 = no, 1 = yes. Default = 1.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#password_min_symbol PasswordPolicy#password_min_symbol}
	PasswordMinSymbol *float64 `field:"optional" json:"passwordMinSymbol" yaml:"passwordMinSymbol"`
	// If a password must contain at least one upper case letter: 0 = no, 1 = yes.
	//
	// Default = 1
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#password_min_uppercase PasswordPolicy#password_min_uppercase}
	PasswordMinUppercase *float64 `field:"optional" json:"passwordMinUppercase" yaml:"passwordMinUppercase"`
	// If a user should be informed when their account is locked.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#password_show_lockout_failures PasswordPolicy#password_show_lockout_failures}
	PasswordShowLockoutFailures interface{} `field:"optional" json:"passwordShowLockoutFailures" yaml:"passwordShowLockoutFailures"`
	// Policy Priority, this attribute can be set to a valid priority.
	//
	// To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last (lowest) if not there.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#priority PasswordPolicy#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Min length of the password recovery question answer.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#question_min_length PasswordPolicy#question_min_length}
	QuestionMinLength *float64 `field:"optional" json:"questionMinLength" yaml:"questionMinLength"`
	// Enable or disable security question password recovery: ACTIVE or INACTIVE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#question_recovery PasswordPolicy#question_recovery}
	QuestionRecovery *string `field:"optional" json:"questionRecovery" yaml:"questionRecovery"`
	// Lifetime in minutes of the recovery email token.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#recovery_email_token PasswordPolicy#recovery_email_token}
	RecoveryEmailToken *float64 `field:"optional" json:"recoveryEmailToken" yaml:"recoveryEmailToken"`
	// When an Active Directory user is locked out of Okta, the Okta unlock operation should also attempt to unlock the user's Windows account.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#skip_unlock PasswordPolicy#skip_unlock}
	SkipUnlock interface{} `field:"optional" json:"skipUnlock" yaml:"skipUnlock"`
	// Enable or disable SMS password recovery: ACTIVE or INACTIVE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#sms_recovery PasswordPolicy#sms_recovery}
	SmsRecovery *string `field:"optional" json:"smsRecovery" yaml:"smsRecovery"`
	// Policy Status: ACTIVE or INACTIVE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/password_policy#status PasswordPolicy#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

