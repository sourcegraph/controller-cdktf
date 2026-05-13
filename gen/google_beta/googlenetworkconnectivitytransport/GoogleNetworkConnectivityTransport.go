package googlenetworkconnectivitytransport

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetworkconnectivitytransport/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_transport google_network_connectivity_transport}.
type GoogleNetworkConnectivityTransport interface {
	cdktf.TerraformResource
	AdminEnabled() interface{}
	SetAdminEnabled(val interface{})
	AdminEnabledInput() interface{}
	AdvertisedRoutes() *[]*string
	SetAdvertisedRoutes(val *[]*string)
	AdvertisedRoutesInput() *[]*string
	AutoAccept() interface{}
	SetAutoAccept(val interface{})
	AutoAcceptInput() interface{}
	Bandwidth() *string
	SetBandwidth(val *string)
	BandwidthInput() *string
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
	GeneratedActivationKey() *string
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
	MtuLimit() *float64
	SetMtuLimit(val *float64)
	MtuLimitInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	// The tree node.
	Node() constructs.Node
	PeeringNetwork() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	ProvidedActivationKey() *string
	SetProvidedActivationKey(val *string)
	ProvidedActivationKeyInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PscRoutingEnabled() interface{}
	SetPscRoutingEnabled(val interface{})
	PscRoutingEnabledInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RemoteAccountId() *string
	SetRemoteAccountId(val *string)
	RemoteAccountIdInput() *string
	RemoteProfile() *string
	SetRemoteProfile(val *string)
	RemoteProfileInput() *string
	StackType() *string
	SetStackType(val *string)
	StackTypeInput() *string
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleNetworkConnectivityTransportTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutTimeouts(value *GoogleNetworkConnectivityTransportTimeouts)
	ResetAdminEnabled()
	ResetAdvertisedRoutes()
	ResetAutoAccept()
	ResetBandwidth()
	ResetDescription()
	ResetHub()
	ResetId()
	ResetLabels()
	ResetMtuLimit()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetProvidedActivationKey()
	ResetPscRoutingEnabled()
	ResetRemoteAccountId()
	ResetStackType()
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

// The jsii proxy struct for GoogleNetworkConnectivityTransport
type jsiiProxy_GoogleNetworkConnectivityTransport struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) AdminEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) AdminEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) AdvertisedRoutes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"advertisedRoutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) AdvertisedRoutesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"advertisedRoutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) AutoAccept() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAccept",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) AutoAcceptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAcceptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) Bandwidth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) BandwidthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) GeneratedActivationKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"generatedActivationKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) Hub() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) HubInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) MtuLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mtuLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) MtuLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mtuLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) PeeringNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peeringNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) ProvidedActivationKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providedActivationKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) ProvidedActivationKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providedActivationKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) PscRoutingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pscRoutingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) PscRoutingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pscRoutingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) RemoteAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) RemoteAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) RemoteProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) RemoteProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) StackType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) StackTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) Timeouts() GoogleNetworkConnectivityTransportTimeoutsOutputReference {
	var returns GoogleNetworkConnectivityTransportTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_transport google_network_connectivity_transport} Resource.
func NewGoogleNetworkConnectivityTransport(scope constructs.Construct, id *string, config *GoogleNetworkConnectivityTransportConfig) GoogleNetworkConnectivityTransport {
	_init_.Initialize()

	if err := validateNewGoogleNetworkConnectivityTransportParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkConnectivityTransport{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetworkConnectivityTransport.GoogleNetworkConnectivityTransport",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_transport google_network_connectivity_transport} Resource.
func NewGoogleNetworkConnectivityTransport_Override(g GoogleNetworkConnectivityTransport, scope constructs.Construct, id *string, config *GoogleNetworkConnectivityTransportConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetworkConnectivityTransport.GoogleNetworkConnectivityTransport",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport)SetAdminEnabled(val interface{}) {
	if err := j.validateSetAdminEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport)SetAdvertisedRoutes(val *[]*string) {
	if err := j.validateSetAdvertisedRoutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advertisedRoutes",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport)SetAutoAccept(val interface{}) {
	if err := j.validateSetAutoAcceptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoAccept",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport)SetBandwidth(val *string) {
	if err := j.validateSetBandwidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidth",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport)SetHub(val *string) {
	if err := j.validateSetHubParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hub",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport)SetMtuLimit(val *float64) {
	if err := j.validateSetMtuLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mtuLimit",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport)SetProvidedActivationKey(val *string) {
	if err := j.validateSetProvidedActivationKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providedActivationKey",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport)SetPscRoutingEnabled(val interface{}) {
	if err := j.validateSetPscRoutingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pscRoutingEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport)SetRemoteAccountId(val *string) {
	if err := j.validateSetRemoteAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteAccountId",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport)SetRemoteProfile(val *string) {
	if err := j.validateSetRemoteProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteProfile",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityTransport)SetStackType(val *string) {
	if err := j.validateSetStackTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackType",
		val,
	)
}

// Generates CDKTF code for importing a GoogleNetworkConnectivityTransport resource upon running "cdktf plan <stack-name>".
func GoogleNetworkConnectivityTransport_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleNetworkConnectivityTransport_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleNetworkConnectivityTransport.GoogleNetworkConnectivityTransport",
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
func GoogleNetworkConnectivityTransport_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleNetworkConnectivityTransport_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleNetworkConnectivityTransport.GoogleNetworkConnectivityTransport",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleNetworkConnectivityTransport_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleNetworkConnectivityTransport_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleNetworkConnectivityTransport.GoogleNetworkConnectivityTransport",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleNetworkConnectivityTransport_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleNetworkConnectivityTransport_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleNetworkConnectivityTransport.GoogleNetworkConnectivityTransport",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleNetworkConnectivityTransport_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleNetworkConnectivityTransport.GoogleNetworkConnectivityTransport",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) PutTimeouts(value *GoogleNetworkConnectivityTransportTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) ResetAdminEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetAdminEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) ResetAdvertisedRoutes() {
	_jsii_.InvokeVoid(
		g,
		"resetAdvertisedRoutes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) ResetAutoAccept() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoAccept",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) ResetBandwidth() {
	_jsii_.InvokeVoid(
		g,
		"resetBandwidth",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) ResetHub() {
	_jsii_.InvokeVoid(
		g,
		"resetHub",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) ResetMtuLimit() {
	_jsii_.InvokeVoid(
		g,
		"resetMtuLimit",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) ResetProvidedActivationKey() {
	_jsii_.InvokeVoid(
		g,
		"resetProvidedActivationKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) ResetPscRoutingEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetPscRoutingEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) ResetRemoteAccountId() {
	_jsii_.InvokeVoid(
		g,
		"resetRemoteAccountId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) ResetStackType() {
	_jsii_.InvokeVoid(
		g,
		"resetStackType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivityTransport) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

