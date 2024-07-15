package googleartifactregistryrepository


type GoogleArtifactRegistryRepositoryRemoteRepositoryConfigYumRepositoryPublicRepository struct {
	// A common public repository base for Yum. Possible values: ["CENTOS", "CENTOS_DEBUG", "CENTOS_VAULT", "CENTOS_STREAM", "ROCKY", "EPEL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#repository_base GoogleArtifactRegistryRepository#repository_base}
	RepositoryBase *string `field:"required" json:"repositoryBase" yaml:"repositoryBase"`
	// Specific repository from the base, e.g. '"pub/rocky/9/BaseOS/x86_64/os"'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#repository_path GoogleArtifactRegistryRepository#repository_path}
	RepositoryPath *string `field:"required" json:"repositoryPath" yaml:"repositoryPath"`
}

