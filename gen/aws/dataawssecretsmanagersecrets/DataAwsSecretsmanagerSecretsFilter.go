package dataawssecretsmanagersecrets


type DataAwsSecretsmanagerSecretsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/data-sources/secretsmanager_secrets#name DataAwsSecretsmanagerSecrets#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/data-sources/secretsmanager_secrets#values DataAwsSecretsmanagerSecrets#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

