package googleartifactregistryrepository


type GoogleArtifactRegistryRepositoryRemoteRepositoryConfig struct {
	// apt_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#apt_repository GoogleArtifactRegistryRepository#apt_repository}
	AptRepository *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigAptRepository `field:"optional" json:"aptRepository" yaml:"aptRepository"`
	// The description of the remote source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#description GoogleArtifactRegistryRepository#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If true, the remote repository upstream and upstream credentials will not be validated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#disable_upstream_validation GoogleArtifactRegistryRepository#disable_upstream_validation}
	DisableUpstreamValidation interface{} `field:"optional" json:"disableUpstreamValidation" yaml:"disableUpstreamValidation"`
	// docker_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#docker_repository GoogleArtifactRegistryRepository#docker_repository}
	DockerRepository *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepository `field:"optional" json:"dockerRepository" yaml:"dockerRepository"`
	// maven_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#maven_repository GoogleArtifactRegistryRepository#maven_repository}
	MavenRepository *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepository `field:"optional" json:"mavenRepository" yaml:"mavenRepository"`
	// npm_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#npm_repository GoogleArtifactRegistryRepository#npm_repository}
	NpmRepository *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepository `field:"optional" json:"npmRepository" yaml:"npmRepository"`
	// python_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#python_repository GoogleArtifactRegistryRepository#python_repository}
	PythonRepository *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepository `field:"optional" json:"pythonRepository" yaml:"pythonRepository"`
	// upstream_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#upstream_credentials GoogleArtifactRegistryRepository#upstream_credentials}
	UpstreamCredentials *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentials `field:"optional" json:"upstreamCredentials" yaml:"upstreamCredentials"`
	// yum_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#yum_repository GoogleArtifactRegistryRepository#yum_repository}
	YumRepository *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigYumRepository `field:"optional" json:"yumRepository" yaml:"yumRepository"`
}

