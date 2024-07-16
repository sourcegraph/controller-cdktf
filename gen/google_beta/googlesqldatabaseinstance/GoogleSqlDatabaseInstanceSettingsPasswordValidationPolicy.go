package googlesqldatabaseinstance


type GoogleSqlDatabaseInstanceSettingsPasswordValidationPolicy struct {
	// Whether the password policy is enabled or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_database_instance#enable_password_policy GoogleSqlDatabaseInstance#enable_password_policy}
	EnablePasswordPolicy interface{} `field:"required" json:"enablePasswordPolicy" yaml:"enablePasswordPolicy"`
	// Password complexity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_database_instance#complexity GoogleSqlDatabaseInstance#complexity}
	Complexity *string `field:"optional" json:"complexity" yaml:"complexity"`
	// Disallow username as a part of the password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_database_instance#disallow_username_substring GoogleSqlDatabaseInstance#disallow_username_substring}
	DisallowUsernameSubstring interface{} `field:"optional" json:"disallowUsernameSubstring" yaml:"disallowUsernameSubstring"`
	// Minimum number of characters allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_database_instance#min_length GoogleSqlDatabaseInstance#min_length}
	MinLength *float64 `field:"optional" json:"minLength" yaml:"minLength"`
	// Minimum interval after which the password can be changed. This flag is only supported for PostgresSQL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_database_instance#password_change_interval GoogleSqlDatabaseInstance#password_change_interval}
	PasswordChangeInterval *string `field:"optional" json:"passwordChangeInterval" yaml:"passwordChangeInterval"`
	// Number of previous passwords that cannot be reused.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_sql_database_instance#reuse_interval GoogleSqlDatabaseInstance#reuse_interval}
	ReuseInterval *float64 `field:"optional" json:"reuseInterval" yaml:"reuseInterval"`
}

