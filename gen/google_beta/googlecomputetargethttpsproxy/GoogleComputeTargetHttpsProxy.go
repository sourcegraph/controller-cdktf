package googlecomputetargethttpsproxy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputetargethttpsproxy/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_target_https_proxy google_compute_target_https_proxy}.
type GoogleComputeTargetHttpsProxy interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CertificateManagerCertificates() *[]*string
	SetCertificateManagerCertificates(val *[]*string)
	CertificateManagerCertificatesInput() *[]*string
	CertificateMap() *string
	SetCertificateMap(val *string)
	CertificateMapInput() *string
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
	Fingerprint() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpKeepAliveTimeoutSec() *float64
	SetHttpKeepAliveTimeoutSec(val *float64)
	HttpKeepAliveTimeoutSecInput() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
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
	ProxyBind() interface{}
	SetProxyBind(val interface{})
	ProxyBindInput() interface{}
	ProxyId() *float64
	QuicOverride() *string
	SetQuicOverride(val *string)
	QuicOverrideInput() *string
	// Experimental.
	RawOverrides() interface{}
	SelfLink() *string
	ServerTlsPolicy() *string
	SetServerTlsPolicy(val *string)
	ServerTlsPolicyInput() *string
	SslCertificates() *[]*string
	SetSslCertificates(val *[]*string)
	SslCertificatesInput() *[]*string
	SslPolicy() *string
	SetSslPolicy(val *string)
	SslPolicyInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleComputeTargetHttpsProxyTimeoutsOutputReference
	TimeoutsInput() interface{}
	TlsEarlyData() *string
	SetTlsEarlyData(val *string)
	TlsEarlyDataInput() *string
	UrlMap() *string
	SetUrlMap(val *string)
	UrlMapInput() *string
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
	PutTimeouts(value *GoogleComputeTargetHttpsProxyTimeouts)
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
	ResetCertificateManagerCertificates()
	ResetCertificateMap()
	ResetDescription()
	ResetHttpKeepAliveTimeoutSec()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetProxyBind()
	ResetQuicOverride()
	ResetServerTlsPolicy()
	ResetSslCertificates()
	ResetSslPolicy()
	ResetTimeouts()
	ResetTlsEarlyData()
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

// The jsii proxy struct for GoogleComputeTargetHttpsProxy
type jsiiProxy_GoogleComputeTargetHttpsProxy struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) CertificateManagerCertificates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"certificateManagerCertificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) CertificateManagerCertificatesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"certificateManagerCertificatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) CertificateMap() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) CertificateMapInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) Fingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) HttpKeepAliveTimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpKeepAliveTimeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) HttpKeepAliveTimeoutSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpKeepAliveTimeoutSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) ProxyBind() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxyBind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) ProxyBindInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxyBindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) ProxyId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"proxyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) QuicOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quicOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) QuicOverrideInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quicOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) ServerTlsPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverTlsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) ServerTlsPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverTlsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) SslCertificates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sslCertificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) SslCertificatesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sslCertificatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) SslPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) SslPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) Timeouts() GoogleComputeTargetHttpsProxyTimeoutsOutputReference {
	var returns GoogleComputeTargetHttpsProxyTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) TlsEarlyData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsEarlyData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) TlsEarlyDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsEarlyDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) UrlMap() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy) UrlMapInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlMapInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_target_https_proxy google_compute_target_https_proxy} Resource.
func NewGoogleComputeTargetHttpsProxy(scope constructs.Construct, id *string, config *GoogleComputeTargetHttpsProxyConfig) GoogleComputeTargetHttpsProxy {
	_init_.Initialize()

	if err := validateNewGoogleComputeTargetHttpsProxyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeTargetHttpsProxy{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeTargetHttpsProxy.GoogleComputeTargetHttpsProxy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_target_https_proxy google_compute_target_https_proxy} Resource.
func NewGoogleComputeTargetHttpsProxy_Override(g GoogleComputeTargetHttpsProxy, scope constructs.Construct, id *string, config *GoogleComputeTargetHttpsProxyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeTargetHttpsProxy.GoogleComputeTargetHttpsProxy",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy)SetCertificateManagerCertificates(val *[]*string) {
	if err := j.validateSetCertificateManagerCertificatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateManagerCertificates",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy)SetCertificateMap(val *string) {
	if err := j.validateSetCertificateMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateMap",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy)SetHttpKeepAliveTimeoutSec(val *float64) {
	if err := j.validateSetHttpKeepAliveTimeoutSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpKeepAliveTimeoutSec",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy)SetProxyBind(val interface{}) {
	if err := j.validateSetProxyBindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyBind",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy)SetQuicOverride(val *string) {
	if err := j.validateSetQuicOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quicOverride",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy)SetServerTlsPolicy(val *string) {
	if err := j.validateSetServerTlsPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverTlsPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy)SetSslCertificates(val *[]*string) {
	if err := j.validateSetSslCertificatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCertificates",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy)SetSslPolicy(val *string) {
	if err := j.validateSetSslPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy)SetTlsEarlyData(val *string) {
	if err := j.validateSetTlsEarlyDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsEarlyData",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeTargetHttpsProxy)SetUrlMap(val *string) {
	if err := j.validateSetUrlMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlMap",
		val,
	)
}

// Generates CDKTN code for importing a GoogleComputeTargetHttpsProxy resource upon running "cdktn plan <stack-name>".
func GoogleComputeTargetHttpsProxy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleComputeTargetHttpsProxy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleComputeTargetHttpsProxy.GoogleComputeTargetHttpsProxy",
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
func GoogleComputeTargetHttpsProxy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeTargetHttpsProxy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleComputeTargetHttpsProxy.GoogleComputeTargetHttpsProxy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeTargetHttpsProxy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeTargetHttpsProxy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleComputeTargetHttpsProxy.GoogleComputeTargetHttpsProxy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeTargetHttpsProxy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeTargetHttpsProxy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleComputeTargetHttpsProxy.GoogleComputeTargetHttpsProxy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputeTargetHttpsProxy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google-beta.googleComputeTargetHttpsProxy.GoogleComputeTargetHttpsProxy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) PutTimeouts(value *GoogleComputeTargetHttpsProxyTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) ResetCertificateManagerCertificates() {
	_jsii_.InvokeVoid(
		g,
		"resetCertificateManagerCertificates",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) ResetCertificateMap() {
	_jsii_.InvokeVoid(
		g,
		"resetCertificateMap",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) ResetHttpKeepAliveTimeoutSec() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpKeepAliveTimeoutSec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) ResetProxyBind() {
	_jsii_.InvokeVoid(
		g,
		"resetProxyBind",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) ResetQuicOverride() {
	_jsii_.InvokeVoid(
		g,
		"resetQuicOverride",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) ResetServerTlsPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetServerTlsPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) ResetSslCertificates() {
	_jsii_.InvokeVoid(
		g,
		"resetSslCertificates",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) ResetSslPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetSslPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) ResetTlsEarlyData() {
	_jsii_.InvokeVoid(
		g,
		"resetTlsEarlyData",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeTargetHttpsProxy) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

