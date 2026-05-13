package googlecontainercluster


type GoogleContainerClusterFleet struct {
	// The type of the cluster's fleet membership.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#membership_type GoogleContainerCluster#membership_type}
	MembershipType *string `field:"optional" json:"membershipType" yaml:"membershipType"`
	// The Fleet host project of the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#project GoogleContainerCluster#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

