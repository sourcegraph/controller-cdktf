package googlenetworkservicesgrpcroute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetworkservicesgrpcroute/internal"
)

type GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference interface {
	cdktf.ComplexObject
	CaseSensitive() interface{}
	SetCaseSensitive(val interface{})
	CaseSensitiveInput() interface{}
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
	GrpcMethod() *string
	SetGrpcMethod(val *string)
	GrpcMethodInput() *string
	GrpcService() *string
	SetGrpcService(val *string)
	GrpcServiceInput() *string
	InternalValue() *GoogleNetworkServicesGrpcRouteRulesMatchesMethod
	SetInternalValue(val *GoogleNetworkServicesGrpcRouteRulesMatchesMethod)
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
	ResetCaseSensitive()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference
type jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) CaseSensitive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caseSensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) CaseSensitiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caseSensitiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) GrpcMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grpcMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) GrpcMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grpcMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) GrpcService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grpcService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) GrpcServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grpcServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) InternalValue() *GoogleNetworkServicesGrpcRouteRulesMatchesMethod {
	var returns *GoogleNetworkServicesGrpcRouteRulesMatchesMethod
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference{}

	_jsii_.Create(
		"google-beta.googleNetworkServicesGrpcRoute.GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference_Override(g GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleNetworkServicesGrpcRoute.GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference)SetCaseSensitive(val interface{}) {
	if err := j.validateSetCaseSensitiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caseSensitive",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference)SetGrpcMethod(val *string) {
	if err := j.validateSetGrpcMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grpcMethod",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference)SetGrpcService(val *string) {
	if err := j.validateSetGrpcServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grpcService",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference)SetInternalValue(val *GoogleNetworkServicesGrpcRouteRulesMatchesMethod) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) ResetCaseSensitive() {
	_jsii_.InvokeVoid(
		g,
		"resetCaseSensitive",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesMatchesMethodOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

