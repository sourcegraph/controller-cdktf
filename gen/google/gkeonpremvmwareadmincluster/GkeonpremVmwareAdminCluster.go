package gkeonpremvmwareadmincluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/gkeonpremvmwareadmincluster/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/gkeonprem_vmware_admin_cluster google_gkeonprem_vmware_admin_cluster}.
type GkeonpremVmwareAdminCluster interface {
	cdktn.TerraformResource
	AddonNode() GkeonpremVmwareAdminClusterAddonNodeOutputReference
	AddonNodeInput() *GkeonpremVmwareAdminClusterAddonNode
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	AntiAffinityGroups() GkeonpremVmwareAdminClusterAntiAffinityGroupsOutputReference
	AntiAffinityGroupsInput() *GkeonpremVmwareAdminClusterAntiAffinityGroups
	Authorization() GkeonpremVmwareAdminClusterAuthorizationOutputReference
	AuthorizationInput() *GkeonpremVmwareAdminClusterAuthorization
	AutoRepairConfig() GkeonpremVmwareAdminClusterAutoRepairConfigOutputReference
	AutoRepairConfigInput() *GkeonpremVmwareAdminClusterAutoRepairConfig
	BootstrapClusterMembership() *string
	SetBootstrapClusterMembership(val *string)
	BootstrapClusterMembershipInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ControlPlaneNode() GkeonpremVmwareAdminClusterControlPlaneNodeOutputReference
	ControlPlaneNodeInput() *GkeonpremVmwareAdminClusterControlPlaneNode
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveAnnotations() cdktn.StringMap
	EnableAdvancedCluster() interface{}
	SetEnableAdvancedCluster(val interface{})
	EnableAdvancedClusterInput() interface{}
	Endpoint() *string
	Etag() *string
	Fleet() GkeonpremVmwareAdminClusterFleetList
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImageType() *string
	SetImageType(val *string)
	ImageTypeInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LoadBalancer() GkeonpremVmwareAdminClusterLoadBalancerOutputReference
	LoadBalancerInput() *GkeonpremVmwareAdminClusterLoadBalancer
	LocalName() *string
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkConfig() GkeonpremVmwareAdminClusterNetworkConfigOutputReference
	NetworkConfigInput() *GkeonpremVmwareAdminClusterNetworkConfig
	// The tree node.
	Node() constructs.Node
	OnPremVersion() *string
	SetOnPremVersion(val *string)
	OnPremVersionInput() *string
	PlatformConfig() GkeonpremVmwareAdminClusterPlatformConfigOutputReference
	PlatformConfigInput() *GkeonpremVmwareAdminClusterPlatformConfig
	PrivateRegistryConfig() GkeonpremVmwareAdminClusterPrivateRegistryConfigOutputReference
	PrivateRegistryConfigInput() *GkeonpremVmwareAdminClusterPrivateRegistryConfig
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
	Proxy() GkeonpremVmwareAdminClusterProxyOutputReference
	ProxyInput() *GkeonpremVmwareAdminClusterProxy
	// Experimental.
	RawOverrides() interface{}
	Reconciling() cdktn.IResolvable
	State() *string
	Status() GkeonpremVmwareAdminClusterStatusList
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GkeonpremVmwareAdminClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uid() *string
	UpdateTime() *string
	Vcenter() GkeonpremVmwareAdminClusterVcenterOutputReference
	VcenterInput() *GkeonpremVmwareAdminClusterVcenter
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
	PutAddonNode(value *GkeonpremVmwareAdminClusterAddonNode)
	PutAntiAffinityGroups(value *GkeonpremVmwareAdminClusterAntiAffinityGroups)
	PutAuthorization(value *GkeonpremVmwareAdminClusterAuthorization)
	PutAutoRepairConfig(value *GkeonpremVmwareAdminClusterAutoRepairConfig)
	PutControlPlaneNode(value *GkeonpremVmwareAdminClusterControlPlaneNode)
	PutLoadBalancer(value *GkeonpremVmwareAdminClusterLoadBalancer)
	PutNetworkConfig(value *GkeonpremVmwareAdminClusterNetworkConfig)
	PutPlatformConfig(value *GkeonpremVmwareAdminClusterPlatformConfig)
	PutPrivateRegistryConfig(value *GkeonpremVmwareAdminClusterPrivateRegistryConfig)
	PutProxy(value *GkeonpremVmwareAdminClusterProxy)
	PutTimeouts(value *GkeonpremVmwareAdminClusterTimeouts)
	PutVcenter(value *GkeonpremVmwareAdminClusterVcenter)
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
	ResetAddonNode()
	ResetAnnotations()
	ResetAntiAffinityGroups()
	ResetAuthorization()
	ResetAutoRepairConfig()
	ResetBootstrapClusterMembership()
	ResetControlPlaneNode()
	ResetDescription()
	ResetEnableAdvancedCluster()
	ResetId()
	ResetImageType()
	ResetLoadBalancer()
	ResetOnPremVersion()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlatformConfig()
	ResetPrivateRegistryConfig()
	ResetProject()
	ResetProxy()
	ResetTimeouts()
	ResetVcenter()
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

// The jsii proxy struct for GkeonpremVmwareAdminCluster
type jsiiProxy_GkeonpremVmwareAdminCluster struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) AddonNode() GkeonpremVmwareAdminClusterAddonNodeOutputReference {
	var returns GkeonpremVmwareAdminClusterAddonNodeOutputReference
	_jsii_.Get(
		j,
		"addonNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) AddonNodeInput() *GkeonpremVmwareAdminClusterAddonNode {
	var returns *GkeonpremVmwareAdminClusterAddonNode
	_jsii_.Get(
		j,
		"addonNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) AntiAffinityGroups() GkeonpremVmwareAdminClusterAntiAffinityGroupsOutputReference {
	var returns GkeonpremVmwareAdminClusterAntiAffinityGroupsOutputReference
	_jsii_.Get(
		j,
		"antiAffinityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) AntiAffinityGroupsInput() *GkeonpremVmwareAdminClusterAntiAffinityGroups {
	var returns *GkeonpremVmwareAdminClusterAntiAffinityGroups
	_jsii_.Get(
		j,
		"antiAffinityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) Authorization() GkeonpremVmwareAdminClusterAuthorizationOutputReference {
	var returns GkeonpremVmwareAdminClusterAuthorizationOutputReference
	_jsii_.Get(
		j,
		"authorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) AuthorizationInput() *GkeonpremVmwareAdminClusterAuthorization {
	var returns *GkeonpremVmwareAdminClusterAuthorization
	_jsii_.Get(
		j,
		"authorizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) AutoRepairConfig() GkeonpremVmwareAdminClusterAutoRepairConfigOutputReference {
	var returns GkeonpremVmwareAdminClusterAutoRepairConfigOutputReference
	_jsii_.Get(
		j,
		"autoRepairConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) AutoRepairConfigInput() *GkeonpremVmwareAdminClusterAutoRepairConfig {
	var returns *GkeonpremVmwareAdminClusterAutoRepairConfig
	_jsii_.Get(
		j,
		"autoRepairConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) BootstrapClusterMembership() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapClusterMembership",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) BootstrapClusterMembershipInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapClusterMembershipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) ControlPlaneNode() GkeonpremVmwareAdminClusterControlPlaneNodeOutputReference {
	var returns GkeonpremVmwareAdminClusterControlPlaneNodeOutputReference
	_jsii_.Get(
		j,
		"controlPlaneNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) ControlPlaneNodeInput() *GkeonpremVmwareAdminClusterControlPlaneNode {
	var returns *GkeonpremVmwareAdminClusterControlPlaneNode
	_jsii_.Get(
		j,
		"controlPlaneNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) EffectiveAnnotations() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"effectiveAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) EnableAdvancedCluster() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAdvancedCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) EnableAdvancedClusterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAdvancedClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) Fleet() GkeonpremVmwareAdminClusterFleetList {
	var returns GkeonpremVmwareAdminClusterFleetList
	_jsii_.Get(
		j,
		"fleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) ImageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) ImageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) LoadBalancer() GkeonpremVmwareAdminClusterLoadBalancerOutputReference {
	var returns GkeonpremVmwareAdminClusterLoadBalancerOutputReference
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) LoadBalancerInput() *GkeonpremVmwareAdminClusterLoadBalancer {
	var returns *GkeonpremVmwareAdminClusterLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) LocalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) NetworkConfig() GkeonpremVmwareAdminClusterNetworkConfigOutputReference {
	var returns GkeonpremVmwareAdminClusterNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) NetworkConfigInput() *GkeonpremVmwareAdminClusterNetworkConfig {
	var returns *GkeonpremVmwareAdminClusterNetworkConfig
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) OnPremVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onPremVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) OnPremVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onPremVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) PlatformConfig() GkeonpremVmwareAdminClusterPlatformConfigOutputReference {
	var returns GkeonpremVmwareAdminClusterPlatformConfigOutputReference
	_jsii_.Get(
		j,
		"platformConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) PlatformConfigInput() *GkeonpremVmwareAdminClusterPlatformConfig {
	var returns *GkeonpremVmwareAdminClusterPlatformConfig
	_jsii_.Get(
		j,
		"platformConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) PrivateRegistryConfig() GkeonpremVmwareAdminClusterPrivateRegistryConfigOutputReference {
	var returns GkeonpremVmwareAdminClusterPrivateRegistryConfigOutputReference
	_jsii_.Get(
		j,
		"privateRegistryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) PrivateRegistryConfigInput() *GkeonpremVmwareAdminClusterPrivateRegistryConfig {
	var returns *GkeonpremVmwareAdminClusterPrivateRegistryConfig
	_jsii_.Get(
		j,
		"privateRegistryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) Proxy() GkeonpremVmwareAdminClusterProxyOutputReference {
	var returns GkeonpremVmwareAdminClusterProxyOutputReference
	_jsii_.Get(
		j,
		"proxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) ProxyInput() *GkeonpremVmwareAdminClusterProxy {
	var returns *GkeonpremVmwareAdminClusterProxy
	_jsii_.Get(
		j,
		"proxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) Reconciling() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"reconciling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) Status() GkeonpremVmwareAdminClusterStatusList {
	var returns GkeonpremVmwareAdminClusterStatusList
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) Timeouts() GkeonpremVmwareAdminClusterTimeoutsOutputReference {
	var returns GkeonpremVmwareAdminClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) Vcenter() GkeonpremVmwareAdminClusterVcenterOutputReference {
	var returns GkeonpremVmwareAdminClusterVcenterOutputReference
	_jsii_.Get(
		j,
		"vcenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster) VcenterInput() *GkeonpremVmwareAdminClusterVcenter {
	var returns *GkeonpremVmwareAdminClusterVcenter
	_jsii_.Get(
		j,
		"vcenterInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/gkeonprem_vmware_admin_cluster google_gkeonprem_vmware_admin_cluster} Resource.
func NewGkeonpremVmwareAdminCluster(scope constructs.Construct, id *string, config *GkeonpremVmwareAdminClusterConfig) GkeonpremVmwareAdminCluster {
	_init_.Initialize()

	if err := validateNewGkeonpremVmwareAdminClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeonpremVmwareAdminCluster{}

	_jsii_.Create(
		"@cdktn/provider-google.gkeonpremVmwareAdminCluster.GkeonpremVmwareAdminCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/gkeonprem_vmware_admin_cluster google_gkeonprem_vmware_admin_cluster} Resource.
func NewGkeonpremVmwareAdminCluster_Override(g GkeonpremVmwareAdminCluster, scope constructs.Construct, id *string, config *GkeonpremVmwareAdminClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.gkeonpremVmwareAdminCluster.GkeonpremVmwareAdminCluster",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster)SetBootstrapClusterMembership(val *string) {
	if err := j.validateSetBootstrapClusterMembershipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootstrapClusterMembership",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster)SetEnableAdvancedCluster(val interface{}) {
	if err := j.validateSetEnableAdvancedClusterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAdvancedCluster",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster)SetImageType(val *string) {
	if err := j.validateSetImageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageType",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster)SetOnPremVersion(val *string) {
	if err := j.validateSetOnPremVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onPremVersion",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a GkeonpremVmwareAdminCluster resource upon running "cdktn plan <stack-name>".
func GkeonpremVmwareAdminCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGkeonpremVmwareAdminCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.gkeonpremVmwareAdminCluster.GkeonpremVmwareAdminCluster",
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
func GkeonpremVmwareAdminCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGkeonpremVmwareAdminCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.gkeonpremVmwareAdminCluster.GkeonpremVmwareAdminCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GkeonpremVmwareAdminCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGkeonpremVmwareAdminCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.gkeonpremVmwareAdminCluster.GkeonpremVmwareAdminCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GkeonpremVmwareAdminCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGkeonpremVmwareAdminCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.gkeonpremVmwareAdminCluster.GkeonpremVmwareAdminCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GkeonpremVmwareAdminCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.gkeonpremVmwareAdminCluster.GkeonpremVmwareAdminCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) PutAddonNode(value *GkeonpremVmwareAdminClusterAddonNode) {
	if err := g.validatePutAddonNodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAddonNode",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) PutAntiAffinityGroups(value *GkeonpremVmwareAdminClusterAntiAffinityGroups) {
	if err := g.validatePutAntiAffinityGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAntiAffinityGroups",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) PutAuthorization(value *GkeonpremVmwareAdminClusterAuthorization) {
	if err := g.validatePutAuthorizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuthorization",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) PutAutoRepairConfig(value *GkeonpremVmwareAdminClusterAutoRepairConfig) {
	if err := g.validatePutAutoRepairConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutoRepairConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) PutControlPlaneNode(value *GkeonpremVmwareAdminClusterControlPlaneNode) {
	if err := g.validatePutControlPlaneNodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putControlPlaneNode",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) PutLoadBalancer(value *GkeonpremVmwareAdminClusterLoadBalancer) {
	if err := g.validatePutLoadBalancerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLoadBalancer",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) PutNetworkConfig(value *GkeonpremVmwareAdminClusterNetworkConfig) {
	if err := g.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) PutPlatformConfig(value *GkeonpremVmwareAdminClusterPlatformConfig) {
	if err := g.validatePutPlatformConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPlatformConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) PutPrivateRegistryConfig(value *GkeonpremVmwareAdminClusterPrivateRegistryConfig) {
	if err := g.validatePutPrivateRegistryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPrivateRegistryConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) PutProxy(value *GkeonpremVmwareAdminClusterProxy) {
	if err := g.validatePutProxyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putProxy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) PutTimeouts(value *GkeonpremVmwareAdminClusterTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) PutVcenter(value *GkeonpremVmwareAdminClusterVcenter) {
	if err := g.validatePutVcenterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVcenter",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) ResetAddonNode() {
	_jsii_.InvokeVoid(
		g,
		"resetAddonNode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) ResetAnnotations() {
	_jsii_.InvokeVoid(
		g,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) ResetAntiAffinityGroups() {
	_jsii_.InvokeVoid(
		g,
		"resetAntiAffinityGroups",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) ResetAuthorization() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthorization",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) ResetAutoRepairConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoRepairConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) ResetBootstrapClusterMembership() {
	_jsii_.InvokeVoid(
		g,
		"resetBootstrapClusterMembership",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) ResetControlPlaneNode() {
	_jsii_.InvokeVoid(
		g,
		"resetControlPlaneNode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) ResetEnableAdvancedCluster() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableAdvancedCluster",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) ResetImageType() {
	_jsii_.InvokeVoid(
		g,
		"resetImageType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) ResetLoadBalancer() {
	_jsii_.InvokeVoid(
		g,
		"resetLoadBalancer",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) ResetOnPremVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetOnPremVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) ResetPlatformConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPlatformConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) ResetPrivateRegistryConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivateRegistryConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) ResetProxy() {
	_jsii_.InvokeVoid(
		g,
		"resetProxy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) ResetVcenter() {
	_jsii_.InvokeVoid(
		g,
		"resetVcenter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareAdminCluster) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

