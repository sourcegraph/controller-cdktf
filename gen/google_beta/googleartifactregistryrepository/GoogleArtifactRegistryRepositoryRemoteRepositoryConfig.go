package googleartifactregistryrepository


type GoogleArtifactRegistryRepositoryRemoteRepositoryConfig struct {
	// The description of the remote source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_artifact_registry_repository#description GoogleArtifactRegistryRepository#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// docker_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_artifact_registry_repository#docker_repository GoogleArtifactRegistryRepository#docker_repository}
	DockerRepository *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepository `field:"optional" json:"dockerRepository" yaml:"dockerRepository"`
	// maven_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_artifact_registry_repository#maven_repository GoogleArtifactRegistryRepository#maven_repository}
	MavenRepository *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepository `field:"optional" json:"mavenRepository" yaml:"mavenRepository"`
	// npm_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_artifact_registry_repository#npm_repository GoogleArtifactRegistryRepository#npm_repository}
	NpmRepository *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepository `field:"optional" json:"npmRepository" yaml:"npmRepository"`
	// python_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_artifact_registry_repository#python_repository GoogleArtifactRegistryRepository#python_repository}
	PythonRepository *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepository `field:"optional" json:"pythonRepository" yaml:"pythonRepository"`
}

