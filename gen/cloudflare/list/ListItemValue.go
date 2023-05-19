package list


type ListItemValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/list#ip List#ip}.
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/list#redirect List#redirect}
	Redirect interface{} `field:"optional" json:"redirect" yaml:"redirect"`
}

