package addressmap


type AddressMapIps struct {
	// An IPv4 or IPv6 address.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/address_map#ip AddressMap#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
}

