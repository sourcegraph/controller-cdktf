package googletranscoderjobtemplate


type GoogleTranscoderJobTemplateConfigMuxStreamsSegmentSettings struct {
	// Duration of the segments in seconds. The default is '6.0s'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job_template#segment_duration GoogleTranscoderJobTemplate#segment_duration}
	SegmentDuration *string `field:"optional" json:"segmentDuration" yaml:"segmentDuration"`
}

