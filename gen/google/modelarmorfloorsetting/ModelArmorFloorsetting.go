package modelarmorfloorsetting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/modelarmorfloorsetting/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/model_armor_floorsetting google_model_armor_floorsetting}.
type ModelArmorFloorsetting interface {
	cdktf.TerraformResource
	AiPlatformFloorSetting() ModelArmorFloorsettingAiPlatformFloorSettingOutputReference
	AiPlatformFloorSettingInput() *ModelArmorFloorsettingAiPlatformFloorSetting
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableFloorSettingEnforcement() any
	SetEnableFloorSettingEnforcement(val any)
	EnableFloorSettingEnforcementInput() any
	FilterConfig() ModelArmorFloorsettingFilterConfigOutputReference
	FilterConfigInput() *ModelArmorFloorsettingFilterConfig
	FloorSettingMetadata() ModelArmorFloorsettingFloorSettingMetadataOutputReference
	FloorSettingMetadataInput() *ModelArmorFloorsettingFloorSettingMetadata
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GoogleMcpServerFloorSetting() ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference
	GoogleMcpServerFloorSettingInput() *ModelArmorFloorsettingGoogleMcpServerFloorSetting
	Id() *string
	SetId(val *string)
	IdInput() *string
	IntegratedServices() *[]*string
	SetIntegratedServices(val *[]*string)
	IntegratedServicesInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	Parent() *string
	SetParent(val *string)
	ParentInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	// Experimental.
	RawOverrides() any
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ModelArmorFloorsettingTimeoutsOutputReference
	TimeoutsInput() any
	UpdateTime() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value any)
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	HasResourceMove() any
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index any)
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAiPlatformFloorSetting(value *ModelArmorFloorsettingAiPlatformFloorSetting)
	PutFilterConfig(value *ModelArmorFloorsettingFilterConfig)
	PutFloorSettingMetadata(value *ModelArmorFloorsettingFloorSettingMetadata)
	PutGoogleMcpServerFloorSetting(value *ModelArmorFloorsettingGoogleMcpServerFloorSetting)
	PutTimeouts(value *ModelArmorFloorsettingTimeouts)
	ResetAiPlatformFloorSetting()
	ResetEnableFloorSettingEnforcement()
	ResetFloorSettingMetadata()
	ResetGoogleMcpServerFloorSetting()
	ResetId()
	ResetIntegratedServices()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]any
	SynthesizeHclAttributes() *map[string]any
	// Experimental.
	ToHclTerraform() any
	// Experimental.
	ToMetadata() any
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() any
}

// The jsii proxy struct for ModelArmorFloorsetting
type jsiiProxy_ModelArmorFloorsetting struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ModelArmorFloorsetting) AiPlatformFloorSetting() ModelArmorFloorsettingAiPlatformFloorSettingOutputReference {
	var returns ModelArmorFloorsettingAiPlatformFloorSettingOutputReference
	_jsii_.Get(
		j,
		"aiPlatformFloorSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) AiPlatformFloorSettingInput() *ModelArmorFloorsettingAiPlatformFloorSetting {
	var returns *ModelArmorFloorsettingAiPlatformFloorSetting
	_jsii_.Get(
		j,
		"aiPlatformFloorSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) EnableFloorSettingEnforcement() any {
	var returns any
	_jsii_.Get(
		j,
		"enableFloorSettingEnforcement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) EnableFloorSettingEnforcementInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableFloorSettingEnforcementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) FilterConfig() ModelArmorFloorsettingFilterConfigOutputReference {
	var returns ModelArmorFloorsettingFilterConfigOutputReference
	_jsii_.Get(
		j,
		"filterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) FilterConfigInput() *ModelArmorFloorsettingFilterConfig {
	var returns *ModelArmorFloorsettingFilterConfig
	_jsii_.Get(
		j,
		"filterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) FloorSettingMetadata() ModelArmorFloorsettingFloorSettingMetadataOutputReference {
	var returns ModelArmorFloorsettingFloorSettingMetadataOutputReference
	_jsii_.Get(
		j,
		"floorSettingMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) FloorSettingMetadataInput() *ModelArmorFloorsettingFloorSettingMetadata {
	var returns *ModelArmorFloorsettingFloorSettingMetadata
	_jsii_.Get(
		j,
		"floorSettingMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) GoogleMcpServerFloorSetting() ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference {
	var returns ModelArmorFloorsettingGoogleMcpServerFloorSettingOutputReference
	_jsii_.Get(
		j,
		"googleMcpServerFloorSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) GoogleMcpServerFloorSettingInput() *ModelArmorFloorsettingGoogleMcpServerFloorSetting {
	var returns *ModelArmorFloorsettingGoogleMcpServerFloorSetting
	_jsii_.Get(
		j,
		"googleMcpServerFloorSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) IntegratedServices() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"integratedServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) IntegratedServicesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"integratedServicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) Parent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) ParentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) Timeouts() ModelArmorFloorsettingTimeoutsOutputReference {
	var returns ModelArmorFloorsettingTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsetting) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/model_armor_floorsetting google_model_armor_floorsetting} Resource.
func NewModelArmorFloorsetting(scope constructs.Construct, id *string, config *ModelArmorFloorsettingConfig) ModelArmorFloorsetting {
	_init_.Initialize()

	if err := validateNewModelArmorFloorsettingParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ModelArmorFloorsetting{}

	_jsii_.Create(
		"@cdktf/provider-google.modelArmorFloorsetting.ModelArmorFloorsetting",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/model_armor_floorsetting google_model_armor_floorsetting} Resource.
func NewModelArmorFloorsetting_Override(m ModelArmorFloorsetting, scope constructs.Construct, id *string, config *ModelArmorFloorsettingConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.modelArmorFloorsetting.ModelArmorFloorsetting",
		[]any{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_ModelArmorFloorsetting) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsetting) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsetting) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsetting) SetEnableFloorSettingEnforcement(val any) {
	if err := j.validateSetEnableFloorSettingEnforcementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableFloorSettingEnforcement",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsetting) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsetting) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsetting) SetIntegratedServices(val *[]*string) {
	if err := j.validateSetIntegratedServicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integratedServices",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsetting) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsetting) SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsetting) SetParent(val *string) {
	if err := j.validateSetParentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parent",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsetting) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsetting) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a ModelArmorFloorsetting resource upon running "cdktf plan <stack-name>".
func ModelArmorFloorsetting_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateModelArmorFloorsetting_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.modelArmorFloorsetting.ModelArmorFloorsetting",
		"generateConfigForImport",
		[]any{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func ModelArmorFloorsetting_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateModelArmorFloorsetting_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.modelArmorFloorsetting.ModelArmorFloorsetting",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func ModelArmorFloorsetting_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateModelArmorFloorsetting_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.modelArmorFloorsetting.ModelArmorFloorsetting",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func ModelArmorFloorsetting_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateModelArmorFloorsetting_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.modelArmorFloorsetting.ModelArmorFloorsetting",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func ModelArmorFloorsetting_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.modelArmorFloorsetting.ModelArmorFloorsetting",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_ModelArmorFloorsetting) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (m *jsiiProxy_ModelArmorFloorsetting) AddOverride(path *string, value any) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]any{path, value},
	)
}

func (m *jsiiProxy_ModelArmorFloorsetting) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsetting) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsetting) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsetting) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsetting) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsetting) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsetting) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsetting) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsetting) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsetting) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsetting) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]any{id, provider},
	)
}

