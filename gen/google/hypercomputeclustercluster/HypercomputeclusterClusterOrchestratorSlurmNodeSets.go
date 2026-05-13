package hypercomputeclustercluster


type HypercomputeclusterClusterOrchestratorSlurmNodeSets struct {
	// Identifier for the nodeset, which allows it to be referenced by partitions.
	//
	// Must conform to
	// [RFC-1034](https://datatracker.ietf.org/doc/html/rfc1034) (lower-case,
	// alphanumeric, and at most 63 characters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#id HypercomputeclusterCluster#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// ID of the compute resource on which this nodeset will run. Must match a key in the cluster's [compute_resources](Cluster.compute_resources).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#compute_id HypercomputeclusterCluster#compute_id}
	ComputeId *string `field:"optional" json:"computeId" yaml:"computeId"`
	// compute_instance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#compute_instance HypercomputeclusterCluster#compute_instance}
	ComputeInstance *HypercomputeclusterClusterOrchestratorSlurmNodeSetsComputeInstance `field:"optional" json:"computeInstance" yaml:"computeInstance"`
	// Controls how many additional nodes a cluster can bring online to handle workloads.
	//
	// Set this value to enable dynamic node creation and limit the
	// number of additional nodes the cluster can bring online. Leave empty if you
	// do not want the cluster to create nodes dynamically, and instead rely only
	// on static nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#max_dynamic_node_count HypercomputeclusterCluster#max_dynamic_node_count}
	MaxDynamicNodeCount *string `field:"optional" json:"maxDynamicNodeCount" yaml:"maxDynamicNodeCount"`
	// Number of nodes to be statically created for this nodeset.
	//
	// The cluster will
	// attempt to ensure that at least this many nodes exist at all times.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#static_node_count HypercomputeclusterCluster#static_node_count}
	StaticNodeCount *string `field:"optional" json:"staticNodeCount" yaml:"staticNodeCount"`
	// storage_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#storage_configs HypercomputeclusterCluster#storage_configs}
	StorageConfigs interface{} `field:"optional" json:"storageConfigs" yaml:"storageConfigs"`
}

