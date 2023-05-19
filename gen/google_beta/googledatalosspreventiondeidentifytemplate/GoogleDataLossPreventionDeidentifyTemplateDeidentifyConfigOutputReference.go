package googledatalosspreventiondeidentifytemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatalosspreventiondeidentifytemplate/internal"
)

type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference interface {
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
	ImageTransformations() GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsOutputReference
	ImageTransformationsInput() *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformations
	InfoTypeTransformations() GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsOutputReference
	InfoTypeTransformationsInput() *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformations
	InternalValue() *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfig
	SetInternalValue(val *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfig)
	RecordTransformations() GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsOutputReference
	RecordTransformationsInput() *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformations
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
	PutImageTransformations(value *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformations)
	PutInfoTypeTransformations(value *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformations)
	PutRecordTransformations(value *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformations)
	ResetImageTransformations()
	ResetInfoTypeTransformations()
	ResetRecordTransformations()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference
type jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) ImageTransformations() GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsOutputReference {
	var returns GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsOutputReference
	_jsii_.Get(
		j,
		"imageTransformations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) ImageTransformationsInput() *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformations {
	var returns *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformations
	_jsii_.Get(
		j,
		"imageTransformationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) InfoTypeTransformations() GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsOutputReference {
	var returns GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsOutputReference
	_jsii_.Get(
		j,
		"infoTypeTransformations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) InfoTypeTransformationsInput() *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformations {
	var returns *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformations
	_jsii_.Get(
		j,
		"infoTypeTransformationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) InternalValue() *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfig {
	var returns *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) RecordTransformations() GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsOutputReference {
	var returns GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsOutputReference
	_jsii_.Get(
		j,
		"recordTransformations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) RecordTransformationsInput() *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformations {
	var returns *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformations
	_jsii_.Get(
		j,
		"recordTransformationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googleDataLossPreventionDeidentifyTemplate.GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference_Override(g GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataLossPreventionDeidentifyTemplate.GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference)SetInternalValue(val *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) PutImageTransformations(value *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformations) {
	if err := g.validatePutImageTransformationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putImageTransformations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) PutInfoTypeTransformations(value *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformations) {
	if err := g.validatePutInfoTypeTransformationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInfoTypeTransformations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) PutRecordTransformations(value *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformations) {
	if err := g.validatePutRecordTransformationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRecordTransformations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) ResetImageTransformations() {
	_jsii_.InvokeVoid(
		g,
		"resetImageTransformations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) ResetInfoTypeTransformations() {
	_jsii_.InvokeVoid(
		g,
		"resetInfoTypeTransformations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) ResetRecordTransformations() {
	_jsii_.InvokeVoid(
		g,
		"resetRecordTransformations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

