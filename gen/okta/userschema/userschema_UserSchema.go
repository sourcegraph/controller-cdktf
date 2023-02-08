package userschema

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/userschema/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/okta/r/user_schema okta_user_schema}.
type UserSchema interface {
	cdktf.TerraformResource
	ArrayEnum() *[]*string
	SetArrayEnum(val *[]*string)
	ArrayEnumInput() *[]*string
	ArrayOneOf() UserSchemaArrayOneOfList
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
	MasterOverridePriority() UserSchemaMasterOverridePriorityList
	MasterOverridePriorityInput() interface{}
	MaxLength() *float64
	SetMaxLength(val *float64)
	MaxLengthInput() *float64
	MinLength() *float64
	SetMinLength(val *float64)
	MinLengthInput() *float64
	// The tree node.
	Node() constructs.Node
	OneOf() UserSchemaOneOfList
	OneOfInput() interface{}
	Pattern() *string
	SetPattern(val *string)
	PatternInput() *string
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
	PutMasterOverridePriority(value interface{})
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
	ResetMasterOverridePriority()
	ResetMaxLength()
	ResetMinLength()
	ResetOneOf()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPattern()
	ResetPermissions()
	ResetRequired()
	ResetScope()
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

// The jsii proxy struct for UserSchema
type jsiiProxy_UserSchema struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_UserSchema) ArrayEnum() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"arrayEnum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) ArrayEnumInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"arrayEnumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) ArrayOneOf() UserSchemaArrayOneOfList {
	var returns UserSchemaArrayOneOfList
	_jsii_.Get(
		j,
		"arrayOneOf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) ArrayOneOfInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"arrayOneOfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) ArrayType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arrayType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) ArrayTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arrayTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) Enum() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) EnumInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) ExternalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) ExternalNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) ExternalNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) ExternalNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) Index() *string {
	var returns *string
	_jsii_.Get(
		j,
		"index",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) IndexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) Master() *string {
	var returns *string
	_jsii_.Get(
		j,
		"master",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) MasterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) MasterOverridePriority() UserSchemaMasterOverridePriorityList {
	var returns UserSchemaMasterOverridePriorityList
	_jsii_.Get(
		j,
		"masterOverridePriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) MasterOverridePriorityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"masterOverridePriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) MaxLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) MaxLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) MinLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) MinLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) OneOf() UserSchemaOneOfList {
	var returns UserSchemaOneOfList
	_jsii_.Get(
		j,
		"oneOf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) OneOfInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oneOfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) Pattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) PatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) Permissions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) PermissionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) Required() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"required",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) RequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) Unique() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unique",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) UniqueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniqueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) UserType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserSchema) UserTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/okta/r/user_schema okta_user_schema} Resource.
func NewUserSchema(scope constructs.Construct, id *string, config *UserSchemaConfig) UserSchema {
	_init_.Initialize()

	if err := validateNewUserSchemaParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_UserSchema{}

	_jsii_.Create(
		"okta.userSchema.UserSchema",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/okta/r/user_schema okta_user_schema} Resource.
func NewUserSchema_Override(u UserSchema, scope constructs.Construct, id *string, config *UserSchemaConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"okta.userSchema.UserSchema",
		[]interface{}{scope, id, config},
		u,
	)
}

func (j *jsiiProxy_UserSchema)SetArrayEnum(val *[]*string) {
	if err := j.validateSetArrayEnumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arrayEnum",
		val,
	)
}

func (j *jsiiProxy_UserSchema)SetArrayType(val *string) {
	if err := j.validateSetArrayTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arrayType",
		val,
	)
}

func (j *jsiiProxy_UserSchema)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_UserSchema)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_UserSchema)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_UserSchema)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_UserSchema)SetEnum(val *[]*string) {
	if err := j.validateSetEnumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enum",
		val,
	)
}

func (j *jsiiProxy_UserSchema)SetExternalName(val *string) {
	if err := j.validateSetExternalNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalName",
		val,
	)
}

func (j *jsiiProxy_UserSchema)SetExternalNamespace(val *string) {
	if err := j.validateSetExternalNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalNamespace",
		val,
	)
}

func (j *jsiiProxy_UserSchema)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_UserSchema)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_UserSchema)SetIndex(val *string) {
	if err := j.validateSetIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"index",
		val,
	)
}

func (j *jsiiProxy_UserSchema)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_UserSchema)SetMaster(val *string) {
	if err := j.validateSetMasterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"master",
		val,
	)
}

