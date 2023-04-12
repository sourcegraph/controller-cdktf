package iplist


type IpListItem struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ip_list#value IpList#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ip_list#comment IpList#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

