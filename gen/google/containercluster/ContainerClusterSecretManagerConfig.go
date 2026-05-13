package containercluster


type ContainerClusterSecretManagerConfig struct {
	// Enable the Secret manager csi component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// rotation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#rotation_config ContainerCluster#rotation_config}
	RotationConfig *ContainerClusterSecretManagerConfigRotationConfig `field:"optional" json:"rotationConfig" yaml:"rotationConfig"`
}

