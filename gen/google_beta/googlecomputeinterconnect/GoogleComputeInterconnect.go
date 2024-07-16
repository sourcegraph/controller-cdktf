package googlecomputeinterconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeinterconnect/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_interconnect google_compute_interconnect}.
type GoogleComputeInterconnect interface {
	cdktf.TerraformResource
	AdminEnabled() interface{}
	SetAdminEnabled(val interface{})
	AdminEnabledInput() interface{}
	AvailableFeatures() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CircuitInfos() GoogleComputeInterconnectCircuitInfosList
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
	CustomerName() *string
	SetCustomerName(val *string)
	CustomerNameInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveLabels() cdktf.StringMap
	ExpectedOutages() GoogleComputeInterconnectExpectedOutagesList
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GoogleIpAddress() *string
	GoogleReferenceId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InterconnectAttachments() *[]*string
	InterconnectType() *string
	SetInterconnectType(val *string)
	InterconnectTypeInput() *string
	LabelFingerprint() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LinkType() *string
	SetLinkType(val *string)
	LinkTypeInput() *string
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Macsec() GoogleComputeInterconnectMacsecOutputReference
	MacsecEnabled() interface{}
	SetMacsecEnabled(val interface{})
	MacsecEnabledInput() interface{}
	MacsecInput() *GoogleComputeInterconnectMacsec
	Name() *string
	SetName(val *string)
	NameInput() *string
	NocContactEmail() *string
	SetNocContactEmail(val *string)
	NocContactEmailInput() *string
	// The tree node.
	Node() constructs.Node
	OperationalStatus() *string
	PeerIpAddress() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	ProvisionedLinkCount() *float64
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RemoteLocation() *string
	SetRemoteLocation(val *string)
	RemoteLocationInput() *string
	RequestedFeatures() *[]*string
	SetRequestedFeatures(val *[]*string)
	RequestedFeaturesInput() *[]*string
	RequestedLinkCount() *float64
	SetRequestedLinkCount(val *float64)
	RequestedLinkCountInput() *float64
	SatisfiesPzs() cdktf.IResolvable
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleComputeInterconnectTimeoutsOutputReference
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
	PutMacsec(value *GoogleComputeInterconnectMacsec)
	PutTimeouts(value *GoogleComputeInterconnectTimeouts)
	ResetAdminEnabled()
	ResetDescription()
	ResetId()
	ResetLabels()
	ResetMacsec()
	ResetMacsecEnabled()
	ResetNocContactEmail()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRemoteLocation()
	ResetRequestedFeatures()
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

// The jsii proxy struct for GoogleComputeInterconnect
type jsiiProxy_GoogleComputeInterconnect struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleComputeInterconnect) AdminEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) AdminEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) AvailableFeatures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) CircuitInfos() GoogleComputeInterconnectCircuitInfosList {
	var returns GoogleComputeInterconnectCircuitInfosList
	_jsii_.Get(
		j,
		"circuitInfos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) CustomerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) CustomerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) ExpectedOutages() GoogleComputeInterconnectExpectedOutagesList {
	var returns GoogleComputeInterconnectExpectedOutagesList
	_jsii_.Get(
		j,
		"expectedOutages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) GoogleIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) GoogleReferenceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleReferenceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) InterconnectAttachments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"interconnectAttachments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) InterconnectType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interconnectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) InterconnectTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interconnectTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) LabelFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) LinkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) LinkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) Macsec() GoogleComputeInterconnectMacsecOutputReference {
	var returns GoogleComputeInterconnectMacsecOutputReference
	_jsii_.Get(
		j,
		"macsec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) MacsecEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"macsecEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) MacsecEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"macsecEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) MacsecInput() *GoogleComputeInterconnectMacsec {
	var returns *GoogleComputeInterconnectMacsec
	_jsii_.Get(
		j,
		"macsecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) NocContactEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nocContactEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) NocContactEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nocContactEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) OperationalStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationalStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) PeerIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) ProvisionedLinkCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedLinkCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) RemoteLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) RemoteLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) RequestedFeatures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requestedFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) RequestedFeaturesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requestedFeaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) RequestedLinkCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestedLinkCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) RequestedLinkCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestedLinkCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) SatisfiesPzs() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"satisfiesPzs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) Timeouts() GoogleComputeInterconnectTimeoutsOutputReference {
	var returns GoogleComputeInterconnectTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInterconnect) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_interconnect google_compute_interconnect} Resource.
