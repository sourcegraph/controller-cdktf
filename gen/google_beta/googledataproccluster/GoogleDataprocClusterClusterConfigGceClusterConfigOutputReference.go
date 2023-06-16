package googledataproccluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataproccluster/internal"
)

type GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalIpOnly() interface{}
	SetInternalIpOnly(val interface{})
	InternalIpOnlyInput() interface{}
	InternalValue() *GoogleDataprocClusterClusterConfigGceClusterConfig
	SetInternalValue(val *GoogleDataprocClusterClusterConfigGceClusterConfig)
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataInput() *map[string]*string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	NodeGroupAffinity() GoogleDataprocClusterClusterConfigGceClusterConfigNodeGroupAffinityOutputReference
	NodeGroupAffinityInput() *GoogleDataprocClusterClusterConfigGceClusterConfigNodeGroupAffinity
	ReservationAffinity() GoogleDataprocClusterClusterConfigGceClusterConfigReservationAffinityOutputReference
	ReservationAffinityInput() *GoogleDataprocClusterClusterConfigGceClusterConfigReservationAffinity
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	ServiceAccountScopes() *[]*string
	SetServiceAccountScopes(val *[]*string)
	ServiceAccountScopesInput() *[]*string
	ShieldedInstanceConfig() GoogleDataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference
	ShieldedInstanceConfigInput() *GoogleDataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfig
	Subnetwork() *string
	SetSubnetwork(val *string)
	SubnetworkInput() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutNodeGroupAffinity(value *GoogleDataprocClusterClusterConfigGceClusterConfigNodeGroupAffinity)
	PutReservationAffinity(value *GoogleDataprocClusterClusterConfigGceClusterConfigReservationAffinity)
	PutShieldedInstanceConfig(value *GoogleDataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfig)
	ResetInternalIpOnly()
	ResetMetadata()
	ResetNetwork()
	ResetNodeGroupAffinity()
	ResetReservationAffinity()
	ResetServiceAccount()
	ResetServiceAccountScopes()
	ResetShieldedInstanceConfig()
	ResetSubnetwork()
	ResetTags()
	ResetZone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference
type jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) InternalIpOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalIpOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) InternalIpOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalIpOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) InternalValue() *GoogleDataprocClusterClusterConfigGceClusterConfig {
	var returns *GoogleDataprocClusterClusterConfigGceClusterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) NodeGroupAffinity() GoogleDataprocClusterClusterConfigGceClusterConfigNodeGroupAffinityOutputReference {
	var returns GoogleDataprocClusterClusterConfigGceClusterConfigNodeGroupAffinityOutputReference
	_jsii_.Get(
		j,
		"nodeGroupAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) NodeGroupAffinityInput() *GoogleDataprocClusterClusterConfigGceClusterConfigNodeGroupAffinity {
	var returns *GoogleDataprocClusterClusterConfigGceClusterConfigNodeGroupAffinity
	_jsii_.Get(
		j,
		"nodeGroupAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) ReservationAffinity() GoogleDataprocClusterClusterConfigGceClusterConfigReservationAffinityOutputReference {
	var returns GoogleDataprocClusterClusterConfigGceClusterConfigReservationAffinityOutputReference
	_jsii_.Get(
		j,
		"reservationAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) ReservationAffinityInput() *GoogleDataprocClusterClusterConfigGceClusterConfigReservationAffinity {
	var returns *GoogleDataprocClusterClusterConfigGceClusterConfigReservationAffinity
	_jsii_.Get(
		j,
		"reservationAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) ServiceAccountScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceAccountScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) ServiceAccountScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceAccountScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) ShieldedInstanceConfig() GoogleDataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference {
	var returns GoogleDataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) ShieldedInstanceConfigInput() *GoogleDataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfig {
	var returns *GoogleDataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfig
	_jsii_.Get(
		j,
		"shieldedInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) Subnetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) SubnetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


func NewGoogleDataprocClusterClusterConfigGceClusterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocClusterClusterConfigGceClusterConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googleDataprocCluster.GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataprocClusterClusterConfigGceClusterConfigOutputReference_Override(g GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataprocCluster.GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference)SetInternalIpOnly(val interface{}) {
	if err := j.validateSetInternalIpOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalIpOnly",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference)SetInternalValue(val *GoogleDataprocClusterClusterConfigGceClusterConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference)SetServiceAccountScopes(val *[]*string) {
	if err := j.validateSetServiceAccountScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountScopes",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference)SetSubnetwork(val *string) {
	if err := j.validateSetSubnetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetwork",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) PutNodeGroupAffinity(value *GoogleDataprocClusterClusterConfigGceClusterConfigNodeGroupAffinity) {
	if err := g.validatePutNodeGroupAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNodeGroupAffinity",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) PutReservationAffinity(value *GoogleDataprocClusterClusterConfigGceClusterConfigReservationAffinity) {
	if err := g.validatePutReservationAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putReservationAffinity",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) PutShieldedInstanceConfig(value *GoogleDataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfig) {
	if err := g.validatePutShieldedInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putShieldedInstanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) ResetInternalIpOnly() {
	_jsii_.InvokeVoid(
		g,
		"resetInternalIpOnly",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) ResetNodeGroupAffinity() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeGroupAffinity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) ResetReservationAffinity() {
	_jsii_.InvokeVoid(
		g,
		"resetReservationAffinity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) ResetServiceAccountScopes() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccountScopes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) ResetShieldedInstanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetShieldedInstanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) ResetSubnetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetSubnetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) ResetZone() {
	_jsii_.InvokeVoid(
		g,
		"resetZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigGceClusterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

