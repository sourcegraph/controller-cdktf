package computerouterpeer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/computerouterpeer/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_router_peer google_compute_router_peer}.
type ComputeRouterPeer interface {
	cdktf.TerraformResource
	AdvertisedGroups() *[]*string
	SetAdvertisedGroups(val *[]*string)
	AdvertisedGroupsInput() *[]*string
	AdvertisedIpRanges() ComputeRouterPeerAdvertisedIpRangesList
	AdvertisedIpRangesInput() any
	AdvertisedRoutePriority() *float64
	SetAdvertisedRoutePriority(val *float64)
	AdvertisedRoutePriorityInput() *float64
	AdvertiseMode() *string
	SetAdvertiseMode(val *string)
	AdvertiseModeInput() *string
	Bfd() ComputeRouterPeerBfdOutputReference
	BfdInput() *ComputeRouterPeerBfd
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
	CustomLearnedIpRanges() ComputeRouterPeerCustomLearnedIpRangesList
	CustomLearnedIpRangesInput() any
	CustomLearnedRoutePriority() *float64
	SetCustomLearnedRoutePriority(val *float64)
	CustomLearnedRoutePriorityInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Enable() any
	SetEnable(val any)
	EnableInput() any
	EnableIpv4() any
	SetEnableIpv4(val any)
	EnableIpv4Input() any
	EnableIpv6() any
	SetEnableIpv6(val any)
	EnableIpv6Input() any
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
	Md5AuthenticationKey() ComputeRouterPeerMd5AuthenticationKeyOutputReference
	Md5AuthenticationKeyInput() *ComputeRouterPeerMd5AuthenticationKey
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
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	// Experimental.
	RawOverrides() any
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
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ComputeRouterPeerTimeoutsOutputReference
	TimeoutsInput() any
	ZeroAdvertisedRoutePriority() any
	SetZeroAdvertisedRoutePriority(val any)
	ZeroAdvertisedRoutePriorityInput() any
	ZeroCustomLearnedRoutePriority() any
	SetZeroCustomLearnedRoutePriority(val any)
	ZeroCustomLearnedRoutePriorityInput() any
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
	PutAdvertisedIpRanges(value any)
	PutBfd(value *ComputeRouterPeerBfd)
	PutCustomLearnedIpRanges(value any)
	PutMd5AuthenticationKey(value *ComputeRouterPeerMd5AuthenticationKey)
	PutTimeouts(value *ComputeRouterPeerTimeouts)
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

// The jsii proxy struct for ComputeRouterPeer
type jsiiProxy_ComputeRouterPeer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComputeRouterPeer) AdvertisedGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"advertisedGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) AdvertisedGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"advertisedGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) AdvertisedIpRanges() ComputeRouterPeerAdvertisedIpRangesList {
	var returns ComputeRouterPeerAdvertisedIpRangesList
	_jsii_.Get(
		j,
		"advertisedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) AdvertisedIpRangesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"advertisedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) AdvertisedRoutePriority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"advertisedRoutePriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) AdvertisedRoutePriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"advertisedRoutePriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) AdvertiseMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"advertiseMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) AdvertiseModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"advertiseModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) Bfd() ComputeRouterPeerBfdOutputReference {
	var returns ComputeRouterPeerBfdOutputReference
	_jsii_.Get(
		j,
		"bfd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) BfdInput() *ComputeRouterPeerBfd {
	var returns *ComputeRouterPeerBfd
	_jsii_.Get(
		j,
		"bfdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) CustomLearnedIpRanges() ComputeRouterPeerCustomLearnedIpRangesList {
	var returns ComputeRouterPeerCustomLearnedIpRangesList
	_jsii_.Get(
		j,
		"customLearnedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) CustomLearnedIpRangesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"customLearnedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) CustomLearnedRoutePriority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customLearnedRoutePriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) CustomLearnedRoutePriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customLearnedRoutePriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) Enable() any {
	var returns any
	_jsii_.Get(
		j,
		"enable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) EnableInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) EnableIpv4() any {
	var returns any
	_jsii_.Get(
		j,
		"enableIpv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) EnableIpv4Input() any {
	var returns any
	_jsii_.Get(
		j,
		"enableIpv4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) EnableIpv6() any {
	var returns any
	_jsii_.Get(
		j,
		"enableIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) EnableIpv6Input() any {
	var returns any
	_jsii_.Get(
		j,
		"enableIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) ExportPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exportPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) ExportPoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exportPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) ImportPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"importPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) ImportPoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"importPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) Interface() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) InterfaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) IpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) Ipv4NexthopAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4NexthopAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) Ipv4NexthopAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4NexthopAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) Ipv6NexthopAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6NexthopAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) Ipv6NexthopAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6NexthopAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) IsAdvertisedRoutePrioritySet() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isAdvertisedRoutePrioritySet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) IsCustomLearnedPrioritySet() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isCustomLearnedPrioritySet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) ManagementType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managementType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) Md5AuthenticationKey() ComputeRouterPeerMd5AuthenticationKeyOutputReference {
	var returns ComputeRouterPeerMd5AuthenticationKeyOutputReference
	_jsii_.Get(
		j,
		"md5AuthenticationKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) Md5AuthenticationKeyInput() *ComputeRouterPeerMd5AuthenticationKey {
	var returns *ComputeRouterPeerMd5AuthenticationKey
	_jsii_.Get(
		j,
		"md5AuthenticationKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) PeerAsn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"peerAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) PeerAsnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"peerAsnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) PeerIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) PeerIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) PeerIpv4NexthopAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIpv4NexthopAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) PeerIpv4NexthopAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIpv4NexthopAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) PeerIpv6NexthopAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIpv6NexthopAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) PeerIpv6NexthopAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIpv6NexthopAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) Router() *string {
	var returns *string
	_jsii_.Get(
		j,
		"router",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) RouterApplianceInstance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routerApplianceInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) RouterApplianceInstanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routerApplianceInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) RouterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) Timeouts() ComputeRouterPeerTimeoutsOutputReference {
	var returns ComputeRouterPeerTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) ZeroAdvertisedRoutePriority() any {
	var returns any
	_jsii_.Get(
		j,
		"zeroAdvertisedRoutePriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) ZeroAdvertisedRoutePriorityInput() any {
	var returns any
	_jsii_.Get(
		j,
		"zeroAdvertisedRoutePriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) ZeroCustomLearnedRoutePriority() any {
	var returns any
	_jsii_.Get(
		j,
		"zeroCustomLearnedRoutePriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeer) ZeroCustomLearnedRoutePriorityInput() any {
	var returns any
	_jsii_.Get(
		j,
		"zeroCustomLearnedRoutePriorityInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_router_peer google_compute_router_peer} Resource.
func NewComputeRouterPeer(scope constructs.Construct, id *string, config *ComputeRouterPeerConfig) ComputeRouterPeer {
	_init_.Initialize()

	if err := validateNewComputeRouterPeerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRouterPeer{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRouterPeer.ComputeRouterPeer",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_router_peer google_compute_router_peer} Resource.
func NewComputeRouterPeer_Override(c ComputeRouterPeer, scope constructs.Construct, id *string, config *ComputeRouterPeerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRouterPeer.ComputeRouterPeer",
		[]any{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetAdvertisedGroups(val *[]*string) {
	if err := j.validateSetAdvertisedGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advertisedGroups",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetAdvertisedRoutePriority(val *float64) {
	if err := j.validateSetAdvertisedRoutePriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advertisedRoutePriority",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetAdvertiseMode(val *string) {
	if err := j.validateSetAdvertiseModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advertiseMode",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetCustomLearnedRoutePriority(val *float64) {
	if err := j.validateSetCustomLearnedRoutePriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customLearnedRoutePriority",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetEnable(val any) {
	if err := j.validateSetEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enable",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetEnableIpv4(val any) {
	if err := j.validateSetEnableIpv4Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableIpv4",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetEnableIpv6(val any) {
	if err := j.validateSetEnableIpv6Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableIpv6",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetExportPolicies(val *[]*string) {
	if err := j.validateSetExportPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportPolicies",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetImportPolicies(val *[]*string) {
	if err := j.validateSetImportPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importPolicies",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetInterface(val *string) {
	if err := j.validateSetInterfaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interface",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetIpAddress(val *string) {
	if err := j.validateSetIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetIpv4NexthopAddress(val *string) {
	if err := j.validateSetIpv4NexthopAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv4NexthopAddress",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetIpv6NexthopAddress(val *string) {
	if err := j.validateSetIpv6NexthopAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6NexthopAddress",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetPeerAsn(val *float64) {
	if err := j.validateSetPeerAsnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerAsn",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetPeerIpAddress(val *string) {
	if err := j.validateSetPeerIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerIpAddress",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetPeerIpv4NexthopAddress(val *string) {
	if err := j.validateSetPeerIpv4NexthopAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerIpv4NexthopAddress",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetPeerIpv6NexthopAddress(val *string) {
	if err := j.validateSetPeerIpv6NexthopAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerIpv6NexthopAddress",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetRouter(val *string) {
	if err := j.validateSetRouterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"router",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetRouterApplianceInstance(val *string) {
	if err := j.validateSetRouterApplianceInstanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routerApplianceInstance",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetZeroAdvertisedRoutePriority(val any) {
	if err := j.validateSetZeroAdvertisedRoutePriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zeroAdvertisedRoutePriority",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeer) SetZeroCustomLearnedRoutePriority(val any) {
	if err := j.validateSetZeroCustomLearnedRoutePriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zeroCustomLearnedRoutePriority",
		val,
	)
}

// Generates CDKTF code for importing a ComputeRouterPeer resource upon running "cdktf plan <stack-name>".
func ComputeRouterPeer_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateComputeRouterPeer_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRouterPeer.ComputeRouterPeer",
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
func ComputeRouterPeer_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateComputeRouterPeer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRouterPeer.ComputeRouterPeer",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeRouterPeer_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateComputeRouterPeer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRouterPeer.ComputeRouterPeer",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeRouterPeer_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateComputeRouterPeer_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRouterPeer.ComputeRouterPeer",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func ComputeRouterPeer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.computeRouterPeer.ComputeRouterPeer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeRouterPeer) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (c *jsiiProxy_ComputeRouterPeer) AddOverride(path *string, value any) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]any{path, value},
	)
}

func (c *jsiiProxy_ComputeRouterPeer) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterPeer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterPeer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterPeer) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterPeer) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterPeer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterPeer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterPeer) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterPeer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterPeer) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterPeer) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]any{id, provider},
	)
}

func (c *jsiiProxy_ComputeRouterPeer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterPeer) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]any{id},
	)
}

func (c *jsiiProxy_ComputeRouterPeer) MoveTo(moveTarget *string, index any) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (c *jsiiProxy_ComputeRouterPeer) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]any{id},
	)
}

