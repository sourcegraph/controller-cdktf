package googletranscoderjob


type GoogleTranscoderJobConfigOutput struct {
	// URI for the output file(s). For example, gs://my-bucket/outputs/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#uri GoogleTranscoderJob#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

