package googlecomputefuturereservation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputefuturereservation/internal"
)

type GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference interface {
	cdktn.ComplexObject
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
	InstanceProperties() GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference
	InstancePropertiesInput() *GoogleComputeFutureReservationSpecificSkuPropertiesInstanceProperties
	InternalValue() *GoogleComputeFutureReservationSpecificSkuProperties
	SetInternalValue(val *GoogleComputeFutureReservationSpecificSkuProperties)
	SourceInstanceTemplate() *string
	SetSourceInstanceTemplate(val *string)
	SourceInstanceTemplateInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TotalCount() *string
	SetTotalCount(val *string)
	TotalCountInput() *string
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
	PutInstanceProperties(value *GoogleComputeFutureReservationSpecificSkuPropertiesInstanceProperties)
	ResetInstanceProperties()
	ResetSourceInstanceTemplate()
	ResetTotalCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference
type jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) InstanceProperties() GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference {
	var returns GoogleComputeFutureReservationSpecificSkuPropertiesInstancePropertiesOutputReference
	_jsii_.Get(
		j,
		"instanceProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) InstancePropertiesInput() *GoogleComputeFutureReservationSpecificSkuPropertiesInstanceProperties {
	var returns *GoogleComputeFutureReservationSpecificSkuPropertiesInstanceProperties
	_jsii_.Get(
		j,
		"instancePropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) InternalValue() *GoogleComputeFutureReservationSpecificSkuProperties {
	var returns *GoogleComputeFutureReservationSpecificSkuProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) SourceInstanceTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInstanceTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) SourceInstanceTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInstanceTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) TotalCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) TotalCountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalCountInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeFutureReservationSpecificSkuPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeFutureReservationSpecificSkuPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeFutureReservation.GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeFutureReservationSpecificSkuPropertiesOutputReference_Override(g GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeFutureReservation.GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference)SetInternalValue(val *GoogleComputeFutureReservationSpecificSkuProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference)SetSourceInstanceTemplate(val *string) {
	if err := j.validateSetSourceInstanceTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceInstanceTemplate",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference)SetTotalCount(val *string) {
	if err := j.validateSetTotalCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalCount",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) PutInstanceProperties(value *GoogleComputeFutureReservationSpecificSkuPropertiesInstanceProperties) {
	if err := g.validatePutInstancePropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInstanceProperties",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) ResetInstanceProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetInstanceProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) ResetSourceInstanceTemplate() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceInstanceTemplate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) ResetTotalCount() {
	_jsii_.InvokeVoid(
		g,
		"resetTotalCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeFutureReservationSpecificSkuPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

