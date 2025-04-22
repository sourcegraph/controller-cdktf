package googlecomputerouterpeer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputerouterpeer/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_router_peer google_compute_router_peer}.
type GoogleComputeRouterPeer interface {
	cdktf.TerraformResource
	AdvertisedGroups() *[]*string
	SetAdvertisedGroups(val *[]*string)
	AdvertisedGroupsInput() *[]*string
	AdvertisedIpRanges() GoogleComputeRouterPeerAdvertisedIpRangesList
	AdvertisedIpRangesInput() interface{}
	AdvertisedRoutePriority() *float64
	SetAdvertisedRoutePriority(val *float64)
	AdvertisedRoutePriorityInput() *float64
	AdvertiseMode() *string
	SetAdvertiseMode(val *string)
	AdvertiseModeInput() *string
	Bfd() GoogleComputeRouterPeerBfdOutputReference
	BfdInput() *GoogleComputeRouterPeerBfd
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
	CustomLearnedIpRanges() GoogleComputeRouterPeerCustomLearnedIpRangesList
	CustomLearnedIpRangesInput() interface{}
	CustomLearnedRoutePriority() *float64
	SetCustomLearnedRoutePriority(val *float64)
	CustomLearnedRoutePriorityInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Enable() interface{}
	SetEnable(val interface{})
	EnableInput() interface{}
	EnableIpv4() interface{}
	SetEnableIpv4(val interface{})
	EnableIpv4Input() interface{}
	EnableIpv6() interface{}
	SetEnableIpv6(val interface{})
	EnableIpv6Input() interface{}
	ExportPolicies() *[]*string
	SetExportPolicies(val *[]*string)
	ExportPoliciesInput() *[]*string
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
	ImportPolicies() *[]*string
	SetImportPolicies(val *[]*string)
	ImportPoliciesInput() *[]*string
	Interface() *string
	SetInterface(val *string)
	InterfaceInput() *string
	IpAddress() *string
	SetIpAddress(val *string)
	IpAddressInput() *string
	Ipv4NexthopAddress() *string
	SetIpv4NexthopAddress(val *string)
	Ipv4NexthopAddressInput() *string
	Ipv6NexthopAddress() *string
	SetIpv6NexthopAddress(val *string)
	Ipv6NexthopAddressInput() *string
	IsAdvertisedRoutePrioritySet() cdktf.IResolvable
	IsCustomLearnedPrioritySet() cdktf.IResolvable
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ManagementType() *string
	Md5AuthenticationKey() GoogleComputeRouterPeerMd5AuthenticationKeyOutputReference
	Md5AuthenticationKeyInput() *GoogleComputeRouterPeerMd5AuthenticationKey
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PeerAsn() *float64
	SetPeerAsn(val *float64)
	PeerAsnInput() *float64
	PeerIpAddress() *string
	SetPeerIpAddress(val *string)
	PeerIpAddressInput() *string
	PeerIpv4NexthopAddress() *string
	SetPeerIpv4NexthopAddress(val *string)
	PeerIpv4NexthopAddressInput() *string
	PeerIpv6NexthopAddress() *string
	SetPeerIpv6NexthopAddress(val *string)
	PeerIpv6NexthopAddressInput() *string
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
	Router() *string
	SetRouter(val *string)
	RouterApplianceInstance() *string
	SetRouterApplianceInstance(val *string)
	RouterApplianceInstanceInput() *string
	RouterInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleComputeRouterPeerTimeoutsOutputReference
	TimeoutsInput() interface{}
	ZeroAdvertisedRoutePriority() interface{}
	SetZeroAdvertisedRoutePriority(val interface{})
	ZeroAdvertisedRoutePriorityInput() interface{}
	ZeroCustomLearnedRoutePriority() interface{}
	SetZeroCustomLearnedRoutePriority(val interface{})
	ZeroCustomLearnedRoutePriorityInput() interface{}
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
	PutAdvertisedIpRanges(value interface{})
	PutBfd(value *GoogleComputeRouterPeerBfd)
	PutCustomLearnedIpRanges(value interface{})
	PutMd5AuthenticationKey(value *GoogleComputeRouterPeerMd5AuthenticationKey)
	PutTimeouts(value *GoogleComputeRouterPeerTimeouts)
	ResetAdvertisedGroups()
	ResetAdvertisedIpRanges()
	ResetAdvertisedRoutePriority()
	ResetAdvertiseMode()
	ResetBfd()
	ResetCustomLearnedIpRanges()
	ResetCustomLearnedRoutePriority()
	ResetEnable()
	ResetEnableIpv4()
	ResetEnableIpv6()
	ResetExportPolicies()
	ResetId()
	ResetImportPolicies()
	ResetIpAddress()
	ResetIpv4NexthopAddress()
	ResetIpv6NexthopAddress()
	ResetMd5AuthenticationKey()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPeerIpAddress()
	ResetPeerIpv4NexthopAddress()
	ResetPeerIpv6NexthopAddress()
	ResetProject()
	ResetRegion()
	ResetRouterApplianceInstance()
	ResetTimeouts()
	ResetZeroAdvertisedRoutePriority()
	ResetZeroCustomLearnedRoutePriority()
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

// The jsii proxy struct for GoogleComputeRouterPeer
type jsiiProxy_GoogleComputeRouterPeer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleComputeRouterPeer) AdvertisedGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"advertisedGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) AdvertisedGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"advertisedGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) AdvertisedIpRanges() GoogleComputeRouterPeerAdvertisedIpRangesList {
	var returns GoogleComputeRouterPeerAdvertisedIpRangesList
	_jsii_.Get(
		j,
		"advertisedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) AdvertisedIpRangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advertisedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) AdvertisedRoutePriority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"advertisedRoutePriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) AdvertisedRoutePriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"advertisedRoutePriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) AdvertiseMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"advertiseMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) AdvertiseModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"advertiseModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) Bfd() GoogleComputeRouterPeerBfdOutputReference {
	var returns GoogleComputeRouterPeerBfdOutputReference
	_jsii_.Get(
		j,
		"bfd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) BfdInput() *GoogleComputeRouterPeerBfd {
	var returns *GoogleComputeRouterPeerBfd
	_jsii_.Get(
		j,
		"bfdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) CustomLearnedIpRanges() GoogleComputeRouterPeerCustomLearnedIpRangesList {
	var returns GoogleComputeRouterPeerCustomLearnedIpRangesList
	_jsii_.Get(
		j,
		"customLearnedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) CustomLearnedIpRangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customLearnedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) CustomLearnedRoutePriority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customLearnedRoutePriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) CustomLearnedRoutePriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customLearnedRoutePriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) Enable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) EnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) EnableIpv4() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIpv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) EnableIpv4Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIpv4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) EnableIpv6() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) EnableIpv6Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) ExportPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exportPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) ExportPoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exportPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) ImportPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"importPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) ImportPoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"importPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) Interface() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) InterfaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) IpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) Ipv4NexthopAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4NexthopAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) Ipv4NexthopAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4NexthopAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) Ipv6NexthopAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6NexthopAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) Ipv6NexthopAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6NexthopAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) IsAdvertisedRoutePrioritySet() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isAdvertisedRoutePrioritySet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) IsCustomLearnedPrioritySet() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isCustomLearnedPrioritySet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) ManagementType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managementType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) Md5AuthenticationKey() GoogleComputeRouterPeerMd5AuthenticationKeyOutputReference {
	var returns GoogleComputeRouterPeerMd5AuthenticationKeyOutputReference
	_jsii_.Get(
		j,
		"md5AuthenticationKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) Md5AuthenticationKeyInput() *GoogleComputeRouterPeerMd5AuthenticationKey {
	var returns *GoogleComputeRouterPeerMd5AuthenticationKey
	_jsii_.Get(
		j,
		"md5AuthenticationKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) PeerAsn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"peerAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) PeerAsnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"peerAsnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) PeerIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) PeerIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) PeerIpv4NexthopAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIpv4NexthopAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) PeerIpv4NexthopAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIpv4NexthopAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) PeerIpv6NexthopAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIpv6NexthopAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) PeerIpv6NexthopAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIpv6NexthopAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) Router() *string {
	var returns *string
	_jsii_.Get(
		j,
		"router",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) RouterApplianceInstance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routerApplianceInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) RouterApplianceInstanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routerApplianceInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) RouterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) Timeouts() GoogleComputeRouterPeerTimeoutsOutputReference {
	var returns GoogleComputeRouterPeerTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) ZeroAdvertisedRoutePriority() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zeroAdvertisedRoutePriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) ZeroAdvertisedRoutePriorityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zeroAdvertisedRoutePriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) ZeroCustomLearnedRoutePriority() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zeroCustomLearnedRoutePriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterPeer) ZeroCustomLearnedRoutePriorityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zeroCustomLearnedRoutePriorityInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_router_peer google_compute_router_peer} Resource.
