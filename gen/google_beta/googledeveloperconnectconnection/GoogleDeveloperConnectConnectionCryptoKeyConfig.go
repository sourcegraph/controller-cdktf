package googledeveloperconnectconnection


type GoogleDeveloperConnectConnectionCryptoKeyConfig struct {
	// Required.
	//
	// The name of the key which is used to encrypt/decrypt customer data. For key
	// in Cloud KMS, the key should be in the format of
	// 'projects/* /locations/* /keyRings/* /cryptoKeys/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_developer_connect_connection#key_reference GoogleDeveloperConnectConnection#key_reference}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	KeyReference *string `field:"required" json:"keyReference" yaml:"keyReference"`
}

