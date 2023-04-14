package teamsaccount


type TeamsAccountPayloadLog struct {
	// Public key used to encrypt matched payloads.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_account#public_key TeamsAccount#public_key}
	PublicKey *string `field:"required" json:"publicKey" yaml:"publicKey"`
}

