package record


type RecordTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#create Record#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/record#update Record#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

