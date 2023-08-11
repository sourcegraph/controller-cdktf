package googlenetworkservicesedgecacheorigin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetworkservicesedgecacheorigin/internal"
)

type GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference interface {
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
	ConnectTimeout() *string
	SetConnectTimeout(val *string)
	ConnectTimeoutInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleNetworkServicesEdgeCacheOriginTimeout
	SetInternalValue(val *GoogleNetworkServicesEdgeCacheOriginTimeout)
	MaxAttemptsTimeout() *string
	SetMaxAttemptsTimeout(val *string)
	MaxAttemptsTimeoutInput() *string
	ReadTimeout() *string
	SetReadTimeout(val *string)
	ReadTimeoutInput() *string
	ResponseTimeout() *string
	SetResponseTimeout(val *string)
	ResponseTimeoutInput() *string
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
	ResetConnectTimeout()
	ResetMaxAttemptsTimeout()
	ResetReadTimeout()
	ResetResponseTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference
type jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) ConnectTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) ConnectTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) InternalValue() *GoogleNetworkServicesEdgeCacheOriginTimeout {
	var returns *GoogleNetworkServicesEdgeCacheOriginTimeout
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) MaxAttemptsTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxAttemptsTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) MaxAttemptsTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxAttemptsTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) ReadTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) ReadTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) ResponseTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) ResponseTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkServicesEdgeCacheOriginTimeoutOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetworkServicesEdgeCacheOrigin.GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference_Override(g GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetworkServicesEdgeCacheOrigin.GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference)SetConnectTimeout(val *string) {
	if err := j.validateSetConnectTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectTimeout",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference)SetInternalValue(val *GoogleNetworkServicesEdgeCacheOriginTimeout) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference)SetMaxAttemptsTimeout(val *string) {
	if err := j.validateSetMaxAttemptsTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAttemptsTimeout",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference)SetReadTimeout(val *string) {
	if err := j.validateSetReadTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readTimeout",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference)SetResponseTimeout(val *string) {
	if err := j.validateSetResponseTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseTimeout",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) ResetConnectTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetConnectTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) ResetMaxAttemptsTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxAttemptsTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) ResetReadTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetReadTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) ResetResponseTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetResponseTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginTimeoutOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

