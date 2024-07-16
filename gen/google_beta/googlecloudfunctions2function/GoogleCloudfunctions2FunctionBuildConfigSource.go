package googlecloudfunctions2function


type GoogleCloudfunctions2FunctionBuildConfigSource struct {
	// repo_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions2_function#repo_source GoogleCloudfunctions2Function#repo_source}
	RepoSource *GoogleCloudfunctions2FunctionBuildConfigSourceRepoSource `field:"optional" json:"repoSource" yaml:"repoSource"`
	// storage_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions2_function#storage_source GoogleCloudfunctions2Function#storage_source}
	StorageSource *GoogleCloudfunctions2FunctionBuildConfigSourceStorageSource `field:"optional" json:"storageSource" yaml:"storageSource"`
}

