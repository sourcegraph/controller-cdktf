package googlesqluser


type GoogleSqlUserTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_sql_user#create GoogleSqlUser#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_sql_user#delete GoogleSqlUser#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_sql_user#update GoogleSqlUser#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

