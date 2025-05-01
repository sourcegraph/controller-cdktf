package googletranscoderjob


type GoogleTranscoderJobConfigA struct {
	// ad_breaks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#ad_breaks GoogleTranscoderJob#ad_breaks}
	AdBreaks interface{} `field:"optional" json:"adBreaks" yaml:"adBreaks"`
	// edit_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#edit_list GoogleTranscoderJob#edit_list}
	EditList interface{} `field:"optional" json:"editList" yaml:"editList"`
	// elementary_streams block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#elementary_streams GoogleTranscoderJob#elementary_streams}
	ElementaryStreams interface{} `field:"optional" json:"elementaryStreams" yaml:"elementaryStreams"`
	// encryptions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#encryptions GoogleTranscoderJob#encryptions}
	Encryptions interface{} `field:"optional" json:"encryptions" yaml:"encryptions"`
	// inputs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#inputs GoogleTranscoderJob#inputs}
	Inputs interface{} `field:"optional" json:"inputs" yaml:"inputs"`
	// manifests block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#manifests GoogleTranscoderJob#manifests}
	Manifests interface{} `field:"optional" json:"manifests" yaml:"manifests"`
	// mux_streams block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#mux_streams GoogleTranscoderJob#mux_streams}
	MuxStreams interface{} `field:"optional" json:"muxStreams" yaml:"muxStreams"`
	// output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#output GoogleTranscoderJob#output}
	Output *GoogleTranscoderJobConfigOutput `field:"optional" json:"output" yaml:"output"`
	// overlays block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#overlays GoogleTranscoderJob#overlays}
	Overlays interface{} `field:"optional" json:"overlays" yaml:"overlays"`
	// pubsub_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_transcoder_job#pubsub_destination GoogleTranscoderJob#pubsub_destination}
	PubsubDestination *GoogleTranscoderJobConfigPubsubDestination `field:"optional" json:"pubsubDestination" yaml:"pubsubDestination"`
}

