package googlefirebasehostingversion


type GoogleFirebaseHostingVersionConfigRedirects struct {
	// The value to put in the HTTP location header of the response.
	//
	// The location can contain capture group values from the pattern using a : prefix to identify
	// the segment and an optional * to capture the rest of the URL. For example:
	//
	// ```hcl
	// redirects {
	//   glob = "/:capture*"
	//   status_code = 302
	//   location = "https://example.com/foo/:capture"
	// }
	// ```
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_version#location GoogleFirebaseHostingVersion#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The status HTTP code to return in the response. It must be a valid 3xx status code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_version#status_code GoogleFirebaseHostingVersion#status_code}
	StatusCode *float64 `field:"required" json:"statusCode" yaml:"statusCode"`
	// The user-supplied glob to match against the request URL path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_version#glob GoogleFirebaseHostingVersion#glob}
	Glob *string `field:"optional" json:"glob" yaml:"glob"`
	// The user-supplied RE2 regular expression to match against the request URL path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_version#regex GoogleFirebaseHostingVersion#regex}
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
}

