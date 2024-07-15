package googledatastreamconnectionprofile


type GoogleDatastreamConnectionProfileForwardSshConnectivity struct {
	// Hostname for the SSH tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#hostname GoogleDatastreamConnectionProfile#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Username for the SSH tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#username GoogleDatastreamConnectionProfile#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// SSH password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#password GoogleDatastreamConnectionProfile#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Port for the SSH tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#port GoogleDatastreamConnectionProfile#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// SSH private key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#private_key GoogleDatastreamConnectionProfile#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
}

