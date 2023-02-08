package appuserschemaproperty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/appuserschemaproperty/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/okta/r/app_user_schema_property okta_app_user_schema_property}.
type AppUserSchemaProperty interface {
	cdktf.TerraformResource
	AppId() *string
	SetAppId(val *string)
	AppIdInput() *string
	ArrayEnum() *[]*string
	SetArrayEnum(val *[]*string)
	ArrayEnumInput() *[]*string
	ArrayOneOf() AppUserSchemaPropertyArrayOneOfList
	ArrayOneOfInput() interface{}
	ArrayType() *string
	SetArrayType(val *string)
	ArrayTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Enum() *[]*string
	SetEnum(val *[]*string)
	EnumInput() *[]*string
	ExternalName() *string
	SetExternalName(val *string)
	ExternalNameInput() *string
	ExternalNamespace() *string
	SetExternalNamespace(val *string)
	ExternalNamespaceInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Index() *string
	SetIndex(val *string)
	IndexInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Master() *string
	SetMaster(val *string)
	MasterInput() *string
	MaxLength() *float64
	SetMaxLength(val *float64)
	MaxLengthInput() *float64
	MinLength() *float64
	SetMinLength(val *float64)
	MinLengthInput() *float64
	// The tree node.
	Node() constructs.Node
	OneOf() AppUserSchemaPropertyOneOfList
	OneOfInput() interface{}
	Permissions() *string
	SetPermissions(val *string)
	PermissionsInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Required() interface{}
	SetRequired(val interface{})
	RequiredInput() interface{}
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Union() interface{}
	SetUnion(val interface{})
	UnionInput() interface{}
	Unique() *string
	SetUnique(val *string)
	UniqueInput() *string
	UserType() *string
	SetUserType(val *string)
	UserTypeInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutArrayOneOf(value interface{})
	PutOneOf(value interface{})
	ResetArrayEnum()
	ResetArrayOneOf()
	ResetArrayType()
	ResetDescription()
	ResetEnum()
	ResetExternalName()
	ResetExternalNamespace()
	ResetId()
	ResetMaster()
	ResetMaxLength()
	ResetMinLength()
	ResetOneOf()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPermissions()
	ResetRequired()
	ResetScope()
	ResetUnion()
	ResetUnique()
	ResetUserType()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AppUserSchemaProperty
type jsiiProxy_AppUserSchemaProperty struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppUserSchemaProperty) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) AppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) ArrayEnum() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"arrayEnum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) ArrayEnumInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"arrayEnumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) ArrayOneOf() AppUserSchemaPropertyArrayOneOfList {
	var returns AppUserSchemaPropertyArrayOneOfList
	_jsii_.Get(
		j,
		"arrayOneOf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) ArrayOneOfInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"arrayOneOfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) ArrayType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arrayType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) ArrayTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arrayTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) Enum() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) EnumInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) ExternalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) ExternalNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) ExternalNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) ExternalNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) Index() *string {
	var returns *string
	_jsii_.Get(
		j,
		"index",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) IndexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) Master() *string {
	var returns *string
	_jsii_.Get(
		j,
		"master",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) MasterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) MaxLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) MaxLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) MinLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) MinLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) OneOf() AppUserSchemaPropertyOneOfList {
	var returns AppUserSchemaPropertyOneOfList
	_jsii_.Get(
		j,
		"oneOf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) OneOfInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oneOfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) Permissions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) PermissionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) Required() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"required",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) RequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) Union() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"union",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) UnionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) Unique() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unique",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) UniqueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniqueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) UserType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppUserSchemaProperty) UserTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/okta/r/app_user_schema_property okta_app_user_schema_property} Resource.
