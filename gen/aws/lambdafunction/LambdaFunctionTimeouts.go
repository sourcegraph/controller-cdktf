package lambdafunction


type LambdaFunctionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/lambda_function#create LambdaFunction#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/lambda_function#update LambdaFunction#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

