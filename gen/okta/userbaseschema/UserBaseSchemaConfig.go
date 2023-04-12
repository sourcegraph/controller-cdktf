package userbaseschema

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserBaseSchemaConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Subschema unique string identifier.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/user_base_schema#index UserBaseSchema#index}
	Index *string `field:"required" json:"index" yaml:"index"`
	// Subschema title (display name).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/user_base_schema#title UserBaseSchema#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// Subschema type: string, boolean, number, integer, array, or object.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/user_base_schema#type UserBaseSchema#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/user_base_schema#id UserBaseSchema#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// SubSchema profile manager, if not set it will inherit its setting.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/user_base_schema#master UserBaseSchema#master}
	Master *string `field:"optional" json:"master" yaml:"master"`
	// The validation pattern to use for the subschema. Must be in form of '.+', or '[<pattern>]+' if present.'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/user_base_schema#pattern UserBaseSchema#pattern}
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// SubSchema permissions: HIDE, READ_ONLY, or READ_WRITE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/user_base_schema#permissions UserBaseSchema#permissions}
	Permissions *string `field:"optional" json:"permissions" yaml:"permissions"`
	// Whether the subschema is required.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/user_base_schema#required UserBaseSchema#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// Custom subschema user type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/user_base_schema#user_type UserBaseSchema#user_type}
	UserType *string `field:"optional" json:"userType" yaml:"userType"`
}

