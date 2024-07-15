package googlekmscryptokeyversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleKmsCryptoKeyVersionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the cryptoKey associated with the CryptoKeyVersions. Format: ''projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}''.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_kms_crypto_key_version#crypto_key GoogleKmsCryptoKeyVersion#crypto_key}
	CryptoKey *string `field:"required" json:"cryptoKey" yaml:"cryptoKey"`
	// external_protection_level_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_kms_crypto_key_version#external_protection_level_options GoogleKmsCryptoKeyVersion#external_protection_level_options}
	ExternalProtectionLevelOptions *GoogleKmsCryptoKeyVersionExternalProtectionLevelOptions `field:"optional" json:"externalProtectionLevelOptions" yaml:"externalProtectionLevelOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_kms_crypto_key_version#id GoogleKmsCryptoKeyVersion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The current state of the CryptoKeyVersion. Possible values: ["PENDING_GENERATION", "ENABLED", "DISABLED", "DESTROYED", "DESTROY_SCHEDULED", "PENDING_IMPORT", "IMPORT_FAILED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_kms_crypto_key_version#state GoogleKmsCryptoKeyVersion#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_kms_crypto_key_version#timeouts GoogleKmsCryptoKeyVersion#timeouts}
	Timeouts *GoogleKmsCryptoKeyVersionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

