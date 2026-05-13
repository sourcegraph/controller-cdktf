package googlechroniclefeed


type GoogleChronicleFeedDetailsTrellixHxBulkAcqsSettingsAuthenticationTrellixIam struct {
	// Client ID generated in Trellix IAM.
	//
	// This is a unique identifier for the user that is generated in Trellix IAM.
	// It has no default, specifically enforced min / max length or character set.
	// It is assumed that the Client ID generated in Trellix IAM is internally
	// consistent with Trellix IAM authentication requirements / validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#client_id GoogleChronicleFeed#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Secret associated with the Client ID.
	//
	// This is the secret generated in Trellix IAM for the Client ID. It has no
	// default, specifically enforced min / max length or character set. It is
	// assumed that the secret generated in Trellix IAM is internally
	// consistent with Trellix IAM authentication requirements / validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#client_secret GoogleChronicleFeed#client_secret}
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// OAUTH 2 scope to request for the authentication token.
	//
	// This is the OAUTH 2 scope to request for the authentication token. It has
	// no default, specifically enforced min / max length or character set. It is
	// assumed that the scope provided is internally consistent with Trellix IAM
	// authentication requirements / validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#scope GoogleChronicleFeed#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

