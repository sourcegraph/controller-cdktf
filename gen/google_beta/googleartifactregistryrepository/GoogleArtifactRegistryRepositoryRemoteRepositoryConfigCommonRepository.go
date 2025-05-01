package googleartifactregistryrepository


type GoogleArtifactRegistryRepositoryRemoteRepositoryConfigCommonRepository struct {
	// One of: a.
	//
	// Artifact Registry Repository resource, e.g. 'projects/UPSTREAM_PROJECT_ID/locations/REGION/repositories/UPSTREAM_REPOSITORY'
	// b. URI to the registry, e.g. '"https://registry-1.docker.io"'
	// c. URI to Artifact Registry Repository, e.g. '"https://REGION-docker.pkg.dev/UPSTREAM_PROJECT_ID/UPSTREAM_REPOSITORY"'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_artifact_registry_repository#uri GoogleArtifactRegistryRepository#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

