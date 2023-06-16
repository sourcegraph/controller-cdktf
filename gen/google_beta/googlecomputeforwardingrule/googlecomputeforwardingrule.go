package googlecomputeforwardingrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeforwardingrule/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_compute_forwarding_rule google_compute_forwarding_rule}.
type GoogleComputeForwardingRule interface {
	cdktf.TerraformResource
	AllowGlobalAccess() interface{}
	SetAllowGlobalAccess(val interface{})
	AllowGlobalAccessInput() interface{}
	AllowPscGlobalAccess() interface{}
	SetAllowPscGlobalAccess(val interface{})
	AllowPscGlobalAccessInput() interface{}
	AllPorts() interface{}
	SetAllPorts(val interface{})
	AllPortsInput() interface{}
	BackendService() *string
	SetBackendService(val *string)
	BackendServiceInput() *string
	BaseForwardingRule() *string
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
	IpAddress() *string
	SetIpAddress(val *string)
	IpAddressInput() *string
	IpProtocol() *string
	SetIpProtocol(val *string)
	IpProtocolInput() *string
	IsMirroringCollector() interface{}
	SetIsMirroringCollector(val interface{})
	IsMirroringCollectorInput() interface{}
	LabelFingerprint() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancingScheme() *string
	SetLoadBalancingScheme(val *string)
	LoadBalancingSchemeInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	NetworkTier() *string
	SetNetworkTier(val *string)
	NetworkTierInput() *string
	// The tree node.
	Node() constructs.Node
	PortRange() *string
	SetPortRange(val *string)
	PortRangeInput() *string
	Ports() *[]*string
	SetPorts(val *[]*string)
	PortsInput() *[]*string
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
	PscConnectionId() *string
	PscConnectionStatus() *string
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SelfLink() *string
	ServiceDirectoryRegistrations() GoogleComputeForwardingRuleServiceDirectoryRegistrationsOutputReference
	ServiceDirectoryRegistrationsInput() *GoogleComputeForwardingRuleServiceDirectoryRegistrations
	ServiceLabel() *string
	SetServiceLabel(val *string)
	ServiceLabelInput() *string
	ServiceName() *string
	SourceIpRanges() *[]*string
	SetSourceIpRanges(val *[]*string)
	SourceIpRangesInput() *[]*string
	Subnetwork() *string
	SetSubnetwork(val *string)
	SubnetworkInput() *string
	Target() *string
	SetTarget(val *string)
	TargetInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleComputeForwardingRuleTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutServiceDirectoryRegistrations(value *GoogleComputeForwardingRuleServiceDirectoryRegistrations)
	PutTimeouts(value *GoogleComputeForwardingRuleTimeouts)
	ResetAllowGlobalAccess()
	ResetAllowPscGlobalAccess()
	ResetAllPorts()
	ResetBackendService()
	ResetDescription()
	ResetId()
	ResetIpAddress()
	ResetIpProtocol()
	ResetIsMirroringCollector()
	ResetLabels()
	ResetLoadBalancingScheme()
	ResetNetwork()
	ResetNetworkTier()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPortRange()
	ResetPorts()
	ResetProject()
	ResetRegion()
	ResetServiceDirectoryRegistrations()
	ResetServiceLabel()
	ResetSourceIpRanges()
	ResetSubnetwork()
	ResetTarget()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleComputeForwardingRule
type jsiiProxy_GoogleComputeForwardingRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleComputeForwardingRule) AllowGlobalAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowGlobalAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) AllowGlobalAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowGlobalAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) AllowPscGlobalAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPscGlobalAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) AllowPscGlobalAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPscGlobalAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) AllPorts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) AllPortsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) BackendService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) BackendServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) BaseForwardingRule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseForwardingRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) IpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) IpProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) IpProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) IsMirroringCollector() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMirroringCollector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) IsMirroringCollectorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMirroringCollectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) LabelFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) LoadBalancingScheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) LoadBalancingSchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingSchemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) NetworkTier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) NetworkTierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) PortRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) PortRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) Ports() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) PortsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"portsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) PscConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pscConnectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) PscConnectionStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pscConnectionStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) ServiceDirectoryRegistrations() GoogleComputeForwardingRuleServiceDirectoryRegistrationsOutputReference {
	var returns GoogleComputeForwardingRuleServiceDirectoryRegistrationsOutputReference
	_jsii_.Get(
		j,
		"serviceDirectoryRegistrations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) ServiceDirectoryRegistrationsInput() *GoogleComputeForwardingRuleServiceDirectoryRegistrations {
	var returns *GoogleComputeForwardingRuleServiceDirectoryRegistrations
	_jsii_.Get(
		j,
		"serviceDirectoryRegistrationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) ServiceLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) ServiceLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) SourceIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) SourceIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) Subnetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) SubnetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) Timeouts() GoogleComputeForwardingRuleTimeoutsOutputReference {
	var returns GoogleComputeForwardingRuleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeForwardingRule) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_compute_forwarding_rule google_compute_forwarding_rule} Resource.
