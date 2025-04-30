package googletranscoderjob


type GoogleTranscoderJobConfigOverlays struct {
	// animations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#animations GoogleTranscoderJob#animations}
	Animations interface{} `field:"optional" json:"animations" yaml:"animations"`
	// image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#image GoogleTranscoderJob#image}
	Image *GoogleTranscoderJobConfigOverlaysImage `field:"optional" json:"image" yaml:"image"`
}

