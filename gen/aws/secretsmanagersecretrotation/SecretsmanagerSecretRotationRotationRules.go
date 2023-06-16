package secretsmanagersecretrotation


type SecretsmanagerSecretRotationRotationRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/secretsmanager_secret_rotation#automatically_after_days SecretsmanagerSecretRotation#automatically_after_days}.
	AutomaticallyAfterDays *float64 `field:"required" json:"automaticallyAfterDays" yaml:"automaticallyAfterDays"`
}

