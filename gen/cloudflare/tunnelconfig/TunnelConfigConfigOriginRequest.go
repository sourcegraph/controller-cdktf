package tunnelconfig


type TunnelConfigConfigOriginRequest struct {
	// Runs as jump host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/tunnel_config#bastion_mode TunnelConfigA#bastion_mode}
	BastionMode interface{} `field:"optional" json:"bastionMode" yaml:"bastionMode"`
	// Path to the certificate authority (CA) for the certificate of your origin.
	//
	// This option should be used only if your certificate is not signed by Cloudflare. Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/tunnel_config#ca_pool TunnelConfigA#ca_pool}
	CaPool *string `field:"optional" json:"caPool" yaml:"caPool"`
	// Timeout for establishing a new TCP connection to your origin server.
	//
	// This excludes the time taken to establish TLS, which is controlled by `tlsTimeout`. Defaults to `30s`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/tunnel_config#connect_timeout TunnelConfigA#connect_timeout}
	ConnectTimeout *string `field:"optional" json:"connectTimeout" yaml:"connectTimeout"`
	// Disables chunked transfer encoding. Useful if you are running a Web Server Gateway Interface (WSGI) server. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/tunnel_config#disable_chunked_encoding TunnelConfigA#disable_chunked_encoding}
	DisableChunkedEncoding interface{} `field:"optional" json:"disableChunkedEncoding" yaml:"disableChunkedEncoding"`
	// Sets the HTTP Host header on requests sent to the local service. Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/tunnel_config#http_host_header TunnelConfigA#http_host_header}
	HttpHostHeader *string `field:"optional" json:"httpHostHeader" yaml:"httpHostHeader"`
	// ip_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/tunnel_config#ip_rules TunnelConfigA#ip_rules}
	IpRules interface{} `field:"optional" json:"ipRules" yaml:"ipRules"`
	// Maximum number of idle keepalive connections between Tunnel and your origin.
	//
	// This does not restrict the total number of concurrent connections. Defaults to `100`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/tunnel_config#keep_alive_connections TunnelConfigA#keep_alive_connections}
	KeepAliveConnections *float64 `field:"optional" json:"keepAliveConnections" yaml:"keepAliveConnections"`
	// Timeout after which an idle keepalive connection can be discarded. Defaults to `1m30s`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/tunnel_config#keep_alive_timeout TunnelConfigA#keep_alive_timeout}
	KeepAliveTimeout *string `field:"optional" json:"keepAliveTimeout" yaml:"keepAliveTimeout"`
	// Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local network has misconfigured one of the protocols.
	//
	// Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/tunnel_config#no_happy_eyeballs TunnelConfigA#no_happy_eyeballs}
	NoHappyEyeballs interface{} `field:"optional" json:"noHappyEyeballs" yaml:"noHappyEyeballs"`
	// Disables TLS verification of the certificate presented by your origin.
	//
	// Will allow any certificate from the origin to be accepted. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/tunnel_config#no_tls_verify TunnelConfigA#no_tls_verify}
	NoTlsVerify interface{} `field:"optional" json:"noTlsVerify" yaml:"noTlsVerify"`
	// Hostname that cloudflared should expect from your origin server certificate. Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/tunnel_config#origin_server_name TunnelConfigA#origin_server_name}
	OriginServerName *string `field:"optional" json:"originServerName" yaml:"originServerName"`
	// cloudflared starts a proxy server to translate HTTP traffic into TCP when proxying, for example, SSH or RDP.
	//
	// This configures the listen address for that proxy. Defaults to `127.0.0.1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/tunnel_config#proxy_address TunnelConfigA#proxy_address}
	ProxyAddress *string `field:"optional" json:"proxyAddress" yaml:"proxyAddress"`
	// cloudflared starts a proxy server to translate HTTP traffic into TCP when proxying, for example, SSH or RDP.
	//
	// This configures the listen port for that proxy. If set to zero, an unused port will randomly be chosen. Defaults to `0`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/tunnel_config#proxy_port TunnelConfigA#proxy_port}
	ProxyPort *float64 `field:"optional" json:"proxyPort" yaml:"proxyPort"`
	// cloudflared starts a proxy server to translate HTTP traffic into TCP when proxying, for example, SSH or RDP.
	//
	// This configures what type of proxy will be started. Available values: `""`, `socks`. Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/tunnel_config#proxy_type TunnelConfigA#proxy_type}
	ProxyType *string `field:"optional" json:"proxyType" yaml:"proxyType"`
	// The timeout after which a TCP keepalive packet is sent on a connection between Tunnel and the origin server.
	//
	// Defaults to `30s`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/tunnel_config#tcp_keep_alive TunnelConfigA#tcp_keep_alive}
	TcpKeepAlive *string `field:"optional" json:"tcpKeepAlive" yaml:"tcpKeepAlive"`
	// Timeout for completing a TLS handshake to your origin server, if you have chosen to connect Tunnel to an HTTPS server.
	//
	// Defaults to `10s`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/tunnel_config#tls_timeout TunnelConfigA#tls_timeout}
	TlsTimeout *string `field:"optional" json:"tlsTimeout" yaml:"tlsTimeout"`
}

