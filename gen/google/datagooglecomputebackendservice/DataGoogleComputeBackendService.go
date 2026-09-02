package datagooglecomputebackendservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/datagooglecomputebackendservice/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/data-sources/compute_backend_service google_compute_backend_service}.
type DataGoogleComputeBackendService interface {
	cdktn.TerraformDataSource
	AffinityCookieTtlSec() *float64
	Backend() DataGoogleComputeBackendServiceBackendList
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CdnPolicy() DataGoogleComputeBackendServiceCdnPolicyList
	CircuitBreakers() DataGoogleComputeBackendServiceCircuitBreakersList
	CompressionMode() *string
	ConnectionDrainingTimeoutSec() *float64
	ConsistentHash() DataGoogleComputeBackendServiceConsistentHashList
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreationTimestamp() *string
	CustomMetrics() DataGoogleComputeBackendServiceCustomMetricsList
	CustomRequestHeaders() *[]*string
	CustomResponseHeaders() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	EdgeSecurityPolicy() *string
	EnableCdn() cdktn.IResolvable
	ExternalManagedMigrationState() *string
	ExternalManagedMigrationTestingPercentage() *float64
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
	HealthChecks() *[]*string
	Iap() DataGoogleComputeBackendServiceIapList
	Id() *string
	SetId(val *string)
	IdInput() *string
	IpAddressSelectionPolicy() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LoadBalancingScheme() *string
	LocalityLbPolicies() DataGoogleComputeBackendServiceLocalityLbPoliciesList
	LocalityLbPolicy() *string
	LogConfig() DataGoogleComputeBackendServiceLogConfigList
	MaxStreamDuration() DataGoogleComputeBackendServiceMaxStreamDurationList
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OutlierDetection() DataGoogleComputeBackendServiceOutlierDetectionList
	Params() DataGoogleComputeBackendServiceParamsList
	PortName() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	Protocol() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	SecurityPolicy() *string
	SecuritySettings() DataGoogleComputeBackendServiceSecuritySettingsList
	SelfLink() *string
	ServiceLbPolicy() *string
	SessionAffinity() *string
	StrongSessionAffinityCookie() DataGoogleComputeBackendServiceStrongSessionAffinityCookieList
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeoutSec() *float64
	TlsSettings() DataGoogleComputeBackendServiceTlsSettingsList
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
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
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
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataGoogleComputeBackendService
type jsiiProxy_DataGoogleComputeBackendService struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataGoogleComputeBackendService) AffinityCookieTtlSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"affinityCookieTtlSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) Backend() DataGoogleComputeBackendServiceBackendList {
	var returns DataGoogleComputeBackendServiceBackendList
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) CdnPolicy() DataGoogleComputeBackendServiceCdnPolicyList {
	var returns DataGoogleComputeBackendServiceCdnPolicyList
	_jsii_.Get(
		j,
		"cdnPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) CircuitBreakers() DataGoogleComputeBackendServiceCircuitBreakersList {
	var returns DataGoogleComputeBackendServiceCircuitBreakersList
	_jsii_.Get(
		j,
		"circuitBreakers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) CompressionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) ConnectionDrainingTimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionDrainingTimeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) ConsistentHash() DataGoogleComputeBackendServiceConsistentHashList {
	var returns DataGoogleComputeBackendServiceConsistentHashList
	_jsii_.Get(
		j,
		"consistentHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) CustomMetrics() DataGoogleComputeBackendServiceCustomMetricsList {
	var returns DataGoogleComputeBackendServiceCustomMetricsList
	_jsii_.Get(
		j,
		"customMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) CustomRequestHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customRequestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) CustomResponseHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customResponseHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) EdgeSecurityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeSecurityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) EnableCdn() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"enableCdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) ExternalManagedMigrationState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalManagedMigrationState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) ExternalManagedMigrationTestingPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"externalManagedMigrationTestingPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) Fingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) GeneratedId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"generatedId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) HealthChecks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"healthChecks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) Iap() DataGoogleComputeBackendServiceIapList {
	var returns DataGoogleComputeBackendServiceIapList
	_jsii_.Get(
		j,
		"iap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) IpAddressSelectionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressSelectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) LoadBalancingScheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) LocalityLbPolicies() DataGoogleComputeBackendServiceLocalityLbPoliciesList {
	var returns DataGoogleComputeBackendServiceLocalityLbPoliciesList
	_jsii_.Get(
		j,
		"localityLbPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) LocalityLbPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localityLbPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) LogConfig() DataGoogleComputeBackendServiceLogConfigList {
	var returns DataGoogleComputeBackendServiceLogConfigList
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) MaxStreamDuration() DataGoogleComputeBackendServiceMaxStreamDurationList {
	var returns DataGoogleComputeBackendServiceMaxStreamDurationList
	_jsii_.Get(
		j,
		"maxStreamDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) OutlierDetection() DataGoogleComputeBackendServiceOutlierDetectionList {
	var returns DataGoogleComputeBackendServiceOutlierDetectionList
	_jsii_.Get(
		j,
		"outlierDetection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) Params() DataGoogleComputeBackendServiceParamsList {
	var returns DataGoogleComputeBackendServiceParamsList
	_jsii_.Get(
		j,
		"params",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) PortName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) SecurityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) SecuritySettings() DataGoogleComputeBackendServiceSecuritySettingsList {
	var returns DataGoogleComputeBackendServiceSecuritySettingsList
	_jsii_.Get(
		j,
		"securitySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) ServiceLbPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLbPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) SessionAffinity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) StrongSessionAffinityCookie() DataGoogleComputeBackendServiceStrongSessionAffinityCookieList {
	var returns DataGoogleComputeBackendServiceStrongSessionAffinityCookieList
	_jsii_.Get(
		j,
		"strongSessionAffinityCookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) TimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeBackendService) TlsSettings() DataGoogleComputeBackendServiceTlsSettingsList {
	var returns DataGoogleComputeBackendServiceTlsSettingsList
	_jsii_.Get(
		j,
		"tlsSettings",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/data-sources/compute_backend_service google_compute_backend_service} Data Source.
func NewDataGoogleComputeBackendService(scope constructs.Construct, id *string, config *DataGoogleComputeBackendServiceConfig) DataGoogleComputeBackendService {
	_init_.Initialize()

	if err := validateNewDataGoogleComputeBackendServiceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleComputeBackendService{}

	_jsii_.Create(
		"@cdktn/provider-google.dataGoogleComputeBackendService.DataGoogleComputeBackendService",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/data-sources/compute_backend_service google_compute_backend_service} Data Source.
func NewDataGoogleComputeBackendService_Override(d DataGoogleComputeBackendService, scope constructs.Construct, id *string, config *DataGoogleComputeBackendServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dataGoogleComputeBackendService.DataGoogleComputeBackendService",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataGoogleComputeBackendService)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeBackendService)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeBackendService)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeBackendService)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeBackendService)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeBackendService)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeBackendService)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeBackendService)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTN code for importing a DataGoogleComputeBackendService resource upon running "cdktn plan <stack-name>".
func DataGoogleComputeBackendService_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataGoogleComputeBackendService_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.dataGoogleComputeBackendService.DataGoogleComputeBackendService",
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
func DataGoogleComputeBackendService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleComputeBackendService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.dataGoogleComputeBackendService.DataGoogleComputeBackendService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGoogleComputeBackendService_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleComputeBackendService_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.dataGoogleComputeBackendService.DataGoogleComputeBackendService",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGoogleComputeBackendService_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleComputeBackendService_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.dataGoogleComputeBackendService.DataGoogleComputeBackendService",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataGoogleComputeBackendService_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.dataGoogleComputeBackendService.DataGoogleComputeBackendService",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataGoogleComputeBackendService) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataGoogleComputeBackendService) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeBackendService) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeBackendService) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeBackendService) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeBackendService) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeBackendService) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeBackendService) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeBackendService) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeBackendService) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeBackendService) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeBackendService) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataGoogleComputeBackendService) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataGoogleComputeBackendService) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleComputeBackendService) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleComputeBackendService) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleComputeBackendService) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeBackendService) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeBackendService) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeBackendService) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeBackendService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeBackendService) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeBackendService) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

