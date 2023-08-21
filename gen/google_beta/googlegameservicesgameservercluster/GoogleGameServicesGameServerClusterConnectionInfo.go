package googlegameservicesgameservercluster


type GoogleGameServicesGameServerClusterConnectionInfo struct {
	// gke_cluster_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_game_services_game_server_cluster#gke_cluster_reference GoogleGameServicesGameServerCluster#gke_cluster_reference}
	GkeClusterReference *GoogleGameServicesGameServerClusterConnectionInfoGkeClusterReference `field:"required" json:"gkeClusterReference" yaml:"gkeClusterReference"`
	// Namespace designated on the game server cluster where the game server instances will be created.
	//
	// The namespace existence will be validated
	// during creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_game_services_game_server_cluster#namespace GoogleGameServicesGameServerCluster#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

