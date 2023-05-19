package spectrumapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/spectrumapplication/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/spectrum_application cloudflare_spectrum_application}.
type SpectrumApplication interface {
	cdktf.TerraformResource
	ArgoSmartRouting() interface{}
	SetArgoSmartRouting(val interface{})
	ArgoSmartRoutingInput() interface{}
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
	Dns() SpectrumApplicationDnsOutputReference
	DnsInput() *SpectrumApplicationDns
	EdgeIps() SpectrumApplicationEdgeIpsOutputReference
	EdgeIpsInput() *SpectrumApplicationEdgeIps
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
	IpFirewall() interface{}
	SetIpFirewall(val interface{})
	IpFirewallInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OriginDirect() *[]*string
	SetOriginDirect(val *[]*string)
	OriginDirectInput() *[]*string
	OriginDns() SpectrumApplicationOriginDnsOutputReference
	OriginDnsInput() *SpectrumApplicationOriginDns
	OriginPort() *float64
	SetOriginPort(val *float64)
	OriginPortInput() *float64
	OriginPortRange() SpectrumApplicationOriginPortRangeOutputReference
	OriginPortRangeInput() *SpectrumApplicationOriginPortRange
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	ProxyProtocol() *string
	SetProxyProtocol(val *string)
	ProxyProtocolInput() *string
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Tls() *string
	SetTls(val *string)
	TlsInput() *string
	TrafficType() *string
	SetTrafficType(val *string)
	TrafficTypeInput() *string
	ZoneId() *string
	SetZoneId(val *string)
	ZoneIdInput() *string
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
	PutDns(value *SpectrumApplicationDns)
	PutEdgeIps(value *SpectrumApplicationEdgeIps)
	PutOriginDns(value *SpectrumApplicationOriginDns)
	PutOriginPortRange(value *SpectrumApplicationOriginPortRange)
	ResetArgoSmartRouting()
	ResetEdgeIps()
	ResetId()
	ResetIpFirewall()
	ResetOriginDirect()
	ResetOriginDns()
	ResetOriginPort()
	ResetOriginPortRange()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProxyProtocol()
	ResetTls()
	ResetTrafficType()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SpectrumApplication
type jsiiProxy_SpectrumApplication struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SpectrumApplication) ArgoSmartRouting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"argoSmartRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) ArgoSmartRoutingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"argoSmartRoutingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) Dns() SpectrumApplicationDnsOutputReference {
	var returns SpectrumApplicationDnsOutputReference
	_jsii_.Get(
		j,
		"dns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) DnsInput() *SpectrumApplicationDns {
	var returns *SpectrumApplicationDns
	_jsii_.Get(
		j,
		"dnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) EdgeIps() SpectrumApplicationEdgeIpsOutputReference {
	var returns SpectrumApplicationEdgeIpsOutputReference
	_jsii_.Get(
		j,
		"edgeIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) EdgeIpsInput() *SpectrumApplicationEdgeIps {
	var returns *SpectrumApplicationEdgeIps
	_jsii_.Get(
		j,
		"edgeIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) IpFirewall() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipFirewall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) IpFirewallInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipFirewallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) OriginDirect() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"originDirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) OriginDirectInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"originDirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) OriginDns() SpectrumApplicationOriginDnsOutputReference {
	var returns SpectrumApplicationOriginDnsOutputReference
	_jsii_.Get(
		j,
		"originDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) OriginDnsInput() *SpectrumApplicationOriginDns {
	var returns *SpectrumApplicationOriginDns
	_jsii_.Get(
		j,
		"originDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) OriginPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"originPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) OriginPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"originPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) OriginPortRange() SpectrumApplicationOriginPortRangeOutputReference {
	var returns SpectrumApplicationOriginPortRangeOutputReference
	_jsii_.Get(
		j,
		"originPortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) OriginPortRangeInput() *SpectrumApplicationOriginPortRange {
	var returns *SpectrumApplicationOriginPortRange
	_jsii_.Get(
		j,
		"originPortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) ProxyProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) ProxyProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) Tls() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) TlsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) TrafficType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) TrafficTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpectrumApplication) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/spectrum_application cloudflare_spectrum_application} Resource.
func NewSpectrumApplication(scope constructs.Construct, id *string, config *SpectrumApplicationConfig) SpectrumApplication {
	_init_.Initialize()

	if err := validateNewSpectrumApplicationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpectrumApplication{}

	_jsii_.Create(
		"cloudflare.spectrumApplication.SpectrumApplication",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/spectrum_application cloudflare_spectrum_application} Resource.
func NewSpectrumApplication_Override(s SpectrumApplication, scope constructs.Construct, id *string, config *SpectrumApplicationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudflare.spectrumApplication.SpectrumApplication",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SpectrumApplication)SetArgoSmartRouting(val interface{}) {
	if err := j.validateSetArgoSmartRoutingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"argoSmartRouting",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication)SetIpFirewall(val interface{}) {
	if err := j.validateSetIpFirewallParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipFirewall",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication)SetOriginDirect(val *[]*string) {
	if err := j.validateSetOriginDirectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originDirect",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication)SetOriginPort(val *float64) {
	if err := j.validateSetOriginPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originPort",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication)SetProxyProtocol(val *string) {
	if err := j.validateSetProxyProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyProtocol",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication)SetTls(val *string) {
	if err := j.validateSetTlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tls",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication)SetTrafficType(val *string) {
	if err := j.validateSetTrafficTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trafficType",
		val,
	)
}

func (j *jsiiProxy_SpectrumApplication)SetZoneId(val *string) {
	if err := j.validateSetZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneId",
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
func SpectrumApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpectrumApplication_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cloudflare.spectrumApplication.SpectrumApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SpectrumApplication_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpectrumApplication_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cloudflare.spectrumApplication.SpectrumApplication",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SpectrumApplication_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpectrumApplication_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cloudflare.spectrumApplication.SpectrumApplication",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SpectrumApplication_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cloudflare.spectrumApplication.SpectrumApplication",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SpectrumApplication) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SpectrumApplication) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SpectrumApplication) PutDns(value *SpectrumApplicationDns) {
	if err := s.validatePutDnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDns",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpectrumApplication) PutEdgeIps(value *SpectrumApplicationEdgeIps) {
	if err := s.validatePutEdgeIpsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEdgeIps",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpectrumApplication) PutOriginDns(value *SpectrumApplicationOriginDns) {
	if err := s.validatePutOriginDnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOriginDns",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpectrumApplication) PutOriginPortRange(value *SpectrumApplicationOriginPortRange) {
	if err := s.validatePutOriginPortRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOriginPortRange",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpectrumApplication) ResetArgoSmartRouting() {
	_jsii_.InvokeVoid(
		s,
		"resetArgoSmartRouting",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpectrumApplication) ResetEdgeIps() {
	_jsii_.InvokeVoid(
		s,
		"resetEdgeIps",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpectrumApplication) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpectrumApplication) ResetIpFirewall() {
	_jsii_.InvokeVoid(
		s,
		"resetIpFirewall",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpectrumApplication) ResetOriginDirect() {
	_jsii_.InvokeVoid(
		s,
		"resetOriginDirect",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpectrumApplication) ResetOriginDns() {
	_jsii_.InvokeVoid(
		s,
		"resetOriginDns",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpectrumApplication) ResetOriginPort() {
	_jsii_.InvokeVoid(
		s,
		"resetOriginPort",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpectrumApplication) ResetOriginPortRange() {
	_jsii_.InvokeVoid(
		s,
		"resetOriginPortRange",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpectrumApplication) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpectrumApplication) ResetProxyProtocol() {
	_jsii_.InvokeVoid(
		s,
		"resetProxyProtocol",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpectrumApplication) ResetTls() {
	_jsii_.InvokeVoid(
		s,
		"resetTls",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpectrumApplication) ResetTrafficType() {
	_jsii_.InvokeVoid(
		s,
		"resetTrafficType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpectrumApplication) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpectrumApplication) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

