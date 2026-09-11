package vpcipampool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/aws/vpcipampool/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/vpc_ipam_pool aws_vpc_ipam_pool}.
type VpcIpamPool interface {
	cdktn.TerraformResource
	AddressFamily() *string
	SetAddressFamily(val *string)
	AddressFamilyInput() *string
	AllocationDefaultNetmaskLength() *float64
	SetAllocationDefaultNetmaskLength(val *float64)
	AllocationDefaultNetmaskLengthInput() *float64
	AllocationMaxNetmaskLength() *float64
	SetAllocationMaxNetmaskLength(val *float64)
	AllocationMaxNetmaskLengthInput() *float64
	AllocationMinNetmaskLength() *float64
	SetAllocationMinNetmaskLength(val *float64)
	AllocationMinNetmaskLengthInput() *float64
	AllocationResourceTags() *map[string]*string
	SetAllocationResourceTags(val *map[string]*string)
	AllocationResourceTagsInput() *map[string]*string
	Arn() *string
	AutoImport() interface{}
	SetAutoImport(val interface{})
	AutoImportInput() interface{}
	AwsService() *string
	SetAwsService(val *string)
	AwsServiceInput() *string
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	IpamScopeId() *string
	SetIpamScopeId(val *string)
	IpamScopeIdInput() *string
	IpamScopeType() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Locale() *string
	SetLocale(val *string)
	LocaleInput() *string
	// The tree node.
	Node() constructs.Node
	PoolDepth() *float64
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicIpSource() *string
	SetPublicIpSource(val *string)
	PublicIpSourceInput() *string
	PubliclyAdvertisable() interface{}
	SetPubliclyAdvertisable(val interface{})
	PubliclyAdvertisableInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	SourceIpamPoolId() *string
	SetSourceIpamPoolId(val *string)
	SourceIpamPoolIdInput() *string
	State() *string
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
	Timeouts() VpcIpamPoolTimeoutsOutputReference
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
	PutTimeouts(value *VpcIpamPoolTimeouts)
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
	ResetAllocationDefaultNetmaskLength()
	ResetAllocationMaxNetmaskLength()
	ResetAllocationMinNetmaskLength()
	ResetAllocationResourceTags()
	ResetAutoImport()
	ResetAwsService()
	ResetDescription()
	ResetId()
	ResetLocale()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicIpSource()
	ResetPubliclyAdvertisable()
	ResetSourceIpamPoolId()
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

// The jsii proxy struct for VpcIpamPool
type jsiiProxy_VpcIpamPool struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_VpcIpamPool) AddressFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AddressFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AllocationDefaultNetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationDefaultNetmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AllocationDefaultNetmaskLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationDefaultNetmaskLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AllocationMaxNetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationMaxNetmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AllocationMaxNetmaskLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationMaxNetmaskLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AllocationMinNetmaskLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationMinNetmaskLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AllocationMinNetmaskLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationMinNetmaskLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AllocationResourceTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"allocationResourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AllocationResourceTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"allocationResourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AutoImport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoImport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AutoImportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoImportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AwsService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) AwsServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) IpamScopeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamScopeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) IpamScopeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamScopeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) IpamScopeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamScopeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Locale() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) LocaleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) PoolDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"poolDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) PublicIpSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) PublicIpSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) PubliclyAdvertisable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAdvertisable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) PubliclyAdvertisableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAdvertisableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) SourceIpamPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIpamPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) SourceIpamPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIpamPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) Timeouts() VpcIpamPoolTimeoutsOutputReference {
	var returns VpcIpamPoolTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcIpamPool) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/vpc_ipam_pool aws_vpc_ipam_pool} Resource.
