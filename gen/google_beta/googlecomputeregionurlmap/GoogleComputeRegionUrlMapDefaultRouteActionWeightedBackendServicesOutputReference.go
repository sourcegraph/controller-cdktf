package googlecomputeregionurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeregionurlmap/internal"
)

type GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference interface {
	cdktf.ComplexObject
	BackendService() *string
	SetBackendService(val *string)
	BackendServiceInput() *string
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
	HeaderAction() GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference
	HeaderActionInput() *GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderAction
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
	Weight() *float64
	SetWeight(val *float64)
	WeightInput() *float64
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
	PutHeaderAction(value *GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderAction)
	ResetBackendService()
	ResetHeaderAction()
	ResetWeight()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference
type jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) BackendService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) BackendServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) HeaderAction() GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference {
	var returns GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference
	_jsii_.Get(
		j,
		"headerAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) HeaderActionInput() *GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderAction {
	var returns *GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderAction
	_jsii_.Get(
		j,
		"headerActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) WeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference{}

	_jsii_.Create(
		"google-beta.googleComputeRegionUrlMap.GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference_Override(g GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeRegionUrlMap.GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference)SetBackendService(val *string) {
	if err := j.validateSetBackendServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backendService",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference)SetWeight(val *float64) {
	if err := j.validateSetWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weight",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) PutHeaderAction(value *GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesHeaderAction) {
	if err := g.validatePutHeaderActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHeaderAction",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) ResetBackendService() {
	_jsii_.InvokeVoid(
		g,
		"resetBackendService",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) ResetHeaderAction() {
	_jsii_.InvokeVoid(
		g,
		"resetHeaderAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) ResetWeight() {
	_jsii_.InvokeVoid(
		g,
		"resetWeight",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

