package medialivemultiplexprogram

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/medialivemultiplexprogram/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/medialive_multiplex_program aws_medialive_multiplex_program}.
type MedialiveMultiplexProgram interface {
	cdktf.TerraformResource
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MultiplexId() *string
	SetMultiplexId(val *string)
	MultiplexIdInput() *string
	MultiplexProgramSettings() MedialiveMultiplexProgramMultiplexProgramSettingsList
	MultiplexProgramSettingsInput() any
	// The tree node.
	Node() constructs.Node
	ProgramName() *string
	SetProgramName(val *string)
	ProgramNameInput() *string
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
	PutMultiplexProgramSettings(value any)
	ResetMultiplexProgramSettings()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for MedialiveMultiplexProgram
type jsiiProxy_MedialiveMultiplexProgram struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MedialiveMultiplexProgram) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgram) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgram) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgram) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgram) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgram) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgram) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgram) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgram) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgram) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgram) MultiplexId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multiplexId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgram) MultiplexIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multiplexIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgram) MultiplexProgramSettings() MedialiveMultiplexProgramMultiplexProgramSettingsList {
	var returns MedialiveMultiplexProgramMultiplexProgramSettingsList
	_jsii_.Get(
		j,
		"multiplexProgramSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgram) MultiplexProgramSettingsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"multiplexProgramSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgram) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgram) ProgramName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"programName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgram) ProgramNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"programNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgram) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgram) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgram) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgram) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgram) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveMultiplexProgram) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/medialive_multiplex_program aws_medialive_multiplex_program} Resource.
func NewMedialiveMultiplexProgram(scope constructs.Construct, id *string, config *MedialiveMultiplexProgramConfig) MedialiveMultiplexProgram {
	_init_.Initialize()

	if err := validateNewMedialiveMultiplexProgramParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveMultiplexProgram{}

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveMultiplexProgram.MedialiveMultiplexProgram",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/medialive_multiplex_program aws_medialive_multiplex_program} Resource.
func NewMedialiveMultiplexProgram_Override(m MedialiveMultiplexProgram, scope constructs.Construct, id *string, config *MedialiveMultiplexProgramConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveMultiplexProgram.MedialiveMultiplexProgram",
		[]any{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MedialiveMultiplexProgram) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexProgram) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexProgram) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexProgram) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexProgram) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexProgram) SetMultiplexId(val *string) {
	if err := j.validateSetMultiplexIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiplexId",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexProgram) SetProgramName(val *string) {
	if err := j.validateSetProgramNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"programName",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexProgram) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MedialiveMultiplexProgram) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a MedialiveMultiplexProgram resource upon running "cdktf plan <stack-name>".
func MedialiveMultiplexProgram_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMedialiveMultiplexProgram_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.medialiveMultiplexProgram.MedialiveMultiplexProgram",
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
func MedialiveMultiplexProgram_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateMedialiveMultiplexProgram_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.medialiveMultiplexProgram.MedialiveMultiplexProgram",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func MedialiveMultiplexProgram_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateMedialiveMultiplexProgram_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.medialiveMultiplexProgram.MedialiveMultiplexProgram",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func MedialiveMultiplexProgram_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateMedialiveMultiplexProgram_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.medialiveMultiplexProgram.MedialiveMultiplexProgram",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func MedialiveMultiplexProgram_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.medialiveMultiplexProgram.MedialiveMultiplexProgram",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MedialiveMultiplexProgram) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (m *jsiiProxy_MedialiveMultiplexProgram) AddOverride(path *string, value any) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]any{path, value},
	)
}

func (m *jsiiProxy_MedialiveMultiplexProgram) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (m *jsiiProxy_MedialiveMultiplexProgram) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveMultiplexProgram) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MedialiveMultiplexProgram) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MedialiveMultiplexProgram) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MedialiveMultiplexProgram) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MedialiveMultiplexProgram) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MedialiveMultiplexProgram) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MedialiveMultiplexProgram) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MedialiveMultiplexProgram) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexProgram) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]any{id, provider},
	)
}

func (m *jsiiProxy_MedialiveMultiplexProgram) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveMultiplexProgram) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]any{id},
	)
}

func (m *jsiiProxy_MedialiveMultiplexProgram) MoveTo(moveTarget *string, index any) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (m *jsiiProxy_MedialiveMultiplexProgram) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]any{id},
	)
}

func (m *jsiiProxy_MedialiveMultiplexProgram) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (m *jsiiProxy_MedialiveMultiplexProgram) PutMultiplexProgramSettings(value any) {
	if err := m.validatePutMultiplexProgramSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMultiplexProgramSettings",
		[]any{value},
	)
}

func (m *jsiiProxy_MedialiveMultiplexProgram) ResetMultiplexProgramSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetMultiplexProgramSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexProgram) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveMultiplexProgram) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexProgram) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexProgram) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexProgram) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexProgram) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveMultiplexProgram) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
