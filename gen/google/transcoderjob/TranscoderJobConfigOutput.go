package transcoderjob


type TranscoderJobConfigOutput struct {
	// URI for the output file(s). For example, gs://my-bucket/outputs/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/transcoder_job#uri TranscoderJob#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

