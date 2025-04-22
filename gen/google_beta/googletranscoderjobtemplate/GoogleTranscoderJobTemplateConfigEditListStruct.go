package googletranscoderjobtemplate


type GoogleTranscoderJobTemplateConfigEditListStruct struct {
	// List of values identifying files that should be used in this atom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job_template#inputs GoogleTranscoderJobTemplate#inputs}
	Inputs *[]*string `field:"optional" json:"inputs" yaml:"inputs"`
	// A unique key for this atom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job_template#key GoogleTranscoderJobTemplate#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Start time in seconds for the atom, relative to the input file timeline.  The default is '0s'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job_template#start_time_offset GoogleTranscoderJobTemplate#start_time_offset}
	StartTimeOffset *string `field:"optional" json:"startTimeOffset" yaml:"startTimeOffset"`
}

