package googlecomputevpntunnel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputevpntunnel/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_compute_vpn_tunnel google_compute_vpn_tunnel}.
type GoogleComputeVpnTunnel interface {
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
	CreationTimestamp() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DetailedStatus() *string
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
	IkeVersion() *float64
	SetIkeVersion(val *float64)
	IkeVersionInput() *float64
	LabelFingerprint() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalTrafficSelector() *[]*string
	SetLocalTrafficSelector(val *[]*string)
	LocalTrafficSelectorInput() *[]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PeerExternalGateway() *string
	SetPeerExternalGateway(val *string)
	PeerExternalGatewayInput() *string
	PeerExternalGatewayInterface() *float64
	SetPeerExternalGatewayInterface(val *float64)
	PeerExternalGatewayInterfaceInput() *float64
	PeerGcpGateway() *string
	SetPeerGcpGateway(val *string)
	PeerGcpGatewayInput() *string
	PeerIp() *string
	SetPeerIp(val *string)
	PeerIpInput() *string
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RemoteTrafficSelector() *[]*string
	SetRemoteTrafficSelector(val *[]*string)
	RemoteTrafficSelectorInput() *[]*string
	Router() *string
	SetRouter(val *string)
	RouterInput() *string
	SelfLink() *string
	SharedSecret() *string
	SetSharedSecret(val *string)
	SharedSecretHash() *string
	SharedSecretInput() *string
	TargetVpnGateway() *string
	SetTargetVpnGateway(val *string)
	TargetVpnGatewayInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleComputeVpnTunnelTimeoutsOutputReference
	TimeoutsInput() interface{}
	TunnelId() *string
	VpnGateway() *string
	SetVpnGateway(val *string)
	VpnGatewayInput() *string
	VpnGatewayInterface() *float64
	SetVpnGatewayInterface(val *float64)
	VpnGatewayInterfaceInput() *float64
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
	PutTimeouts(value *GoogleComputeVpnTunnelTimeouts)
	ResetDescription()
	ResetId()
	ResetIkeVersion()
	ResetLabels()
	ResetLocalTrafficSelector()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPeerExternalGateway()
	ResetPeerExternalGatewayInterface()
	ResetPeerGcpGateway()
	ResetPeerIp()
	ResetProject()
	ResetRegion()
	ResetRemoteTrafficSelector()
	ResetRouter()
	ResetTargetVpnGateway()
	ResetTimeouts()
	ResetVpnGateway()
	ResetVpnGatewayInterface()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleComputeVpnTunnel
type jsiiProxy_GoogleComputeVpnTunnel struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) DetailedStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detailedStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) IkeVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ikeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) IkeVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ikeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) LabelFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) LocalTrafficSelector() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localTrafficSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) LocalTrafficSelectorInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localTrafficSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) PeerExternalGateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerExternalGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) PeerExternalGatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerExternalGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) PeerExternalGatewayInterface() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"peerExternalGatewayInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) PeerExternalGatewayInterfaceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"peerExternalGatewayInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) PeerGcpGateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerGcpGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) PeerGcpGatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerGcpGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) PeerIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) PeerIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) RemoteTrafficSelector() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"remoteTrafficSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) RemoteTrafficSelectorInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"remoteTrafficSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) Router() *string {
	var returns *string
	_jsii_.Get(
		j,
		"router",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) RouterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) SharedSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) SharedSecretHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedSecretHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) SharedSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) TargetVpnGateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetVpnGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) TargetVpnGatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetVpnGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) Timeouts() GoogleComputeVpnTunnelTimeoutsOutputReference {
	var returns GoogleComputeVpnTunnelTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) TunnelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) VpnGateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) VpnGatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) VpnGatewayInterface() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vpnGatewayInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeVpnTunnel) VpnGatewayInterfaceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vpnGatewayInterfaceInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_compute_vpn_tunnel google_compute_vpn_tunnel} Resource.
