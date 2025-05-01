package googlenetworkconnectivitygroup


type GoogleNetworkConnectivityGroupAutoAccept struct {
	// A list of project ids or project numbers for which you want to enable auto-accept.
	//
	// The auto-accept setting is applied to spokes being created or updated in these projects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_connectivity_group#auto_accept_projects GoogleNetworkConnectivityGroup#auto_accept_projects}
	AutoAcceptProjects *[]*string `field:"required" json:"autoAcceptProjects" yaml:"autoAcceptProjects"`
}

