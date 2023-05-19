package devicemanagednetworks


type DeviceManagedNetworksConfigA struct {
	// The SHA-256 hash of the TLS certificate presented by the host found at tls_sockaddr.
	//
	// If absent, regular certificate verification (trusted roots, valid timestamp, etc) will be used to validate the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_managed_networks#sha256 DeviceManagedNetworks#sha256}
	Sha256 *string `field:"required" json:"sha256" yaml:"sha256"`
	// A network address of the form "host:port" that the WARP client will use to detect the presence of a TLS host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/device_managed_networks#tls_sockaddr DeviceManagedNetworks#tls_sockaddr}
	TlsSockaddr *string `field:"required" json:"tlsSockaddr" yaml:"tlsSockaddr"`
}

