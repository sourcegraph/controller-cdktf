package teamslocation


type TeamsLocationNetworks struct {
	// CIDR notation representation of the network IP.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_location#network TeamsLocation#network}
	Network *string `field:"required" json:"network" yaml:"network"`
}

