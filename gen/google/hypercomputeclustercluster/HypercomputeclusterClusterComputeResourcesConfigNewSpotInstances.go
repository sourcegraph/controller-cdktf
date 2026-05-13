package hypercomputeclustercluster


type HypercomputeclusterClusterComputeResourcesConfigNewSpotInstances struct {
	// Name of the Compute Engine [machine type](https://cloud.google.com/compute/docs/machine-resource) to use, e.g. 'n2-standard-2'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#machine_type HypercomputeclusterCluster#machine_type}
	MachineType *string `field:"required" json:"machineType" yaml:"machineType"`
	// Name of the zone in which VM instances should run, e.g., 'us-central1-a'. Must be in the same region as the cluster, and must match the zone of any other resources specified in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#zone HypercomputeclusterCluster#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// Specifies the termination action of the instance Possible values: STOP DELETE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#termination_action HypercomputeclusterCluster#termination_action}
	TerminationAction *string `field:"optional" json:"terminationAction" yaml:"terminationAction"`
}