func NewVpcIpamPool(scope constructs.Construct, id *string, config *VpcIpamPoolConfig) VpcIpamPool {
	_init_.Initialize()

	if err := validateNewVpcIpamPoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpcIpamPool{}

	_jsii_.Create(
		"@cdktn/provider-aws.vpcIpamPool.VpcIpamPool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/vpc_ipam_pool aws_vpc_ipam_pool} Resource.
func NewVpcIpamPool_Override(v VpcIpamPool, scope constructs.Construct, id *string, config *VpcIpamPoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.vpcIpamPool.VpcIpamPool",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetAddressFamily(val *string) {
	if err := j.validateSetAddressFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressFamily",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetAllocationDefaultNetmaskLength(val *float64) {
	if err := j.validateSetAllocationDefaultNetmaskLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationDefaultNetmaskLength",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetAllocationMaxNetmaskLength(val *float64) {
	if err := j.validateSetAllocationMaxNetmaskLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationMaxNetmaskLength",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetAllocationMinNetmaskLength(val *float64) {
	if err := j.validateSetAllocationMinNetmaskLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationMinNetmaskLength",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetAllocationResourceTags(val *map[string]*string) {
	if err := j.validateSetAllocationResourceTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationResourceTags",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetAutoImport(val interface{}) {
	if err := j.validateSetAutoImportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoImport",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetAwsService(val *string) {
	if err := j.validateSetAwsServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsService",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetIpamScopeId(val *string) {
	if err := j.validateSetIpamScopeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipamScopeId",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetLocale(val *string) {
	if err := j.validateSetLocaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locale",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetPublicIpSource(val *string) {
	if err := j.validateSetPublicIpSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicIpSource",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetPubliclyAdvertisable(val interface{}) {
	if err := j.validateSetPubliclyAdvertisableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publiclyAdvertisable",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetSourceIpamPoolId(val *string) {
	if err := j.validateSetSourceIpamPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceIpamPoolId",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_VpcIpamPool)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a VpcIpamPool resource upon running "cdktn plan <stack-name>".
func VpcIpamPool_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateVpcIpamPool_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.vpcIpamPool.VpcIpamPool",
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
func VpcIpamPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcIpamPool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.vpcIpamPool.VpcIpamPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpcIpamPool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcIpamPool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.vpcIpamPool.VpcIpamPool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpcIpamPool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcIpamPool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-aws.vpcIpamPool.VpcIpamPool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VpcIpamPool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-aws.vpcIpamPool.VpcIpamPool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VpcIpamPool) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VpcIpamPool) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VpcIpamPool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPool) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPool) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPool) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPool) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPool) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPool) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VpcIpamPool) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPool) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := v.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPool) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VpcIpamPool) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VpcIpamPool) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VpcIpamPool) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VpcIpamPool) PutTimeouts(value *VpcIpamPoolTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpcIpamPool) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := v.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetAllocationDefaultNetmaskLength() {
	_jsii_.InvokeVoid(
		v,
		"resetAllocationDefaultNetmaskLength",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetAllocationMaxNetmaskLength() {
	_jsii_.InvokeVoid(
		v,
		"resetAllocationMaxNetmaskLength",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetAllocationMinNetmaskLength() {
	_jsii_.InvokeVoid(
		v,
		"resetAllocationMinNetmaskLength",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetAllocationResourceTags() {
	_jsii_.InvokeVoid(
		v,
		"resetAllocationResourceTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetAutoImport() {
	_jsii_.InvokeVoid(
		v,
		"resetAutoImport",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetAwsService() {
	_jsii_.InvokeVoid(
		v,
		"resetAwsService",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetDescription() {
	_jsii_.InvokeVoid(
		v,
		"resetDescription",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetLocale() {
	_jsii_.InvokeVoid(
		v,
		"resetLocale",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetPublicIpSource() {
	_jsii_.InvokeVoid(
		v,
		"resetPublicIpSource",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetPubliclyAdvertisable() {
	_jsii_.InvokeVoid(
		v,
		"resetPubliclyAdvertisable",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetSourceIpamPoolId() {
	_jsii_.InvokeVoid(
		v,
		"resetSourceIpamPoolId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetTagsAll() {
	_jsii_.InvokeVoid(
		v,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpcIpamPool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPool) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPool) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcIpamPool) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		v,
		"with",
		args,
		&returns,
	)

	return returns
}

