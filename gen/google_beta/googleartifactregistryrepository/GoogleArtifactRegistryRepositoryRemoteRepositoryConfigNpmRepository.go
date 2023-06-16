package googleartifactregistryrepository


type GoogleArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepository struct {
	// Address of the remote repository. Default value: "NPMJS" Possible values: ["NPMJS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_artifact_registry_repository#public_repository GoogleArtifactRegistryRepository#public_repository}
	PublicRepository *string `field:"optional" json:"publicRepository" yaml:"publicRepository"`
}

