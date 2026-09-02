package googlecomputeregionbackendservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeregionbackendservice/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_region_backend_service google_compute_region_backend_service}.
type GoogleComputeRegionBackendService interface {
	cdktn.TerraformResource
	AffinityCookieTtlSec() *float64
	SetAffinityCookieTtlSec(val *float64)
	AffinityCookieTtlSecInput() *float64
	Backend() GoogleComputeRegionBackendServiceBackendList
	BackendInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CdnPolicy() GoogleComputeRegionBackendServiceCdnPolicyOutputReference
	CdnPolicyInput() *GoogleComputeRegionBackendServiceCdnPolicy
	CircuitBreakers() GoogleComputeRegionBackendServiceCircuitBreakersOutputReference
	CircuitBreakersInput() *GoogleComputeRegionBackendServiceCircuitBreakers
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionDrainingTimeoutSec() *float64
	SetConnectionDrainingTimeoutSec(val *float64)
	ConnectionDrainingTimeoutSecInput() *float64
	ConnectionTrackingPolicy() GoogleComputeRegionBackendServiceConnectionTrackingPolicyOutputReference
	ConnectionTrackingPolicyInput() *GoogleComputeRegionBackendServiceConnectionTrackingPolicy
	ConsistentHash() GoogleComputeRegionBackendServiceConsistentHashOutputReference
	ConsistentHashInput() *GoogleComputeRegionBackendServiceConsistentHash
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreationTimestamp() *string
	CustomMetrics() GoogleComputeRegionBackendServiceCustomMetricsList
	CustomMetricsInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DynamicForwarding() GoogleComputeRegionBackendServiceDynamicForwardingOutputReference
	DynamicForwardingInput() *GoogleComputeRegionBackendServiceDynamicForwarding
	EnableCdn() interface{}
	SetEnableCdn(val interface{})
	EnableCdnInput() interface{}
	FailoverPolicy() GoogleComputeRegionBackendServiceFailoverPolicyOutputReference
	FailoverPolicyInput() *GoogleComputeRegionBackendServiceFailoverPolicy
	Fingerprint() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GeneratedId() *float64
	HaPolicy() GoogleComputeRegionBackendServiceHaPolicyOutputReference
	HaPolicyInput() *GoogleComputeRegionBackendServiceHaPolicy
	HealthChecks() *[]*string
	SetHealthChecks(val *[]*string)
	HealthChecksInput() *[]*string
	Iap() GoogleComputeRegionBackendServiceIapOutputReference
	IapInput() *GoogleComputeRegionBackendServiceIap
	Id() *string
	SetId(val *string)
	IdInput() *string
	IpAddressSelectionPolicy() *string
	SetIpAddressSelectionPolicy(val *string)
	IpAddressSelectionPolicyInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LoadBalancingScheme() *string
	SetLoadBalancingScheme(val *string)
	LoadBalancingSchemeInput() *string
	LocalityLbPolicy() *string
	SetLocalityLbPolicy(val *string)
	LocalityLbPolicyInput() *string
	LogConfig() GoogleComputeRegionBackendServiceLogConfigOutputReference
	LogConfigInput() *GoogleComputeRegionBackendServiceLogConfig
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	NetworkPassThroughLbTrafficPolicy() GoogleComputeRegionBackendServiceNetworkPassThroughLbTrafficPolicyOutputReference
	NetworkPassThroughLbTrafficPolicyInput() *GoogleComputeRegionBackendServiceNetworkPassThroughLbTrafficPolicy
	// The tree node.
	Node() constructs.Node
	OutlierDetection() GoogleComputeRegionBackendServiceOutlierDetectionOutputReference
	OutlierDetectionInput() *GoogleComputeRegionBackendServiceOutlierDetection
	Params() GoogleComputeRegionBackendServiceParamsOutputReference
	ParamsInput() *GoogleComputeRegionBackendServiceParams
	PortName() *string
	SetPortName(val *string)
	PortNameInput() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SecurityPolicy() *string
	SetSecurityPolicy(val *string)
	SecurityPolicyInput() *string
	SelfLink() *string
	SessionAffinity() *string
	SetSessionAffinity(val *string)
	SessionAffinityInput() *string
	StrongSessionAffinityCookie() GoogleComputeRegionBackendServiceStrongSessionAffinityCookieOutputReference
	StrongSessionAffinityCookieInput() *GoogleComputeRegionBackendServiceStrongSessionAffinityCookie
	Subsetting() GoogleComputeRegionBackendServiceSubsettingOutputReference
	SubsettingInput() *GoogleComputeRegionBackendServiceSubsetting
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleComputeRegionBackendServiceTimeoutsOutputReference
	TimeoutSec() *float64
	SetTimeoutSec(val *float64)
	TimeoutSecInput() *float64
	TimeoutsInput() interface{}
	TlsSettings() GoogleComputeRegionBackendServiceTlsSettingsOutputReference
	TlsSettingsInput() *GoogleComputeRegionBackendServiceTlsSettings
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
	PutBackend(value interface{})
	PutCdnPolicy(value *GoogleComputeRegionBackendServiceCdnPolicy)
	PutCircuitBreakers(value *GoogleComputeRegionBackendServiceCircuitBreakers)
	PutConnectionTrackingPolicy(value *GoogleComputeRegionBackendServiceConnectionTrackingPolicy)
	PutConsistentHash(value *GoogleComputeRegionBackendServiceConsistentHash)
	PutCustomMetrics(value interface{})
	PutDynamicForwarding(value *GoogleComputeRegionBackendServiceDynamicForwarding)
	PutFailoverPolicy(value *GoogleComputeRegionBackendServiceFailoverPolicy)
	PutHaPolicy(value *GoogleComputeRegionBackendServiceHaPolicy)
	PutIap(value *GoogleComputeRegionBackendServiceIap)
	PutLogConfig(value *GoogleComputeRegionBackendServiceLogConfig)
	PutNetworkPassThroughLbTrafficPolicy(value *GoogleComputeRegionBackendServiceNetworkPassThroughLbTrafficPolicy)
	PutOutlierDetection(value *GoogleComputeRegionBackendServiceOutlierDetection)
	PutParams(value *GoogleComputeRegionBackendServiceParams)
	PutStrongSessionAffinityCookie(value *GoogleComputeRegionBackendServiceStrongSessionAffinityCookie)
	PutSubsetting(value *GoogleComputeRegionBackendServiceSubsetting)
	PutTimeouts(value *GoogleComputeRegionBackendServiceTimeouts)
	PutTlsSettings(value *GoogleComputeRegionBackendServiceTlsSettings)
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
	ResetAffinityCookieTtlSec()
	ResetBackend()
	ResetCdnPolicy()
	ResetCircuitBreakers()
	ResetConnectionDrainingTimeoutSec()
	ResetConnectionTrackingPolicy()
	ResetConsistentHash()
	ResetCustomMetrics()
	ResetDescription()
	ResetDynamicForwarding()
	ResetEnableCdn()
	ResetFailoverPolicy()
	ResetHaPolicy()
	ResetHealthChecks()
	ResetIap()
	ResetId()
	ResetIpAddressSelectionPolicy()
	ResetLoadBalancingScheme()
	ResetLocalityLbPolicy()
	ResetLogConfig()
	ResetNetwork()
	ResetNetworkPassThroughLbTrafficPolicy()
	ResetOutlierDetection()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParams()
	ResetPortName()
	ResetProject()
	ResetProtocol()
	ResetRegion()
	ResetSecurityPolicy()
	ResetSessionAffinity()
	ResetStrongSessionAffinityCookie()
	ResetSubsetting()
	ResetTimeouts()
	ResetTimeoutSec()
	ResetTlsSettings()
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

// The jsii proxy struct for GoogleComputeRegionBackendService
type jsiiProxy_GoogleComputeRegionBackendService struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) AffinityCookieTtlSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"affinityCookieTtlSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) AffinityCookieTtlSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"affinityCookieTtlSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) Backend() GoogleComputeRegionBackendServiceBackendList {
	var returns GoogleComputeRegionBackendServiceBackendList
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) BackendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) CdnPolicy() GoogleComputeRegionBackendServiceCdnPolicyOutputReference {
	var returns GoogleComputeRegionBackendServiceCdnPolicyOutputReference
	_jsii_.Get(
		j,
		"cdnPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) CdnPolicyInput() *GoogleComputeRegionBackendServiceCdnPolicy {
	var returns *GoogleComputeRegionBackendServiceCdnPolicy
	_jsii_.Get(
		j,
		"cdnPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) CircuitBreakers() GoogleComputeRegionBackendServiceCircuitBreakersOutputReference {
	var returns GoogleComputeRegionBackendServiceCircuitBreakersOutputReference
	_jsii_.Get(
		j,
		"circuitBreakers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) CircuitBreakersInput() *GoogleComputeRegionBackendServiceCircuitBreakers {
	var returns *GoogleComputeRegionBackendServiceCircuitBreakers
	_jsii_.Get(
		j,
		"circuitBreakersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) ConnectionDrainingTimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionDrainingTimeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) ConnectionDrainingTimeoutSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionDrainingTimeoutSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) ConnectionTrackingPolicy() GoogleComputeRegionBackendServiceConnectionTrackingPolicyOutputReference {
	var returns GoogleComputeRegionBackendServiceConnectionTrackingPolicyOutputReference
	_jsii_.Get(
		j,
		"connectionTrackingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) ConnectionTrackingPolicyInput() *GoogleComputeRegionBackendServiceConnectionTrackingPolicy {
	var returns *GoogleComputeRegionBackendServiceConnectionTrackingPolicy
	_jsii_.Get(
		j,
		"connectionTrackingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) ConsistentHash() GoogleComputeRegionBackendServiceConsistentHashOutputReference {
	var returns GoogleComputeRegionBackendServiceConsistentHashOutputReference
	_jsii_.Get(
		j,
		"consistentHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) ConsistentHashInput() *GoogleComputeRegionBackendServiceConsistentHash {
	var returns *GoogleComputeRegionBackendServiceConsistentHash
	_jsii_.Get(
		j,
		"consistentHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) CustomMetrics() GoogleComputeRegionBackendServiceCustomMetricsList {
	var returns GoogleComputeRegionBackendServiceCustomMetricsList
	_jsii_.Get(
		j,
		"customMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) CustomMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) DynamicForwarding() GoogleComputeRegionBackendServiceDynamicForwardingOutputReference {
	var returns GoogleComputeRegionBackendServiceDynamicForwardingOutputReference
	_jsii_.Get(
		j,
		"dynamicForwarding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) DynamicForwardingInput() *GoogleComputeRegionBackendServiceDynamicForwarding {
	var returns *GoogleComputeRegionBackendServiceDynamicForwarding
	_jsii_.Get(
		j,
		"dynamicForwardingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) EnableCdn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) EnableCdnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) FailoverPolicy() GoogleComputeRegionBackendServiceFailoverPolicyOutputReference {
	var returns GoogleComputeRegionBackendServiceFailoverPolicyOutputReference
	_jsii_.Get(
		j,
		"failoverPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) FailoverPolicyInput() *GoogleComputeRegionBackendServiceFailoverPolicy {
	var returns *GoogleComputeRegionBackendServiceFailoverPolicy
	_jsii_.Get(
		j,
		"failoverPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) Fingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) GeneratedId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"generatedId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) HaPolicy() GoogleComputeRegionBackendServiceHaPolicyOutputReference {
	var returns GoogleComputeRegionBackendServiceHaPolicyOutputReference
	_jsii_.Get(
		j,
		"haPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) HaPolicyInput() *GoogleComputeRegionBackendServiceHaPolicy {
	var returns *GoogleComputeRegionBackendServiceHaPolicy
	_jsii_.Get(
		j,
		"haPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) HealthChecks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"healthChecks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) HealthChecksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"healthChecksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) Iap() GoogleComputeRegionBackendServiceIapOutputReference {
	var returns GoogleComputeRegionBackendServiceIapOutputReference
	_jsii_.Get(
		j,
		"iap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) IapInput() *GoogleComputeRegionBackendServiceIap {
	var returns *GoogleComputeRegionBackendServiceIap
	_jsii_.Get(
		j,
		"iapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) IpAddressSelectionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressSelectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) IpAddressSelectionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressSelectionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) LoadBalancingScheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) LoadBalancingSchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingSchemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) LocalityLbPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localityLbPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) LocalityLbPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localityLbPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) LogConfig() GoogleComputeRegionBackendServiceLogConfigOutputReference {
	var returns GoogleComputeRegionBackendServiceLogConfigOutputReference
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) LogConfigInput() *GoogleComputeRegionBackendServiceLogConfig {
	var returns *GoogleComputeRegionBackendServiceLogConfig
	_jsii_.Get(
		j,
		"logConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) NetworkPassThroughLbTrafficPolicy() GoogleComputeRegionBackendServiceNetworkPassThroughLbTrafficPolicyOutputReference {
	var returns GoogleComputeRegionBackendServiceNetworkPassThroughLbTrafficPolicyOutputReference
	_jsii_.Get(
		j,
		"networkPassThroughLbTrafficPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) NetworkPassThroughLbTrafficPolicyInput() *GoogleComputeRegionBackendServiceNetworkPassThroughLbTrafficPolicy {
	var returns *GoogleComputeRegionBackendServiceNetworkPassThroughLbTrafficPolicy
	_jsii_.Get(
		j,
		"networkPassThroughLbTrafficPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) OutlierDetection() GoogleComputeRegionBackendServiceOutlierDetectionOutputReference {
	var returns GoogleComputeRegionBackendServiceOutlierDetectionOutputReference
	_jsii_.Get(
		j,
		"outlierDetection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) OutlierDetectionInput() *GoogleComputeRegionBackendServiceOutlierDetection {
	var returns *GoogleComputeRegionBackendServiceOutlierDetection
	_jsii_.Get(
		j,
		"outlierDetectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) Params() GoogleComputeRegionBackendServiceParamsOutputReference {
	var returns GoogleComputeRegionBackendServiceParamsOutputReference
	_jsii_.Get(
		j,
		"params",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) ParamsInput() *GoogleComputeRegionBackendServiceParams {
	var returns *GoogleComputeRegionBackendServiceParams
	_jsii_.Get(
		j,
		"paramsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) PortName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) PortNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) SecurityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) SecurityPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) SessionAffinity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) SessionAffinityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) StrongSessionAffinityCookie() GoogleComputeRegionBackendServiceStrongSessionAffinityCookieOutputReference {
	var returns GoogleComputeRegionBackendServiceStrongSessionAffinityCookieOutputReference
	_jsii_.Get(
		j,
		"strongSessionAffinityCookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) StrongSessionAffinityCookieInput() *GoogleComputeRegionBackendServiceStrongSessionAffinityCookie {
	var returns *GoogleComputeRegionBackendServiceStrongSessionAffinityCookie
	_jsii_.Get(
		j,
		"strongSessionAffinityCookieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) Subsetting() GoogleComputeRegionBackendServiceSubsettingOutputReference {
	var returns GoogleComputeRegionBackendServiceSubsettingOutputReference
	_jsii_.Get(
		j,
		"subsetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) SubsettingInput() *GoogleComputeRegionBackendServiceSubsetting {
	var returns *GoogleComputeRegionBackendServiceSubsetting
	_jsii_.Get(
		j,
		"subsettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) Timeouts() GoogleComputeRegionBackendServiceTimeoutsOutputReference {
	var returns GoogleComputeRegionBackendServiceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) TimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) TimeoutSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) TlsSettings() GoogleComputeRegionBackendServiceTlsSettingsOutputReference {
	var returns GoogleComputeRegionBackendServiceTlsSettingsOutputReference
	_jsii_.Get(
		j,
		"tlsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendService) TlsSettingsInput() *GoogleComputeRegionBackendServiceTlsSettings {
	var returns *GoogleComputeRegionBackendServiceTlsSettings
	_jsii_.Get(
		j,
		"tlsSettingsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_region_backend_service google_compute_region_backend_service} Resource.
func NewGoogleComputeRegionBackendService(scope constructs.Construct, id *string, config *GoogleComputeRegionBackendServiceConfig) GoogleComputeRegionBackendService {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionBackendServiceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionBackendService{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeRegionBackendService.GoogleComputeRegionBackendService",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_compute_region_backend_service google_compute_region_backend_service} Resource.
func NewGoogleComputeRegionBackendService_Override(g GoogleComputeRegionBackendService, scope constructs.Construct, id *string, config *GoogleComputeRegionBackendServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeRegionBackendService.GoogleComputeRegionBackendService",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendService)SetAffinityCookieTtlSec(val *float64) {
	if err := j.validateSetAffinityCookieTtlSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"affinityCookieTtlSec",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendService)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendService)SetConnectionDrainingTimeoutSec(val *float64) {
	if err := j.validateSetConnectionDrainingTimeoutSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionDrainingTimeoutSec",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendService)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendService)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendService)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendService)SetEnableCdn(val interface{}) {
	if err := j.validateSetEnableCdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCdn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendService)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendService)SetHealthChecks(val *[]*string) {
	if err := j.validateSetHealthChecksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthChecks",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendService)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendService)SetIpAddressSelectionPolicy(val *string) {
	if err := j.validateSetIpAddressSelectionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddressSelectionPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendService)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendService)SetLoadBalancingScheme(val *string) {
	if err := j.validateSetLoadBalancingSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancingScheme",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendService)SetLocalityLbPolicy(val *string) {
	if err := j.validateSetLocalityLbPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localityLbPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendService)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendService)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendService)SetPortName(val *string) {
	if err := j.validateSetPortNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portName",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendService)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendService)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendService)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendService)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendService)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendService)SetSecurityPolicy(val *string) {
	if err := j.validateSetSecurityPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendService)SetSessionAffinity(val *string) {
	if err := j.validateSetSessionAffinityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionAffinity",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendService)SetTimeoutSec(val *float64) {
	if err := j.validateSetTimeoutSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSec",
		val,
	)
}

// Generates CDKTN code for importing a GoogleComputeRegionBackendService resource upon running "cdktn plan <stack-name>".
func GoogleComputeRegionBackendService_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleComputeRegionBackendService_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleComputeRegionBackendService.GoogleComputeRegionBackendService",
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
func GoogleComputeRegionBackendService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeRegionBackendService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleComputeRegionBackendService.GoogleComputeRegionBackendService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeRegionBackendService_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeRegionBackendService_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleComputeRegionBackendService.GoogleComputeRegionBackendService",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeRegionBackendService_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeRegionBackendService_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleComputeRegionBackendService.GoogleComputeRegionBackendService",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputeRegionBackendService_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google-beta.googleComputeRegionBackendService.GoogleComputeRegionBackendService",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionBackendService) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionBackendService) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionBackendService) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionBackendService) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionBackendService) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionBackendService) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionBackendService) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionBackendService) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionBackendService) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionBackendService) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionBackendService) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) PutBackend(value interface{}) {
	if err := g.validatePutBackendParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBackend",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) PutCdnPolicy(value *GoogleComputeRegionBackendServiceCdnPolicy) {
	if err := g.validatePutCdnPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCdnPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) PutCircuitBreakers(value *GoogleComputeRegionBackendServiceCircuitBreakers) {
	if err := g.validatePutCircuitBreakersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCircuitBreakers",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) PutConnectionTrackingPolicy(value *GoogleComputeRegionBackendServiceConnectionTrackingPolicy) {
	if err := g.validatePutConnectionTrackingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConnectionTrackingPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) PutConsistentHash(value *GoogleComputeRegionBackendServiceConsistentHash) {
	if err := g.validatePutConsistentHashParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConsistentHash",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) PutCustomMetrics(value interface{}) {
	if err := g.validatePutCustomMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCustomMetrics",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) PutDynamicForwarding(value *GoogleComputeRegionBackendServiceDynamicForwarding) {
	if err := g.validatePutDynamicForwardingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDynamicForwarding",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) PutFailoverPolicy(value *GoogleComputeRegionBackendServiceFailoverPolicy) {
	if err := g.validatePutFailoverPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFailoverPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) PutHaPolicy(value *GoogleComputeRegionBackendServiceHaPolicy) {
	if err := g.validatePutHaPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHaPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) PutIap(value *GoogleComputeRegionBackendServiceIap) {
	if err := g.validatePutIapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIap",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) PutLogConfig(value *GoogleComputeRegionBackendServiceLogConfig) {
	if err := g.validatePutLogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLogConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) PutNetworkPassThroughLbTrafficPolicy(value *GoogleComputeRegionBackendServiceNetworkPassThroughLbTrafficPolicy) {
	if err := g.validatePutNetworkPassThroughLbTrafficPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkPassThroughLbTrafficPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) PutOutlierDetection(value *GoogleComputeRegionBackendServiceOutlierDetection) {
	if err := g.validatePutOutlierDetectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOutlierDetection",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) PutParams(value *GoogleComputeRegionBackendServiceParams) {
	if err := g.validatePutParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putParams",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) PutStrongSessionAffinityCookie(value *GoogleComputeRegionBackendServiceStrongSessionAffinityCookie) {
	if err := g.validatePutStrongSessionAffinityCookieParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStrongSessionAffinityCookie",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) PutSubsetting(value *GoogleComputeRegionBackendServiceSubsetting) {
	if err := g.validatePutSubsettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSubsetting",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) PutTimeouts(value *GoogleComputeRegionBackendServiceTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) PutTlsSettings(value *GoogleComputeRegionBackendServiceTlsSettings) {
	if err := g.validatePutTlsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTlsSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetAffinityCookieTtlSec() {
	_jsii_.InvokeVoid(
		g,
		"resetAffinityCookieTtlSec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetBackend() {
	_jsii_.InvokeVoid(
		g,
		"resetBackend",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetCdnPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetCdnPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetCircuitBreakers() {
	_jsii_.InvokeVoid(
		g,
		"resetCircuitBreakers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetConnectionDrainingTimeoutSec() {
	_jsii_.InvokeVoid(
		g,
		"resetConnectionDrainingTimeoutSec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetConnectionTrackingPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetConnectionTrackingPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetConsistentHash() {
	_jsii_.InvokeVoid(
		g,
		"resetConsistentHash",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetCustomMetrics() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomMetrics",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetDynamicForwarding() {
	_jsii_.InvokeVoid(
		g,
		"resetDynamicForwarding",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetEnableCdn() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableCdn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetFailoverPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetFailoverPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetHaPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetHaPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetHealthChecks() {
	_jsii_.InvokeVoid(
		g,
		"resetHealthChecks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetIap() {
	_jsii_.InvokeVoid(
		g,
		"resetIap",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetIpAddressSelectionPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetIpAddressSelectionPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetLoadBalancingScheme() {
	_jsii_.InvokeVoid(
		g,
		"resetLoadBalancingScheme",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetLocalityLbPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetLocalityLbPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetLogConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLogConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetNetworkPassThroughLbTrafficPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkPassThroughLbTrafficPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetOutlierDetection() {
	_jsii_.InvokeVoid(
		g,
		"resetOutlierDetection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetParams() {
	_jsii_.InvokeVoid(
		g,
		"resetParams",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetPortName() {
	_jsii_.InvokeVoid(
		g,
		"resetPortName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetProtocol() {
	_jsii_.InvokeVoid(
		g,
		"resetProtocol",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetSecurityPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetSessionAffinity() {
	_jsii_.InvokeVoid(
		g,
		"resetSessionAffinity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetStrongSessionAffinityCookie() {
	_jsii_.InvokeVoid(
		g,
		"resetStrongSessionAffinityCookie",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetSubsetting() {
	_jsii_.InvokeVoid(
		g,
		"resetSubsetting",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetTimeoutSec() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeoutSec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ResetTlsSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetTlsSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionBackendService) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

