package googlesqluser


type GoogleSqlUserSqlServerUserDetails struct {
	// If the user has been disabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_sql_user#disabled GoogleSqlUser#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// The server roles for this user in the database.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_sql_user#server_roles GoogleSqlUser#server_roles}
	ServerRoles *[]*string `field:"optional" json:"serverRoles" yaml:"serverRoles"`
}

