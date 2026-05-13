package hypercomputeclustercluster


type HypercomputeclusterClusterOrchestratorSlurmNodeSetsStorageConfigs struct {
	// ID of the storage resource to mount, which must match a key in the cluster's [storage_resources](Cluster.storage_resources).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#id HypercomputeclusterCluster#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// A directory inside the VM instance's file system where the storage resource should be mounted (e.g., '/mnt/share').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#local_mount HypercomputeclusterCluster#local_mount}
	LocalMount *string `field:"required" json:"localMount" yaml:"localMount"`
}

