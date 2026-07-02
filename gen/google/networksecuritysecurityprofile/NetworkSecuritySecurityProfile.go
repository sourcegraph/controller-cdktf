package networksecuritysecurityprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/networksecuritysecurityprofile/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/network_security_security_profile google_network_security_security_profile}.
type NetworkSecuritySecurityProfile interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	CreateTime() *string
	CustomInterceptProfile() NetworkSecuritySecurityProfileCustomInterceptProfileOutputReference
	CustomInterceptProfileInput() *NetworkSecuritySecurityProfileCustomInterceptProfile
	CustomMirroringProfile() NetworkSecuritySecurityProfileCustomMirroringProfileOutputReference
	CustomMirroringProfileInput() *NetworkSecuritySecurityProfileCustomMirroringProfile
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveLabels() cdktf.StringMap
	Etag() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Parent() *string
	SetParent(val *string)
	ParentInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	// Experimental.
	RawOverrides() any
	SelfLink() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	ThreatPreventionProfile() NetworkSecuritySecurityProfileThreatPreventionProfileOutputReference
	ThreatPreventionProfileInput() *NetworkSecuritySecurityProfileThreatPreventionProfile
	Timeouts() NetworkSecuritySecurityProfileTimeoutsOutputReference
	TimeoutsInput() any
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UpdateTime() *string
	UrlFilteringProfile() NetworkSecuritySecurityProfileUrlFilteringProfileOutputReference
	UrlFilteringProfileInput() *NetworkSecuritySecurityProfileUrlFilteringProfile
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value any)
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
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
	HasResourceMove() any
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index any)
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCustomInterceptProfile(value *NetworkSecuritySecurityProfileCustomInterceptProfile)
	PutCustomMirroringProfile(value *NetworkSecuritySecurityProfileCustomMirroringProfile)
	PutThreatPreventionProfile(value *NetworkSecuritySecurityProfileThreatPreventionProfile)
	PutTimeouts(value *NetworkSecuritySecurityProfileTimeouts)
	PutUrlFilteringProfile(value *NetworkSecuritySecurityProfileUrlFilteringProfile)
	ResetCustomInterceptProfile()
	ResetCustomMirroringProfile()
	ResetDescription()
	ResetId()
	ResetLabels()
	ResetLocation()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParent()
	ResetThreatPreventionProfile()
	ResetTimeouts()
	ResetUrlFilteringProfile()
	SynthesizeAttributes() *map[string]any
	SynthesizeHclAttributes() *map[string]any
	// Experimental.
	ToHclTerraform() any
	// Experimental.
	ToMetadata() any
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() any
}

