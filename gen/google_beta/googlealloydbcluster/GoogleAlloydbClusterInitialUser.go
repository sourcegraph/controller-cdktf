package googlealloydbcluster


type GoogleAlloydbClusterInitialUser struct {
	// The initial password for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#password GoogleAlloydbCluster#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The database username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#user GoogleAlloydbCluster#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

