package teamslocation


type TeamsLocationNetworks struct {
	// CIDR notation representation of the network IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_location#network TeamsLocation#network}
	Network *string `field:"required" json:"network" yaml:"network"`
}

