package googlegameservicesgameservercluster


type GoogleGameServicesGameServerClusterConnectionInfoGkeClusterReference struct {
	// The full or partial name of a GKE cluster, using one of the following forms:.
	//
	// 'projects/{project_id}/locations/{location}/clusters/{cluster_id}'
	// 'locations/{location}/clusters/{cluster_id}'
	// '{cluster_id}'
	//
	// If project and location are not specified, the project and location of the
	// GameServerCluster resource are used to generate the full name of the
	// GKE cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_game_services_game_server_cluster#cluster GoogleGameServicesGameServerCluster#cluster}
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
}

