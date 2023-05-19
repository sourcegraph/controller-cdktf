package googlenetworkservicesedgecacheorigin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetworkservicesedgecacheorigin/internal"
)

type GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference interface {
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
	HeaderAction() GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionHeaderActionOutputReference
	HeaderActionInput() *GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionHeaderAction
	InternalValue() *GoogleNetworkServicesEdgeCacheOriginOriginOverrideAction
	SetInternalValue(val *GoogleNetworkServicesEdgeCacheOriginOriginOverrideAction)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UrlRewrite() GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionUrlRewriteOutputReference
	UrlRewriteInput() *GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionUrlRewrite
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
	PutHeaderAction(value *GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionHeaderAction)
	PutUrlRewrite(value *GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionUrlRewrite)
	ResetHeaderAction()
	ResetUrlRewrite()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference
type jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) HeaderAction() GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionHeaderActionOutputReference {
	var returns GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionHeaderActionOutputReference
	_jsii_.Get(
		j,
		"headerAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) HeaderActionInput() *GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionHeaderAction {
	var returns *GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionHeaderAction
	_jsii_.Get(
		j,
		"headerActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) InternalValue() *GoogleNetworkServicesEdgeCacheOriginOriginOverrideAction {
	var returns *GoogleNetworkServicesEdgeCacheOriginOriginOverrideAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) UrlRewrite() GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionUrlRewriteOutputReference {
	var returns GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionUrlRewriteOutputReference
	_jsii_.Get(
		j,
		"urlRewrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) UrlRewriteInput() *GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionUrlRewrite {
	var returns *GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionUrlRewrite
	_jsii_.Get(
		j,
		"urlRewriteInput",
		&returns,
	)
	return returns
}


func NewGoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference{}

	_jsii_.Create(
		"google-beta.googleNetworkServicesEdgeCacheOrigin.GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference_Override(g GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleNetworkServicesEdgeCacheOrigin.GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference)SetInternalValue(val *GoogleNetworkServicesEdgeCacheOriginOriginOverrideAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) PutHeaderAction(value *GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionHeaderAction) {
	if err := g.validatePutHeaderActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHeaderAction",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) PutUrlRewrite(value *GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionUrlRewrite) {
	if err := g.validatePutUrlRewriteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUrlRewrite",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) ResetHeaderAction() {
	_jsii_.InvokeVoid(
		g,
		"resetHeaderAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) ResetUrlRewrite() {
	_jsii_.InvokeVoid(
		g,
		"resetUrlRewrite",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

