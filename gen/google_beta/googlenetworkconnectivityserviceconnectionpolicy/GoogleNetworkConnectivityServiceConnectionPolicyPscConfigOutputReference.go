package googlenetworkconnectivityserviceconnectionpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetworkconnectivityserviceconnectionpolicy/internal"
)

type GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference interface {
	cdktn.ComplexObject
	AllowedGoogleProducersResourceHierarchyLevel() *[]*string
	SetAllowedGoogleProducersResourceHierarchyLevel(val *[]*string)
	AllowedGoogleProducersResourceHierarchyLevelInput() *[]*string
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
	InternalValue() *GoogleNetworkConnectivityServiceConnectionPolicyPscConfig
	SetInternalValue(val *GoogleNetworkConnectivityServiceConnectionPolicyPscConfig)
	Limit() *string
	SetLimit(val *string)
	LimitInput() *string
	ProducerInstanceLocation() *string
	SetProducerInstanceLocation(val *string)
	ProducerInstanceLocationInput() *string
	Subnetworks() *[]*string
	SetSubnetworks(val *[]*string)
	SubnetworksInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	ResetAllowedGoogleProducersResourceHierarchyLevel()
	ResetLimit()
	ResetProducerInstanceLocation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference
type jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) AllowedGoogleProducersResourceHierarchyLevel() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedGoogleProducersResourceHierarchyLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) AllowedGoogleProducersResourceHierarchyLevelInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedGoogleProducersResourceHierarchyLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) InternalValue() *GoogleNetworkConnectivityServiceConnectionPolicyPscConfig {
	var returns *GoogleNetworkConnectivityServiceConnectionPolicyPscConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) Limit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) LimitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) ProducerInstanceLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"producerInstanceLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) ProducerInstanceLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"producerInstanceLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) Subnetworks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetworks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) SubnetworksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetworksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleNetworkConnectivityServiceConnectionPolicy.GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference_Override(g GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleNetworkConnectivityServiceConnectionPolicy.GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference)SetAllowedGoogleProducersResourceHierarchyLevel(val *[]*string) {
	if err := j.validateSetAllowedGoogleProducersResourceHierarchyLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedGoogleProducersResourceHierarchyLevel",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference)SetInternalValue(val *GoogleNetworkConnectivityServiceConnectionPolicyPscConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference)SetLimit(val *string) {
	if err := j.validateSetLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limit",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference)SetProducerInstanceLocation(val *string) {
	if err := j.validateSetProducerInstanceLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"producerInstanceLocation",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference)SetSubnetworks(val *[]*string) {
	if err := j.validateSetSubnetworksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetworks",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) ResetAllowedGoogleProducersResourceHierarchyLevel() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowedGoogleProducersResourceHierarchyLevel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) ResetLimit() {
	_jsii_.InvokeVoid(
		g,
		"resetLimit",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) ResetProducerInstanceLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetProducerInstanceLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivityServiceConnectionPolicyPscConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

