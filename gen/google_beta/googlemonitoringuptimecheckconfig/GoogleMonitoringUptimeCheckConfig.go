package googlemonitoringuptimecheckconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlemonitoringuptimecheckconfig/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_monitoring_uptime_check_config google_monitoring_uptime_check_config}.
type GoogleMonitoringUptimeCheckConfig interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CheckerType() *string
	SetCheckerType(val *string)
	CheckerTypeInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContentMatchers() GoogleMonitoringUptimeCheckConfigContentMatchersList
	ContentMatchersInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpCheck() GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference
	HttpCheckInput() *GoogleMonitoringUptimeCheckConfigHttpCheck
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LogCheckFailures() interface{}
	SetLogCheckFailures(val interface{})
	LogCheckFailuresInput() interface{}
	MonitoredResource() GoogleMonitoringUptimeCheckConfigMonitoredResourceOutputReference
	MonitoredResourceInput() *GoogleMonitoringUptimeCheckConfigMonitoredResource
	Name() *string
	// The tree node.
	Node() constructs.Node
	Period() *string
	SetPeriod(val *string)
	PeriodInput() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
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
	ResourceGroup() GoogleMonitoringUptimeCheckConfigResourceGroupOutputReference
	ResourceGroupInput() *GoogleMonitoringUptimeCheckConfigResourceGroup
	SelectedRegions() *[]*string
	SetSelectedRegions(val *[]*string)
	SelectedRegionsInput() *[]*string
	SyntheticMonitor() GoogleMonitoringUptimeCheckConfigSyntheticMonitorOutputReference
	SyntheticMonitorInput() *GoogleMonitoringUptimeCheckConfigSyntheticMonitor
	TcpCheck() GoogleMonitoringUptimeCheckConfigTcpCheckOutputReference
	TcpCheckInput() *GoogleMonitoringUptimeCheckConfigTcpCheck
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeout() *string
	SetTimeout(val *string)
	TimeoutInput() *string
	Timeouts() GoogleMonitoringUptimeCheckConfigTimeoutsOutputReference
	TimeoutsInput() interface{}
	UptimeCheckId() *string
	UserLabels() *map[string]*string
	SetUserLabels(val *map[string]*string)
	UserLabelsInput() *map[string]*string
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
	PutContentMatchers(value interface{})
	PutHttpCheck(value *GoogleMonitoringUptimeCheckConfigHttpCheck)
	PutMonitoredResource(value *GoogleMonitoringUptimeCheckConfigMonitoredResource)
	PutResourceGroup(value *GoogleMonitoringUptimeCheckConfigResourceGroup)
	PutSyntheticMonitor(value *GoogleMonitoringUptimeCheckConfigSyntheticMonitor)
	PutTcpCheck(value *GoogleMonitoringUptimeCheckConfigTcpCheck)
	PutTimeouts(value *GoogleMonitoringUptimeCheckConfigTimeouts)
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
	ResetCheckerType()
	ResetContentMatchers()
	ResetHttpCheck()
	ResetId()
	ResetLogCheckFailures()
	ResetMonitoredResource()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPeriod()
	ResetProject()
	ResetResourceGroup()
	ResetSelectedRegions()
	ResetSyntheticMonitor()
	ResetTcpCheck()
	ResetTimeouts()
	ResetUserLabels()
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

