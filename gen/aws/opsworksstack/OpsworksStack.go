package opsworksstack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/aws/opsworksstack/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/opsworks_stack aws_opsworks_stack}.
type OpsworksStack interface {
	cdktn.TerraformResource
	AgentVersion() *string
	SetAgentVersion(val *string)
	AgentVersionInput() *string
	Arn() *string
	BerkshelfVersion() *string
	SetBerkshelfVersion(val *string)
	BerkshelfVersionInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Color() *string
	SetColor(val *string)
	ColorInput() *string
	ConfigurationManagerName() *string
	SetConfigurationManagerName(val *string)
	ConfigurationManagerNameInput() *string
	ConfigurationManagerVersion() *string
	SetConfigurationManagerVersion(val *string)
	ConfigurationManagerVersionInput() *string
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
	CustomCookbooksSource() OpsworksStackCustomCookbooksSourceOutputReference
	CustomCookbooksSourceInput() *OpsworksStackCustomCookbooksSource
	CustomJson() *string
	SetCustomJson(val *string)
	CustomJsonInput() *string
	DefaultAvailabilityZone() *string
	SetDefaultAvailabilityZone(val *string)
	DefaultAvailabilityZoneInput() *string
	DefaultInstanceProfileArn() *string
	SetDefaultInstanceProfileArn(val *string)
	DefaultInstanceProfileArnInput() *string
	DefaultOs() *string
	SetDefaultOs(val *string)
	DefaultOsInput() *string
	DefaultRootDeviceType() *string
	SetDefaultRootDeviceType(val *string)
	DefaultRootDeviceTypeInput() *string
	DefaultSshKeyName() *string
	SetDefaultSshKeyName(val *string)
	DefaultSshKeyNameInput() *string
	DefaultSubnetId() *string
	SetDefaultSubnetId(val *string)
	DefaultSubnetIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HostnameTheme() *string
	SetHostnameTheme(val *string)
	HostnameThemeInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	ManageBerkshelf() interface{}
	SetManageBerkshelf(val interface{})
	ManageBerkshelfInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
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
	ServiceRoleArn() *string
	SetServiceRoleArn(val *string)
	ServiceRoleArnInput() *string
	StackEndpoint() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() OpsworksStackTimeoutsOutputReference
	TimeoutsInput() interface{}
	UseCustomCookbooks() interface{}
	SetUseCustomCookbooks(val interface{})
	UseCustomCookbooksInput() interface{}
	UseOpsworksSecurityGroups() interface{}
	SetUseOpsworksSecurityGroups(val interface{})
	UseOpsworksSecurityGroupsInput() interface{}
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
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
	PutCustomCookbooksSource(value *OpsworksStackCustomCookbooksSource)
	PutTimeouts(value *OpsworksStackTimeouts)
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
	ResetAgentVersion()
	ResetBerkshelfVersion()
	ResetColor()
	ResetConfigurationManagerName()
	ResetConfigurationManagerVersion()
	ResetCustomCookbooksSource()
	ResetCustomJson()
	ResetDefaultAvailabilityZone()
	ResetDefaultOs()
	ResetDefaultRootDeviceType()
	ResetDefaultSshKeyName()
	ResetDefaultSubnetId()
	ResetHostnameTheme()
	ResetId()
	ResetManageBerkshelf()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetUseCustomCookbooks()
	ResetUseOpsworksSecurityGroups()
	ResetVpcId()
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

// The jsii proxy struct for OpsworksStack
type jsiiProxy_OpsworksStack struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_OpsworksStack) AgentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) AgentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) BerkshelfVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"berkshelfVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) BerkshelfVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"berkshelfVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) Color() *string {
	var returns *string
	_jsii_.Get(
		j,
		"color",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) ColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) ConfigurationManagerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationManagerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) ConfigurationManagerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationManagerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) ConfigurationManagerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationManagerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) ConfigurationManagerVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationManagerVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) CustomCookbooksSource() OpsworksStackCustomCookbooksSourceOutputReference {
	var returns OpsworksStackCustomCookbooksSourceOutputReference
	_jsii_.Get(
		j,
		"customCookbooksSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) CustomCookbooksSourceInput() *OpsworksStackCustomCookbooksSource {
	var returns *OpsworksStackCustomCookbooksSource
	_jsii_.Get(
		j,
		"customCookbooksSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) CustomJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) CustomJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DefaultAvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DefaultAvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAvailabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DefaultInstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DefaultInstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultInstanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DefaultOs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultOs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DefaultOsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultOsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DefaultRootDeviceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootDeviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DefaultRootDeviceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootDeviceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DefaultSshKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSshKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DefaultSshKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSshKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DefaultSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DefaultSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) HostnameTheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameTheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) HostnameThemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameThemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) ManageBerkshelf() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageBerkshelf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) ManageBerkshelfInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageBerkshelfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) ServiceRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) StackEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) Timeouts() OpsworksStackTimeoutsOutputReference {
	var returns OpsworksStackTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) UseCustomCookbooks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCustomCookbooks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) UseCustomCookbooksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCustomCookbooksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) UseOpsworksSecurityGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useOpsworksSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) UseOpsworksSecurityGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useOpsworksSecurityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksStack) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/opsworks_stack aws_opsworks_stack} Resource.
