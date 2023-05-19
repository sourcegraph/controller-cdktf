package googlecomputesubnetwork

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputesubnetwork/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_compute_subnetwork google_compute_subnetwork}.
type GoogleComputeSubnetwork interface {
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
	ExternalIpv6Prefix() *string
	Fingerprint() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GatewayAddress() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IpCidrRange() *string
	SetIpCidrRange(val *string)
	IpCidrRangeInput() *string
	Ipv6AccessType() *string
	SetIpv6AccessType(val *string)
	Ipv6AccessTypeInput() *string
	Ipv6CidrRange() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogConfig() GoogleComputeSubnetworkLogConfigOutputReference
	LogConfigInput() *GoogleComputeSubnetworkLogConfig
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	// The tree node.
	Node() constructs.Node
	PrivateIpGoogleAccess() interface{}
	SetPrivateIpGoogleAccess(val interface{})
	PrivateIpGoogleAccessInput() interface{}
	PrivateIpv6GoogleAccess() *string
	SetPrivateIpv6GoogleAccess(val *string)
	PrivateIpv6GoogleAccessInput() *string
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
	Purpose() *string
	SetPurpose(val *string)
	PurposeInput() *string
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	SecondaryIpRange() GoogleComputeSubnetworkSecondaryIpRangeList
	SecondaryIpRangeInput() interface{}
	SelfLink() *string
	StackType() *string
	SetStackType(val *string)
	StackTypeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleComputeSubnetworkTimeoutsOutputReference
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
	PutLogConfig(value *GoogleComputeSubnetworkLogConfig)
	PutSecondaryIpRange(value interface{})
	PutTimeouts(value *GoogleComputeSubnetworkTimeouts)
	ResetDescription()
	ResetId()
	ResetIpv6AccessType()
	ResetLogConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateIpGoogleAccess()
	ResetPrivateIpv6GoogleAccess()
	ResetProject()
	ResetPurpose()
	ResetRegion()
	ResetRole()
	ResetSecondaryIpRange()
	ResetStackType()
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

// The jsii proxy struct for GoogleComputeSubnetwork
type jsiiProxy_GoogleComputeSubnetwork struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleComputeSubnetwork) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) ExternalIpv6Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIpv6Prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) Fingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) GatewayAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) IpCidrRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipCidrRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) IpCidrRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipCidrRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) Ipv6AccessType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6AccessType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) Ipv6AccessTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6AccessTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) Ipv6CidrRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6CidrRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) LogConfig() GoogleComputeSubnetworkLogConfigOutputReference {
	var returns GoogleComputeSubnetworkLogConfigOutputReference
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) LogConfigInput() *GoogleComputeSubnetworkLogConfig {
	var returns *GoogleComputeSubnetworkLogConfig
	_jsii_.Get(
		j,
		"logConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) PrivateIpGoogleAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateIpGoogleAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) PrivateIpGoogleAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateIpGoogleAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) PrivateIpv6GoogleAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpv6GoogleAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) PrivateIpv6GoogleAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpv6GoogleAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) Purpose() *string {
	var returns *string
	_jsii_.Get(
		j,
		"purpose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) PurposeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"purposeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) SecondaryIpRange() GoogleComputeSubnetworkSecondaryIpRangeList {
	var returns GoogleComputeSubnetworkSecondaryIpRangeList
	_jsii_.Get(
		j,
		"secondaryIpRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) SecondaryIpRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryIpRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) StackType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) StackTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) Timeouts() GoogleComputeSubnetworkTimeoutsOutputReference {
	var returns GoogleComputeSubnetworkTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSubnetwork) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_compute_subnetwork google_compute_subnetwork} Resource.
func NewGoogleComputeSubnetwork(scope constructs.Construct, id *string, config *GoogleComputeSubnetworkConfig) GoogleComputeSubnetwork {
	_init_.Initialize()

	if err := validateNewGoogleComputeSubnetworkParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeSubnetwork{}

	_jsii_.Create(
		"google-beta.googleComputeSubnetwork.GoogleComputeSubnetwork",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_compute_subnetwork google_compute_subnetwork} Resource.
func NewGoogleComputeSubnetwork_Override(g GoogleComputeSubnetwork, scope constructs.Construct, id *string, config *GoogleComputeSubnetworkConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeSubnetwork.GoogleComputeSubnetwork",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeSubnetwork)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSubnetwork)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSubnetwork)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSubnetwork)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSubnetwork)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSubnetwork)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSubnetwork)SetIpCidrRange(val *string) {
	if err := j.validateSetIpCidrRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipCidrRange",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSubnetwork)SetIpv6AccessType(val *string) {
	if err := j.validateSetIpv6AccessTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6AccessType",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSubnetwork)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSubnetwork)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSubnetwork)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSubnetwork)SetPrivateIpGoogleAccess(val interface{}) {
	if err := j.validateSetPrivateIpGoogleAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpGoogleAccess",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSubnetwork)SetPrivateIpv6GoogleAccess(val *string) {
	if err := j.validateSetPrivateIpv6GoogleAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpv6GoogleAccess",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSubnetwork)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSubnetwork)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSubnetwork)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSubnetwork)SetPurpose(val *string) {
	if err := j.validateSetPurposeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"purpose",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSubnetwork)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSubnetwork)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSubnetwork)SetStackType(val *string) {
	if err := j.validateSetStackTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackType",
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
func GoogleComputeSubnetwork_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeSubnetwork_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleComputeSubnetwork.GoogleComputeSubnetwork",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeSubnetwork_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeSubnetwork_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleComputeSubnetwork.GoogleComputeSubnetwork",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeSubnetwork_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeSubnetwork_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleComputeSubnetwork.GoogleComputeSubnetwork",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputeSubnetwork_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"google-beta.googleComputeSubnetwork.GoogleComputeSubnetwork",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeSubnetwork) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeSubnetwork) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeSubnetwork) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeSubnetwork) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeSubnetwork) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeSubnetwork) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeSubnetwork) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeSubnetwork) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeSubnetwork) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeSubnetwork) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeSubnetwork) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeSubnetwork) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeSubnetwork) PutLogConfig(value *GoogleComputeSubnetworkLogConfig) {
	if err := g.validatePutLogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLogConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeSubnetwork) PutSecondaryIpRange(value interface{}) {
	if err := g.validatePutSecondaryIpRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecondaryIpRange",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeSubnetwork) PutTimeouts(value *GoogleComputeSubnetworkTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeSubnetwork) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSubnetwork) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSubnetwork) ResetIpv6AccessType() {
	_jsii_.InvokeVoid(
		g,
		"resetIpv6AccessType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSubnetwork) ResetLogConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLogConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSubnetwork) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSubnetwork) ResetPrivateIpGoogleAccess() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivateIpGoogleAccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSubnetwork) ResetPrivateIpv6GoogleAccess() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivateIpv6GoogleAccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSubnetwork) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSubnetwork) ResetPurpose() {
	_jsii_.InvokeVoid(
		g,
		"resetPurpose",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSubnetwork) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSubnetwork) ResetRole() {
	_jsii_.InvokeVoid(
		g,
		"resetRole",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSubnetwork) ResetSecondaryIpRange() {
	_jsii_.InvokeVoid(
		g,
		"resetSecondaryIpRange",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSubnetwork) ResetStackType() {
	_jsii_.InvokeVoid(
		g,
		"resetStackType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSubnetwork) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSubnetwork) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeSubnetwork) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeSubnetwork) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeSubnetwork) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

