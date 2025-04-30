package googletranscoderjobtemplate


type GoogleTranscoderJobTemplateConfigAdBreaks struct {
	// Start time in seconds for the ad break, relative to the output file timeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job_template#start_time_offset GoogleTranscoderJobTemplate#start_time_offset}
	StartTimeOffset *string `field:"optional" json:"startTimeOffset" yaml:"startTimeOffset"`
}