func (c *jsiiProxy_ComputeRouterPeer) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeRouterPeer) PutAdvertisedIpRanges(value any) {
	if err := c.validatePutAdvertisedIpRangesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAdvertisedIpRanges",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeRouterPeer) PutBfd(value *ComputeRouterPeerBfd) {
	if err := c.validatePutBfdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBfd",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeRouterPeer) PutCustomLearnedIpRanges(value any) {
	if err := c.validatePutCustomLearnedIpRangesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCustomLearnedIpRanges",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeRouterPeer) PutMd5AuthenticationKey(value *ComputeRouterPeerMd5AuthenticationKey) {
	if err := c.validatePutMd5AuthenticationKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMd5AuthenticationKey",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeRouterPeer) PutTimeouts(value *ComputeRouterPeerTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetAdvertisedGroups() {
	_jsii_.InvokeVoid(
		c,
		"resetAdvertisedGroups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetAdvertisedIpRanges() {
	_jsii_.InvokeVoid(
		c,
		"resetAdvertisedIpRanges",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetAdvertisedRoutePriority() {
	_jsii_.InvokeVoid(
		c,
		"resetAdvertisedRoutePriority",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetAdvertiseMode() {
	_jsii_.InvokeVoid(
		c,
		"resetAdvertiseMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetBfd() {
	_jsii_.InvokeVoid(
		c,
		"resetBfd",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetCustomLearnedIpRanges() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomLearnedIpRanges",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetCustomLearnedRoutePriority() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomLearnedRoutePriority",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetEnable() {
	_jsii_.InvokeVoid(
		c,
		"resetEnable",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetEnableIpv4() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableIpv4",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetEnableIpv6() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableIpv6",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetExportPolicies() {
	_jsii_.InvokeVoid(
		c,
		"resetExportPolicies",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetImportPolicies() {
	_jsii_.InvokeVoid(
		c,
		"resetImportPolicies",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetIpAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetIpv4NexthopAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetIpv4NexthopAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetIpv6NexthopAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetIpv6NexthopAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetMd5AuthenticationKey() {
	_jsii_.InvokeVoid(
		c,
		"resetMd5AuthenticationKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetPeerIpAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetPeerIpAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetPeerIpv4NexthopAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetPeerIpv4NexthopAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetPeerIpv6NexthopAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetPeerIpv6NexthopAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetRouterApplianceInstance() {
	_jsii_.InvokeVoid(
		c,
		"resetRouterApplianceInstance",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetZeroAdvertisedRoutePriority() {
	_jsii_.InvokeVoid(
		c,
		"resetZeroAdvertisedRoutePriority",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) ResetZeroCustomLearnedRoutePriority() {
	_jsii_.InvokeVoid(
		c,
		"resetZeroCustomLearnedRoutePriority",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeer) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterPeer) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterPeer) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterPeer) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterPeer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterPeer) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
