package googleedgecontainercluster


type GoogleEdgecontainerClusterAuthorization struct {
	// admin_users block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgecontainer_cluster#admin_users GoogleEdgecontainerCluster#admin_users}
	AdminUsers *GoogleEdgecontainerClusterAuthorizationAdminUsers `field:"required" json:"adminUsers" yaml:"adminUsers"`
}

