package googletranscoderjob


type GoogleTranscoderJobConfigManifests struct {
	// The name of the generated file. The default is 'manifest'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#file_name GoogleTranscoderJob#file_name}
	FileName *string `field:"optional" json:"fileName" yaml:"fileName"`
	// List of user supplied MuxStream.key values that should appear in this manifest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#mux_streams GoogleTranscoderJob#mux_streams}
	MuxStreams *[]*string `field:"optional" json:"muxStreams" yaml:"muxStreams"`
	// Type of the manifest. Possible values: ["MANIFEST_TYPE_UNSPECIFIED", "HLS", "DASH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#type GoogleTranscoderJob#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

