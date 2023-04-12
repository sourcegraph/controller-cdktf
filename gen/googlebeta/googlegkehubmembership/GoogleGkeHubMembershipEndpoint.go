package googlegkehubmembership


type GoogleGkeHubMembershipEndpoint struct {
	// gke_cluster block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_hub_membership#gke_cluster GoogleGkeHubMembership#gke_cluster}
	GkeCluster *GoogleGkeHubMembershipEndpointGkeCluster `field:"optional" json:"gkeCluster" yaml:"gkeCluster"`
}

