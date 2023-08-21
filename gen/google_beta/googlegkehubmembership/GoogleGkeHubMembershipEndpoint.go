package googlegkehubmembership


type GoogleGkeHubMembershipEndpoint struct {
	// gke_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_gke_hub_membership#gke_cluster GoogleGkeHubMembership#gke_cluster}
	GkeCluster *GoogleGkeHubMembershipEndpointGkeCluster `field:"optional" json:"gkeCluster" yaml:"gkeCluster"`
}

