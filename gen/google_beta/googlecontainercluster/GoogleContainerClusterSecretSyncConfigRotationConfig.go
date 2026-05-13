package googlecontainercluster


type GoogleContainerClusterSecretSyncConfigRotationConfig struct {
	// Enable the Secret sync auto rotation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#enabled GoogleContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The interval between two consecutive rotations. Default rotation interval is 2 minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#rotation_interval GoogleContainerCluster#rotation_interval}
	RotationInterval *string `field:"optional" json:"rotationInterval" yaml:"rotationInterval"`
}

