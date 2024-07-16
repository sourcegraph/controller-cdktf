package googledataproccluster


type GoogleDataprocClusterClusterConfigAuxiliaryNodeGroups struct {
	// node_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#node_group GoogleDataprocCluster#node_group}
	NodeGroup interface{} `field:"required" json:"nodeGroup" yaml:"nodeGroup"`
	// A node group ID.
	//
	// Generated if not specified. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of from 3 to 33 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#node_group_id GoogleDataprocCluster#node_group_id}
	NodeGroupId *string `field:"optional" json:"nodeGroupId" yaml:"nodeGroupId"`
}

