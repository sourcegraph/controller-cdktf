package googlefirebasehostingversion


type GoogleFirebaseHostingVersionConfigHeaders struct {
	// The additional headers to add to the response. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_firebase_hosting_version#headers GoogleFirebaseHostingVersion#headers}
	Headers *map[string]*string `field:"required" json:"headers" yaml:"headers"`
	// The user-supplied glob to match against the request URL path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_firebase_hosting_version#glob GoogleFirebaseHostingVersion#glob}
	Glob *string `field:"optional" json:"glob" yaml:"glob"`
	// The user-supplied RE2 regular expression to match against the request URL path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_firebase_hosting_version#regex GoogleFirebaseHostingVersion#regex}
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
}

