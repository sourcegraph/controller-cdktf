package googleartifactregistryrepository


type GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepository struct {
	// custom_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#custom_repository GoogleArtifactRegistryRepository#custom_repository}
	CustomRepository *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepositoryCustomRepository `field:"optional" json:"customRepository" yaml:"customRepository"`
	// Address of the remote repository. Default value: "DOCKER_HUB" Possible values: ["DOCKER_HUB"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#public_repository GoogleArtifactRegistryRepository#public_repository}
	PublicRepository *string `field:"optional" json:"publicRepository" yaml:"publicRepository"`
}

