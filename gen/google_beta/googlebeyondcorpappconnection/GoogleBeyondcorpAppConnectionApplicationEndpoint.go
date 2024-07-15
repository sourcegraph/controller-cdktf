package googlebeyondcorpappconnection


type GoogleBeyondcorpAppConnectionApplicationEndpoint struct {
	// Hostname or IP address of the remote application endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_beyondcorp_app_connection#host GoogleBeyondcorpAppConnection#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// Port of the remote application endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_beyondcorp_app_connection#port GoogleBeyondcorpAppConnection#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