func NewOpsworksStack(scope constructs.Construct, id *string, config *OpsworksStackConfig) OpsworksStack {
	_init_.Initialize()

	if err := validateNewOpsworksStackParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpsworksStack{}

	_jsii_.Create(
		"@cdktn/provider-aws.opsworksStack.OpsworksStack",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/opsworks_stack aws_opsworks_stack} Resource.
func NewOpsworksStack_Override(o OpsworksStack, scope constructs.Construct, id *string, config *OpsworksStackConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.opsworksStack.OpsworksStack",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpsworksStack)SetAgentVersion(val *string) {
	if err := j.validateSetAgentVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentVersion",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetBerkshelfVersion(val *string) {
	if err := j.validateSetBerkshelfVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"berkshelfVersion",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetColor(val *string) {
	if err := j.validateSetColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"color",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetConfigurationManagerName(val *string) {
	if err := j.validateSetConfigurationManagerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationManagerName",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetConfigurationManagerVersion(val *string) {
	if err := j.validateSetConfigurationManagerVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationManagerVersion",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetCustomJson(val *string) {
	if err := j.validateSetCustomJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customJson",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetDefaultAvailabilityZone(val *string) {
	if err := j.validateSetDefaultAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultAvailabilityZone",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetDefaultInstanceProfileArn(val *string) {
	if err := j.validateSetDefaultInstanceProfileArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultInstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetDefaultOs(val *string) {
	if err := j.validateSetDefaultOsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultOs",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetDefaultRootDeviceType(val *string) {
	if err := j.validateSetDefaultRootDeviceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRootDeviceType",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetDefaultSshKeyName(val *string) {
	if err := j.validateSetDefaultSshKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSshKeyName",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetDefaultSubnetId(val *string) {
	if err := j.validateSetDefaultSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSubnetId",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetHostnameTheme(val *string) {
	if err := j.validateSetHostnameThemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostnameTheme",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetManageBerkshelf(val interface{}) {
	if err := j.validateSetManageBerkshelfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageBerkshelf",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetServiceRoleArn(val *string) {
	if err := j.validateSetServiceRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRoleArn",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetUseCustomCookbooks(val interface{}) {
	if err := j.validateSetUseCustomCookbooksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCustomCookbooks",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetUseOpsworksSecurityGroups(val interface{}) {
	if err := j.validateSetUseOpsworksSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useOpsworksSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_OpsworksStack)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTN code for importing a OpsworksStack resource upon running "cdktn plan <stack-name>".
func OpsworksStack_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateOpsworksStack_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.opsworksStack.OpsworksStack",
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
func OpsworksStack_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpsworksStack_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.opsworksStack.OpsworksStack",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OpsworksStack_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpsworksStack_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.opsworksStack.OpsworksStack",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OpsworksStack_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpsworksStack_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.opsworksStack.OpsworksStack",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpsworksStack_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-aws.opsworksStack.OpsworksStack",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OpsworksStack) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OpsworksStack) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OpsworksStack) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStack) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStack) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStack) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStack) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStack) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStack) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStack) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStack) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStack) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStack) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OpsworksStack) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStack) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := o.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStack) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OpsworksStack) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OpsworksStack) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OpsworksStack) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OpsworksStack) PutCustomCookbooksSource(value *OpsworksStackCustomCookbooksSource) {
	if err := o.validatePutCustomCookbooksSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCustomCookbooksSource",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpsworksStack) PutTimeouts(value *OpsworksStackTimeouts) {
	if err := o.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpsworksStack) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := o.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (o *jsiiProxy_OpsworksStack) ResetAgentVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetAgentVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetBerkshelfVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetBerkshelfVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetColor() {
	_jsii_.InvokeVoid(
		o,
		"resetColor",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetConfigurationManagerName() {
	_jsii_.InvokeVoid(
		o,
		"resetConfigurationManagerName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetConfigurationManagerVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetConfigurationManagerVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetCustomCookbooksSource() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomCookbooksSource",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetCustomJson() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomJson",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetDefaultAvailabilityZone() {
	_jsii_.InvokeVoid(
		o,
		"resetDefaultAvailabilityZone",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetDefaultOs() {
	_jsii_.InvokeVoid(
		o,
		"resetDefaultOs",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetDefaultRootDeviceType() {
	_jsii_.InvokeVoid(
		o,
		"resetDefaultRootDeviceType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetDefaultSshKeyName() {
	_jsii_.InvokeVoid(
		o,
		"resetDefaultSshKeyName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetDefaultSubnetId() {
	_jsii_.InvokeVoid(
		o,
		"resetDefaultSubnetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetHostnameTheme() {
	_jsii_.InvokeVoid(
		o,
		"resetHostnameTheme",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetManageBerkshelf() {
	_jsii_.InvokeVoid(
		o,
		"resetManageBerkshelf",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetTagsAll() {
	_jsii_.InvokeVoid(
		o,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetTimeouts() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetUseCustomCookbooks() {
	_jsii_.InvokeVoid(
		o,
		"resetUseCustomCookbooks",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetUseOpsworksSecurityGroups() {
	_jsii_.InvokeVoid(
		o,
		"resetUseOpsworksSecurityGroups",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) ResetVpcId() {
	_jsii_.InvokeVoid(
		o,
		"resetVpcId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksStack) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStack) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStack) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStack) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStack) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStack) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksStack) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		o,
		"with",
		args,
		&returns,
	)

	return returns
}

