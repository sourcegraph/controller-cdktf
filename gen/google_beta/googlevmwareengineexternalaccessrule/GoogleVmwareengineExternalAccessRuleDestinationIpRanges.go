package googlevmwareengineexternalaccessrule


type GoogleVmwareengineExternalAccessRuleDestinationIpRanges struct {
	// The name of an 'ExternalAddress' resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vmwareengine_external_access_rule#external_address GoogleVmwareengineExternalAccessRule#external_address}
	ExternalAddress *string `field:"optional" json:"externalAddress" yaml:"externalAddress"`
	// An IP address range in the CIDR format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vmwareengine_external_access_rule#ip_address_range GoogleVmwareengineExternalAccessRule#ip_address_range}
	IpAddressRange *string `field:"optional" json:"ipAddressRange" yaml:"ipAddressRange"`
}

