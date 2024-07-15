package googleartifactregistryrepository


type GoogleArtifactRegistryRepositoryCleanupPoliciesCondition struct {
	// Match versions newer than a duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#newer_than GoogleArtifactRegistryRepository#newer_than}
	NewerThan *string `field:"optional" json:"newerThan" yaml:"newerThan"`
	// Match versions older than a duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#older_than GoogleArtifactRegistryRepository#older_than}
	OlderThan *string `field:"optional" json:"olderThan" yaml:"olderThan"`
	// Match versions by package prefix. Applied on any prefix match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#package_name_prefixes GoogleArtifactRegistryRepository#package_name_prefixes}
	PackageNamePrefixes *[]*string `field:"optional" json:"packageNamePrefixes" yaml:"packageNamePrefixes"`
	// Match versions by tag prefix. Applied on any prefix match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#tag_prefixes GoogleArtifactRegistryRepository#tag_prefixes}
	TagPrefixes *[]*string `field:"optional" json:"tagPrefixes" yaml:"tagPrefixes"`
	// Match versions by tag status. Default value: "ANY" Possible values: ["TAGGED", "UNTAGGED", "ANY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#tag_state GoogleArtifactRegistryRepository#tag_state}
	TagState *string `field:"optional" json:"tagState" yaml:"tagState"`
	// Match versions by version name prefix. Applied on any prefix match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#version_name_prefixes GoogleArtifactRegistryRepository#version_name_prefixes}
	VersionNamePrefixes *[]*string `field:"optional" json:"versionNamePrefixes" yaml:"versionNamePrefixes"`
}

