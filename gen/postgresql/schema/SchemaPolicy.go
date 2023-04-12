package schema


type SchemaPolicy struct {
	// If true, allow the specified ROLEs to CREATE new objects within the schema(s).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/schema#create Schema#create}
	Create interface{} `field:"optional" json:"create" yaml:"create"`
	// If true, allow the specified ROLEs to CREATE new objects within the schema(s) and GRANT the same CREATE privilege to different ROLEs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/schema#create_with_grant Schema#create_with_grant}
	CreateWithGrant interface{} `field:"optional" json:"createWithGrant" yaml:"createWithGrant"`
	// ROLE who will receive this policy (default: PUBLIC).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/schema#role Schema#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// If true, allow the specified ROLEs to use objects within the schema(s).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/schema#usage Schema#usage}
	Usage interface{} `field:"optional" json:"usage" yaml:"usage"`
	// If true, allow the specified ROLEs to use objects within the schema(s) and GRANT the same USAGE privilege to different ROLEs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/schema#usage_with_grant Schema#usage_with_grant}
	UsageWithGrant interface{} `field:"optional" json:"usageWithGrant" yaml:"usageWithGrant"`
}

