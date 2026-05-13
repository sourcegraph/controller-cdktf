package googlefirebaseapphostingdomain


type GoogleFirebaseAppHostingDomainServeRedirect struct {
	// The URI of the redirect's intended destination.
	//
	// This URI will be
	// prepended to the original request path. URI without a scheme are
	// assumed to be HTTPS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_app_hosting_domain#uri GoogleFirebaseAppHostingDomain#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// The status code to use in a redirect response.
	//
	// Must be a valid HTTP 3XX
	// status code. Defaults to 302 if not present.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_app_hosting_domain#status GoogleFirebaseAppHostingDomain#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