// The jsii proxy struct for NetworkSecuritySecurityProfile
type jsiiProxy_NetworkSecuritySecurityProfile struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) CustomInterceptProfile() NetworkSecuritySecurityProfileCustomInterceptProfileOutputReference {
	var returns NetworkSecuritySecurityProfileCustomInterceptProfileOutputReference
	_jsii_.Get(
		j,
		"customInterceptProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) CustomInterceptProfileInput() *NetworkSecuritySecurityProfileCustomInterceptProfile {
	var returns *NetworkSecuritySecurityProfileCustomInterceptProfile
	_jsii_.Get(
		j,
		"customInterceptProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) CustomMirroringProfile() NetworkSecuritySecurityProfileCustomMirroringProfileOutputReference {
	var returns NetworkSecuritySecurityProfileCustomMirroringProfileOutputReference
	_jsii_.Get(
		j,
		"customMirroringProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) CustomMirroringProfileInput() *NetworkSecuritySecurityProfileCustomMirroringProfile {
	var returns *NetworkSecuritySecurityProfileCustomMirroringProfile
	_jsii_.Get(
		j,
		"customMirroringProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) Parent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) ParentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) ThreatPreventionProfile() NetworkSecuritySecurityProfileThreatPreventionProfileOutputReference {
	var returns NetworkSecuritySecurityProfileThreatPreventionProfileOutputReference
	_jsii_.Get(
		j,
		"threatPreventionProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) ThreatPreventionProfileInput() *NetworkSecuritySecurityProfileThreatPreventionProfile {
	var returns *NetworkSecuritySecurityProfileThreatPreventionProfile
	_jsii_.Get(
		j,
		"threatPreventionProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) Timeouts() NetworkSecuritySecurityProfileTimeoutsOutputReference {
	var returns NetworkSecuritySecurityProfileTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) UrlFilteringProfile() NetworkSecuritySecurityProfileUrlFilteringProfileOutputReference {
	var returns NetworkSecuritySecurityProfileUrlFilteringProfileOutputReference
	_jsii_.Get(
		j,
		"urlFilteringProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) UrlFilteringProfileInput() *NetworkSecuritySecurityProfileUrlFilteringProfile {
	var returns *NetworkSecuritySecurityProfileUrlFilteringProfile
	_jsii_.Get(
		j,
		"urlFilteringProfileInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/network_security_security_profile google_network_security_security_profile} Resource.
func NewNetworkSecuritySecurityProfile(scope constructs.Construct, id *string, config *NetworkSecuritySecurityProfileConfig) NetworkSecuritySecurityProfile {
	_init_.Initialize()

	if err := validateNewNetworkSecuritySecurityProfileParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkSecuritySecurityProfile{}

	_jsii_.Create(
		"@cdktf/provider-google.networkSecuritySecurityProfile.NetworkSecuritySecurityProfile",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/network_security_security_profile google_network_security_security_profile} Resource.
func NewNetworkSecuritySecurityProfile_Override(n NetworkSecuritySecurityProfile, scope constructs.Construct, id *string, config *NetworkSecuritySecurityProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.networkSecuritySecurityProfile.NetworkSecuritySecurityProfile",
		[]any{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) SetParent(val *string) {
	if err := j.validateSetParentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parent",
		val,
	)
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NetworkSecuritySecurityProfile) SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTF code for importing a NetworkSecuritySecurityProfile resource upon running "cdktf plan <stack-name>".
func NetworkSecuritySecurityProfile_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNetworkSecuritySecurityProfile_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.networkSecuritySecurityProfile.NetworkSecuritySecurityProfile",
		"generateConfigForImport",
		[]any{scope, importToId, importFromId, provider},
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
func NetworkSecuritySecurityProfile_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateNetworkSecuritySecurityProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.networkSecuritySecurityProfile.NetworkSecuritySecurityProfile",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkSecuritySecurityProfile_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateNetworkSecuritySecurityProfile_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.networkSecuritySecurityProfile.NetworkSecuritySecurityProfile",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkSecuritySecurityProfile_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateNetworkSecuritySecurityProfile_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.networkSecuritySecurityProfile.NetworkSecuritySecurityProfile",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func NetworkSecuritySecurityProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.networkSecuritySecurityProfile.NetworkSecuritySecurityProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) AddOverride(path *string, value any) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]any{path, value},
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]any{id, provider},
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]any{id},
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) MoveTo(moveTarget *string, index any) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]any{id},
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) PutCustomInterceptProfile(value *NetworkSecuritySecurityProfileCustomInterceptProfile) {
	if err := n.validatePutCustomInterceptProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putCustomInterceptProfile",
		[]any{value},
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) PutCustomMirroringProfile(value *NetworkSecuritySecurityProfileCustomMirroringProfile) {
	if err := n.validatePutCustomMirroringProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putCustomMirroringProfile",
		[]any{value},
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) PutThreatPreventionProfile(value *NetworkSecuritySecurityProfileThreatPreventionProfile) {
	if err := n.validatePutThreatPreventionProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putThreatPreventionProfile",
		[]any{value},
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) PutTimeouts(value *NetworkSecuritySecurityProfileTimeouts) {
	if err := n.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeouts",
		[]any{value},
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) PutUrlFilteringProfile(value *NetworkSecuritySecurityProfileUrlFilteringProfile) {
	if err := n.validatePutUrlFilteringProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putUrlFilteringProfile",
		[]any{value},
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) ResetCustomInterceptProfile() {
	_jsii_.InvokeVoid(
		n,
		"resetCustomInterceptProfile",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) ResetCustomMirroringProfile() {
	_jsii_.InvokeVoid(
		n,
		"resetCustomMirroringProfile",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) ResetLabels() {
	_jsii_.InvokeVoid(
		n,
		"resetLabels",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) ResetLocation() {
	_jsii_.InvokeVoid(
		n,
		"resetLocation",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) ResetParent() {
	_jsii_.InvokeVoid(
		n,
		"resetParent",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) ResetThreatPreventionProfile() {
	_jsii_.InvokeVoid(
		n,
		"resetThreatPreventionProfile",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) ResetTimeouts() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) ResetUrlFilteringProfile() {
	_jsii_.InvokeVoid(
		n,
		"resetUrlFilteringProfile",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecuritySecurityProfile) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
