package googlesqlsslcert


type GoogleSqlSslCertTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_sql_ssl_cert#create GoogleSqlSslCert#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_sql_ssl_cert#delete GoogleSqlSslCert#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

