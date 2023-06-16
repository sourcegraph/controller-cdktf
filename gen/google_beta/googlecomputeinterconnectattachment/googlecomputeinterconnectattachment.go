package googlecomputeinterconnectattachment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeinterconnectattachment/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_compute_interconnect_attachment google_compute_interconnect_attachment}.
type GoogleComputeInterconnectAttachment interface {
	cdktf.TerraformResource
	AdminEnabled() interface{}
	SetAdminEnabled(val interface{})
	AdminEnabledInput() interface{}
	Bandwidth() *string
	SetBandwidth(val *string)
	BandwidthInput() *string
	CandidateSubnets() *[]*string
	SetCandidateSubnets(val *[]*string)
	CandidateSubnetsInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudRouterIpAddress() *string
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
	CustomerRouterIpAddress() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EdgeAvailabilityDomain() *string
	SetEdgeAvailabilityDomain(val *string)
	EdgeAvailabilityDomainInput() *string
	Encryption() *string
	SetEncryption(val *string)
	EncryptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GoogleReferenceId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Interconnect() *string
	SetInterconnect(val *string)
	InterconnectInput() *string
	IpsecInternalAddresses() *[]*string
	SetIpsecInternalAddresses(val *[]*string)
	IpsecInternalAddressesInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Mtu() *string
	SetMtu(val *string)
	MtuInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PairingKey() *string
	PartnerAsn() *string
	PrivateInterconnectInfo() GoogleComputeInterconnectAttachmentPrivateInterconnectInfoList
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
	RouterInput() *string
	SelfLink() *string
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleComputeInterconnectAttachmentTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	VlanTag8021Q() *float64
	SetVlanTag8021Q(val *float64)
	VlanTag8021QInput() *float64
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
	PutTimeouts(value *GoogleComputeInterconnectAttachmentTimeouts)
	ResetAdminEnabled()
	ResetBandwidth()
	ResetCandidateSubnets()
	ResetDescription()
	ResetEdgeAvailabilityDomain()
	ResetEncryption()
	ResetId()
	ResetInterconnect()
	ResetIpsecInternalAddresses()
	ResetMtu()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRegion()
	ResetTimeouts()
	ResetType()
	ResetVlanTag8021Q()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleComputeInterconnectAttachment
type jsiiProxy_GoogleComputeInterconnectAttachment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) AdminEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) AdminEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) Bandwidth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) BandwidthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) CandidateSubnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"candidateSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) CandidateSubnetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"candidateSubnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) CloudRouterIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudRouterIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) CustomerRouterIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerRouterIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) EdgeAvailabilityDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeAvailabilityDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) EdgeAvailabilityDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeAvailabilityDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) Encryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) EncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) GoogleReferenceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleReferenceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) Interconnect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interconnect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) InterconnectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interconnectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) IpsecInternalAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipsecInternalAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) IpsecInternalAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipsecInternalAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) Mtu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mtu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) MtuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mtuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) PairingKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pairingKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) PartnerAsn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) PrivateInterconnectInfo() GoogleComputeInterconnectAttachmentPrivateInterconnectInfoList {
	var returns GoogleComputeInterconnectAttachmentPrivateInterconnectInfoList
	_jsii_.Get(
		j,
		"privateInterconnectInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) Router() *string {
	var returns *string
	_jsii_.Get(
		j,
		"router",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) RouterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) Timeouts() GoogleComputeInterconnectAttachmentTimeoutsOutputReference {
	var returns GoogleComputeInterconnectAttachmentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) VlanTag8021Q() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vlanTag8021Q",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment) VlanTag8021QInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vlanTag8021QInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_compute_interconnect_attachment google_compute_interconnect_attachment} Resource.
func NewGoogleComputeInterconnectAttachment(scope constructs.Construct, id *string, config *GoogleComputeInterconnectAttachmentConfig) GoogleComputeInterconnectAttachment {
	_init_.Initialize()

	if err := validateNewGoogleComputeInterconnectAttachmentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeInterconnectAttachment{}

	_jsii_.Create(
		"google-beta.googleComputeInterconnectAttachment.GoogleComputeInterconnectAttachment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_compute_interconnect_attachment google_compute_interconnect_attachment} Resource.
func NewGoogleComputeInterconnectAttachment_Override(g GoogleComputeInterconnectAttachment, scope constructs.Construct, id *string, config *GoogleComputeInterconnectAttachmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeInterconnectAttachment.GoogleComputeInterconnectAttachment",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment)SetAdminEnabled(val interface{}) {
	if err := j.validateSetAdminEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment)SetBandwidth(val *string) {
	if err := j.validateSetBandwidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidth",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment)SetCandidateSubnets(val *[]*string) {
	if err := j.validateSetCandidateSubnetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"candidateSubnets",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment)SetEdgeAvailabilityDomain(val *string) {
	if err := j.validateSetEdgeAvailabilityDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeAvailabilityDomain",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment)SetEncryption(val *string) {
	if err := j.validateSetEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryption",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment)SetInterconnect(val *string) {
	if err := j.validateSetInterconnectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interconnect",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment)SetIpsecInternalAddresses(val *[]*string) {
	if err := j.validateSetIpsecInternalAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipsecInternalAddresses",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment)SetMtu(val *string) {
	if err := j.validateSetMtuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mtu",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment)SetRouter(val *string) {
	if err := j.validateSetRouterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"router",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnectAttachment)SetVlanTag8021Q(val *float64) {
	if err := j.validateSetVlanTag8021QParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vlanTag8021Q",
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
func GoogleComputeInterconnectAttachment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeInterconnectAttachment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleComputeInterconnectAttachment.GoogleComputeInterconnectAttachment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeInterconnectAttachment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeInterconnectAttachment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleComputeInterconnectAttachment.GoogleComputeInterconnectAttachment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeInterconnectAttachment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeInterconnectAttachment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleComputeInterconnectAttachment.GoogleComputeInterconnectAttachment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputeInterconnectAttachment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"google-beta.googleComputeInterconnectAttachment.GoogleComputeInterconnectAttachment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) PutTimeouts(value *GoogleComputeInterconnectAttachmentTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) ResetAdminEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetAdminEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) ResetBandwidth() {
	_jsii_.InvokeVoid(
		g,
		"resetBandwidth",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) ResetCandidateSubnets() {
	_jsii_.InvokeVoid(
		g,
		"resetCandidateSubnets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) ResetEdgeAvailabilityDomain() {
	_jsii_.InvokeVoid(
		g,
		"resetEdgeAvailabilityDomain",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) ResetEncryption() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryption",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) ResetInterconnect() {
	_jsii_.InvokeVoid(
		g,
		"resetInterconnect",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) ResetIpsecInternalAddresses() {
	_jsii_.InvokeVoid(
		g,
		"resetIpsecInternalAddresses",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) ResetMtu() {
	_jsii_.InvokeVoid(
		g,
		"resetMtu",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) ResetType() {
	_jsii_.InvokeVoid(
		g,
		"resetType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) ResetVlanTag8021Q() {
	_jsii_.InvokeVoid(
		g,
		"resetVlanTag8021Q",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInterconnectAttachment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

