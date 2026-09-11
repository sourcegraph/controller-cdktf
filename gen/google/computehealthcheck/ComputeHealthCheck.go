package computehealthcheck

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/computehealthcheck/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_health_check google_compute_health_check}.
type ComputeHealthCheck interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CheckIntervalSec() *float64
	SetCheckIntervalSec(val *float64)
	CheckIntervalSecInput() *float64
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
	CreationTimestamp() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GrpcHealthCheck() ComputeHealthCheckGrpcHealthCheckOutputReference
	GrpcHealthCheckInput() *ComputeHealthCheckGrpcHealthCheck
	GrpcTlsHealthCheck() ComputeHealthCheckGrpcTlsHealthCheckOutputReference
	GrpcTlsHealthCheckInput() *ComputeHealthCheckGrpcTlsHealthCheck
	HealthyThreshold() *float64
	SetHealthyThreshold(val *float64)
	HealthyThresholdInput() *float64
	Http2HealthCheck() ComputeHealthCheckHttp2HealthCheckOutputReference
	Http2HealthCheckInput() *ComputeHealthCheckHttp2HealthCheck
	HttpHealthCheck() ComputeHealthCheckHttpHealthCheckOutputReference
	HttpHealthCheckInput() *ComputeHealthCheckHttpHealthCheck
	HttpsHealthCheck() ComputeHealthCheckHttpsHealthCheckOutputReference
	HttpsHealthCheckInput() *ComputeHealthCheckHttpsHealthCheck
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LogConfig() ComputeHealthCheckLogConfigOutputReference
	LogConfigInput() *ComputeHealthCheckLogConfig
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
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
	SelfLink() *string
	SourceRegions() *[]*string
	SetSourceRegions(val *[]*string)
	SourceRegionsInput() *[]*string
	SslHealthCheck() ComputeHealthCheckSslHealthCheckOutputReference
	SslHealthCheckInput() *ComputeHealthCheckSslHealthCheck
	TcpHealthCheck() ComputeHealthCheckTcpHealthCheckOutputReference
	TcpHealthCheckInput() *ComputeHealthCheckTcpHealthCheck
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ComputeHealthCheckTimeoutsOutputReference
	TimeoutSec() *float64
	SetTimeoutSec(val *float64)
	TimeoutSecInput() *float64
	TimeoutsInput() interface{}
	Type() *string
	UnhealthyThreshold() *float64
	SetUnhealthyThreshold(val *float64)
	UnhealthyThresholdInput() *float64
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
	PutGrpcHealthCheck(value *ComputeHealthCheckGrpcHealthCheck)
	PutGrpcTlsHealthCheck(value *ComputeHealthCheckGrpcTlsHealthCheck)
	PutHttp2HealthCheck(value *ComputeHealthCheckHttp2HealthCheck)
	PutHttpHealthCheck(value *ComputeHealthCheckHttpHealthCheck)
	PutHttpsHealthCheck(value *ComputeHealthCheckHttpsHealthCheck)
	PutLogConfig(value *ComputeHealthCheckLogConfig)
	PutSslHealthCheck(value *ComputeHealthCheckSslHealthCheck)
	PutTcpHealthCheck(value *ComputeHealthCheckTcpHealthCheck)
	PutTimeouts(value *ComputeHealthCheckTimeouts)
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
	ResetCheckIntervalSec()
	ResetDescription()
	ResetGrpcHealthCheck()
	ResetGrpcTlsHealthCheck()
	ResetHealthyThreshold()
	ResetHttp2HealthCheck()
	ResetHttpHealthCheck()
	ResetHttpsHealthCheck()
	ResetId()
	ResetLogConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetSourceRegions()
	ResetSslHealthCheck()
	ResetTcpHealthCheck()
	ResetTimeouts()
	ResetTimeoutSec()
	ResetUnhealthyThreshold()
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

