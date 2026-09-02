package googlecomputehealthcheck

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputehealthcheck/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_health_check google_compute_health_check}.
type GoogleComputeHealthCheck interface {
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
	GrpcHealthCheck() GoogleComputeHealthCheckGrpcHealthCheckOutputReference
	GrpcHealthCheckInput() *GoogleComputeHealthCheckGrpcHealthCheck
	GrpcTlsHealthCheck() GoogleComputeHealthCheckGrpcTlsHealthCheckOutputReference
	GrpcTlsHealthCheckInput() *GoogleComputeHealthCheckGrpcTlsHealthCheck
	HealthyThreshold() *float64
	SetHealthyThreshold(val *float64)
	HealthyThresholdInput() *float64
	Http2HealthCheck() GoogleComputeHealthCheckHttp2HealthCheckOutputReference
	Http2HealthCheckInput() *GoogleComputeHealthCheckHttp2HealthCheck
	HttpHealthCheck() GoogleComputeHealthCheckHttpHealthCheckOutputReference
	HttpHealthCheckInput() *GoogleComputeHealthCheckHttpHealthCheck
	HttpsHealthCheck() GoogleComputeHealthCheckHttpsHealthCheckOutputReference
	HttpsHealthCheckInput() *GoogleComputeHealthCheckHttpsHealthCheck
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LogConfig() GoogleComputeHealthCheckLogConfigOutputReference
	LogConfigInput() *GoogleComputeHealthCheckLogConfig
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
	SslHealthCheck() GoogleComputeHealthCheckSslHealthCheckOutputReference
	SslHealthCheckInput() *GoogleComputeHealthCheckSslHealthCheck
	TcpHealthCheck() GoogleComputeHealthCheckTcpHealthCheckOutputReference
	TcpHealthCheckInput() *GoogleComputeHealthCheckTcpHealthCheck
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleComputeHealthCheckTimeoutsOutputReference
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
	PutGrpcHealthCheck(value *GoogleComputeHealthCheckGrpcHealthCheck)
	PutGrpcTlsHealthCheck(value *GoogleComputeHealthCheckGrpcTlsHealthCheck)
	PutHttp2HealthCheck(value *GoogleComputeHealthCheckHttp2HealthCheck)
	PutHttpHealthCheck(value *GoogleComputeHealthCheckHttpHealthCheck)
	PutHttpsHealthCheck(value *GoogleComputeHealthCheckHttpsHealthCheck)
	PutLogConfig(value *GoogleComputeHealthCheckLogConfig)
	PutSslHealthCheck(value *GoogleComputeHealthCheckSslHealthCheck)
	PutTcpHealthCheck(value *GoogleComputeHealthCheckTcpHealthCheck)
	PutTimeouts(value *GoogleComputeHealthCheckTimeouts)
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

// The jsii proxy struct for GoogleComputeHealthCheck
type jsiiProxy_GoogleComputeHealthCheck struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GoogleComputeHealthCheck) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) CheckIntervalSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"checkIntervalSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) CheckIntervalSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"checkIntervalSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) GrpcHealthCheck() GoogleComputeHealthCheckGrpcHealthCheckOutputReference {
	var returns GoogleComputeHealthCheckGrpcHealthCheckOutputReference
	_jsii_.Get(
		j,
		"grpcHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) GrpcHealthCheckInput() *GoogleComputeHealthCheckGrpcHealthCheck {
	var returns *GoogleComputeHealthCheckGrpcHealthCheck
	_jsii_.Get(
		j,
		"grpcHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) GrpcTlsHealthCheck() GoogleComputeHealthCheckGrpcTlsHealthCheckOutputReference {
	var returns GoogleComputeHealthCheckGrpcTlsHealthCheckOutputReference
	_jsii_.Get(
		j,
		"grpcTlsHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) GrpcTlsHealthCheckInput() *GoogleComputeHealthCheckGrpcTlsHealthCheck {
	var returns *GoogleComputeHealthCheckGrpcTlsHealthCheck
	_jsii_.Get(
		j,
		"grpcTlsHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) HealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) HealthyThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) Http2HealthCheck() GoogleComputeHealthCheckHttp2HealthCheckOutputReference {
	var returns GoogleComputeHealthCheckHttp2HealthCheckOutputReference
	_jsii_.Get(
		j,
		"http2HealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) Http2HealthCheckInput() *GoogleComputeHealthCheckHttp2HealthCheck {
	var returns *GoogleComputeHealthCheckHttp2HealthCheck
	_jsii_.Get(
		j,
		"http2HealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) HttpHealthCheck() GoogleComputeHealthCheckHttpHealthCheckOutputReference {
	var returns GoogleComputeHealthCheckHttpHealthCheckOutputReference
	_jsii_.Get(
		j,
		"httpHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) HttpHealthCheckInput() *GoogleComputeHealthCheckHttpHealthCheck {
	var returns *GoogleComputeHealthCheckHttpHealthCheck
	_jsii_.Get(
		j,
		"httpHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) HttpsHealthCheck() GoogleComputeHealthCheckHttpsHealthCheckOutputReference {
	var returns GoogleComputeHealthCheckHttpsHealthCheckOutputReference
	_jsii_.Get(
		j,
		"httpsHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) HttpsHealthCheckInput() *GoogleComputeHealthCheckHttpsHealthCheck {
	var returns *GoogleComputeHealthCheckHttpsHealthCheck
	_jsii_.Get(
		j,
		"httpsHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) LogConfig() GoogleComputeHealthCheckLogConfigOutputReference {
	var returns GoogleComputeHealthCheckLogConfigOutputReference
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) LogConfigInput() *GoogleComputeHealthCheckLogConfig {
	var returns *GoogleComputeHealthCheckLogConfig
	_jsii_.Get(
		j,
		"logConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) SourceRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) SourceRegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceRegionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) SslHealthCheck() GoogleComputeHealthCheckSslHealthCheckOutputReference {
	var returns GoogleComputeHealthCheckSslHealthCheckOutputReference
	_jsii_.Get(
		j,
		"sslHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) SslHealthCheckInput() *GoogleComputeHealthCheckSslHealthCheck {
	var returns *GoogleComputeHealthCheckSslHealthCheck
	_jsii_.Get(
		j,
		"sslHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) TcpHealthCheck() GoogleComputeHealthCheckTcpHealthCheckOutputReference {
	var returns GoogleComputeHealthCheckTcpHealthCheckOutputReference
	_jsii_.Get(
		j,
		"tcpHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) TcpHealthCheckInput() *GoogleComputeHealthCheckTcpHealthCheck {
	var returns *GoogleComputeHealthCheckTcpHealthCheck
	_jsii_.Get(
		j,
		"tcpHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) Timeouts() GoogleComputeHealthCheckTimeoutsOutputReference {
	var returns GoogleComputeHealthCheckTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) TimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) TimeoutSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) UnhealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeHealthCheck) UnhealthyThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThresholdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_health_check google_compute_health_check} Resource.
func NewGoogleComputeHealthCheck(scope constructs.Construct, id *string, config *GoogleComputeHealthCheckConfig) GoogleComputeHealthCheck {
	_init_.Initialize()

	if err := validateNewGoogleComputeHealthCheckParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeHealthCheck{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeHealthCheck.GoogleComputeHealthCheck",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_health_check google_compute_health_check} Resource.
func NewGoogleComputeHealthCheck_Override(g GoogleComputeHealthCheck, scope constructs.Construct, id *string, config *GoogleComputeHealthCheckConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeHealthCheck.GoogleComputeHealthCheck",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheck)SetCheckIntervalSec(val *float64) {
	if err := j.validateSetCheckIntervalSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkIntervalSec",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheck)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheck)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheck)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheck)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheck)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheck)SetHealthyThreshold(val *float64) {
	if err := j.validateSetHealthyThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthyThreshold",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheck)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheck)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheck)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheck)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheck)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheck)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheck)SetSourceRegions(val *[]*string) {
	if err := j.validateSetSourceRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceRegions",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheck)SetTimeoutSec(val *float64) {
	if err := j.validateSetTimeoutSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSec",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeHealthCheck)SetUnhealthyThreshold(val *float64) {
	if err := j.validateSetUnhealthyThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unhealthyThreshold",
		val,
	)
}

// Generates CDKTN code for importing a GoogleComputeHealthCheck resource upon running "cdktn plan <stack-name>".
func GoogleComputeHealthCheck_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleComputeHealthCheck_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleComputeHealthCheck.GoogleComputeHealthCheck",
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
func GoogleComputeHealthCheck_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeHealthCheck_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleComputeHealthCheck.GoogleComputeHealthCheck",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeHealthCheck_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeHealthCheck_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleComputeHealthCheck.GoogleComputeHealthCheck",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeHealthCheck_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeHealthCheck_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleComputeHealthCheck.GoogleComputeHealthCheck",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputeHealthCheck_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google-beta.googleComputeHealthCheck.GoogleComputeHealthCheck",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeHealthCheck) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeHealthCheck) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeHealthCheck) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeHealthCheck) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeHealthCheck) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeHealthCheck) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeHealthCheck) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeHealthCheck) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeHealthCheck) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeHealthCheck) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeHealthCheck) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeHealthCheck) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (g *jsiiProxy_GoogleComputeHealthCheck) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) PutGrpcHealthCheck(value *GoogleComputeHealthCheckGrpcHealthCheck) {
	if err := g.validatePutGrpcHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGrpcHealthCheck",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) PutGrpcTlsHealthCheck(value *GoogleComputeHealthCheckGrpcTlsHealthCheck) {
	if err := g.validatePutGrpcTlsHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGrpcTlsHealthCheck",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) PutHttp2HealthCheck(value *GoogleComputeHealthCheckHttp2HealthCheck) {
	if err := g.validatePutHttp2HealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHttp2HealthCheck",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) PutHttpHealthCheck(value *GoogleComputeHealthCheckHttpHealthCheck) {
	if err := g.validatePutHttpHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHttpHealthCheck",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) PutHttpsHealthCheck(value *GoogleComputeHealthCheckHttpsHealthCheck) {
	if err := g.validatePutHttpsHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHttpsHealthCheck",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) PutLogConfig(value *GoogleComputeHealthCheckLogConfig) {
	if err := g.validatePutLogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLogConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) PutSslHealthCheck(value *GoogleComputeHealthCheckSslHealthCheck) {
	if err := g.validatePutSslHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSslHealthCheck",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) PutTcpHealthCheck(value *GoogleComputeHealthCheckTcpHealthCheck) {
	if err := g.validatePutTcpHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTcpHealthCheck",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) PutTimeouts(value *GoogleComputeHealthCheckTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) ResetCheckIntervalSec() {
	_jsii_.InvokeVoid(
		g,
		"resetCheckIntervalSec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) ResetGrpcHealthCheck() {
	_jsii_.InvokeVoid(
		g,
		"resetGrpcHealthCheck",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) ResetGrpcTlsHealthCheck() {
	_jsii_.InvokeVoid(
		g,
		"resetGrpcTlsHealthCheck",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) ResetHealthyThreshold() {
	_jsii_.InvokeVoid(
		g,
		"resetHealthyThreshold",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) ResetHttp2HealthCheck() {
	_jsii_.InvokeVoid(
		g,
		"resetHttp2HealthCheck",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) ResetHttpHealthCheck() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpHealthCheck",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) ResetHttpsHealthCheck() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpsHealthCheck",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) ResetLogConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLogConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) ResetSourceRegions() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceRegions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) ResetSslHealthCheck() {
	_jsii_.InvokeVoid(
		g,
		"resetSslHealthCheck",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) ResetTcpHealthCheck() {
	_jsii_.InvokeVoid(
		g,
		"resetTcpHealthCheck",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) ResetTimeoutSec() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeoutSec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) ResetUnhealthyThreshold() {
	_jsii_.InvokeVoid(
		g,
		"resetUnhealthyThreshold",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeHealthCheck) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeHealthCheck) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeHealthCheck) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeHealthCheck) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeHealthCheck) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeHealthCheck) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeHealthCheck) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

