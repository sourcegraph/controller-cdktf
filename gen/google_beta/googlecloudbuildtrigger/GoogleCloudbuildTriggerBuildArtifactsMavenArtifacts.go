package googlecloudbuildtrigger


type GoogleCloudbuildTriggerBuildArtifactsMavenArtifacts struct {
	// Maven artifactId value used when uploading the artifact to Artifact Registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#artifact_id GoogleCloudbuildTrigger#artifact_id}
	ArtifactId *string `field:"optional" json:"artifactId" yaml:"artifactId"`
	// Maven groupId value used when uploading the artifact to Artifact Registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#group_id GoogleCloudbuildTrigger#group_id}
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// Path to an artifact in the build's workspace to be uploaded to Artifact Registry.
	//
	// This can be either an absolute path, e.g. /workspace/my-app/target/my-app-1.0.SNAPSHOT.jar or a relative path from /workspace, e.g. my-app/target/my-app-1.0.SNAPSHOT.jar.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#path GoogleCloudbuildTrigger#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Artifact Registry repository, in the form "https://$REGION-maven.pkg.dev/$PROJECT/$REPOSITORY".
	//
	// Artifact in the workspace specified by path will be uploaded to Artifact Registry with this location as a prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#repository GoogleCloudbuildTrigger#repository}
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// Maven version value used when uploading the artifact to Artifact Registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#version GoogleCloudbuildTrigger#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

