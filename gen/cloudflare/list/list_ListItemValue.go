package list


type ListItemValue struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/list#ip List#ip}.
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// redirect block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/list#redirect List#redirect}
	Redirect interface{} `field:"optional" json:"redirect" yaml:"redirect"`
}

