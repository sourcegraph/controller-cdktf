package googlecomputeperinstanceconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeperinstanceconfig/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_compute_per_instance_config google_compute_per_instance_config}.
type GoogleComputePerInstanceConfig interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
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
	SetId(val *string)
	IdInput() *string
	InstanceGroupManager() *string
	SetInstanceGroupManager(val *string)
	InstanceGroupManagerInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MinimalAction() *string
	SetMinimalAction(val *string)
	MinimalActionInput() *string
	MostDisruptiveAllowedAction() *string
	SetMostDisruptiveAllowedAction(val *string)
	MostDisruptiveAllowedActionInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PreservedState() GoogleComputePerInstanceConfigPreservedStateOutputReference
	PreservedStateInput() *GoogleComputePerInstanceConfigPreservedState
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
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
	RemoveInstanceOnDestroy() interface{}
	SetRemoveInstanceOnDestroy(val interface{})
	RemoveInstanceOnDestroyInput() interface{}
	RemoveInstanceStateOnDestroy() interface{}
	SetRemoveInstanceStateOnDestroy(val interface{})
	RemoveInstanceStateOnDestroyInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleComputePerInstanceConfigTimeoutsOutputReference
	TimeoutsInput() interface{}
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
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
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutPreservedState(value *GoogleComputePerInstanceConfigPreservedState)
	PutTimeouts(value *GoogleComputePerInstanceConfigTimeouts)
	ResetId()
	ResetMinimalAction()
	ResetMostDisruptiveAllowedAction()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreservedState()
	ResetProject()
	ResetRemoveInstanceOnDestroy()
	ResetRemoveInstanceStateOnDestroy()
	ResetTimeouts()
	ResetZone()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleComputePerInstanceConfig
type jsiiProxy_GoogleComputePerInstanceConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) InstanceGroupManager() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceGroupManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) InstanceGroupManagerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceGroupManagerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) MinimalAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimalAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) MinimalActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimalActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) MostDisruptiveAllowedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mostDisruptiveAllowedAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) MostDisruptiveAllowedActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mostDisruptiveAllowedActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) PreservedState() GoogleComputePerInstanceConfigPreservedStateOutputReference {
	var returns GoogleComputePerInstanceConfigPreservedStateOutputReference
	_jsii_.Get(
		j,
		"preservedState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) PreservedStateInput() *GoogleComputePerInstanceConfigPreservedState {
	var returns *GoogleComputePerInstanceConfigPreservedState
	_jsii_.Get(
		j,
		"preservedStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) RemoveInstanceOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeInstanceOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) RemoveInstanceOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeInstanceOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) RemoveInstanceStateOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeInstanceStateOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) RemoveInstanceStateOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeInstanceStateOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) Timeouts() GoogleComputePerInstanceConfigTimeoutsOutputReference {
	var returns GoogleComputePerInstanceConfigTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_compute_per_instance_config google_compute_per_instance_config} Resource.
func NewGoogleComputePerInstanceConfig(scope constructs.Construct, id *string, config *GoogleComputePerInstanceConfigConfig) GoogleComputePerInstanceConfig {
	_init_.Initialize()

	if err := validateNewGoogleComputePerInstanceConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputePerInstanceConfig{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputePerInstanceConfig.GoogleComputePerInstanceConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_compute_per_instance_config google_compute_per_instance_config} Resource.
func NewGoogleComputePerInstanceConfig_Override(g GoogleComputePerInstanceConfig, scope constructs.Construct, id *string, config *GoogleComputePerInstanceConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputePerInstanceConfig.GoogleComputePerInstanceConfig",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig)SetInstanceGroupManager(val *string) {
	if err := j.validateSetInstanceGroupManagerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceGroupManager",
		val,
	)
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig)SetMinimalAction(val *string) {
	if err := j.validateSetMinimalActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimalAction",
		val,
	)
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig)SetMostDisruptiveAllowedAction(val *string) {
	if err := j.validateSetMostDisruptiveAllowedActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mostDisruptiveAllowedAction",
		val,
	)
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig)SetRemoveInstanceOnDestroy(val interface{}) {
	if err := j.validateSetRemoveInstanceOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"removeInstanceOnDestroy",
		val,
	)
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig)SetRemoveInstanceStateOnDestroy(val interface{}) {
	if err := j.validateSetRemoveInstanceStateOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"removeInstanceStateOnDestroy",
		val,
	)
}

func (j *jsiiProxy_GoogleComputePerInstanceConfig)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

// Generates CDKTF code for importing a GoogleComputePerInstanceConfig resource upon running "cdktf plan <stack-name>".
func GoogleComputePerInstanceConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleComputePerInstanceConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputePerInstanceConfig.GoogleComputePerInstanceConfig",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
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
func GoogleComputePerInstanceConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputePerInstanceConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputePerInstanceConfig.GoogleComputePerInstanceConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputePerInstanceConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputePerInstanceConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputePerInstanceConfig.GoogleComputePerInstanceConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputePerInstanceConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputePerInstanceConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputePerInstanceConfig.GoogleComputePerInstanceConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputePerInstanceConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleComputePerInstanceConfig.GoogleComputePerInstanceConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputePerInstanceConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputePerInstanceConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputePerInstanceConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputePerInstanceConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputePerInstanceConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputePerInstanceConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputePerInstanceConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputePerInstanceConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputePerInstanceConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputePerInstanceConfig) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) PutPreservedState(value *GoogleComputePerInstanceConfigPreservedState) {
	if err := g.validatePutPreservedStateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPreservedState",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) PutTimeouts(value *GoogleComputePerInstanceConfigTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) ResetMinimalAction() {
	_jsii_.InvokeVoid(
		g,
		"resetMinimalAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) ResetMostDisruptiveAllowedAction() {
	_jsii_.InvokeVoid(
		g,
		"resetMostDisruptiveAllowedAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) ResetPreservedState() {
	_jsii_.InvokeVoid(
		g,
		"resetPreservedState",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) ResetRemoveInstanceOnDestroy() {
	_jsii_.InvokeVoid(
		g,
		"resetRemoveInstanceOnDestroy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) ResetRemoveInstanceStateOnDestroy() {
	_jsii_.InvokeVoid(
		g,
		"resetRemoveInstanceStateOnDestroy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) ResetZone() {
	_jsii_.InvokeVoid(
		g,
		"resetZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputePerInstanceConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