// The jsii proxy struct for ComputeHealthCheck
type jsiiProxy_ComputeHealthCheck struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_ComputeHealthCheck) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) CheckIntervalSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"checkIntervalSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) CheckIntervalSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"checkIntervalSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) GrpcHealthCheck() ComputeHealthCheckGrpcHealthCheckOutputReference {
	var returns ComputeHealthCheckGrpcHealthCheckOutputReference
	_jsii_.Get(
		j,
		"grpcHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) GrpcHealthCheckInput() *ComputeHealthCheckGrpcHealthCheck {
	var returns *ComputeHealthCheckGrpcHealthCheck
	_jsii_.Get(
		j,
		"grpcHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) GrpcTlsHealthCheck() ComputeHealthCheckGrpcTlsHealthCheckOutputReference {
	var returns ComputeHealthCheckGrpcTlsHealthCheckOutputReference
	_jsii_.Get(
		j,
		"grpcTlsHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) GrpcTlsHealthCheckInput() *ComputeHealthCheckGrpcTlsHealthCheck {
	var returns *ComputeHealthCheckGrpcTlsHealthCheck
	_jsii_.Get(
		j,
		"grpcTlsHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) HealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) HealthyThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Http2HealthCheck() ComputeHealthCheckHttp2HealthCheckOutputReference {
	var returns ComputeHealthCheckHttp2HealthCheckOutputReference
	_jsii_.Get(
		j,
		"http2HealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Http2HealthCheckInput() *ComputeHealthCheckHttp2HealthCheck {
	var returns *ComputeHealthCheckHttp2HealthCheck
	_jsii_.Get(
		j,
		"http2HealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) HttpHealthCheck() ComputeHealthCheckHttpHealthCheckOutputReference {
	var returns ComputeHealthCheckHttpHealthCheckOutputReference
	_jsii_.Get(
		j,
		"httpHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) HttpHealthCheckInput() *ComputeHealthCheckHttpHealthCheck {
	var returns *ComputeHealthCheckHttpHealthCheck
	_jsii_.Get(
		j,
		"httpHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) HttpsHealthCheck() ComputeHealthCheckHttpsHealthCheckOutputReference {
	var returns ComputeHealthCheckHttpsHealthCheckOutputReference
	_jsii_.Get(
		j,
		"httpsHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) HttpsHealthCheckInput() *ComputeHealthCheckHttpsHealthCheck {
	var returns *ComputeHealthCheckHttpsHealthCheck
	_jsii_.Get(
		j,
		"httpsHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) LogConfig() ComputeHealthCheckLogConfigOutputReference {
	var returns ComputeHealthCheckLogConfigOutputReference
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) LogConfigInput() *ComputeHealthCheckLogConfig {
	var returns *ComputeHealthCheckLogConfig
	_jsii_.Get(
		j,
		"logConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) SourceRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) SourceRegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceRegionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) SslHealthCheck() ComputeHealthCheckSslHealthCheckOutputReference {
	var returns ComputeHealthCheckSslHealthCheckOutputReference
	_jsii_.Get(
		j,
		"sslHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) SslHealthCheckInput() *ComputeHealthCheckSslHealthCheck {
	var returns *ComputeHealthCheckSslHealthCheck
	_jsii_.Get(
		j,
		"sslHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) TcpHealthCheck() ComputeHealthCheckTcpHealthCheckOutputReference {
	var returns ComputeHealthCheckTcpHealthCheckOutputReference
	_jsii_.Get(
		j,
		"tcpHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) TcpHealthCheckInput() *ComputeHealthCheckTcpHealthCheck {
	var returns *ComputeHealthCheckTcpHealthCheck
	_jsii_.Get(
		j,
		"tcpHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Timeouts() ComputeHealthCheckTimeoutsOutputReference {
	var returns ComputeHealthCheckTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) TimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) TimeoutSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) UnhealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheck) UnhealthyThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThresholdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_health_check google_compute_health_check} Resource.
func NewComputeHealthCheck(scope constructs.Construct, id *string, config *ComputeHealthCheckConfig) ComputeHealthCheck {
	_init_.Initialize()

	if err := validateNewComputeHealthCheckParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeHealthCheck{}

	_jsii_.Create(
		"@cdktn/provider-google.computeHealthCheck.ComputeHealthCheck",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_health_check google_compute_health_check} Resource.
func NewComputeHealthCheck_Override(c ComputeHealthCheck, scope constructs.Construct, id *string, config *ComputeHealthCheckConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.computeHealthCheck.ComputeHealthCheck",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetCheckIntervalSec(val *float64) {
	if err := j.validateSetCheckIntervalSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkIntervalSec",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetHealthyThreshold(val *float64) {
	if err := j.validateSetHealthyThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthyThreshold",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetSourceRegions(val *[]*string) {
	if err := j.validateSetSourceRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceRegions",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetTimeoutSec(val *float64) {
	if err := j.validateSetTimeoutSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSec",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheck)SetUnhealthyThreshold(val *float64) {
	if err := j.validateSetUnhealthyThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unhealthyThreshold",
		val,
	)
}

// Generates CDKTN code for importing a ComputeHealthCheck resource upon running "cdktn plan <stack-name>".
func ComputeHealthCheck_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateComputeHealthCheck_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.computeHealthCheck.ComputeHealthCheck",
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
func ComputeHealthCheck_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeHealthCheck_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.computeHealthCheck.ComputeHealthCheck",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeHealthCheck_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeHealthCheck_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.computeHealthCheck.ComputeHealthCheck",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeHealthCheck_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeHealthCheck_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.computeHealthCheck.ComputeHealthCheck",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ComputeHealthCheck_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.computeHealthCheck.ComputeHealthCheck",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := c.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) PutGrpcHealthCheck(value *ComputeHealthCheckGrpcHealthCheck) {
	if err := c.validatePutGrpcHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGrpcHealthCheck",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) PutGrpcTlsHealthCheck(value *ComputeHealthCheckGrpcTlsHealthCheck) {
	if err := c.validatePutGrpcTlsHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGrpcTlsHealthCheck",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) PutHttp2HealthCheck(value *ComputeHealthCheckHttp2HealthCheck) {
	if err := c.validatePutHttp2HealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttp2HealthCheck",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) PutHttpHealthCheck(value *ComputeHealthCheckHttpHealthCheck) {
	if err := c.validatePutHttpHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttpHealthCheck",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) PutHttpsHealthCheck(value *ComputeHealthCheckHttpsHealthCheck) {
	if err := c.validatePutHttpsHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttpsHealthCheck",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) PutLogConfig(value *ComputeHealthCheckLogConfig) {
	if err := c.validatePutLogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLogConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) PutSslHealthCheck(value *ComputeHealthCheckSslHealthCheck) {
	if err := c.validatePutSslHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSslHealthCheck",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) PutTcpHealthCheck(value *ComputeHealthCheckTcpHealthCheck) {
	if err := c.validatePutTcpHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTcpHealthCheck",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) PutTimeouts(value *ComputeHealthCheckTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := c.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetCheckIntervalSec() {
	_jsii_.InvokeVoid(
		c,
		"resetCheckIntervalSec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetGrpcHealthCheck() {
	_jsii_.InvokeVoid(
		c,
		"resetGrpcHealthCheck",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetGrpcTlsHealthCheck() {
	_jsii_.InvokeVoid(
		c,
		"resetGrpcTlsHealthCheck",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetHealthyThreshold() {
	_jsii_.InvokeVoid(
		c,
		"resetHealthyThreshold",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetHttp2HealthCheck() {
	_jsii_.InvokeVoid(
		c,
		"resetHttp2HealthCheck",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetHttpHealthCheck() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpHealthCheck",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetHttpsHealthCheck() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpsHealthCheck",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetLogConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetLogConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetSourceRegions() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceRegions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetSslHealthCheck() {
	_jsii_.InvokeVoid(
		c,
		"resetSslHealthCheck",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetTcpHealthCheck() {
	_jsii_.InvokeVoid(
		c,
		"resetTcpHealthCheck",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetTimeoutSec() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeoutSec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) ResetUnhealthyThreshold() {
	_jsii_.InvokeVoid(
		c,
		"resetUnhealthyThreshold",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheck) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheck) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

