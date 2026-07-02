package transcoderjobtemplate

type TranscoderJobTemplateConfigA struct {
	// ad_breaks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/transcoder_job_template#ad_breaks TranscoderJobTemplate#ad_breaks}
	AdBreaks any `field:"optional" json:"adBreaks" yaml:"adBreaks"`
	// edit_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/transcoder_job_template#edit_list TranscoderJobTemplate#edit_list}
	EditList any `field:"optional" json:"editList" yaml:"editList"`
	// elementary_streams block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/transcoder_job_template#elementary_streams TranscoderJobTemplate#elementary_streams}
	ElementaryStreams any `field:"optional" json:"elementaryStreams" yaml:"elementaryStreams"`
	// encryptions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/transcoder_job_template#encryptions TranscoderJobTemplate#encryptions}
	Encryptions any `field:"optional" json:"encryptions" yaml:"encryptions"`
	// inputs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/transcoder_job_template#inputs TranscoderJobTemplate#inputs}
	Inputs any `field:"optional" json:"inputs" yaml:"inputs"`
	// manifests block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/transcoder_job_template#manifests TranscoderJobTemplate#manifests}
	Manifests any `field:"optional" json:"manifests" yaml:"manifests"`
	// mux_streams block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/transcoder_job_template#mux_streams TranscoderJobTemplate#mux_streams}
	MuxStreams any `field:"optional" json:"muxStreams" yaml:"muxStreams"`
	// output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/transcoder_job_template#output TranscoderJobTemplate#output}
	Output *TranscoderJobTemplateConfigOutput `field:"optional" json:"output" yaml:"output"`
	// overlays block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/transcoder_job_template#overlays TranscoderJobTemplate#overlays}
	Overlays any `field:"optional" json:"overlays" yaml:"overlays"`
	// pubsub_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/transcoder_job_template#pubsub_destination TranscoderJobTemplate#pubsub_destination}
	PubsubDestination *TranscoderJobTemplateConfigPubsubDestination `field:"optional" json:"pubsubDestination" yaml:"pubsubDestination"`
}
