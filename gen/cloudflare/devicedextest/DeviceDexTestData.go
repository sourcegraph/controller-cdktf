package devicedextest


type DeviceDexTestData struct {
	// The host URL for `http` test `kind`. For `traceroute`, it must be a valid hostname or IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_dex_test#host DeviceDexTest#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// The type of Device Dex Test. Available values: `http`, `traceroute`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_dex_test#kind DeviceDexTest#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// The http request method. Available values: `GET`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_dex_test#method DeviceDexTest#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
}

