package groupschemaproperty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/groupschemaproperty/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/okta/r/group_schema_property okta_group_schema_property}.
type GroupSchemaProperty interface {
	cdktf.TerraformResource
	ArrayEnum() *[]*string
	SetArrayEnum(val *[]*string)
	ArrayEnumInput() *[]*string
	ArrayOneOf() GroupSchemaPropertyArrayOneOfList
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
	MasterOverridePriority() GroupSchemaPropertyMasterOverridePriorityList
	MasterOverridePriorityInput() interface{}
	MaxLength() *float64
	SetMaxLength(val *float64)
	MaxLengthInput() *float64
	MinLength() *float64
	SetMinLength(val *float64)
	MinLengthInput() *float64
	// The tree node.
	Node() constructs.Node
	OneOf() GroupSchemaPropertyOneOfList
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
	Unique() *string
	SetUnique(val *string)
	UniqueInput() *string
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
	ResetPermissions()
	ResetRequired()
	ResetScope()
	ResetUnique()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GroupSchemaProperty
type jsiiProxy_GroupSchemaProperty struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GroupSchemaProperty) ArrayEnum() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"arrayEnum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) ArrayEnumInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"arrayEnumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) ArrayOneOf() GroupSchemaPropertyArrayOneOfList {
	var returns GroupSchemaPropertyArrayOneOfList
	_jsii_.Get(
		j,
		"arrayOneOf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) ArrayOneOfInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"arrayOneOfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) ArrayType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arrayType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) ArrayTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arrayTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) Enum() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) EnumInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) ExternalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) ExternalNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) ExternalNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) ExternalNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) Index() *string {
	var returns *string
	_jsii_.Get(
		j,
		"index",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) IndexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) Master() *string {
	var returns *string
	_jsii_.Get(
		j,
		"master",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) MasterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) MasterOverridePriority() GroupSchemaPropertyMasterOverridePriorityList {
	var returns GroupSchemaPropertyMasterOverridePriorityList
	_jsii_.Get(
		j,
		"masterOverridePriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) MasterOverridePriorityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"masterOverridePriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) MaxLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) MaxLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) MinLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) MinLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) OneOf() GroupSchemaPropertyOneOfList {
	var returns GroupSchemaPropertyOneOfList
	_jsii_.Get(
		j,
		"oneOf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) OneOfInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oneOfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) Permissions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) PermissionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) Required() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"required",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) RequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) Unique() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unique",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSchemaProperty) UniqueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniqueInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/okta/r/group_schema_property okta_group_schema_property} Resource.
func NewGroupSchemaProperty(scope constructs.Construct, id *string, config *GroupSchemaPropertyConfig) GroupSchemaProperty {
	_init_.Initialize()

	if err := validateNewGroupSchemaPropertyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GroupSchemaProperty{}

	_jsii_.Create(
		"@cdktf/provider-okta.groupSchemaProperty.GroupSchemaProperty",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/okta/r/group_schema_property okta_group_schema_property} Resource.
func NewGroupSchemaProperty_Override(g GroupSchemaProperty, scope constructs.Construct, id *string, config *GroupSchemaPropertyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.groupSchemaProperty.GroupSchemaProperty",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GroupSchemaProperty)SetArrayEnum(val *[]*string) {
	if err := j.validateSetArrayEnumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arrayEnum",
		val,
	)
}

func (j *jsiiProxy_GroupSchemaProperty)SetArrayType(val *string) {
	if err := j.validateSetArrayTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arrayType",
		val,
	)
}

func (j *jsiiProxy_GroupSchemaProperty)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GroupSchemaProperty)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GroupSchemaProperty)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GroupSchemaProperty)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GroupSchemaProperty)SetEnum(val *[]*string) {
	if err := j.validateSetEnumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enum",
		val,
	)
}

func (j *jsiiProxy_GroupSchemaProperty)SetExternalName(val *string) {
	if err := j.validateSetExternalNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalName",
		val,
	)
}

func (j *jsiiProxy_GroupSchemaProperty)SetExternalNamespace(val *string) {
	if err := j.validateSetExternalNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalNamespace",
		val,
	)
}

func (j *jsiiProxy_GroupSchemaProperty)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GroupSchemaProperty)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GroupSchemaProperty)SetIndex(val *string) {
	if err := j.validateSetIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"index",
		val,
	)
}

