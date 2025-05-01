package googlenetworkconnectivityspoke

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetworkconnectivityspoke/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_connectivity_spoke google_network_connectivity_spoke}.
type GoogleNetworkConnectivitySpoke interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveLabels() cdktf.StringMap
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Group() *string
	SetGroup(val *string)
	GroupInput() *string
	Hub() *string
	SetHub(val *string)
	HubInput() *string
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
	LinkedInterconnectAttachments() GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference
	LinkedInterconnectAttachmentsInput() *GoogleNetworkConnectivitySpokeLinkedInterconnectAttachments
	LinkedProducerVpcNetwork() GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference
	LinkedProducerVpcNetworkInput() *GoogleNetworkConnectivitySpokeLinkedProducerVpcNetwork
	LinkedRouterApplianceInstances() GoogleNetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference
	LinkedRouterApplianceInstancesInput() *GoogleNetworkConnectivitySpokeLinkedRouterApplianceInstances
	LinkedVpcNetwork() GoogleNetworkConnectivitySpokeLinkedVpcNetworkOutputReference
	LinkedVpcNetworkInput() *GoogleNetworkConnectivitySpokeLinkedVpcNetwork
	LinkedVpnTunnels() GoogleNetworkConnectivitySpokeLinkedVpnTunnelsOutputReference
	LinkedVpnTunnelsInput() *GoogleNetworkConnectivitySpokeLinkedVpnTunnels
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleNetworkConnectivitySpokeTimeoutsOutputReference
	TimeoutsInput() interface{}
	UniqueId() *string
	UpdateTime() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
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
	HasResourceMove() interface{}
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
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutLinkedInterconnectAttachments(value *GoogleNetworkConnectivitySpokeLinkedInterconnectAttachments)
	PutLinkedProducerVpcNetwork(value *GoogleNetworkConnectivitySpokeLinkedProducerVpcNetwork)
	PutLinkedRouterApplianceInstances(value *GoogleNetworkConnectivitySpokeLinkedRouterApplianceInstances)
	PutLinkedVpcNetwork(value *GoogleNetworkConnectivitySpokeLinkedVpcNetwork)
	PutLinkedVpnTunnels(value *GoogleNetworkConnectivitySpokeLinkedVpnTunnels)
	PutTimeouts(value *GoogleNetworkConnectivitySpokeTimeouts)
	ResetDescription()
	ResetGroup()
	ResetId()
	ResetLabels()
	ResetLinkedInterconnectAttachments()
	ResetLinkedProducerVpcNetwork()
	ResetLinkedRouterApplianceInstances()
	ResetLinkedVpcNetwork()
	ResetLinkedVpnTunnels()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
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
}

