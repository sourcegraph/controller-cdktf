package googlecomputerouternat

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputerouternat/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_compute_router_nat google_compute_router_nat}.
type GoogleComputeRouterNat interface {
	cdktf.TerraformResource
	AutoNetworkTier() *string
	SetAutoNetworkTier(val *string)
	AutoNetworkTierInput() *string
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
	DrainNatIps() *[]*string
	SetDrainNatIps(val *[]*string)
	DrainNatIpsInput() *[]*string
	EnableDynamicPortAllocation() interface{}
	SetEnableDynamicPortAllocation(val interface{})
	EnableDynamicPortAllocationInput() interface{}
	EnableEndpointIndependentMapping() interface{}
	SetEnableEndpointIndependentMapping(val interface{})
	EnableEndpointIndependentMappingInput() interface{}
	EndpointTypes() *[]*string
	SetEndpointTypes(val *[]*string)
	EndpointTypesInput() *[]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IcmpIdleTimeoutSec() *float64
	SetIcmpIdleTimeoutSec(val *float64)
	IcmpIdleTimeoutSecInput() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogConfig() GoogleComputeRouterNatLogConfigOutputReference
	LogConfigInput() *GoogleComputeRouterNatLogConfig
	MaxPortsPerVm() *float64
	SetMaxPortsPerVm(val *float64)
	MaxPortsPerVmInput() *float64
	MinPortsPerVm() *float64
	SetMinPortsPerVm(val *float64)
	MinPortsPerVmInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NatIpAllocateOption() *string
	SetNatIpAllocateOption(val *string)
	NatIpAllocateOptionInput() *string
	NatIps() *[]*string
	SetNatIps(val *[]*string)
	NatIpsInput() *[]*string
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	Router() *string
	SetRouter(val *string)
	RouterInput() *string
	Rules() GoogleComputeRouterNatRulesList
	RulesInput() interface{}
	SourceSubnetworkIpRangesToNat() *string
	SetSourceSubnetworkIpRangesToNat(val *string)
	SourceSubnetworkIpRangesToNatInput() *string
	Subnetwork() GoogleComputeRouterNatSubnetworkList
	SubnetworkInput() interface{}
	TcpEstablishedIdleTimeoutSec() *float64
	SetTcpEstablishedIdleTimeoutSec(val *float64)
	TcpEstablishedIdleTimeoutSecInput() *float64
	TcpTimeWaitTimeoutSec() *float64
	SetTcpTimeWaitTimeoutSec(val *float64)
	TcpTimeWaitTimeoutSecInput() *float64
	TcpTransitoryIdleTimeoutSec() *float64
	SetTcpTransitoryIdleTimeoutSec(val *float64)
	TcpTransitoryIdleTimeoutSecInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleComputeRouterNatTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UdpIdleTimeoutSec() *float64
	SetUdpIdleTimeoutSec(val *float64)
	UdpIdleTimeoutSecInput() *float64
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
	PutLogConfig(value *GoogleComputeRouterNatLogConfig)
	PutRules(value interface{})
	PutSubnetwork(value interface{})
	PutTimeouts(value *GoogleComputeRouterNatTimeouts)
	ResetAutoNetworkTier()
	ResetDrainNatIps()
	ResetEnableDynamicPortAllocation()
	ResetEnableEndpointIndependentMapping()
	ResetEndpointTypes()
	ResetIcmpIdleTimeoutSec()
	ResetId()
	ResetLogConfig()
	ResetMaxPortsPerVm()
	ResetMinPortsPerVm()
	ResetNatIpAllocateOption()
	ResetNatIps()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRegion()
	ResetRules()
	ResetSubnetwork()
	ResetTcpEstablishedIdleTimeoutSec()
	ResetTcpTimeWaitTimeoutSec()
	ResetTcpTransitoryIdleTimeoutSec()
	ResetTimeouts()
	ResetType()
	ResetUdpIdleTimeoutSec()
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

// The jsii proxy struct for GoogleComputeRouterNat
type jsiiProxy_GoogleComputeRouterNat struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleComputeRouterNat) AutoNetworkTier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoNetworkTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) AutoNetworkTierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoNetworkTierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) DrainNatIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"drainNatIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) DrainNatIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"drainNatIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) EnableDynamicPortAllocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDynamicPortAllocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) EnableDynamicPortAllocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDynamicPortAllocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) EnableEndpointIndependentMapping() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEndpointIndependentMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) EnableEndpointIndependentMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEndpointIndependentMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) EndpointTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"endpointTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) EndpointTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"endpointTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) IcmpIdleTimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"icmpIdleTimeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) IcmpIdleTimeoutSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"icmpIdleTimeoutSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) LogConfig() GoogleComputeRouterNatLogConfigOutputReference {
	var returns GoogleComputeRouterNatLogConfigOutputReference
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) LogConfigInput() *GoogleComputeRouterNatLogConfig {
	var returns *GoogleComputeRouterNatLogConfig
	_jsii_.Get(
		j,
		"logConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) MaxPortsPerVm() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPortsPerVm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) MaxPortsPerVmInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPortsPerVmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) MinPortsPerVm() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minPortsPerVm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) MinPortsPerVmInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minPortsPerVmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) NatIpAllocateOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natIpAllocateOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) NatIpAllocateOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natIpAllocateOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) NatIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"natIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) NatIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"natIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) Router() *string {
	var returns *string
	_jsii_.Get(
		j,
		"router",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) RouterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) Rules() GoogleComputeRouterNatRulesList {
	var returns GoogleComputeRouterNatRulesList
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) RulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) SourceSubnetworkIpRangesToNat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSubnetworkIpRangesToNat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) SourceSubnetworkIpRangesToNatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSubnetworkIpRangesToNatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) Subnetwork() GoogleComputeRouterNatSubnetworkList {
	var returns GoogleComputeRouterNatSubnetworkList
	_jsii_.Get(
		j,
		"subnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) SubnetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subnetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) TcpEstablishedIdleTimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpEstablishedIdleTimeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) TcpEstablishedIdleTimeoutSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpEstablishedIdleTimeoutSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) TcpTimeWaitTimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpTimeWaitTimeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) TcpTimeWaitTimeoutSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpTimeWaitTimeoutSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) TcpTransitoryIdleTimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpTransitoryIdleTimeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) TcpTransitoryIdleTimeoutSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpTransitoryIdleTimeoutSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) Timeouts() GoogleComputeRouterNatTimeoutsOutputReference {
	var returns GoogleComputeRouterNatTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) UdpIdleTimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"udpIdleTimeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRouterNat) UdpIdleTimeoutSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"udpIdleTimeoutSecInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_compute_router_nat google_compute_router_nat} Resource.
