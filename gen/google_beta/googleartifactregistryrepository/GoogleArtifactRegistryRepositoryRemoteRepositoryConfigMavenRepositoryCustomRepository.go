package googleartifactregistryrepository


type GoogleArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepositoryCustomRepository struct {
	// Specific uri to the registry, e.g. '"https://repo.maven.apache.org/maven2"'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#uri GoogleArtifactRegistryRepository#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

