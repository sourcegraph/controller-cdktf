package appuserschemaproperty

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppUserSchemaPropertyConfig struct {
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
	// The Application's ID the user custom schema property should be assigned to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user_schema_property#app_id AppUserSchemaProperty#app_id}
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// Subschema unique string identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user_schema_property#index AppUserSchemaProperty#index}
	Index *string `field:"required" json:"index" yaml:"index"`
	// Subschema title (display name).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user_schema_property#title AppUserSchemaProperty#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// The type of the schema property. It can be `string`, `boolean`, `number`, `integer`, `array`, or `object`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user_schema_property#type AppUserSchemaProperty#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Array of values that an array property's items can be set to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user_schema_property#array_enum AppUserSchemaProperty#array_enum}
	ArrayEnum *[]*string `field:"optional" json:"arrayEnum" yaml:"arrayEnum"`
	// array_one_of block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user_schema_property#array_one_of AppUserSchemaProperty#array_one_of}
	ArrayOneOf interface{} `field:"optional" json:"arrayOneOf" yaml:"arrayOneOf"`
	// The type of the array elements if `type` is set to `array`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user_schema_property#array_type AppUserSchemaProperty#array_type}
	ArrayType *string `field:"optional" json:"arrayType" yaml:"arrayType"`
	// The description of the user schema property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user_schema_property#description AppUserSchemaProperty#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Array of values a primitive property can be set to. See `array_enum` for arrays.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user_schema_property#enum AppUserSchemaProperty#enum}
	Enum *[]*string `field:"optional" json:"enum" yaml:"enum"`
	// External name of the user schema property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user_schema_property#external_name AppUserSchemaProperty#external_name}
	ExternalName *string `field:"optional" json:"externalName" yaml:"externalName"`
	// External namespace of the user schema property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user_schema_property#external_namespace AppUserSchemaProperty#external_namespace}
	ExternalNamespace *string `field:"optional" json:"externalNamespace" yaml:"externalNamespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user_schema_property#id AppUserSchemaProperty#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Master priority for the user schema property. It can be set to `PROFILE_MASTER` or `OKTA`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user_schema_property#master AppUserSchemaProperty#master}
	Master *string `field:"optional" json:"master" yaml:"master"`
	// The maximum length of the user property value. Only applies to type `string`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user_schema_property#max_length AppUserSchemaProperty#max_length}
	MaxLength *float64 `field:"optional" json:"maxLength" yaml:"maxLength"`
	// The minimum length of the user property value. Only applies to type `string`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user_schema_property#min_length AppUserSchemaProperty#min_length}
	MinLength *float64 `field:"optional" json:"minLength" yaml:"minLength"`
	// one_of block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user_schema_property#one_of AppUserSchemaProperty#one_of}
	OneOf interface{} `field:"optional" json:"oneOf" yaml:"oneOf"`
	// Access control permissions for the property. It can be set to `READ_WRITE`, `READ_ONLY`, `HIDE`. Default: `READ_ONLY`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user_schema_property#permissions AppUserSchemaProperty#permissions}
	Permissions *string `field:"optional" json:"permissions" yaml:"permissions"`
	// Whether the subschema is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user_schema_property#required AppUserSchemaProperty#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// determines whether an app user attribute can be set at the Personal `SELF` or Group `NONE` level.
	//
	// Default value is `NONE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user_schema_property#scope AppUserSchemaProperty#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// If `type` is set to `array`, used to set whether attribute value is determined by group priority `false`, or combine values across groups `true`.
	//
	// Can not be set to `true` if `scope` is set to `SELF`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user_schema_property#union AppUserSchemaProperty#union}
	Union interface{} `field:"optional" json:"union" yaml:"union"`
	// Whether the property should be unique. It can be set to `UNIQUE_VALIDATED` or `NOT_UNIQUE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user_schema_property#unique AppUserSchemaProperty#unique}
	Unique *string `field:"optional" json:"unique" yaml:"unique"`
	// User type ID. By default, it is `default`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_user_schema_property#user_type AppUserSchemaProperty#user_type}
	UserType *string `field:"optional" json:"userType" yaml:"userType"`
}

