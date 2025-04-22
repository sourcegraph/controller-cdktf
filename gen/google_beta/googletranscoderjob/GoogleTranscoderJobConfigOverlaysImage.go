package googletranscoderjob


type GoogleTranscoderJobConfigOverlaysImage struct {
	// URI of the image in Cloud Storage. For example, gs://bucket/inputs/image.png.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job#uri GoogleTranscoderJob#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

