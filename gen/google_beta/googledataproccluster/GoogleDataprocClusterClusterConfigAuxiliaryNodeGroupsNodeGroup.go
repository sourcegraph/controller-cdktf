package googledataproccluster


type GoogleDataprocClusterClusterConfigAuxiliaryNodeGroupsNodeGroup struct {
	// Node group roles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#roles GoogleDataprocCluster#roles}
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
	// node_group_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#node_group_config GoogleDataprocCluster#node_group_config}
	NodeGroupConfig *GoogleDataprocClusterClusterConfigAuxiliaryNodeGroupsNodeGroupNodeGroupConfig `field:"optional" json:"nodeGroupConfig" yaml:"nodeGroupConfig"`
}

