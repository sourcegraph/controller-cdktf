package googlenetworkconnectivitypolicybasedroute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetworkconnectivitypolicybasedroute/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_network_connectivity_policy_based_route google_network_connectivity_policy_based_route}.
type GoogleNetworkConnectivityPolicyBasedRoute interface {
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
	Filter() GoogleNetworkConnectivityPolicyBasedRouteFilterOutputReference
	FilterInput() *GoogleNetworkConnectivityPolicyBasedRouteFilter
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
	InterconnectAttachment() GoogleNetworkConnectivityPolicyBasedRouteInterconnectAttachmentOutputReference
	InterconnectAttachmentInput() *GoogleNetworkConnectivityPolicyBasedRouteInterconnectAttachment
	Kind() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	NextHopIlbIp() *string
	SetNextHopIlbIp(val *string)
	NextHopIlbIpInput() *string
	NextHopOtherRoutes() *string
	SetNextHopOtherRoutes(val *string)
	NextHopOtherRoutesInput() *string
	// The tree node.
	Node() constructs.Node
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleNetworkConnectivityPolicyBasedRouteTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateTime() *string
	VirtualMachine() GoogleNetworkConnectivityPolicyBasedRouteVirtualMachineOutputReference
	VirtualMachineInput() *GoogleNetworkConnectivityPolicyBasedRouteVirtualMachine
	Warnings() GoogleNetworkConnectivityPolicyBasedRouteWarningsList
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutFilter(value *GoogleNetworkConnectivityPolicyBasedRouteFilter)
	PutInterconnectAttachment(value *GoogleNetworkConnectivityPolicyBasedRouteInterconnectAttachment)
	PutTimeouts(value *GoogleNetworkConnectivityPolicyBasedRouteTimeouts)
	PutVirtualMachine(value *GoogleNetworkConnectivityPolicyBasedRouteVirtualMachine)
	ResetDescription()
	ResetId()
	ResetInterconnectAttachment()
	ResetLabels()
	ResetNextHopIlbIp()
	ResetNextHopOtherRoutes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPriority()
	ResetProject()
	ResetTimeouts()
	ResetVirtualMachine()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleNetworkConnectivityPolicyBasedRoute
type jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) Filter() GoogleNetworkConnectivityPolicyBasedRouteFilterOutputReference {
	var returns GoogleNetworkConnectivityPolicyBasedRouteFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) FilterInput() *GoogleNetworkConnectivityPolicyBasedRouteFilter {
	var returns *GoogleNetworkConnectivityPolicyBasedRouteFilter
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) InterconnectAttachment() GoogleNetworkConnectivityPolicyBasedRouteInterconnectAttachmentOutputReference {
	var returns GoogleNetworkConnectivityPolicyBasedRouteInterconnectAttachmentOutputReference
	_jsii_.Get(
		j,
		"interconnectAttachment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) InterconnectAttachmentInput() *GoogleNetworkConnectivityPolicyBasedRouteInterconnectAttachment {
	var returns *GoogleNetworkConnectivityPolicyBasedRouteInterconnectAttachment
	_jsii_.Get(
		j,
		"interconnectAttachmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) NextHopIlbIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextHopIlbIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) NextHopIlbIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextHopIlbIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) NextHopOtherRoutes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextHopOtherRoutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) NextHopOtherRoutesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextHopOtherRoutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) Timeouts() GoogleNetworkConnectivityPolicyBasedRouteTimeoutsOutputReference {
	var returns GoogleNetworkConnectivityPolicyBasedRouteTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) VirtualMachine() GoogleNetworkConnectivityPolicyBasedRouteVirtualMachineOutputReference {
	var returns GoogleNetworkConnectivityPolicyBasedRouteVirtualMachineOutputReference
	_jsii_.Get(
		j,
		"virtualMachine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) VirtualMachineInput() *GoogleNetworkConnectivityPolicyBasedRouteVirtualMachine {
	var returns *GoogleNetworkConnectivityPolicyBasedRouteVirtualMachine
	_jsii_.Get(
		j,
		"virtualMachineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) Warnings() GoogleNetworkConnectivityPolicyBasedRouteWarningsList {
	var returns GoogleNetworkConnectivityPolicyBasedRouteWarningsList
	_jsii_.Get(
		j,
		"warnings",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_network_connectivity_policy_based_route google_network_connectivity_policy_based_route} Resource.
func NewGoogleNetworkConnectivityPolicyBasedRoute(scope constructs.Construct, id *string, config *GoogleNetworkConnectivityPolicyBasedRouteConfig) GoogleNetworkConnectivityPolicyBasedRoute {
	_init_.Initialize()

	if err := validateNewGoogleNetworkConnectivityPolicyBasedRouteParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetworkConnectivityPolicyBasedRoute.GoogleNetworkConnectivityPolicyBasedRoute",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_network_connectivity_policy_based_route google_network_connectivity_policy_based_route} Resource.
func NewGoogleNetworkConnectivityPolicyBasedRoute_Override(g GoogleNetworkConnectivityPolicyBasedRoute, scope constructs.Construct, id *string, config *GoogleNetworkConnectivityPolicyBasedRouteConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetworkConnectivityPolicyBasedRoute.GoogleNetworkConnectivityPolicyBasedRoute",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute)SetNextHopIlbIp(val *string) {
	if err := j.validateSetNextHopIlbIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nextHopIlbIp",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute)SetNextHopOtherRoutes(val *string) {
	if err := j.validateSetNextHopOtherRoutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nextHopOtherRoutes",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
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
func GoogleNetworkConnectivityPolicyBasedRoute_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleNetworkConnectivityPolicyBasedRoute_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleNetworkConnectivityPolicyBasedRoute.GoogleNetworkConnectivityPolicyBasedRoute",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleNetworkConnectivityPolicyBasedRoute_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleNetworkConnectivityPolicyBasedRoute_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleNetworkConnectivityPolicyBasedRoute.GoogleNetworkConnectivityPolicyBasedRoute",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleNetworkConnectivityPolicyBasedRoute_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleNetworkConnectivityPolicyBasedRoute_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleNetworkConnectivityPolicyBasedRoute.GoogleNetworkConnectivityPolicyBasedRoute",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleNetworkConnectivityPolicyBasedRoute_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleNetworkConnectivityPolicyBasedRoute.GoogleNetworkConnectivityPolicyBasedRoute",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) PutFilter(value *GoogleNetworkConnectivityPolicyBasedRouteFilter) {
	if err := g.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFilter",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) PutInterconnectAttachment(value *GoogleNetworkConnectivityPolicyBasedRouteInterconnectAttachment) {
	if err := g.validatePutInterconnectAttachmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInterconnectAttachment",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) PutTimeouts(value *GoogleNetworkConnectivityPolicyBasedRouteTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) PutVirtualMachine(value *GoogleNetworkConnectivityPolicyBasedRouteVirtualMachine) {
	if err := g.validatePutVirtualMachineParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVirtualMachine",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) ResetInterconnectAttachment() {
	_jsii_.InvokeVoid(
		g,
		"resetInterconnectAttachment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) ResetNextHopIlbIp() {
	_jsii_.InvokeVoid(
		g,
		"resetNextHopIlbIp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) ResetNextHopOtherRoutes() {
	_jsii_.InvokeVoid(
		g,
		"resetNextHopOtherRoutes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) ResetPriority() {
	_jsii_.InvokeVoid(
		g,
		"resetPriority",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) ResetVirtualMachine() {
	_jsii_.InvokeVoid(
		g,
		"resetVirtualMachine",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivityPolicyBasedRoute) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

