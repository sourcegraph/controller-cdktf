package teamsaccount


type TeamsAccountSshSessionLog struct {
	// Public key used to encrypt ssh session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/teams_account#public_key TeamsAccount#public_key}
	PublicKey *string `field:"required" json:"publicKey" yaml:"publicKey"`
}

