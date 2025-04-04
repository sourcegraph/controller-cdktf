package dynamodbglobaltable


type DynamodbGlobalTableTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/dynamodb_global_table#create DynamodbGlobalTable#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/dynamodb_global_table#delete DynamodbGlobalTable#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/dynamodb_global_table#update DynamodbGlobalTable#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

