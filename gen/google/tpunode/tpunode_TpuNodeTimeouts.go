package tpunode


type TpuNodeTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/tpu_node#create TpuNode#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/tpu_node#delete TpuNode#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/tpu_node#update TpuNode#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
