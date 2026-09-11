package globalacceleratorendpointgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/aws/globalacceleratorendpointgroup/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/globalaccelerator_endpoint_group aws_globalaccelerator_endpoint_group}.
type GlobalacceleratorEndpointGroup interface {
	cdktn.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
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
	EndpointConfiguration() GlobalacceleratorEndpointGroupEndpointConfigurationList
	EndpointConfigurationInput() interface{}
	EndpointGroupRegion() *string
	SetEndpointGroupRegion(val *string)
	EndpointGroupRegionInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HealthCheckIntervalSeconds() *float64
	SetHealthCheckIntervalSeconds(val *float64)
	HealthCheckIntervalSecondsInput() *float64
	HealthCheckPath() *string
	SetHealthCheckPath(val *string)
	HealthCheckPathInput() *string
	HealthCheckPort() *float64
	SetHealthCheckPort(val *float64)
	HealthCheckPortInput() *float64
	HealthCheckProtocol() *string
	SetHealthCheckProtocol(val *string)
	HealthCheckProtocolInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	ListenerArn() *string
	SetListenerArn(val *string)
	ListenerArnInput() *string
	// The tree node.
	Node() constructs.Node
	PortOverride() GlobalacceleratorEndpointGroupPortOverrideList
	PortOverrideInput() interface{}
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ThresholdCount() *float64
	SetThresholdCount(val *float64)
	ThresholdCountInput() *float64
	Timeouts() GlobalacceleratorEndpointGroupTimeoutsOutputReference
	TimeoutsInput() interface{}
	TrafficDialPercentage() *float64
	SetTrafficDialPercentage(val *float64)
	TrafficDialPercentageInput() *float64
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
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
	PutEndpointConfiguration(value interface{})
	PutPortOverride(value interface{})
	PutTimeouts(value *GlobalacceleratorEndpointGroupTimeouts)
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetEndpointConfiguration()
	ResetEndpointGroupRegion()
	ResetHealthCheckIntervalSeconds()
	ResetHealthCheckPath()
	ResetHealthCheckPort()
	ResetHealthCheckProtocol()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPortOverride()
	ResetThresholdCount()
	ResetTimeouts()
	ResetTrafficDialPercentage()
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
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for GlobalacceleratorEndpointGroup
type jsiiProxy_GlobalacceleratorEndpointGroup struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) EndpointConfiguration() GlobalacceleratorEndpointGroupEndpointConfigurationList {
	var returns GlobalacceleratorEndpointGroupEndpointConfigurationList
	_jsii_.Get(
		j,
		"endpointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) EndpointConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) EndpointGroupRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointGroupRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) EndpointGroupRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointGroupRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) HealthCheckIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) HealthCheckIntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckIntervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) HealthCheckPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) HealthCheckPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) HealthCheckPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) HealthCheckPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) HealthCheckProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) HealthCheckProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) ListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) ListenerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) PortOverride() GlobalacceleratorEndpointGroupPortOverrideList {
	var returns GlobalacceleratorEndpointGroupPortOverrideList
	_jsii_.Get(
		j,
		"portOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) PortOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) ThresholdCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) ThresholdCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) Timeouts() GlobalacceleratorEndpointGroupTimeoutsOutputReference {
	var returns GlobalacceleratorEndpointGroupTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) TrafficDialPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"trafficDialPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup) TrafficDialPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"trafficDialPercentageInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/globalaccelerator_endpoint_group aws_globalaccelerator_endpoint_group} Resource.
func NewGlobalacceleratorEndpointGroup(scope constructs.Construct, id *string, config *GlobalacceleratorEndpointGroupConfig) GlobalacceleratorEndpointGroup {
	_init_.Initialize()

	if err := validateNewGlobalacceleratorEndpointGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlobalacceleratorEndpointGroup{}

	_jsii_.Create(
		"@cdktn/provider-aws.globalacceleratorEndpointGroup.GlobalacceleratorEndpointGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/globalaccelerator_endpoint_group aws_globalaccelerator_endpoint_group} Resource.
func NewGlobalacceleratorEndpointGroup_Override(g GlobalacceleratorEndpointGroup, scope constructs.Construct, id *string, config *GlobalacceleratorEndpointGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.globalacceleratorEndpointGroup.GlobalacceleratorEndpointGroup",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetEndpointGroupRegion(val *string) {
	if err := j.validateSetEndpointGroupRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointGroupRegion",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetHealthCheckIntervalSeconds(val *float64) {
	if err := j.validateSetHealthCheckIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetHealthCheckPath(val *string) {
	if err := j.validateSetHealthCheckPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckPath",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetHealthCheckPort(val *float64) {
	if err := j.validateSetHealthCheckPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckPort",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetHealthCheckProtocol(val *string) {
	if err := j.validateSetHealthCheckProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckProtocol",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetListenerArn(val *string) {
	if err := j.validateSetListenerArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listenerArn",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetThresholdCount(val *float64) {
	if err := j.validateSetThresholdCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdCount",
		val,
	)
}

func (j *jsiiProxy_GlobalacceleratorEndpointGroup)SetTrafficDialPercentage(val *float64) {
	if err := j.validateSetTrafficDialPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trafficDialPercentage",
		val,
	)
}

// Generates CDKTN code for importing a GlobalacceleratorEndpointGroup resource upon running "cdktn plan <stack-name>".
func GlobalacceleratorEndpointGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGlobalacceleratorEndpointGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.globalacceleratorEndpointGroup.GlobalacceleratorEndpointGroup",
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
func GlobalacceleratorEndpointGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGlobalacceleratorEndpointGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.globalacceleratorEndpointGroup.GlobalacceleratorEndpointGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GlobalacceleratorEndpointGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGlobalacceleratorEndpointGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.globalacceleratorEndpointGroup.GlobalacceleratorEndpointGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GlobalacceleratorEndpointGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGlobalacceleratorEndpointGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.globalacceleratorEndpointGroup.GlobalacceleratorEndpointGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GlobalacceleratorEndpointGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-aws.globalacceleratorEndpointGroup.GlobalacceleratorEndpointGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := g.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) PutEndpointConfiguration(value interface{}) {
	if err := g.validatePutEndpointConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEndpointConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) PutPortOverride(value interface{}) {
	if err := g.validatePutPortOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPortOverride",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) PutTimeouts(value *GlobalacceleratorEndpointGroupTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetEndpointConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetEndpointConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetEndpointGroupRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetEndpointGroupRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetHealthCheckIntervalSeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetHealthCheckIntervalSeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetHealthCheckPath() {
	_jsii_.InvokeVoid(
		g,
		"resetHealthCheckPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetHealthCheckPort() {
	_jsii_.InvokeVoid(
		g,
		"resetHealthCheckPort",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetHealthCheckProtocol() {
	_jsii_.InvokeVoid(
		g,
		"resetHealthCheckProtocol",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetPortOverride() {
	_jsii_.InvokeVoid(
		g,
		"resetPortOverride",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetThresholdCount() {
	_jsii_.InvokeVoid(
		g,
		"resetThresholdCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ResetTrafficDialPercentage() {
	_jsii_.InvokeVoid(
		g,
		"resetTrafficDialPercentage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlobalacceleratorEndpointGroup) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		g,
		"with",
		args,
		&returns,
	)

	return returns
}