func NewAppUserSchemaProperty(scope constructs.Construct, id *string, config *AppUserSchemaPropertyConfig) AppUserSchemaProperty {
	_init_.Initialize()

	if err := validateNewAppUserSchemaPropertyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppUserSchemaProperty{}

	_jsii_.Create(
		"okta.appUserSchemaProperty.AppUserSchemaProperty",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/okta/r/app_user_schema_property okta_app_user_schema_property} Resource.
func NewAppUserSchemaProperty_Override(a AppUserSchemaProperty, scope constructs.Construct, id *string, config *AppUserSchemaPropertyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"okta.appUserSchemaProperty.AppUserSchemaProperty",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetAppId(val *string) {
	if err := j.validateSetAppIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appId",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetArrayEnum(val *[]*string) {
	if err := j.validateSetArrayEnumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arrayEnum",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetArrayType(val *string) {
	if err := j.validateSetArrayTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arrayType",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetEnum(val *[]*string) {
	if err := j.validateSetEnumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enum",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetExternalName(val *string) {
	if err := j.validateSetExternalNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalName",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetExternalNamespace(val *string) {
	if err := j.validateSetExternalNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalNamespace",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetIndex(val *string) {
	if err := j.validateSetIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"index",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetMaster(val *string) {
	if err := j.validateSetMasterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"master",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetMaxLength(val *float64) {
	if err := j.validateSetMaxLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxLength",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetMinLength(val *float64) {
	if err := j.validateSetMinLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minLength",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetPermissions(val *string) {
	if err := j.validateSetPermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissions",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetRequired(val interface{}) {
	if err := j.validateSetRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"required",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetUnion(val interface{}) {
	if err := j.validateSetUnionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"union",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetUnique(val *string) {
	if err := j.validateSetUniqueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unique",
		val,
	)
}

func (j *jsiiProxy_AppUserSchemaProperty)SetUserType(val *string) {
	if err := j.validateSetUserTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userType",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func AppUserSchemaProperty_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppUserSchemaProperty_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"okta.appUserSchemaProperty.AppUserSchemaProperty",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppUserSchemaProperty_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"okta.appUserSchemaProperty.AppUserSchemaProperty",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AppUserSchemaProperty) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AppUserSchemaProperty) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppUserSchemaProperty) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppUserSchemaProperty) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppUserSchemaProperty) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppUserSchemaProperty) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppUserSchemaProperty) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppUserSchemaProperty) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppUserSchemaProperty) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppUserSchemaProperty) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppUserSchemaProperty) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppUserSchemaProperty) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppUserSchemaProperty) PutArrayOneOf(value interface{}) {
	if err := a.validatePutArrayOneOfParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putArrayOneOf",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppUserSchemaProperty) PutOneOf(value interface{}) {
	if err := a.validatePutOneOfParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOneOf",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppUserSchemaProperty) ResetArrayEnum() {
	_jsii_.InvokeVoid(
		a,
		"resetArrayEnum",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppUserSchemaProperty) ResetArrayOneOf() {
	_jsii_.InvokeVoid(
		a,
		"resetArrayOneOf",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppUserSchemaProperty) ResetArrayType() {
	_jsii_.InvokeVoid(
		a,
		"resetArrayType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppUserSchemaProperty) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppUserSchemaProperty) ResetEnum() {
	_jsii_.InvokeVoid(
		a,
		"resetEnum",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppUserSchemaProperty) ResetExternalName() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppUserSchemaProperty) ResetExternalNamespace() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalNamespace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppUserSchemaProperty) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppUserSchemaProperty) ResetMaster() {
	_jsii_.InvokeVoid(
		a,
		"resetMaster",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppUserSchemaProperty) ResetMaxLength() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxLength",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppUserSchemaProperty) ResetMinLength() {
	_jsii_.InvokeVoid(
		a,
		"resetMinLength",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppUserSchemaProperty) ResetOneOf() {
	_jsii_.InvokeVoid(
		a,
		"resetOneOf",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppUserSchemaProperty) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppUserSchemaProperty) ResetPermissions() {
	_jsii_.InvokeVoid(
		a,
		"resetPermissions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppUserSchemaProperty) ResetRequired() {
	_jsii_.InvokeVoid(
		a,
		"resetRequired",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppUserSchemaProperty) ResetScope() {
	_jsii_.InvokeVoid(
		a,
		"resetScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppUserSchemaProperty) ResetUnion() {
	_jsii_.InvokeVoid(
		a,
		"resetUnion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppUserSchemaProperty) ResetUnique() {
	_jsii_.InvokeVoid(
		a,
		"resetUnique",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppUserSchemaProperty) ResetUserType() {
	_jsii_.InvokeVoid(
		a,
		"resetUserType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppUserSchemaProperty) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppUserSchemaProperty) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppUserSchemaProperty) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppUserSchemaProperty) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

