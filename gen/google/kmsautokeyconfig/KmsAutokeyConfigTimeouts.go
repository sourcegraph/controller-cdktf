package kmsautokeyconfig


type KmsAutokeyConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/kms_autokey_config#create KmsAutokeyConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/kms_autokey_config#delete KmsAutokeyConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/kms_autokey_config#update KmsAutokeyConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

