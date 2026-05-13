package googlenetworksecuritysacrealm


type GoogleNetworkSecuritySacRealmSymantecOptions struct {
	// API Key used to call Symantec APIs on the user's behalf.
	//
	// Required if using Symantec Cloud SWG. P4SA account needs permissions granted to read this secret.
	// A secret ID, secret name, or secret URI can be specified, but it will be parsed and stored as a secret URI in the form projects/{projectNumber}/secrets/my-secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_sac_realm#secret_path GoogleNetworkSecuritySacRealm#secret_path}
	SecretPath *string `field:"optional" json:"secretPath" yaml:"secretPath"`
}

