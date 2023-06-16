package teamsaccount


type TeamsAccountProxy struct {
	// Whether gateway proxy is enabled on gateway devices for TCP traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#tcp TeamsAccount#tcp}
	Tcp interface{} `field:"required" json:"tcp" yaml:"tcp"`
	// Whether gateway proxy is enabled on gateway devices for UDP traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#udp TeamsAccount#udp}
	Udp interface{} `field:"required" json:"udp" yaml:"udp"`
}