// The jsii proxy struct for GoogleNetworkConnectivitySpoke
type jsiiProxy_GoogleNetworkConnectivitySpoke struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) Group() *string {
	var returns *string
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) GroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) Hub() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) HubInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) LinkedInterconnectAttachments() GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference {
	var returns GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference
	_jsii_.Get(
		j,
		"linkedInterconnectAttachments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) LinkedInterconnectAttachmentsInput() *GoogleNetworkConnectivitySpokeLinkedInterconnectAttachments {
	var returns *GoogleNetworkConnectivitySpokeLinkedInterconnectAttachments
	_jsii_.Get(
		j,
		"linkedInterconnectAttachmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) LinkedProducerVpcNetwork() GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference {
	var returns GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference
	_jsii_.Get(
		j,
		"linkedProducerVpcNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) LinkedProducerVpcNetworkInput() *GoogleNetworkConnectivitySpokeLinkedProducerVpcNetwork {
	var returns *GoogleNetworkConnectivitySpokeLinkedProducerVpcNetwork
	_jsii_.Get(
		j,
		"linkedProducerVpcNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) LinkedRouterApplianceInstances() GoogleNetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference {
	var returns GoogleNetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference
	_jsii_.Get(
		j,
		"linkedRouterApplianceInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) LinkedRouterApplianceInstancesInput() *GoogleNetworkConnectivitySpokeLinkedRouterApplianceInstances {
	var returns *GoogleNetworkConnectivitySpokeLinkedRouterApplianceInstances
	_jsii_.Get(
		j,
		"linkedRouterApplianceInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) LinkedVpcNetwork() GoogleNetworkConnectivitySpokeLinkedVpcNetworkOutputReference {
	var returns GoogleNetworkConnectivitySpokeLinkedVpcNetworkOutputReference
	_jsii_.Get(
		j,
		"linkedVpcNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) LinkedVpcNetworkInput() *GoogleNetworkConnectivitySpokeLinkedVpcNetwork {
	var returns *GoogleNetworkConnectivitySpokeLinkedVpcNetwork
	_jsii_.Get(
		j,
		"linkedVpcNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) LinkedVpnTunnels() GoogleNetworkConnectivitySpokeLinkedVpnTunnelsOutputReference {
	var returns GoogleNetworkConnectivitySpokeLinkedVpnTunnelsOutputReference
	_jsii_.Get(
		j,
		"linkedVpnTunnels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) LinkedVpnTunnelsInput() *GoogleNetworkConnectivitySpokeLinkedVpnTunnels {
	var returns *GoogleNetworkConnectivitySpokeLinkedVpnTunnels
	_jsii_.Get(
		j,
		"linkedVpnTunnelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) Timeouts() GoogleNetworkConnectivitySpokeTimeoutsOutputReference {
	var returns GoogleNetworkConnectivitySpokeTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) UniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_connectivity_spoke google_network_connectivity_spoke} Resource.
func NewGoogleNetworkConnectivitySpoke(scope constructs.Construct, id *string, config *GoogleNetworkConnectivitySpokeConfig) GoogleNetworkConnectivitySpoke {
	_init_.Initialize()

	if err := validateNewGoogleNetworkConnectivitySpokeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkConnectivitySpoke{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetworkConnectivitySpoke.GoogleNetworkConnectivitySpoke",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_connectivity_spoke google_network_connectivity_spoke} Resource.
func NewGoogleNetworkConnectivitySpoke_Override(g GoogleNetworkConnectivitySpoke, scope constructs.Construct, id *string, config *GoogleNetworkConnectivitySpokeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetworkConnectivitySpoke.GoogleNetworkConnectivitySpoke",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke)SetGroup(val *string) {
	if err := j.validateSetGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"group",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke)SetHub(val *string) {
	if err := j.validateSetHubParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hub",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpoke)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a GoogleNetworkConnectivitySpoke resource upon running "cdktf plan <stack-name>".
func GoogleNetworkConnectivitySpoke_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleNetworkConnectivitySpoke_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleNetworkConnectivitySpoke.GoogleNetworkConnectivitySpoke",
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
func GoogleNetworkConnectivitySpoke_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleNetworkConnectivitySpoke_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleNetworkConnectivitySpoke.GoogleNetworkConnectivitySpoke",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleNetworkConnectivitySpoke_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleNetworkConnectivitySpoke_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleNetworkConnectivitySpoke.GoogleNetworkConnectivitySpoke",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleNetworkConnectivitySpoke_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleNetworkConnectivitySpoke_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleNetworkConnectivitySpoke.GoogleNetworkConnectivitySpoke",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleNetworkConnectivitySpoke_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleNetworkConnectivitySpoke.GoogleNetworkConnectivitySpoke",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) PutLinkedInterconnectAttachments(value *GoogleNetworkConnectivitySpokeLinkedInterconnectAttachments) {
	if err := g.validatePutLinkedInterconnectAttachmentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLinkedInterconnectAttachments",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) PutLinkedProducerVpcNetwork(value *GoogleNetworkConnectivitySpokeLinkedProducerVpcNetwork) {
	if err := g.validatePutLinkedProducerVpcNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLinkedProducerVpcNetwork",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) PutLinkedRouterApplianceInstances(value *GoogleNetworkConnectivitySpokeLinkedRouterApplianceInstances) {
	if err := g.validatePutLinkedRouterApplianceInstancesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLinkedRouterApplianceInstances",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) PutLinkedVpcNetwork(value *GoogleNetworkConnectivitySpokeLinkedVpcNetwork) {
	if err := g.validatePutLinkedVpcNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLinkedVpcNetwork",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) PutLinkedVpnTunnels(value *GoogleNetworkConnectivitySpokeLinkedVpnTunnels) {
	if err := g.validatePutLinkedVpnTunnelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLinkedVpnTunnels",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) PutTimeouts(value *GoogleNetworkConnectivitySpokeTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) ResetGroup() {
	_jsii_.InvokeVoid(
		g,
		"resetGroup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) ResetLinkedInterconnectAttachments() {
	_jsii_.InvokeVoid(
		g,
		"resetLinkedInterconnectAttachments",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) ResetLinkedProducerVpcNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetLinkedProducerVpcNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) ResetLinkedRouterApplianceInstances() {
	_jsii_.InvokeVoid(
		g,
		"resetLinkedRouterApplianceInstances",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) ResetLinkedVpcNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetLinkedVpcNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) ResetLinkedVpnTunnels() {
	_jsii_.InvokeVoid(
		g,
		"resetLinkedVpnTunnels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpoke) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

