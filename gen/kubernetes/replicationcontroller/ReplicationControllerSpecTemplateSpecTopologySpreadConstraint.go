package replicationcontroller


type ReplicationControllerSpecTemplateSpecTopologySpreadConstraint struct {
	// label_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/replication_controller#label_selector ReplicationController#label_selector}
	LabelSelector interface{} `field:"optional" json:"labelSelector" yaml:"labelSelector"`
	// describes the degree to which pods may be unevenly distributed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/replication_controller#max_skew ReplicationController#max_skew}
	MaxSkew *float64 `field:"optional" json:"maxSkew" yaml:"maxSkew"`
	// the key of node labels.
	//
	// Nodes that have a label with this key and identical values are considered to be in the same topology.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/replication_controller#topology_key ReplicationController#topology_key}
	TopologyKey *string `field:"optional" json:"topologyKey" yaml:"topologyKey"`
	// indicates how to deal with a pod if it doesn't satisfy the spread constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/replication_controller#when_unsatisfiable ReplicationController#when_unsatisfiable}
	WhenUnsatisfiable *string `field:"optional" json:"whenUnsatisfiable" yaml:"whenUnsatisfiable"`
}

