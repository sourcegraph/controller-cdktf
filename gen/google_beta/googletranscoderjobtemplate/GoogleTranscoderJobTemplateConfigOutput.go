package googletranscoderjobtemplate


type GoogleTranscoderJobTemplateConfigOutput struct {
	// URI for the output file(s). For example, gs://my-bucket/outputs/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job_template#uri GoogleTranscoderJobTemplate#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

