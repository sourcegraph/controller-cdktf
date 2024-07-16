package googlecloudbuildtrigger


type GoogleCloudbuildTriggerBuildArtifactsPythonPackages struct {
	// Path globs used to match files in the build's workspace.
	//
	// For Python/ Twine, this is usually dist/*, and sometimes additionally an .asc file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#paths GoogleCloudbuildTrigger#paths}
	Paths *[]*string `field:"optional" json:"paths" yaml:"paths"`
	// Artifact Registry repository, in the form "https://$REGION-python.pkg.dev/$PROJECT/$REPOSITORY".
	//
	// Files in the workspace matching any path pattern will be uploaded to Artifact Registry with this location as a prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#repository GoogleCloudbuildTrigger#repository}
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
}

