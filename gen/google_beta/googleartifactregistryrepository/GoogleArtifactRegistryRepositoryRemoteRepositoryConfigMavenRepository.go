package googleartifactregistryrepository


type GoogleArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepository struct {
	// Address of the remote repository. Default value: "MAVEN_CENTRAL" Possible values: ["MAVEN_CENTRAL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_artifact_registry_repository#public_repository GoogleArtifactRegistryRepository#public_repository}
	PublicRepository *string `field:"optional" json:"publicRepository" yaml:"publicRepository"`
}

