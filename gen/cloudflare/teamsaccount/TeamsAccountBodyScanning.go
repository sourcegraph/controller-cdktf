package teamsaccount


type TeamsAccountBodyScanning struct {
	// Body scanning inspection mode. Available values: `deep`, `shallow`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/teams_account#inspection_mode TeamsAccount#inspection_mode}
	InspectionMode *string `field:"required" json:"inspectionMode" yaml:"inspectionMode"`
}