func (m *jsiiProxy_ModelArmorFloorsetting) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsetting) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]any{id},
	)
}

func (m *jsiiProxy_ModelArmorFloorsetting) MoveTo(moveTarget *string, index any) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (m *jsiiProxy_ModelArmorFloorsetting) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]any{id},
	)
}

func (m *jsiiProxy_ModelArmorFloorsetting) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (m *jsiiProxy_ModelArmorFloorsetting) PutAiPlatformFloorSetting(value *ModelArmorFloorsettingAiPlatformFloorSetting) {
	if err := m.validatePutAiPlatformFloorSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAiPlatformFloorSetting",
		[]any{value},
	)
}

func (m *jsiiProxy_ModelArmorFloorsetting) PutFilterConfig(value *ModelArmorFloorsettingFilterConfig) {
	if err := m.validatePutFilterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFilterConfig",
		[]any{value},
	)
}

func (m *jsiiProxy_ModelArmorFloorsetting) PutFloorSettingMetadata(value *ModelArmorFloorsettingFloorSettingMetadata) {
	if err := m.validatePutFloorSettingMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFloorSettingMetadata",
		[]any{value},
	)
}

func (m *jsiiProxy_ModelArmorFloorsetting) PutGoogleMcpServerFloorSetting(value *ModelArmorFloorsettingGoogleMcpServerFloorSetting) {
	if err := m.validatePutGoogleMcpServerFloorSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putGoogleMcpServerFloorSetting",
		[]any{value},
	)
}

func (m *jsiiProxy_ModelArmorFloorsetting) PutTimeouts(value *ModelArmorFloorsettingTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]any{value},
	)
}

func (m *jsiiProxy_ModelArmorFloorsetting) ResetAiPlatformFloorSetting() {
	_jsii_.InvokeVoid(
		m,
		"resetAiPlatformFloorSetting",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorFloorsetting) ResetEnableFloorSettingEnforcement() {
	_jsii_.InvokeVoid(
		m,
		"resetEnableFloorSettingEnforcement",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorFloorsetting) ResetFloorSettingMetadata() {
	_jsii_.InvokeVoid(
		m,
		"resetFloorSettingMetadata",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorFloorsetting) ResetGoogleMcpServerFloorSetting() {
	_jsii_.InvokeVoid(
		m,
		"resetGoogleMcpServerFloorSetting",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorFloorsetting) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorFloorsetting) ResetIntegratedServices() {
	_jsii_.InvokeVoid(
		m,
		"resetIntegratedServices",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorFloorsetting) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorFloorsetting) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorFloorsetting) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsetting) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsetting) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsetting) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsetting) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsetting) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
