package googlebigqueryconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebigqueryconnection/internal"
)

type GoogleBigqueryConnectionConfigurationOutputReference interface {
	cdktn.ComplexObject
	Asset() GoogleBigqueryConnectionConfigurationAssetOutputReference
	AssetInput() *GoogleBigqueryConnectionConfigurationAsset
	Authentication() GoogleBigqueryConnectionConfigurationAuthenticationOutputReference
	AuthenticationInput() *GoogleBigqueryConnectionConfigurationAuthentication
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
	ConnectorId() *string
	SetConnectorId(val *string)
	ConnectorIdInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Endpoint() GoogleBigqueryConnectionConfigurationEndpointOutputReference
	EndpointInput() *GoogleBigqueryConnectionConfigurationEndpoint
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleBigqueryConnectionConfiguration
	SetInternalValue(val *GoogleBigqueryConnectionConfiguration)
	Network() GoogleBigqueryConnectionConfigurationNetworkOutputReference
	NetworkInput() *GoogleBigqueryConnectionConfigurationNetwork
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
	PutAsset(value *GoogleBigqueryConnectionConfigurationAsset)
	PutAuthentication(value *GoogleBigqueryConnectionConfigurationAuthentication)
	PutEndpoint(value *GoogleBigqueryConnectionConfigurationEndpoint)
	PutNetwork(value *GoogleBigqueryConnectionConfigurationNetwork)
	ResetAuthentication()
	ResetEndpoint()
	ResetNetwork()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBigqueryConnectionConfigurationOutputReference
type jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) Asset() GoogleBigqueryConnectionConfigurationAssetOutputReference {
	var returns GoogleBigqueryConnectionConfigurationAssetOutputReference
	_jsii_.Get(
		j,
		"asset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) AssetInput() *GoogleBigqueryConnectionConfigurationAsset {
	var returns *GoogleBigqueryConnectionConfigurationAsset
	_jsii_.Get(
		j,
		"assetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) Authentication() GoogleBigqueryConnectionConfigurationAuthenticationOutputReference {
	var returns GoogleBigqueryConnectionConfigurationAuthenticationOutputReference
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) AuthenticationInput() *GoogleBigqueryConnectionConfigurationAuthentication {
	var returns *GoogleBigqueryConnectionConfigurationAuthentication
	_jsii_.Get(
		j,
		"authenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) ConnectorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) ConnectorIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) Endpoint() GoogleBigqueryConnectionConfigurationEndpointOutputReference {
	var returns GoogleBigqueryConnectionConfigurationEndpointOutputReference
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) EndpointInput() *GoogleBigqueryConnectionConfigurationEndpoint {
	var returns *GoogleBigqueryConnectionConfigurationEndpoint
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) InternalValue() *GoogleBigqueryConnectionConfiguration {
	var returns *GoogleBigqueryConnectionConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) Network() GoogleBigqueryConnectionConfigurationNetworkOutputReference {
	var returns GoogleBigqueryConnectionConfigurationNetworkOutputReference
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) NetworkInput() *GoogleBigqueryConnectionConfigurationNetwork {
	var returns *GoogleBigqueryConnectionConfigurationNetwork
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleBigqueryConnectionConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleBigqueryConnectionConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBigqueryConnectionConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleBigqueryConnection.GoogleBigqueryConnectionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBigqueryConnectionConfigurationOutputReference_Override(g GoogleBigqueryConnectionConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleBigqueryConnection.GoogleBigqueryConnectionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference)SetConnectorId(val *string) {
	if err := j.validateSetConnectorIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectorId",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference)SetInternalValue(val *GoogleBigqueryConnectionConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) PutAsset(value *GoogleBigqueryConnectionConfigurationAsset) {
	if err := g.validatePutAssetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAsset",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) PutAuthentication(value *GoogleBigqueryConnectionConfigurationAuthentication) {
	if err := g.validatePutAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuthentication",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) PutEndpoint(value *GoogleBigqueryConnectionConfigurationEndpoint) {
	if err := g.validatePutEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEndpoint",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) PutNetwork(value *GoogleBigqueryConnectionConfigurationNetwork) {
	if err := g.validatePutNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetwork",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) ResetAuthentication() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthentication",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) ResetEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBigqueryConnectionConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