func NewGoogleComputeForwardingRule(scope constructs.Construct, id *string, config *GoogleComputeForwardingRuleConfig) GoogleComputeForwardingRule {
	_init_.Initialize()

	if err := validateNewGoogleComputeForwardingRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeForwardingRule{}

	_jsii_.Create(
		"google-beta.googleComputeForwardingRule.GoogleComputeForwardingRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_compute_forwarding_rule google_compute_forwarding_rule} Resource.
func NewGoogleComputeForwardingRule_Override(g GoogleComputeForwardingRule, scope constructs.Construct, id *string, config *GoogleComputeForwardingRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeForwardingRule.GoogleComputeForwardingRule",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetAllowGlobalAccess(val interface{}) {
	if err := j.validateSetAllowGlobalAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowGlobalAccess",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetAllowPscGlobalAccess(val interface{}) {
	if err := j.validateSetAllowPscGlobalAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowPscGlobalAccess",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetAllPorts(val interface{}) {
	if err := j.validateSetAllPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allPorts",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetBackendService(val *string) {
	if err := j.validateSetBackendServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backendService",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetIpAddress(val *string) {
	if err := j.validateSetIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetIpProtocol(val *string) {
	if err := j.validateSetIpProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipProtocol",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetIsMirroringCollector(val interface{}) {
	if err := j.validateSetIsMirroringCollectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isMirroringCollector",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetLoadBalancingScheme(val *string) {
	if err := j.validateSetLoadBalancingSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancingScheme",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetNetworkTier(val *string) {
	if err := j.validateSetNetworkTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkTier",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetPortRange(val *string) {
	if err := j.validateSetPortRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portRange",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetPorts(val *[]*string) {
	if err := j.validateSetPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ports",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetServiceLabel(val *string) {
	if err := j.validateSetServiceLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceLabel",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetSourceIpRanges(val *[]*string) {
	if err := j.validateSetSourceIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceIpRanges",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetSubnetwork(val *string) {
	if err := j.validateSetSubnetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetwork",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeForwardingRule)SetTarget(val *string) {
	if err := j.validateSetTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"target",
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
func GoogleComputeForwardingRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeForwardingRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleComputeForwardingRule.GoogleComputeForwardingRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeForwardingRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeForwardingRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleComputeForwardingRule.GoogleComputeForwardingRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeForwardingRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeForwardingRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleComputeForwardingRule.GoogleComputeForwardingRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputeForwardingRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"google-beta.googleComputeForwardingRule.GoogleComputeForwardingRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeForwardingRule) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeForwardingRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeForwardingRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeForwardingRule) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeForwardingRule) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeForwardingRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeForwardingRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeForwardingRule) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeForwardingRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeForwardingRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeForwardingRule) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) PutServiceDirectoryRegistrations(value *GoogleComputeForwardingRuleServiceDirectoryRegistrations) {
	if err := g.validatePutServiceDirectoryRegistrationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServiceDirectoryRegistrations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) PutTimeouts(value *GoogleComputeForwardingRuleTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ResetAllowGlobalAccess() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowGlobalAccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ResetAllowPscGlobalAccess() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowPscGlobalAccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ResetAllPorts() {
	_jsii_.InvokeVoid(
		g,
		"resetAllPorts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ResetBackendService() {
	_jsii_.InvokeVoid(
		g,
		"resetBackendService",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ResetIpAddress() {
	_jsii_.InvokeVoid(
		g,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ResetIpProtocol() {
	_jsii_.InvokeVoid(
		g,
		"resetIpProtocol",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ResetIsMirroringCollector() {
	_jsii_.InvokeVoid(
		g,
		"resetIsMirroringCollector",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ResetLoadBalancingScheme() {
	_jsii_.InvokeVoid(
		g,
		"resetLoadBalancingScheme",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ResetNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ResetNetworkTier() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkTier",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ResetPortRange() {
	_jsii_.InvokeVoid(
		g,
		"resetPortRange",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ResetPorts() {
	_jsii_.InvokeVoid(
		g,
		"resetPorts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ResetServiceDirectoryRegistrations() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceDirectoryRegistrations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ResetServiceLabel() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceLabel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ResetSourceIpRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceIpRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ResetSubnetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetSubnetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ResetTarget() {
	_jsii_.InvokeVoid(
		g,
		"resetTarget",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeForwardingRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeForwardingRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

