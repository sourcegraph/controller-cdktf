package googleiamworkforcepoolproviderkey


type GoogleIamWorkforcePoolProviderKeyKeyData struct {
	// The specifications for the key. Possible values: ["RSA_2048", "RSA_3072", "RSA_4096"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workforce_pool_provider_key#key_spec GoogleIamWorkforcePoolProviderKey#key_spec}
	KeySpec *string `field:"required" json:"keySpec" yaml:"keySpec"`
}

