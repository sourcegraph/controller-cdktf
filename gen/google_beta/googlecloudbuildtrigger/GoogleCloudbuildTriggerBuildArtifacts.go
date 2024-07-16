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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#images GoogleCloudbuildTrigger#images}
	Images *[]*string `field:"optional" json:"images" yaml:"images"`
	// maven_artifacts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#maven_artifacts GoogleCloudbuildTrigger#maven_artifacts}
	MavenArtifacts interface{} `field:"optional" json:"mavenArtifacts" yaml:"mavenArtifacts"`
	// npm_packages block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#npm_packages GoogleCloudbuildTrigger#npm_packages}
	NpmPackages interface{} `field:"optional" json:"npmPackages" yaml:"npmPackages"`
	// objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#objects GoogleCloudbuildTrigger#objects}
	Objects *GoogleCloudbuildTriggerBuildArtifactsObjects `field:"optional" json:"objects" yaml:"objects"`
	// python_packages block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#python_packages GoogleCloudbuildTrigger#python_packages}
	PythonPackages interface{} `field:"optional" json:"pythonPackages" yaml:"pythonPackages"`
}