func (j *jsiiProxy_UserSchema)SetMaxLength(val *float64) {
	if err := j.validateSetMaxLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxLength",
		val,
	)
}

func (j *jsiiProxy_UserSchema)SetMinLength(val *float64) {
	if err := j.validateSetMinLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minLength",
		val,
	)
}

func (j *jsiiProxy_UserSchema)SetPattern(val *string) {
	if err := j.validateSetPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pattern",
		val,
	)
}

func (j *jsiiProxy_UserSchema)SetPermissions(val *string) {
	if err := j.validateSetPermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissions",
		val,
	)
}

func (j *jsiiProxy_UserSchema)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_UserSchema)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_UserSchema)SetRequired(val interface{}) {
	if err := j.validateSetRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"required",
		val,
	)
}

func (j *jsiiProxy_UserSchema)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_UserSchema)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_UserSchema)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_UserSchema)SetUnique(val *string) {
	if err := j.validateSetUniqueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unique",
		val,
	)
}

func (j *jsiiProxy_UserSchema)SetUserType(val *string) {
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
func UserSchema_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUserSchema_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"okta.userSchema.UserSchema",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func UserSchema_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"okta.userSchema.UserSchema",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (u *jsiiProxy_UserSchema) AddOverride(path *string, value interface{}) {
	if err := u.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (u *jsiiProxy_UserSchema) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := u.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchema) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := u.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchema) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := u.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		u,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchema) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := u.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchema) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := u.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		u,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchema) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := u.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		u,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchema) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := u.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		u,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchema) GetStringAttribute(terraformAttribute *string) *string {
	if err := u.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		u,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchema) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := u.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		u,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchema) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := u.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchema) OverrideLogicalId(newLogicalId *string) {
	if err := u.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (u *jsiiProxy_UserSchema) PutArrayOneOf(value interface{}) {
	if err := u.validatePutArrayOneOfParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putArrayOneOf",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_UserSchema) PutMasterOverridePriority(value interface{}) {
	if err := u.validatePutMasterOverridePriorityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putMasterOverridePriority",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_UserSchema) PutOneOf(value interface{}) {
	if err := u.validatePutOneOfParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putOneOf",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_UserSchema) ResetArrayEnum() {
	_jsii_.InvokeVoid(
		u,
		"resetArrayEnum",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchema) ResetArrayOneOf() {
	_jsii_.InvokeVoid(
		u,
		"resetArrayOneOf",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchema) ResetArrayType() {
	_jsii_.InvokeVoid(
		u,
		"resetArrayType",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchema) ResetDescription() {
	_jsii_.InvokeVoid(
		u,
		"resetDescription",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchema) ResetEnum() {
	_jsii_.InvokeVoid(
		u,
		"resetEnum",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchema) ResetExternalName() {
	_jsii_.InvokeVoid(
		u,
		"resetExternalName",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchema) ResetExternalNamespace() {
	_jsii_.InvokeVoid(
		u,
		"resetExternalNamespace",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchema) ResetId() {
	_jsii_.InvokeVoid(
		u,
		"resetId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchema) ResetMaster() {
	_jsii_.InvokeVoid(
		u,
		"resetMaster",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchema) ResetMasterOverridePriority() {
	_jsii_.InvokeVoid(
		u,
		"resetMasterOverridePriority",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchema) ResetMaxLength() {
	_jsii_.InvokeVoid(
		u,
		"resetMaxLength",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchema) ResetMinLength() {
	_jsii_.InvokeVoid(
		u,
		"resetMinLength",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchema) ResetOneOf() {
	_jsii_.InvokeVoid(
		u,
		"resetOneOf",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchema) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		u,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchema) ResetPattern() {
	_jsii_.InvokeVoid(
		u,
		"resetPattern",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchema) ResetPermissions() {
	_jsii_.InvokeVoid(
		u,
		"resetPermissions",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchema) ResetRequired() {
	_jsii_.InvokeVoid(
		u,
		"resetRequired",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchema) ResetScope() {
	_jsii_.InvokeVoid(
		u,
		"resetScope",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchema) ResetUnique() {
	_jsii_.InvokeVoid(
		u,
		"resetUnique",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchema) ResetUserType() {
	_jsii_.InvokeVoid(
		u,
		"resetUserType",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserSchema) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchema) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchema) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserSchema) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