func NewGoogleComputeVpnTunnel(scope constructs.Construct, id *string, config *GoogleComputeVpnTunnelConfig) GoogleComputeVpnTunnel {
	_init_.Initialize()

	if err := validateNewGoogleComputeVpnTunnelParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeVpnTunnel{}

	_jsii_.Create(
		"google-beta.googleComputeVpnTunnel.GoogleComputeVpnTunnel",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_compute_vpn_tunnel google_compute_vpn_tunnel} Resource.
func NewGoogleComputeVpnTunnel_Override(g GoogleComputeVpnTunnel, scope constructs.Construct, id *string, config *GoogleComputeVpnTunnelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeVpnTunnel.GoogleComputeVpnTunnel",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnel)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnel)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnel)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnel)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnel)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnel)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnel)SetIkeVersion(val *float64) {
	if err := j.validateSetIkeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ikeVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnel)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnel)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnel)SetLocalTrafficSelector(val *[]*string) {
	if err := j.validateSetLocalTrafficSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localTrafficSelector",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnel)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnel)SetPeerExternalGateway(val *string) {
	if err := j.validateSetPeerExternalGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerExternalGateway",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnel)SetPeerExternalGatewayInterface(val *float64) {
	if err := j.validateSetPeerExternalGatewayInterfaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerExternalGatewayInterface",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnel)SetPeerGcpGateway(val *string) {
	if err := j.validateSetPeerGcpGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerGcpGateway",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnel)SetPeerIp(val *string) {
	if err := j.validateSetPeerIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerIp",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnel)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnel)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnel)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnel)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnel)SetRemoteTrafficSelector(val *[]*string) {
	if err := j.validateSetRemoteTrafficSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteTrafficSelector",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnel)SetRouter(val *string) {
	if err := j.validateSetRouterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"router",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnel)SetSharedSecret(val *string) {
	if err := j.validateSetSharedSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedSecret",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnel)SetTargetVpnGateway(val *string) {
	if err := j.validateSetTargetVpnGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetVpnGateway",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnel)SetVpnGateway(val *string) {
	if err := j.validateSetVpnGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnGateway",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeVpnTunnel)SetVpnGatewayInterface(val *float64) {
	if err := j.validateSetVpnGatewayInterfaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnGatewayInterface",
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
func GoogleComputeVpnTunnel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeVpnTunnel_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleComputeVpnTunnel.GoogleComputeVpnTunnel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeVpnTunnel_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeVpnTunnel_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleComputeVpnTunnel.GoogleComputeVpnTunnel",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeVpnTunnel_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeVpnTunnel_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleComputeVpnTunnel.GoogleComputeVpnTunnel",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputeVpnTunnel_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"google-beta.googleComputeVpnTunnel.GoogleComputeVpnTunnel",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeVpnTunnel) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnel) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeVpnTunnel) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeVpnTunnel) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeVpnTunnel) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeVpnTunnel) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeVpnTunnel) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeVpnTunnel) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeVpnTunnel) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeVpnTunnel) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeVpnTunnel) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeVpnTunnel) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnel) PutTimeouts(value *GoogleComputeVpnTunnelTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnel) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnel) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnel) ResetIkeVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetIkeVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnel) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnel) ResetLocalTrafficSelector() {
	_jsii_.InvokeVoid(
		g,
		"resetLocalTrafficSelector",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnel) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnel) ResetPeerExternalGateway() {
	_jsii_.InvokeVoid(
		g,
		"resetPeerExternalGateway",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnel) ResetPeerExternalGatewayInterface() {
	_jsii_.InvokeVoid(
		g,
		"resetPeerExternalGatewayInterface",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnel) ResetPeerGcpGateway() {
	_jsii_.InvokeVoid(
		g,
		"resetPeerGcpGateway",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnel) ResetPeerIp() {
	_jsii_.InvokeVoid(
		g,
		"resetPeerIp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnel) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnel) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnel) ResetRemoteTrafficSelector() {
	_jsii_.InvokeVoid(
		g,
		"resetRemoteTrafficSelector",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnel) ResetRouter() {
	_jsii_.InvokeVoid(
		g,
		"resetRouter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnel) ResetTargetVpnGateway() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetVpnGateway",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnel) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnel) ResetVpnGateway() {
	_jsii_.InvokeVoid(
		g,
		"resetVpnGateway",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnel) ResetVpnGatewayInterface() {
	_jsii_.InvokeVoid(
		g,
		"resetVpnGatewayInterface",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeVpnTunnel) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeVpnTunnel) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeVpnTunnel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeVpnTunnel) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

