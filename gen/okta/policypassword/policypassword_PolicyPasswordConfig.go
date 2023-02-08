package policypassword

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PolicyPasswordConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#name PolicyPassword#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Authentication Provider: OKTA, ACTIVE_DIRECTORY or LDAP.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#auth_provider PolicyPassword#auth_provider}
	AuthProvider *string `field:"optional" json:"authProvider" yaml:"authProvider"`
	// Enable or disable voice call recovery: ACTIVE or INACTIVE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#call_recovery PolicyPassword#call_recovery}
	CallRecovery *string `field:"optional" json:"callRecovery" yaml:"callRecovery"`
	// Policy Description.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#description PolicyPassword#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Enable or disable email password recovery: ACTIVE or INACTIVE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#email_recovery PolicyPassword#email_recovery}
	EmailRecovery *string `field:"optional" json:"emailRecovery" yaml:"emailRecovery"`
	// List of Group IDs to Include.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#groups_included PolicyPassword#groups_included}
	GroupsIncluded *[]*string `field:"optional" json:"groupsIncluded" yaml:"groupsIncluded"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#id PolicyPassword#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Number of minutes before a locked account is unlocked: 0 = no limit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#password_auto_unlock_minutes PolicyPassword#password_auto_unlock_minutes}
	PasswordAutoUnlockMinutes *float64 `field:"optional" json:"passwordAutoUnlockMinutes" yaml:"passwordAutoUnlockMinutes"`
	// Check Passwords Against Common Password Dictionary.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#password_dictionary_lookup PolicyPassword#password_dictionary_lookup}
	PasswordDictionaryLookup interface{} `field:"optional" json:"passwordDictionaryLookup" yaml:"passwordDictionaryLookup"`
	// User firstName attribute must be excluded from the password.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#password_exclude_first_name PolicyPassword#password_exclude_first_name}
	PasswordExcludeFirstName interface{} `field:"optional" json:"passwordExcludeFirstName" yaml:"passwordExcludeFirstName"`
	// User lastName attribute must be excluded from the password.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#password_exclude_last_name PolicyPassword#password_exclude_last_name}
	PasswordExcludeLastName interface{} `field:"optional" json:"passwordExcludeLastName" yaml:"passwordExcludeLastName"`
	// If the user name must be excluded from the password.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#password_exclude_username PolicyPassword#password_exclude_username}
	PasswordExcludeUsername interface{} `field:"optional" json:"passwordExcludeUsername" yaml:"passwordExcludeUsername"`
	// Length in days a user will be warned before password expiry: 0 = no warning.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#password_expire_warn_days PolicyPassword#password_expire_warn_days}
	PasswordExpireWarnDays *float64 `field:"optional" json:"passwordExpireWarnDays" yaml:"passwordExpireWarnDays"`
	// Number of distinct passwords that must be created before they can be reused: 0 = none.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#password_history_count PolicyPassword#password_history_count}
	PasswordHistoryCount *float64 `field:"optional" json:"passwordHistoryCount" yaml:"passwordHistoryCount"`
	// Notification channels to use to notify a user when their account has been locked.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#password_lockout_notification_channels PolicyPassword#password_lockout_notification_channels}
	PasswordLockoutNotificationChannels *[]*string `field:"optional" json:"passwordLockoutNotificationChannels" yaml:"passwordLockoutNotificationChannels"`
	// Length in days a password is valid before expiry: 0 = no limit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#password_max_age_days PolicyPassword#password_max_age_days}
	PasswordMaxAgeDays *float64 `field:"optional" json:"passwordMaxAgeDays" yaml:"passwordMaxAgeDays"`
	// Number of unsuccessful login attempts allowed before lockout: 0 = no limit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#password_max_lockout_attempts PolicyPassword#password_max_lockout_attempts}
	PasswordMaxLockoutAttempts *float64 `field:"optional" json:"passwordMaxLockoutAttempts" yaml:"passwordMaxLockoutAttempts"`
	// Minimum time interval in minutes between password changes: 0 = no limit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#password_min_age_minutes PolicyPassword#password_min_age_minutes}
	PasswordMinAgeMinutes *float64 `field:"optional" json:"passwordMinAgeMinutes" yaml:"passwordMinAgeMinutes"`
	// Minimum password length.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#password_min_length PolicyPassword#password_min_length}
	PasswordMinLength *float64 `field:"optional" json:"passwordMinLength" yaml:"passwordMinLength"`
	// If a password must contain at least one lower case letter: 0 = no, 1 = yes.
	//
	// Default = 1
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#password_min_lowercase PolicyPassword#password_min_lowercase}
	PasswordMinLowercase *float64 `field:"optional" json:"passwordMinLowercase" yaml:"passwordMinLowercase"`
	// If a password must contain at least one number: 0 = no, 1 = yes. Default = 1.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#password_min_number PolicyPassword#password_min_number}
	PasswordMinNumber *float64 `field:"optional" json:"passwordMinNumber" yaml:"passwordMinNumber"`
	// If a password must contain at least one symbol (!@#$%^&*): 0 = no, 1 = yes. Default = 1.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#password_min_symbol PolicyPassword#password_min_symbol}
	PasswordMinSymbol *float64 `field:"optional" json:"passwordMinSymbol" yaml:"passwordMinSymbol"`
	// If a password must contain at least one upper case letter: 0 = no, 1 = yes.
	//
	// Default = 1
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#password_min_uppercase PolicyPassword#password_min_uppercase}
	PasswordMinUppercase *float64 `field:"optional" json:"passwordMinUppercase" yaml:"passwordMinUppercase"`
	// If a user should be informed when their account is locked.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#password_show_lockout_failures PolicyPassword#password_show_lockout_failures}
	PasswordShowLockoutFailures interface{} `field:"optional" json:"passwordShowLockoutFailures" yaml:"passwordShowLockoutFailures"`
	// Policy Priority, this attribute can be set to a valid priority.
	//
	// To avoid endless diff situation we error if an invalid priority is provided. API defaults it to the last (lowest) if not there.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#priority PolicyPassword#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Min length of the password recovery question answer.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#question_min_length PolicyPassword#question_min_length}
	QuestionMinLength *float64 `field:"optional" json:"questionMinLength" yaml:"questionMinLength"`
	// Enable or disable security question password recovery: ACTIVE or INACTIVE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#question_recovery PolicyPassword#question_recovery}
	QuestionRecovery *string `field:"optional" json:"questionRecovery" yaml:"questionRecovery"`
	// Lifetime in minutes of the recovery email token.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#recovery_email_token PolicyPassword#recovery_email_token}
	RecoveryEmailToken *float64 `field:"optional" json:"recoveryEmailToken" yaml:"recoveryEmailToken"`
	// When an Active Directory user is locked out of Okta, the Okta unlock operation should also attempt to unlock the user's Windows account.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#skip_unlock PolicyPassword#skip_unlock}
	SkipUnlock interface{} `field:"optional" json:"skipUnlock" yaml:"skipUnlock"`
	// Enable or disable SMS password recovery: ACTIVE or INACTIVE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#sms_recovery PolicyPassword#sms_recovery}
	SmsRecovery *string `field:"optional" json:"smsRecovery" yaml:"smsRecovery"`
	// Policy Status: ACTIVE or INACTIVE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/policy_password#status PolicyPassword#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

