package containercluster


type ContainerClusterSecretManagerConfigRotationConfig struct {
	// Enable the Secret manager auto rotation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The interval between two consecutive rotations. Default rotation interval is 2 minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#rotation_interval ContainerCluster#rotation_interval}
	RotationInterval *string `field:"optional" json:"rotationInterval" yaml:"rotationInterval"`
}

