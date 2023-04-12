package googleartifactregistryrepository


type GoogleArtifactRegistryRepositoryTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_artifact_registry_repository#create GoogleArtifactRegistryRepository#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_artifact_registry_repository#delete GoogleArtifactRegistryRepository#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_artifact_registry_repository#update GoogleArtifactRegistryRepository#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

