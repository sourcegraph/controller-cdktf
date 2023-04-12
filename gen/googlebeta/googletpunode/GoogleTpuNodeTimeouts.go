package googletpunode


type GoogleTpuNodeTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_tpu_node#create GoogleTpuNode#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_tpu_node#delete GoogleTpuNode#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_tpu_node#update GoogleTpuNode#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

