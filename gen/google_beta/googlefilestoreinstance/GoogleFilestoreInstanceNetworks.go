package googlefilestoreinstance


type GoogleFilestoreInstanceNetworks struct {
	// IP versions for which the instance has IP addresses assigned. Possible values: ["ADDRESS_MODE_UNSPECIFIED", "MODE_IPV4", "MODE_IPV6"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_filestore_instance#modes GoogleFilestoreInstance#modes}
	Modes *[]*string `field:"required" json:"modes" yaml:"modes"`
	// The name of the GCE VPC network to which the instance is connected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_filestore_instance#network GoogleFilestoreInstance#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// The network connect mode of the Filestore instance.
	//
	// If not provided, the connect mode defaults to
	// DIRECT_PEERING. Default value: "DIRECT_PEERING" Possible values: ["DIRECT_PEERING", "PRIVATE_SERVICE_ACCESS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_filestore_instance#connect_mode GoogleFilestoreInstance#connect_mode}
	ConnectMode *string `field:"optional" json:"connectMode" yaml:"connectMode"`
	// A /29 CIDR block that identifies the range of IP addresses reserved for this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_filestore_instance#reserved_ip_range GoogleFilestoreInstance#reserved_ip_range}
	ReservedIpRange *string `field:"optional" json:"reservedIpRange" yaml:"reservedIpRange"`
}

