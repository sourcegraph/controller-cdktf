package googletranscoderjobtemplate


type GoogleTranscoderJobTemplateConfigA struct {
	// ad_breaks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job_template#ad_breaks GoogleTranscoderJobTemplate#ad_breaks}
	AdBreaks interface{} `field:"optional" json:"adBreaks" yaml:"adBreaks"`
	// edit_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job_template#edit_list GoogleTranscoderJobTemplate#edit_list}
	EditList interface{} `field:"optional" json:"editList" yaml:"editList"`
	// elementary_streams block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job_template#elementary_streams GoogleTranscoderJobTemplate#elementary_streams}
	ElementaryStreams interface{} `field:"optional" json:"elementaryStreams" yaml:"elementaryStreams"`
	// encryptions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job_template#encryptions GoogleTranscoderJobTemplate#encryptions}
	Encryptions interface{} `field:"optional" json:"encryptions" yaml:"encryptions"`
	// inputs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job_template#inputs GoogleTranscoderJobTemplate#inputs}
	Inputs interface{} `field:"optional" json:"inputs" yaml:"inputs"`
	// manifests block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job_template#manifests GoogleTranscoderJobTemplate#manifests}
	Manifests interface{} `field:"optional" json:"manifests" yaml:"manifests"`
	// mux_streams block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job_template#mux_streams GoogleTranscoderJobTemplate#mux_streams}
	MuxStreams interface{} `field:"optional" json:"muxStreams" yaml:"muxStreams"`
	// output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job_template#output GoogleTranscoderJobTemplate#output}
	Output *GoogleTranscoderJobTemplateConfigOutput `field:"optional" json:"output" yaml:"output"`
	// overlays block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job_template#overlays GoogleTranscoderJobTemplate#overlays}
	Overlays interface{} `field:"optional" json:"overlays" yaml:"overlays"`
	// pubsub_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_transcoder_job_template#pubsub_destination GoogleTranscoderJobTemplate#pubsub_destination}
	PubsubDestination *GoogleTranscoderJobTemplateConfigPubsubDestination `field:"optional" json:"pubsubDestination" yaml:"pubsubDestination"`
}

