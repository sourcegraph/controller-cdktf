package hypercomputeclustercluster


type HypercomputeclusterClusterComputeResourcesConfigNewReservedInstances struct {
	// Name of the reservation from which VM instances should be created, in the format 'projects/{project}/zones/{zone}/reservations/{reservation}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#reservation HypercomputeclusterCluster#reservation}
	Reservation *string `field:"optional" json:"reservation" yaml:"reservation"`
}

