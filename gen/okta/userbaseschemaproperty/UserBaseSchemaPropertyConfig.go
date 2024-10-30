package userbaseschemaproperty

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserBaseSchemaPropertyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_base_schema_property#index UserBaseSchemaProperty#index}
	Index *string `field:"required" json:"index" yaml:"index"`
	// Subschema title (display name).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_base_schema_property#title UserBaseSchemaProperty#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// The type of the schema property. It can be `string`, `boolean`, `number`, `integer`, `array`, or `object`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_base_schema_property#type UserBaseSchemaProperty#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_base_schema_property#id UserBaseSchemaProperty#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Master priority for the user schema property. It can be set to `PROFILE_MASTER` or `OKTA`. Default: `PROFILE_MASTER`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_base_schema_property#master UserBaseSchemaProperty#master}
	Master *string `field:"optional" json:"master" yaml:"master"`
	// The validation pattern to use for the subschema. Must be in form of '.+', or '[<pattern>]+' if present.'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_base_schema_property#pattern UserBaseSchemaProperty#pattern}
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// Access control permissions for the property. It can be set to `READ_WRITE`, `READ_ONLY`, `HIDE`. Default: `READ_ONLY`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_base_schema_property#permissions UserBaseSchemaProperty#permissions}
	Permissions *string `field:"optional" json:"permissions" yaml:"permissions"`
	// Whether the subschema is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_base_schema_property#required UserBaseSchemaProperty#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// User type ID. By default, it is `default`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_base_schema_property#user_type UserBaseSchemaProperty#user_type}
	UserType *string `field:"optional" json:"userType" yaml:"userType"`
}

