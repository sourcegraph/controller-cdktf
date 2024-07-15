package googleartifactregistryrepository


type GoogleArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentials struct {
	// username_password_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#username_password_credentials GoogleArtifactRegistryRepository#username_password_credentials}
	UsernamePasswordCredentials *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigUpstreamCredentialsUsernamePasswordCredentials `field:"optional" json:"usernamePasswordCredentials" yaml:"usernamePasswordCredentials"`
}

