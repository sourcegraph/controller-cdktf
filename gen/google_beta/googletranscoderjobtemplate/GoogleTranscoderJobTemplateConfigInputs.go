package googletranscoderjobtemplate


type GoogleTranscoderJobTemplateConfigInputs struct {
	// A unique key for this input. Must be specified when using advanced mapping and edit lists.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job_template#key GoogleTranscoderJobTemplate#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// URI of the media.
	//
	// Input files must be at least 5 seconds in duration and stored in Cloud Storage (for example, gs://bucket/inputs/file.mp4).
	// If empty, the value is populated from Job.input_uri.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job_template#uri GoogleTranscoderJobTemplate#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

