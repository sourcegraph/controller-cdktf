package googletranscoderjobtemplate


type GoogleTranscoderJobTemplateConfigManifests struct {
	// The name of the generated file. The default is 'manifest'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job_template#file_name GoogleTranscoderJobTemplate#file_name}
	FileName *string `field:"optional" json:"fileName" yaml:"fileName"`
	// List of user supplied MuxStream.key values that should appear in this manifest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job_template#mux_streams GoogleTranscoderJobTemplate#mux_streams}
	MuxStreams *[]*string `field:"optional" json:"muxStreams" yaml:"muxStreams"`
	// Type of the manifest. Possible values: ["MANIFEST_TYPE_UNSPECIFIED", "HLS", "DASH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job_template#type GoogleTranscoderJobTemplate#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

