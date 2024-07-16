package googlecloudbuildtrigger


type GoogleCloudbuildTriggerBuildSource struct {
	// repo_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#repo_source GoogleCloudbuildTrigger#repo_source}
	RepoSource *GoogleCloudbuildTriggerBuildSourceRepoSource `field:"optional" json:"repoSource" yaml:"repoSource"`
	// storage_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#storage_source GoogleCloudbuildTrigger#storage_source}
	StorageSource *GoogleCloudbuildTriggerBuildSourceStorageSource `field:"optional" json:"storageSource" yaml:"storageSource"`
}

