package googleartifactregistryrepository


type GoogleArtifactRegistryRepositoryCleanupPoliciesMostRecentVersions struct {
	// Minimum number of versions to keep.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#keep_count GoogleArtifactRegistryRepository#keep_count}
	KeepCount *float64 `field:"optional" json:"keepCount" yaml:"keepCount"`
	// Match versions by package prefix. Applied on any prefix match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_artifact_registry_repository#package_name_prefixes GoogleArtifactRegistryRepository#package_name_prefixes}
	PackageNamePrefixes *[]*string `field:"optional" json:"packageNamePrefixes" yaml:"packageNamePrefixes"`
}

