package managedheaders


type ManagedHeadersManagedResponseHeaders struct {
	// Whether the headers rule is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/managed_headers#enabled ManagedHeaders#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Unique headers rule identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/managed_headers#id ManagedHeaders#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
}

