package cognitouserpool


type CognitoUserPoolAccountRecoverySetting struct {
	// recovery_mechanism block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/cognito_user_pool#recovery_mechanism CognitoUserPool#recovery_mechanism}
	RecoveryMechanism interface{} `field:"required" json:"recoveryMechanism" yaml:"recoveryMechanism"`
}