func NewGoogleComputeInterconnect(scope constructs.Construct, id *string, config *GoogleComputeInterconnectConfig) GoogleComputeInterconnect {
	_init_.Initialize()

	if err := validateNewGoogleComputeInterconnectParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeInterconnect{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeInterconnect.GoogleComputeInterconnect",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_interconnect google_compute_interconnect} Resource.
func NewGoogleComputeInterconnect_Override(g GoogleComputeInterconnect, scope constructs.Construct, id *string, config *GoogleComputeInterconnectConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeInterconnect.GoogleComputeInterconnect",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnect)SetAdminEnabled(val interface{}) {
	if err := j.validateSetAdminEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnect)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnect)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnect)SetCustomerName(val *string) {
	if err := j.validateSetCustomerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerName",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnect)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnect)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnect)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnect)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnect)SetInterconnectType(val *string) {
	if err := j.validateSetInterconnectTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interconnectType",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnect)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnect)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnect)SetLinkType(val *string) {
	if err := j.validateSetLinkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linkType",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnect)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnect)SetMacsecEnabled(val interface{}) {
	if err := j.validateSetMacsecEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"macsecEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnect)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnect)SetNocContactEmail(val *string) {
	if err := j.validateSetNocContactEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nocContactEmail",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnect)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnect)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnect)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnect)SetRemoteLocation(val *string) {
	if err := j.validateSetRemoteLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteLocation",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnect)SetRequestedFeatures(val *[]*string) {
	if err := j.validateSetRequestedFeaturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestedFeatures",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInterconnect)SetRequestedLinkCount(val *float64) {
	if err := j.validateSetRequestedLinkCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestedLinkCount",
		val,
	)
}

// Generates CDKTF code for importing a GoogleComputeInterconnect resource upon running "cdktf plan <stack-name>".
func GoogleComputeInterconnect_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleComputeInterconnect_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeInterconnect.GoogleComputeInterconnect",
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
func GoogleComputeInterconnect_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeInterconnect_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeInterconnect.GoogleComputeInterconnect",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeInterconnect_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeInterconnect_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeInterconnect.GoogleComputeInterconnect",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeInterconnect_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeInterconnect_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeInterconnect.GoogleComputeInterconnect",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputeInterconnect_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleComputeInterconnect.GoogleComputeInterconnect",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeInterconnect) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleComputeInterconnect) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeInterconnect) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeInterconnect) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInterconnect) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeInterconnect) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeInterconnect) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeInterconnect) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeInterconnect) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeInterconnect) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeInterconnect) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeInterconnect) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInterconnect) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleComputeInterconnect) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInterconnect) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeInterconnect) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleComputeInterconnect) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeInterconnect) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeInterconnect) PutMacsec(value *GoogleComputeInterconnectMacsec) {
	if err := g.validatePutMacsecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMacsec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInterconnect) PutTimeouts(value *GoogleComputeInterconnectTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInterconnect) ResetAdminEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetAdminEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnect) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnect) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnect) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnect) ResetMacsec() {
	_jsii_.InvokeVoid(
		g,
		"resetMacsec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnect) ResetMacsecEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetMacsecEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnect) ResetNocContactEmail() {
	_jsii_.InvokeVoid(
		g,
		"resetNocContactEmail",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnect) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnect) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnect) ResetRemoteLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetRemoteLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnect) ResetRequestedFeatures() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestedFeatures",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnect) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInterconnect) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInterconnect) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInterconnect) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInterconnect) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInterconnect) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInterconnect) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

