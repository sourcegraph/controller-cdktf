package cloudbuildtrigger


type CloudbuildTriggerBuildArtifacts struct {
	// A list of images to be pushed upon the successful completion of all build steps.
	//
	// The images will be pushed using the builder service account's credentials.
	//
	// The digests of the pushed images will be stored in the Build resource's results field.
	//
	// If any of the images fail to be pushed, the build is marked FAILURE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#images CloudbuildTrigger#images}
	Images *[]*string `field:"optional" json:"images" yaml:"images"`
	// objects block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#objects CloudbuildTrigger#objects}
	Objects *CloudbuildTriggerBuildArtifactsObjects `field:"optional" json:"objects" yaml:"objects"`
}

