package userschemaproperty

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserSchemaPropertyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_schema_property#index UserSchemaProperty#index}
	Index *string `field:"required" json:"index" yaml:"index"`
	// Subschema title (display name).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_schema_property#title UserSchemaProperty#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// The type of the schema property. It can be `string`, `boolean`, `number`, `integer`, `array`, or `object`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_schema_property#type UserSchemaProperty#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Array of values that an array property's items can be set to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_schema_property#array_enum UserSchemaProperty#array_enum}
	ArrayEnum *[]*string `field:"optional" json:"arrayEnum" yaml:"arrayEnum"`
	// array_one_of block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_schema_property#array_one_of UserSchemaProperty#array_one_of}
	ArrayOneOf interface{} `field:"optional" json:"arrayOneOf" yaml:"arrayOneOf"`
	// The type of the array elements if `type` is set to `array`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_schema_property#array_type UserSchemaProperty#array_type}
	ArrayType *string `field:"optional" json:"arrayType" yaml:"arrayType"`
	// The description of the user schema property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_schema_property#description UserSchemaProperty#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Array of values a primitive property can be set to. See `array_enum` for arrays.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_schema_property#enum UserSchemaProperty#enum}
	Enum *[]*string `field:"optional" json:"enum" yaml:"enum"`
	// External name of the user schema property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_schema_property#external_name UserSchemaProperty#external_name}
	ExternalName *string `field:"optional" json:"externalName" yaml:"externalName"`
	// External namespace of the user schema property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_schema_property#external_namespace UserSchemaProperty#external_namespace}
	ExternalNamespace *string `field:"optional" json:"externalNamespace" yaml:"externalNamespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_schema_property#id UserSchemaProperty#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Master priority for the user schema property. It can be set to `PROFILE_MASTER`, `OVERRIDE` or `OKTA`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_schema_property#master UserSchemaProperty#master}
	Master *string `field:"optional" json:"master" yaml:"master"`
	// master_override_priority block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_schema_property#master_override_priority UserSchemaProperty#master_override_priority}
	MasterOverridePriority interface{} `field:"optional" json:"masterOverridePriority" yaml:"masterOverridePriority"`
	// The maximum length of the user property value. Only applies to type `string`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_schema_property#max_length UserSchemaProperty#max_length}
	MaxLength *float64 `field:"optional" json:"maxLength" yaml:"maxLength"`
	// The minimum length of the user property value. Only applies to type `string`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_schema_property#min_length UserSchemaProperty#min_length}
	MinLength *float64 `field:"optional" json:"minLength" yaml:"minLength"`
	// one_of block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_schema_property#one_of UserSchemaProperty#one_of}
	OneOf interface{} `field:"optional" json:"oneOf" yaml:"oneOf"`
	// The validation pattern to use for the subschema. Must be in form of '.+', or '[<pattern>]+' if present.'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_schema_property#pattern UserSchemaProperty#pattern}
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// Access control permissions for the property. It can be set to `READ_WRITE`, `READ_ONLY`, `HIDE`. Default: `READ_ONLY`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_schema_property#permissions UserSchemaProperty#permissions}
	Permissions *string `field:"optional" json:"permissions" yaml:"permissions"`
	// Whether the subschema is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_schema_property#required UserSchemaProperty#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// determines whether an app user attribute can be set at the Individual or Group Level. Default: `NONE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_schema_property#scope UserSchemaProperty#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Whether the property should be unique. It can be set to `UNIQUE_VALIDATED` or `NOT_UNIQUE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_schema_property#unique UserSchemaProperty#unique}
	Unique *string `field:"optional" json:"unique" yaml:"unique"`
	// User type ID. By default, it is `default`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/user_schema_property#user_type UserSchemaProperty#user_type}
	UserType *string `field:"optional" json:"userType" yaml:"userType"`
}

