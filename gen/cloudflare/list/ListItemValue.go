package list


type ListItemValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/list#asn List#asn}.
	Asn *float64 `field:"optional" json:"asn" yaml:"asn"`
	// hostname block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/list#hostname List#hostname}
	Hostname interface{} `field:"optional" json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/list#ip List#ip}.
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/list#redirect List#redirect}
	Redirect interface{} `field:"optional" json:"redirect" yaml:"redirect"`
}

