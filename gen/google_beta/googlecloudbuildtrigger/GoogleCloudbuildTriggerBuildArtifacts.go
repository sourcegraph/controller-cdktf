package googlecloudbuildtrigger


type GoogleCloudbuildTriggerBuildArtifacts struct {
	// A list of images to be pushed upon the successful completion of all build steps.
	//
	// The images will be pushed using the builder service account's credentials.
	//
	// The digests of the pushed images will be stored in the Build resource's results field.
	//
	// If any of the images fail to be pushed, the build is marked FAILURE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_cloudbuild_trigger#images GoogleCloudbuildTrigger#images}
	Images *[]*string `field:"optional" json:"images" yaml:"images"`
	// objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_cloudbuild_trigger#objects GoogleCloudbuildTrigger#objects}
	Objects *GoogleCloudbuildTriggerBuildArtifactsObjects `field:"optional" json:"objects" yaml:"objects"`
}

