package googlecloudbuildtrigger


type GoogleCloudbuildTriggerBuildSource struct {
	// repo_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudbuild_trigger#repo_source GoogleCloudbuildTrigger#repo_source}
	RepoSource *GoogleCloudbuildTriggerBuildSourceRepoSource `field:"optional" json:"repoSource" yaml:"repoSource"`
	// storage_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudbuild_trigger#storage_source GoogleCloudbuildTrigger#storage_source}
	StorageSource *GoogleCloudbuildTriggerBuildSourceStorageSource `field:"optional" json:"storageSource" yaml:"storageSource"`
}

