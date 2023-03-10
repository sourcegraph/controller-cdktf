package dataawsidentitystoreuser


type DataAwsIdentitystoreUserAlternateIdentifier struct {
	// external_id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/identitystore_user#external_id DataAwsIdentitystoreUser#external_id}
	ExternalId *DataAwsIdentitystoreUserAlternateIdentifierExternalId `field:"optional" json:"externalId" yaml:"externalId"`
	// unique_attribute block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/identitystore_user#unique_attribute DataAwsIdentitystoreUser#unique_attribute}
	UniqueAttribute *DataAwsIdentitystoreUserAlternateIdentifierUniqueAttribute `field:"optional" json:"uniqueAttribute" yaml:"uniqueAttribute"`
}

