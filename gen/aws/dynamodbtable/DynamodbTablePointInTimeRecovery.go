package dynamodbtable


type DynamodbTablePointInTimeRecovery struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/dynamodb_table#enabled DynamodbTable#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