func NewGoogleComputeRouterNat(scope constructs.Construct, id *string, config *GoogleComputeRouterNatConfig) GoogleComputeRouterNat {
	_init_.Initialize()

	if err := validateNewGoogleComputeRouterNatParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRouterNat{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeRouterNat.GoogleComputeRouterNat",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_compute_router_nat google_compute_router_nat} Resource.
func NewGoogleComputeRouterNat_Override(g GoogleComputeRouterNat, scope constructs.Construct, id *string, config *GoogleComputeRouterNatConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeRouterNat.GoogleComputeRouterNat",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetAutoNetworkTier(val *string) {
	if err := j.validateSetAutoNetworkTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoNetworkTier",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetDrainNatIps(val *[]*string) {
	if err := j.validateSetDrainNatIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drainNatIps",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetEnableDynamicPortAllocation(val interface{}) {
	if err := j.validateSetEnableDynamicPortAllocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDynamicPortAllocation",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetEnableEndpointIndependentMapping(val interface{}) {
	if err := j.validateSetEnableEndpointIndependentMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableEndpointIndependentMapping",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetEndpointTypes(val *[]*string) {
	if err := j.validateSetEndpointTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointTypes",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetIcmpIdleTimeoutSec(val *float64) {
	if err := j.validateSetIcmpIdleTimeoutSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"icmpIdleTimeoutSec",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetMaxPortsPerVm(val *float64) {
	if err := j.validateSetMaxPortsPerVmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPortsPerVm",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetMinPortsPerVm(val *float64) {
	if err := j.validateSetMinPortsPerVmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minPortsPerVm",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetNatIpAllocateOption(val *string) {
	if err := j.validateSetNatIpAllocateOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"natIpAllocateOption",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetNatIps(val *[]*string) {
	if err := j.validateSetNatIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"natIps",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetRouter(val *string) {
	if err := j.validateSetRouterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"router",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetSourceSubnetworkIpRangesToNat(val *string) {
	if err := j.validateSetSourceSubnetworkIpRangesToNatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceSubnetworkIpRangesToNat",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetTcpEstablishedIdleTimeoutSec(val *float64) {
	if err := j.validateSetTcpEstablishedIdleTimeoutSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcpEstablishedIdleTimeoutSec",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetTcpTimeWaitTimeoutSec(val *float64) {
	if err := j.validateSetTcpTimeWaitTimeoutSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcpTimeWaitTimeoutSec",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetTcpTransitoryIdleTimeoutSec(val *float64) {
	if err := j.validateSetTcpTransitoryIdleTimeoutSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcpTransitoryIdleTimeoutSec",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRouterNat)SetUdpIdleTimeoutSec(val *float64) {
	if err := j.validateSetUdpIdleTimeoutSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"udpIdleTimeoutSec",
		val,
	)
}

// Generates CDKTF code for importing a GoogleComputeRouterNat resource upon running "cdktf plan <stack-name>".
func GoogleComputeRouterNat_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleComputeRouterNat_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeRouterNat.GoogleComputeRouterNat",
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
func GoogleComputeRouterNat_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeRouterNat_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeRouterNat.GoogleComputeRouterNat",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeRouterNat_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeRouterNat_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeRouterNat.GoogleComputeRouterNat",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeRouterNat_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeRouterNat_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeRouterNat.GoogleComputeRouterNat",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputeRouterNat_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleComputeRouterNat.GoogleComputeRouterNat",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeRouterNat) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRouterNat) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRouterNat) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRouterNat) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRouterNat) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRouterNat) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRouterNat) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRouterNat) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRouterNat) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRouterNat) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRouterNat) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRouterNat) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) PutLogConfig(value *GoogleComputeRouterNatLogConfig) {
	if err := g.validatePutLogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLogConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) PutRules(value interface{}) {
	if err := g.validatePutRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRules",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) PutSubnetwork(value interface{}) {
	if err := g.validatePutSubnetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSubnetwork",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) PutTimeouts(value *GoogleComputeRouterNatTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) ResetAutoNetworkTier() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoNetworkTier",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) ResetDrainNatIps() {
	_jsii_.InvokeVoid(
		g,
		"resetDrainNatIps",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) ResetEnableDynamicPortAllocation() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableDynamicPortAllocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) ResetEnableEndpointIndependentMapping() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableEndpointIndependentMapping",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) ResetEndpointTypes() {
	_jsii_.InvokeVoid(
		g,
		"resetEndpointTypes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) ResetIcmpIdleTimeoutSec() {
	_jsii_.InvokeVoid(
		g,
		"resetIcmpIdleTimeoutSec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) ResetLogConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLogConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) ResetMaxPortsPerVm() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxPortsPerVm",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) ResetMinPortsPerVm() {
	_jsii_.InvokeVoid(
		g,
		"resetMinPortsPerVm",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) ResetNatIpAllocateOption() {
	_jsii_.InvokeVoid(
		g,
		"resetNatIpAllocateOption",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) ResetNatIps() {
	_jsii_.InvokeVoid(
		g,
		"resetNatIps",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) ResetRules() {
	_jsii_.InvokeVoid(
		g,
		"resetRules",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) ResetSubnetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetSubnetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) ResetTcpEstablishedIdleTimeoutSec() {
	_jsii_.InvokeVoid(
		g,
		"resetTcpEstablishedIdleTimeoutSec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) ResetTcpTimeWaitTimeoutSec() {
	_jsii_.InvokeVoid(
		g,
		"resetTcpTimeWaitTimeoutSec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) ResetTcpTransitoryIdleTimeoutSec() {
	_jsii_.InvokeVoid(
		g,
		"resetTcpTransitoryIdleTimeoutSec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) ResetType() {
	_jsii_.InvokeVoid(
		g,
		"resetType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) ResetUdpIdleTimeoutSec() {
	_jsii_.InvokeVoid(
		g,
		"resetUdpIdleTimeoutSec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRouterNat) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRouterNat) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRouterNat) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRouterNat) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRouterNat) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRouterNat) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

