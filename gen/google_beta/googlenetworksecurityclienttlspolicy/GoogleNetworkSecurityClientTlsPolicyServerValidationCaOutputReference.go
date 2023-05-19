package googlenetworksecurityclienttlspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetworksecurityclienttlspolicy/internal"
)

type GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference interface {
	cdktf.ComplexObject
	CertificateProviderInstance() GoogleNetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstanceOutputReference
	CertificateProviderInstanceInput() *GoogleNetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstance
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
	GrpcEndpoint() GoogleNetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpointOutputReference
	GrpcEndpointInput() *GoogleNetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpoint
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutCertificateProviderInstance(value *GoogleNetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstance)
	PutGrpcEndpoint(value *GoogleNetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpoint)
	ResetCertificateProviderInstance()
	ResetGrpcEndpoint()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference
type jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) CertificateProviderInstance() GoogleNetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstanceOutputReference {
	var returns GoogleNetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstanceOutputReference
	_jsii_.Get(
		j,
		"certificateProviderInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) CertificateProviderInstanceInput() *GoogleNetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstance {
	var returns *GoogleNetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstance
	_jsii_.Get(
		j,
		"certificateProviderInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) GrpcEndpoint() GoogleNetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpointOutputReference {
	var returns GoogleNetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpointOutputReference
	_jsii_.Get(
		j,
		"grpcEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) GrpcEndpointInput() *GoogleNetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpoint {
	var returns *GoogleNetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpoint
	_jsii_.Get(
		j,
		"grpcEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference{}

	_jsii_.Create(
		"google-beta.googleNetworkSecurityClientTlsPolicy.GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference_Override(g GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleNetworkSecurityClientTlsPolicy.GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) PutCertificateProviderInstance(value *GoogleNetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstance) {
	if err := g.validatePutCertificateProviderInstanceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCertificateProviderInstance",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) PutGrpcEndpoint(value *GoogleNetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpoint) {
	if err := g.validatePutGrpcEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGrpcEndpoint",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) ResetCertificateProviderInstance() {
	_jsii_.InvokeVoid(
		g,
		"resetCertificateProviderInstance",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) ResetGrpcEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetGrpcEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyServerValidationCaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

