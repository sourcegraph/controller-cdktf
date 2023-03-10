package cognitoresourceserver


type CognitoResourceServerScope struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_resource_server#scope_description CognitoResourceServer#scope_description}.
	ScopeDescription *string `field:"required" json:"scopeDescription" yaml:"scopeDescription"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_resource_server#scope_name CognitoResourceServer#scope_name}.
	ScopeName *string `field:"required" json:"scopeName" yaml:"scopeName"`
}