// The jsii proxy struct for GoogleMonitoringUptimeCheckConfig
type jsiiProxy_GoogleMonitoringUptimeCheckConfig struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) CheckerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checkerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) CheckerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checkerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ContentMatchers() GoogleMonitoringUptimeCheckConfigContentMatchersList {
	var returns GoogleMonitoringUptimeCheckConfigContentMatchersList
	_jsii_.Get(
		j,
		"contentMatchers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ContentMatchersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentMatchersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) HttpCheck() GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference {
	var returns GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference
	_jsii_.Get(
		j,
		"httpCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) HttpCheckInput() *GoogleMonitoringUptimeCheckConfigHttpCheck {
	var returns *GoogleMonitoringUptimeCheckConfigHttpCheck
	_jsii_.Get(
		j,
		"httpCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) LogCheckFailures() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logCheckFailures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) LogCheckFailuresInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logCheckFailuresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) MonitoredResource() GoogleMonitoringUptimeCheckConfigMonitoredResourceOutputReference {
	var returns GoogleMonitoringUptimeCheckConfigMonitoredResourceOutputReference
	_jsii_.Get(
		j,
		"monitoredResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) MonitoredResourceInput() *GoogleMonitoringUptimeCheckConfigMonitoredResource {
	var returns *GoogleMonitoringUptimeCheckConfigMonitoredResource
	_jsii_.Get(
		j,
		"monitoredResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Period() *string {
	var returns *string
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) PeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"periodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResourceGroup() GoogleMonitoringUptimeCheckConfigResourceGroupOutputReference {
	var returns GoogleMonitoringUptimeCheckConfigResourceGroupOutputReference
	_jsii_.Get(
		j,
		"resourceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResourceGroupInput() *GoogleMonitoringUptimeCheckConfigResourceGroup {
	var returns *GoogleMonitoringUptimeCheckConfigResourceGroup
	_jsii_.Get(
		j,
		"resourceGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) SelectedRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"selectedRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) SelectedRegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"selectedRegionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) SyntheticMonitor() GoogleMonitoringUptimeCheckConfigSyntheticMonitorOutputReference {
	var returns GoogleMonitoringUptimeCheckConfigSyntheticMonitorOutputReference
	_jsii_.Get(
		j,
		"syntheticMonitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) SyntheticMonitorInput() *GoogleMonitoringUptimeCheckConfigSyntheticMonitor {
	var returns *GoogleMonitoringUptimeCheckConfigSyntheticMonitor
	_jsii_.Get(
		j,
		"syntheticMonitorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) TcpCheck() GoogleMonitoringUptimeCheckConfigTcpCheckOutputReference {
	var returns GoogleMonitoringUptimeCheckConfigTcpCheckOutputReference
	_jsii_.Get(
		j,
		"tcpCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) TcpCheckInput() *GoogleMonitoringUptimeCheckConfigTcpCheck {
	var returns *GoogleMonitoringUptimeCheckConfigTcpCheck
	_jsii_.Get(
		j,
		"tcpCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) TimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Timeouts() GoogleMonitoringUptimeCheckConfigTimeoutsOutputReference {
	var returns GoogleMonitoringUptimeCheckConfigTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) UptimeCheckId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uptimeCheckId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) UserLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"userLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) UserLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"userLabelsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_monitoring_uptime_check_config google_monitoring_uptime_check_config} Resource.
func NewGoogleMonitoringUptimeCheckConfig(scope constructs.Construct, id *string, config *GoogleMonitoringUptimeCheckConfigConfig) GoogleMonitoringUptimeCheckConfig {
	_init_.Initialize()

	if err := validateNewGoogleMonitoringUptimeCheckConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleMonitoringUptimeCheckConfig{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleMonitoringUptimeCheckConfig.GoogleMonitoringUptimeCheckConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_monitoring_uptime_check_config google_monitoring_uptime_check_config} Resource.
func NewGoogleMonitoringUptimeCheckConfig_Override(g GoogleMonitoringUptimeCheckConfig, scope constructs.Construct, id *string, config *GoogleMonitoringUptimeCheckConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleMonitoringUptimeCheckConfig.GoogleMonitoringUptimeCheckConfig",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetCheckerType(val *string) {
	if err := j.validateSetCheckerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkerType",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetLogCheckFailures(val interface{}) {
	if err := j.validateSetLogCheckFailuresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logCheckFailures",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetPeriod(val *string) {
	if err := j.validateSetPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetSelectedRegions(val *[]*string) {
	if err := j.validateSetSelectedRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selectedRegions",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetTimeout(val *string) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetUserLabels(val *map[string]*string) {
	if err := j.validateSetUserLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userLabels",
		val,
	)
}

// Generates CDKTN code for importing a GoogleMonitoringUptimeCheckConfig resource upon running "cdktn plan <stack-name>".
func GoogleMonitoringUptimeCheckConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleMonitoringUptimeCheckConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleMonitoringUptimeCheckConfig.GoogleMonitoringUptimeCheckConfig",
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
func GoogleMonitoringUptimeCheckConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleMonitoringUptimeCheckConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleMonitoringUptimeCheckConfig.GoogleMonitoringUptimeCheckConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleMonitoringUptimeCheckConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleMonitoringUptimeCheckConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleMonitoringUptimeCheckConfig.GoogleMonitoringUptimeCheckConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleMonitoringUptimeCheckConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleMonitoringUptimeCheckConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleMonitoringUptimeCheckConfig.GoogleMonitoringUptimeCheckConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleMonitoringUptimeCheckConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google-beta.googleMonitoringUptimeCheckConfig.GoogleMonitoringUptimeCheckConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) PutContentMatchers(value interface{}) {
	if err := g.validatePutContentMatchersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putContentMatchers",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) PutHttpCheck(value *GoogleMonitoringUptimeCheckConfigHttpCheck) {
	if err := g.validatePutHttpCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHttpCheck",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) PutMonitoredResource(value *GoogleMonitoringUptimeCheckConfigMonitoredResource) {
	if err := g.validatePutMonitoredResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMonitoredResource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) PutResourceGroup(value *GoogleMonitoringUptimeCheckConfigResourceGroup) {
	if err := g.validatePutResourceGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putResourceGroup",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) PutSyntheticMonitor(value *GoogleMonitoringUptimeCheckConfigSyntheticMonitor) {
	if err := g.validatePutSyntheticMonitorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSyntheticMonitor",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) PutTcpCheck(value *GoogleMonitoringUptimeCheckConfigTcpCheck) {
	if err := g.validatePutTcpCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTcpCheck",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) PutTimeouts(value *GoogleMonitoringUptimeCheckConfigTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetCheckerType() {
	_jsii_.InvokeVoid(
		g,
		"resetCheckerType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetContentMatchers() {
	_jsii_.InvokeVoid(
		g,
		"resetContentMatchers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetHttpCheck() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpCheck",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetLogCheckFailures() {
	_jsii_.InvokeVoid(
		g,
		"resetLogCheckFailures",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetMonitoredResource() {
	_jsii_.InvokeVoid(
		g,
		"resetMonitoredResource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetPeriod() {
	_jsii_.InvokeVoid(
		g,
		"resetPeriod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetResourceGroup() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceGroup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetSelectedRegions() {
	_jsii_.InvokeVoid(
		g,
		"resetSelectedRegions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetSyntheticMonitor() {
	_jsii_.InvokeVoid(
		g,
		"resetSyntheticMonitor",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetTcpCheck() {
	_jsii_.InvokeVoid(
		g,
		"resetTcpCheck",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetUserLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetUserLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

