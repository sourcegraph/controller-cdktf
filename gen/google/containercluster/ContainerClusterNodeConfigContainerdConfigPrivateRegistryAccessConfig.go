package containercluster

type ContainerClusterNodeConfigContainerdConfigPrivateRegistryAccessConfig struct {
	// Whether or not private registries are configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/container_cluster#enabled ContainerCluster#enabled}
	Enabled any `field:"required" json:"enabled" yaml:"enabled"`
	// certificate_authority_domain_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/container_cluster#certificate_authority_domain_config ContainerCluster#certificate_authority_domain_config}
	CertificateAuthorityDomainConfig any `field:"optional" json:"certificateAuthorityDomainConfig" yaml:"certificateAuthorityDomainConfig"`
}
