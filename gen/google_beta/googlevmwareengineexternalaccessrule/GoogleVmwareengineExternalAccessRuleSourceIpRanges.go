package googlevmwareengineexternalaccessrule


type GoogleVmwareengineExternalAccessRuleSourceIpRanges struct {
	// A single IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vmwareengine_external_access_rule#ip_address GoogleVmwareengineExternalAccessRule#ip_address}
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// An IP address range in the CIDR format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vmwareengine_external_access_rule#ip_address_range GoogleVmwareengineExternalAccessRule#ip_address_range}
	IpAddressRange *string `field:"optional" json:"ipAddressRange" yaml:"ipAddressRange"`
}

