package googlesqluser


type GoogleSqlUserPasswordPolicy struct {
	// Number of failed attempts allowed before the user get locked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_user#allowed_failed_attempts GoogleSqlUser#allowed_failed_attempts}
	AllowedFailedAttempts *float64 `field:"optional" json:"allowedFailedAttempts" yaml:"allowedFailedAttempts"`
	// If true, the check that will lock user after too many failed login attempts will be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_user#enable_failed_attempts_check GoogleSqlUser#enable_failed_attempts_check}
	EnableFailedAttemptsCheck interface{} `field:"optional" json:"enableFailedAttemptsCheck" yaml:"enableFailedAttemptsCheck"`
	// If true, the user must specify the current password before changing the password.
	//
	// This flag is supported only for MySQL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_user#enable_password_verification GoogleSqlUser#enable_password_verification}
	EnablePasswordVerification interface{} `field:"optional" json:"enablePasswordVerification" yaml:"enablePasswordVerification"`
	// Password expiration duration with one week grace period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_user#password_expiration_duration GoogleSqlUser#password_expiration_duration}
	PasswordExpirationDuration *string `field:"optional" json:"passwordExpirationDuration" yaml:"passwordExpirationDuration"`
}

