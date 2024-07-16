package googlecloudbuildtrigger


type GoogleCloudbuildTriggerBuildSourceStorageSource struct {
	// Google Cloud Storage bucket containing the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#bucket GoogleCloudbuildTrigger#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Google Cloud Storage object containing the source. This object must be a gzipped archive file (.tar.gz) containing source to build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#object GoogleCloudbuildTrigger#object}
	Object *string `field:"required" json:"object" yaml:"object"`
	// Google Cloud Storage generation for the object. If the generation is omitted, the latest generation will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#generation GoogleCloudbuildTrigger#generation}
	Generation *string `field:"optional" json:"generation" yaml:"generation"`
}