func NewGoogleComputeRouterPeer(scope constructs.Construct, id *string, config *GoogleComputeRouterPeerConfig) GoogleComputeRouterPeer {
	_init_.Initialize()

	if err := validateNewGoogleComputeRouterPeerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRouterPeer{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeRouterPeer.GoogleComputeRouterPeer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_router_peer google_compute_router_peer} Resource.
func NewGoogleComputeRouterPeer_Override(g GoogleComputeRouterPeer, scope constructs.Construct, id *string, config *GoogleComputeRouterPeerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeRouterPeer.GoogleComputeRouterPeer",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetAdvertisedGroups(val *[]*string) {
	if err := j.validateSetAdvertisedGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advertisedGroups",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetAdvertisedRoutePriority(val *float64) {
	if err := j.validateSetAdvertisedRoutePriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advertisedRoutePriority",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetAdvertiseMode(val *string) {
	if err := j.validateSetAdvertiseModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advertiseMode",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetCustomLearnedRoutePriority(val *float64) {
	if err := j.validateSetCustomLearnedRoutePriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customLearnedRoutePriority",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetEnable(val interface{}) {
	if err := j.validateSetEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enable",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetEnableIpv4(val interface{}) {
	if err := j.validateSetEnableIpv4Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableIpv4",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetEnableIpv6(val interface{}) {
	if err := j.validateSetEnableIpv6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableIpv6",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetExportPolicies(val *[]*string) {
	if err := j.validateSetExportPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportPolicies",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetImportPolicies(val *[]*string) {
	if err := j.validateSetImportPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importPolicies",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetInterface(val *string) {
	if err := j.validateSetInterfaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interface",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetIpAddress(val *string) {
	if err := j.validateSetIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetIpv4NexthopAddress(val *string) {
	if err := j.validateSetIpv4NexthopAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4NexthopAddress",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetIpv6NexthopAddress(val *string) {
	if err := j.validateSetIpv6NexthopAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6NexthopAddress",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetPeerAsn(val *float64) {
	if err := j.validateSetPeerAsnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerAsn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetPeerIpAddress(val *string) {
	if err := j.validateSetPeerIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerIpAddress",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetPeerIpv4NexthopAddress(val *string) {
	if err := j.validateSetPeerIpv4NexthopAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerIpv4NexthopAddress",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetPeerIpv6NexthopAddress(val *string) {
	if err := j.validateSetPeerIpv6NexthopAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerIpv6NexthopAddress",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetRouter(val *string) {
	if err := j.validateSetRouterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"router",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetRouterApplianceInstance(val *string) {
	if err := j.validateSetRouterApplianceInstanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routerApplianceInstance",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetZeroAdvertisedRoutePriority(val interface{}) {
	if err := j.validateSetZeroAdvertisedRoutePriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zeroAdvertisedRoutePriority",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterPeer)SetZeroCustomLearnedRoutePriority(val interface{}) {
	if err := j.validateSetZeroCustomLearnedRoutePriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zeroCustomLearnedRoutePriority",
		val,
	)
}

// Generates CDKTF code for importing a GoogleComputeRouterPeer resource upon running "cdktf plan <stack-name>".
func GoogleComputeRouterPeer_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleComputeRouterPeer_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeRouterPeer.GoogleComputeRouterPeer",
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
func GoogleComputeRouterPeer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeRouterPeer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeRouterPeer.GoogleComputeRouterPeer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeRouterPeer_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeRouterPeer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeRouterPeer.GoogleComputeRouterPeer",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeRouterPeer_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeRouterPeer_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeRouterPeer.GoogleComputeRouterPeer",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputeRouterPeer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleComputeRouterPeer.GoogleComputeRouterPeer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeRouterPeer) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRouterPeer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRouterPeer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRouterPeer) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRouterPeer) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRouterPeer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRouterPeer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRouterPeer) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRouterPeer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRouterPeer) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRouterPeer) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) PutAdvertisedIpRanges(value interface{}) {
	if err := g.validatePutAdvertisedIpRangesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdvertisedIpRanges",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) PutBfd(value *GoogleComputeRouterPeerBfd) {
	if err := g.validatePutBfdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBfd",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) PutCustomLearnedIpRanges(value interface{}) {
	if err := g.validatePutCustomLearnedIpRangesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCustomLearnedIpRanges",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) PutMd5AuthenticationKey(value *GoogleComputeRouterPeerMd5AuthenticationKey) {
	if err := g.validatePutMd5AuthenticationKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMd5AuthenticationKey",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) PutTimeouts(value *GoogleComputeRouterPeerTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetAdvertisedGroups() {
	_jsii_.InvokeVoid(
		g,
		"resetAdvertisedGroups",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetAdvertisedIpRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetAdvertisedIpRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetAdvertisedRoutePriority() {
	_jsii_.InvokeVoid(
		g,
		"resetAdvertisedRoutePriority",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetAdvertiseMode() {
	_jsii_.InvokeVoid(
		g,
		"resetAdvertiseMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetBfd() {
	_jsii_.InvokeVoid(
		g,
		"resetBfd",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetCustomLearnedIpRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomLearnedIpRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetCustomLearnedRoutePriority() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomLearnedRoutePriority",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetEnable() {
	_jsii_.InvokeVoid(
		g,
		"resetEnable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetEnableIpv4() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableIpv4",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetEnableIpv6() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableIpv6",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetExportPolicies() {
	_jsii_.InvokeVoid(
		g,
		"resetExportPolicies",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetImportPolicies() {
	_jsii_.InvokeVoid(
		g,
		"resetImportPolicies",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetIpAddress() {
	_jsii_.InvokeVoid(
		g,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetIpv4NexthopAddress() {
	_jsii_.InvokeVoid(
		g,
		"resetIpv4NexthopAddress",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetIpv6NexthopAddress() {
	_jsii_.InvokeVoid(
		g,
		"resetIpv6NexthopAddress",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetMd5AuthenticationKey() {
	_jsii_.InvokeVoid(
		g,
		"resetMd5AuthenticationKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetPeerIpAddress() {
	_jsii_.InvokeVoid(
		g,
		"resetPeerIpAddress",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetPeerIpv4NexthopAddress() {
	_jsii_.InvokeVoid(
		g,
		"resetPeerIpv4NexthopAddress",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetPeerIpv6NexthopAddress() {
	_jsii_.InvokeVoid(
		g,
		"resetPeerIpv6NexthopAddress",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetRouterApplianceInstance() {
	_jsii_.InvokeVoid(
		g,
		"resetRouterApplianceInstance",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetZeroAdvertisedRoutePriority() {
	_jsii_.InvokeVoid(
		g,
		"resetZeroAdvertisedRoutePriority",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ResetZeroCustomLearnedRoutePriority() {
	_jsii_.InvokeVoid(
		g,
		"resetZeroCustomLearnedRoutePriority",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterPeer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRouterPeer) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRouterPeer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

