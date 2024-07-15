package googlecloudbuildtrigger


type GoogleCloudbuildTriggerBuildArtifactsObjects struct {
	// Cloud Storage bucket and optional object path, in the form "gs://bucket/path/to/somewhere/".
	//
	// Files in the workspace matching any path pattern will be uploaded to Cloud Storage with
	// this location as a prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#location GoogleCloudbuildTrigger#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Path globs used to match files in the build's workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#paths GoogleCloudbuildTrigger#paths}
	Paths *[]*string `field:"optional" json:"paths" yaml:"paths"`
}

