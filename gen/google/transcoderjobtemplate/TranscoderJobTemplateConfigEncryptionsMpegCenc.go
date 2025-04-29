package transcoderjobtemplate


type TranscoderJobTemplateConfigEncryptionsMpegCenc struct {
	// Specify the encryption scheme.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/transcoder_job_template#scheme TranscoderJobTemplate#scheme}
	Scheme *string `field:"required" json:"scheme" yaml:"scheme"`
}

