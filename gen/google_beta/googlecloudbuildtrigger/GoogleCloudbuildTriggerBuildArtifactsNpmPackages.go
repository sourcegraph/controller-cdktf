package googlecloudbuildtrigger


type GoogleCloudbuildTriggerBuildArtifactsNpmPackages struct {
	// Path to the package.json. e.g. workspace/path/to/package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#package_path GoogleCloudbuildTrigger#package_path}
	PackagePath *string `field:"optional" json:"packagePath" yaml:"packagePath"`
	// Artifact Registry repository, in the form "https://$REGION-npm.pkg.dev/$PROJECT/$REPOSITORY".
	//
	// Npm package in the workspace specified by path will be zipped and uploaded to Artifact Registry with this location as a prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#repository GoogleCloudbuildTrigger#repository}
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
}

