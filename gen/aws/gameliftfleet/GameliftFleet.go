package gameliftfleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/aws/gameliftfleet/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/gamelift_fleet aws_gamelift_fleet}.
type GameliftFleet interface {
	cdktn.TerraformResource
	Arn() *string
	BuildArn() *string
	BuildId() *string
	SetBuildId(val *string)
	BuildIdInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CertificateConfiguration() GameliftFleetCertificateConfigurationOutputReference
	CertificateConfigurationInput() *GameliftFleetCertificateConfiguration
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Ec2InboundPermission() GameliftFleetEc2InboundPermissionList
	Ec2InboundPermissionInput() interface{}
	Ec2InstanceType() *string
	SetEc2InstanceType(val *string)
	Ec2InstanceTypeInput() *string
	FleetType() *string
	SetFleetType(val *string)
	FleetTypeInput() *string
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
	InstanceRoleArn() *string
	SetInstanceRoleArn(val *string)
	InstanceRoleArnInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LogPaths() *[]*string
	MetricGroups() *[]*string
	SetMetricGroups(val *[]*string)
	MetricGroupsInput() *[]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NewGameSessionProtectionPolicy() *string
	SetNewGameSessionProtectionPolicy(val *string)
	NewGameSessionProtectionPolicyInput() *string
	// The tree node.
	Node() constructs.Node
	OperatingSystem() *string
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
	ResourceCreationLimitPolicy() GameliftFleetResourceCreationLimitPolicyOutputReference
	ResourceCreationLimitPolicyInput() *GameliftFleetResourceCreationLimitPolicy
	RuntimeConfiguration() GameliftFleetRuntimeConfigurationOutputReference
	RuntimeConfigurationInput() *GameliftFleetRuntimeConfiguration
	ScriptArn() *string
	ScriptId() *string
	SetScriptId(val *string)
	ScriptIdInput() *string
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
	Timeouts() GameliftFleetTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutCertificateConfiguration(value *GameliftFleetCertificateConfiguration)
	PutEc2InboundPermission(value interface{})
	PutResourceCreationLimitPolicy(value *GameliftFleetResourceCreationLimitPolicy)
	PutRuntimeConfiguration(value *GameliftFleetRuntimeConfiguration)
	PutTimeouts(value *GameliftFleetTimeouts)
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
	ResetBuildId()
	ResetCertificateConfiguration()
	ResetDescription()
	ResetEc2InboundPermission()
	ResetFleetType()
	ResetId()
	ResetInstanceRoleArn()
	ResetMetricGroups()
	ResetNewGameSessionProtectionPolicy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetResourceCreationLimitPolicy()
	ResetRuntimeConfiguration()
	ResetScriptId()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
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

// The jsii proxy struct for GameliftFleet
type jsiiProxy_GameliftFleet struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GameliftFleet) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) BuildArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) BuildId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) BuildIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) CertificateConfiguration() GameliftFleetCertificateConfigurationOutputReference {
	var returns GameliftFleetCertificateConfigurationOutputReference
	_jsii_.Get(
		j,
		"certificateConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) CertificateConfigurationInput() *GameliftFleetCertificateConfiguration {
	var returns *GameliftFleetCertificateConfiguration
	_jsii_.Get(
		j,
		"certificateConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Ec2InboundPermission() GameliftFleetEc2InboundPermissionList {
	var returns GameliftFleetEc2InboundPermissionList
	_jsii_.Get(
		j,
		"ec2InboundPermission",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Ec2InboundPermissionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2InboundPermissionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Ec2InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2InstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Ec2InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2InstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) FleetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) FleetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) InstanceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) InstanceRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) LogPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) MetricGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metricGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) MetricGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metricGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) NewGameSessionProtectionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newGameSessionProtectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) NewGameSessionProtectionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newGameSessionProtectionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) OperatingSystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ResourceCreationLimitPolicy() GameliftFleetResourceCreationLimitPolicyOutputReference {
	var returns GameliftFleetResourceCreationLimitPolicyOutputReference
	_jsii_.Get(
		j,
		"resourceCreationLimitPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ResourceCreationLimitPolicyInput() *GameliftFleetResourceCreationLimitPolicy {
	var returns *GameliftFleetResourceCreationLimitPolicy
	_jsii_.Get(
		j,
		"resourceCreationLimitPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) RuntimeConfiguration() GameliftFleetRuntimeConfigurationOutputReference {
	var returns GameliftFleetRuntimeConfigurationOutputReference
	_jsii_.Get(
		j,
		"runtimeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) RuntimeConfigurationInput() *GameliftFleetRuntimeConfiguration {
	var returns *GameliftFleetRuntimeConfiguration
	_jsii_.Get(
		j,
		"runtimeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ScriptArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ScriptId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) ScriptIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) Timeouts() GameliftFleetTimeoutsOutputReference {
	var returns GameliftFleetTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameliftFleet) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/gamelift_fleet aws_gamelift_fleet} Resource.
func NewGameliftFleet(scope constructs.Construct, id *string, config *GameliftFleetConfig) GameliftFleet {
	_init_.Initialize()

	if err := validateNewGameliftFleetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GameliftFleet{}

	_jsii_.Create(
		"@cdktn/provider-aws.gameliftFleet.GameliftFleet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/gamelift_fleet aws_gamelift_fleet} Resource.
func NewGameliftFleet_Override(g GameliftFleet, scope constructs.Construct, id *string, config *GameliftFleetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.gameliftFleet.GameliftFleet",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GameliftFleet)SetBuildId(val *string) {
	if err := j.validateSetBuildIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildId",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetEc2InstanceType(val *string) {
	if err := j.validateSetEc2InstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ec2InstanceType",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetFleetType(val *string) {
	if err := j.validateSetFleetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fleetType",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetInstanceRoleArn(val *string) {
	if err := j.validateSetInstanceRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceRoleArn",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetMetricGroups(val *[]*string) {
	if err := j.validateSetMetricGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricGroups",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetNewGameSessionProtectionPolicy(val *string) {
	if err := j.validateSetNewGameSessionProtectionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newGameSessionProtectionPolicy",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetScriptId(val *string) {
	if err := j.validateSetScriptIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scriptId",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GameliftFleet)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a GameliftFleet resource upon running "cdktn plan <stack-name>".
func GameliftFleet_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGameliftFleet_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.gameliftFleet.GameliftFleet",
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
func GameliftFleet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftFleet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.gameliftFleet.GameliftFleet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GameliftFleet_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftFleet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.gameliftFleet.GameliftFleet",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GameliftFleet_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameliftFleet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.gameliftFleet.GameliftFleet",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GameliftFleet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-aws.gameliftFleet.GameliftFleet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GameliftFleet) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GameliftFleet) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GameliftFleet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GameliftFleet) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GameliftFleet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GameliftFleet) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GameliftFleet) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GameliftFleet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GameliftFleet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GameliftFleet) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GameliftFleet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GameliftFleet) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleet) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GameliftFleet) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GameliftFleet) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (g *jsiiProxy_GameliftFleet) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GameliftFleet) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GameliftFleet) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GameliftFleet) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GameliftFleet) PutCertificateConfiguration(value *GameliftFleetCertificateConfiguration) {
	if err := g.validatePutCertificateConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCertificateConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftFleet) PutEc2InboundPermission(value interface{}) {
	if err := g.validatePutEc2InboundPermissionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEc2InboundPermission",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftFleet) PutResourceCreationLimitPolicy(value *GameliftFleetResourceCreationLimitPolicy) {
	if err := g.validatePutResourceCreationLimitPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putResourceCreationLimitPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftFleet) PutRuntimeConfiguration(value *GameliftFleetRuntimeConfiguration) {
	if err := g.validatePutRuntimeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRuntimeConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftFleet) PutTimeouts(value *GameliftFleetTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GameliftFleet) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GameliftFleet) ResetBuildId() {
	_jsii_.InvokeVoid(
		g,
		"resetBuildId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetCertificateConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetCertificateConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetEc2InboundPermission() {
	_jsii_.InvokeVoid(
		g,
		"resetEc2InboundPermission",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetFleetType() {
	_jsii_.InvokeVoid(
		g,
		"resetFleetType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetInstanceRoleArn() {
	_jsii_.InvokeVoid(
		g,
		"resetInstanceRoleArn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetMetricGroups() {
	_jsii_.InvokeVoid(
		g,
		"resetMetricGroups",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetNewGameSessionProtectionPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetNewGameSessionProtectionPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetResourceCreationLimitPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceCreationLimitPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetRuntimeConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetRuntimeConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetScriptId() {
	_jsii_.InvokeVoid(
		g,
		"resetScriptId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetTagsAll() {
	_jsii_.InvokeVoid(
		g,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GameliftFleet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleet) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleet) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameliftFleet) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

