package hypercomputeclustercluster


type HypercomputeclusterClusterOrchestratorSlurmPartitions struct {
	// ID of the partition, which is how users will identify it.
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
	// IDs of the nodesets that make up this partition. Values must match SlurmNodeSet.id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#node_set_ids HypercomputeclusterCluster#node_set_ids}
	NodeSetIds *[]*string `field:"required" json:"nodeSetIds" yaml:"nodeSetIds"`
}

