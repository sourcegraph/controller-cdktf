package provider


type PostgresqlProviderClientcert struct {
	// The SSL client certificate file path. The file must contain PEM encoded data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.18.0/docs#cert PostgresqlProvider#cert}
	Cert *string `field:"required" json:"cert" yaml:"cert"`
	// The SSL client certificate private key file path. The file must contain PEM encoded data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.18.0/docs#key PostgresqlProvider#key}
	Key *string `field:"required" json:"key" yaml:"key"`
}

