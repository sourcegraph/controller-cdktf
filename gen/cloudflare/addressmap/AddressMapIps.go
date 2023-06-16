package addressmap


type AddressMapIps struct {
	// An IPv4 or IPv6 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/address_map#ip AddressMap#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
}