func (j *jsiiProxy_GroupSchemaProperty)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GroupSchemaProperty)SetMaster(val *string) {
	if err := j.validateSetMasterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"master",
		val,
	)
}

func (j *jsiiProxy_GroupSchemaProperty)SetMaxLength(val *float64) {
	if err := j.validateSetMaxLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxLength",
		val,
	)
}

func (j *jsiiProxy_GroupSchemaProperty)SetMinLength(val *float64) {
	if err := j.validateSetMinLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minLength",
		val,
	)
}

func (j *jsiiProxy_GroupSchemaProperty)SetPermissions(val *string) {
	if err := j.validateSetPermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissions",
		val,
	)
}

func (j *jsiiProxy_GroupSchemaProperty)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GroupSchemaProperty)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GroupSchemaProperty)SetRequired(val interface{}) {
	if err := j.validateSetRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"required",
		val,
	)
}

func (j *jsiiProxy_GroupSchemaProperty)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_GroupSchemaProperty)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_GroupSchemaProperty)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_GroupSchemaProperty)SetUnique(val *string) {
	if err := j.validateSetUniqueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unique",
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
func GroupSchemaProperty_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGroupSchemaProperty_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.groupSchemaProperty.GroupSchemaProperty",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GroupSchemaProperty_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGroupSchemaProperty_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.groupSchemaProperty.GroupSchemaProperty",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GroupSchemaProperty_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGroupSchemaProperty_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.groupSchemaProperty.GroupSchemaProperty",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GroupSchemaProperty_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.groupSchemaProperty.GroupSchemaProperty",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GroupSchemaProperty) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GroupSchemaProperty) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupSchemaProperty) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupSchemaProperty) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupSchemaProperty) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupSchemaProperty) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupSchemaProperty) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupSchemaProperty) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupSchemaProperty) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupSchemaProperty) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupSchemaProperty) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupSchemaProperty) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GroupSchemaProperty) PutArrayOneOf(value interface{}) {
	if err := g.validatePutArrayOneOfParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putArrayOneOf",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroupSchemaProperty) PutMasterOverridePriority(value interface{}) {
	if err := g.validatePutMasterOverridePriorityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMasterOverridePriority",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroupSchemaProperty) PutOneOf(value interface{}) {
	if err := g.validatePutOneOfParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOneOf",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroupSchemaProperty) ResetArrayEnum() {
	_jsii_.InvokeVoid(
		g,
		"resetArrayEnum",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSchemaProperty) ResetArrayOneOf() {
	_jsii_.InvokeVoid(
		g,
		"resetArrayOneOf",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSchemaProperty) ResetArrayType() {
	_jsii_.InvokeVoid(
		g,
		"resetArrayType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSchemaProperty) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSchemaProperty) ResetEnum() {
	_jsii_.InvokeVoid(
		g,
		"resetEnum",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSchemaProperty) ResetExternalName() {
	_jsii_.InvokeVoid(
		g,
		"resetExternalName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSchemaProperty) ResetExternalNamespace() {
	_jsii_.InvokeVoid(
		g,
		"resetExternalNamespace",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSchemaProperty) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSchemaProperty) ResetMaster() {
	_jsii_.InvokeVoid(
		g,
		"resetMaster",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSchemaProperty) ResetMasterOverridePriority() {
	_jsii_.InvokeVoid(
		g,
		"resetMasterOverridePriority",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSchemaProperty) ResetMaxLength() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxLength",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSchemaProperty) ResetMinLength() {
	_jsii_.InvokeVoid(
		g,
		"resetMinLength",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSchemaProperty) ResetOneOf() {
	_jsii_.InvokeVoid(
		g,
		"resetOneOf",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSchemaProperty) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSchemaProperty) ResetPermissions() {
	_jsii_.InvokeVoid(
		g,
		"resetPermissions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSchemaProperty) ResetRequired() {
	_jsii_.InvokeVoid(
		g,
		"resetRequired",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSchemaProperty) ResetScope() {
	_jsii_.InvokeVoid(
		g,
		"resetScope",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSchemaProperty) ResetUnique() {
	_jsii_.InvokeVoid(
		g,
		"resetUnique",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSchemaProperty) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupSchemaProperty) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupSchemaProperty) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupSchemaProperty) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

