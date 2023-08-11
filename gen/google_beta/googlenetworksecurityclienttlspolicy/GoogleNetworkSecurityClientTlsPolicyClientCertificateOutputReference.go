package googlenetworksecurityclienttlspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetworksecurityclienttlspolicy/internal"
)

type GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference interface {
	cdktf.ComplexObject
	CertificateProviderInstance() GoogleNetworkSecurityClientTlsPolicyClientCertificateCertificateProviderInstanceOutputReference
	CertificateProviderInstanceInput() *GoogleNetworkSecurityClientTlsPolicyClientCertificateCertificateProviderInstance
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
	GrpcEndpoint() GoogleNetworkSecurityClientTlsPolicyClientCertificateGrpcEndpointOutputReference
	GrpcEndpointInput() *GoogleNetworkSecurityClientTlsPolicyClientCertificateGrpcEndpoint
	InternalValue() *GoogleNetworkSecurityClientTlsPolicyClientCertificate
	SetInternalValue(val *GoogleNetworkSecurityClientTlsPolicyClientCertificate)
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
	PutCertificateProviderInstance(value *GoogleNetworkSecurityClientTlsPolicyClientCertificateCertificateProviderInstance)
	PutGrpcEndpoint(value *GoogleNetworkSecurityClientTlsPolicyClientCertificateGrpcEndpoint)
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

// The jsii proxy struct for GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference
type jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) CertificateProviderInstance() GoogleNetworkSecurityClientTlsPolicyClientCertificateCertificateProviderInstanceOutputReference {
	var returns GoogleNetworkSecurityClientTlsPolicyClientCertificateCertificateProviderInstanceOutputReference
	_jsii_.Get(
		j,
		"certificateProviderInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) CertificateProviderInstanceInput() *GoogleNetworkSecurityClientTlsPolicyClientCertificateCertificateProviderInstance {
	var returns *GoogleNetworkSecurityClientTlsPolicyClientCertificateCertificateProviderInstance
	_jsii_.Get(
		j,
		"certificateProviderInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) GrpcEndpoint() GoogleNetworkSecurityClientTlsPolicyClientCertificateGrpcEndpointOutputReference {
	var returns GoogleNetworkSecurityClientTlsPolicyClientCertificateGrpcEndpointOutputReference
	_jsii_.Get(
		j,
		"grpcEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) GrpcEndpointInput() *GoogleNetworkSecurityClientTlsPolicyClientCertificateGrpcEndpoint {
	var returns *GoogleNetworkSecurityClientTlsPolicyClientCertificateGrpcEndpoint
	_jsii_.Get(
		j,
		"grpcEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) InternalValue() *GoogleNetworkSecurityClientTlsPolicyClientCertificate {
	var returns *GoogleNetworkSecurityClientTlsPolicyClientCertificate
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetworkSecurityClientTlsPolicy.GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference_Override(g GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetworkSecurityClientTlsPolicy.GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference)SetInternalValue(val *GoogleNetworkSecurityClientTlsPolicyClientCertificate) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) PutCertificateProviderInstance(value *GoogleNetworkSecurityClientTlsPolicyClientCertificateCertificateProviderInstance) {
	if err := g.validatePutCertificateProviderInstanceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCertificateProviderInstance",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) PutGrpcEndpoint(value *GoogleNetworkSecurityClientTlsPolicyClientCertificateGrpcEndpoint) {
	if err := g.validatePutGrpcEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGrpcEndpoint",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) ResetCertificateProviderInstance() {
	_jsii_.InvokeVoid(
		g,
		"resetCertificateProviderInstance",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) ResetGrpcEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetGrpcEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetworkSecurityClientTlsPolicyClientCertificateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

