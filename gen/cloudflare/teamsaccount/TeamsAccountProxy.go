package teamsaccount


type TeamsAccountProxy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_account#tcp TeamsAccount#tcp}.
	Tcp interface{} `field:"required" json:"tcp" yaml:"tcp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_account#udp TeamsAccount#udp}.
	Udp interface{} `field:"required" json:"udp" yaml:"udp"`
}

