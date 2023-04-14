package teamsaccount


type TeamsAccountFips struct {
	// Only allow FIPS-compliant TLS configuration.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_account#tls TeamsAccount#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

