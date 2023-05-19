package secretsmanagersecret


type SecretsmanagerSecretRotationRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/secretsmanager_secret#automatically_after_days SecretsmanagerSecret#automatically_after_days}.
	AutomaticallyAfterDays *float64 `field:"required" json:"automaticallyAfterDays" yaml:"automaticallyAfterDays"`
}

