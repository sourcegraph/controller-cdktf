package googlefirebasehostingversion


type GoogleFirebaseHostingVersionConfigRewrites struct {
	// The function to proxy requests to. Must match the exported function name exactly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_version#function GoogleFirebaseHostingVersion#function}
	Function *string `field:"optional" json:"function" yaml:"function"`
	// The user-supplied glob to match against the request URL path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_version#glob GoogleFirebaseHostingVersion#glob}
	Glob *string `field:"optional" json:"glob" yaml:"glob"`
	// The URL path to rewrite the request to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_version#path GoogleFirebaseHostingVersion#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The user-supplied RE2 regular expression to match against the request URL path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_version#regex GoogleFirebaseHostingVersion#regex}
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// run block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_version#run GoogleFirebaseHostingVersion#run}
	Run *GoogleFirebaseHostingVersionConfigRewritesRun `field:"optional" json:"run" yaml:"run"`
}

