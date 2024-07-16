package googleartifactregistryrepository


type GoogleArtifactRegistryRepositoryCleanupPolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#id GoogleArtifactRegistryRepository#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Policy action. Possible values: ["DELETE", "KEEP"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#action GoogleArtifactRegistryRepository#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#condition GoogleArtifactRegistryRepository#condition}
	Condition *GoogleArtifactRegistryRepositoryCleanupPoliciesCondition `field:"optional" json:"condition" yaml:"condition"`
	// most_recent_versions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#most_recent_versions GoogleArtifactRegistryRepository#most_recent_versions}
	MostRecentVersions *GoogleArtifactRegistryRepositoryCleanupPoliciesMostRecentVersions `field:"optional" json:"mostRecentVersions" yaml:"mostRecentVersions"`
}

