package googledatalosspreventiondeidentifytemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatalosspreventiondeidentifytemplate/internal"
)

type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference interface {
	cdktf.ComplexObject
	AllInfoTypes() GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllInfoTypesOutputReference
	AllInfoTypesInput() *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllInfoTypes
	AllText() GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllTextOutputReference
	AllTextInput() *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllText
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RedactionColor() GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsRedactionColorOutputReference
	RedactionColorInput() *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsRedactionColor
	SelectedInfoTypes() GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsSelectedInfoTypesOutputReference
	SelectedInfoTypesInput() *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsSelectedInfoTypes
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
	PutAllInfoTypes(value *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllInfoTypes)
	PutAllText(value *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllText)
	PutRedactionColor(value *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsRedactionColor)
	PutSelectedInfoTypes(value *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsSelectedInfoTypes)
	ResetAllInfoTypes()
	ResetAllText()
	ResetRedactionColor()
	ResetSelectedInfoTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference
type jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) AllInfoTypes() GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllInfoTypesOutputReference {
	var returns GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllInfoTypesOutputReference
	_jsii_.Get(
		j,
		"allInfoTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) AllInfoTypesInput() *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllInfoTypes {
	var returns *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllInfoTypes
	_jsii_.Get(
		j,
		"allInfoTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) AllText() GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllTextOutputReference {
	var returns GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllTextOutputReference
	_jsii_.Get(
		j,
		"allText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) AllTextInput() *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllText {
	var returns *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllText
	_jsii_.Get(
		j,
		"allTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) RedactionColor() GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsRedactionColorOutputReference {
	var returns GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsRedactionColorOutputReference
	_jsii_.Get(
		j,
		"redactionColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) RedactionColorInput() *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsRedactionColor {
	var returns *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsRedactionColor
	_jsii_.Get(
		j,
		"redactionColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) SelectedInfoTypes() GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsSelectedInfoTypesOutputReference {
	var returns GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsSelectedInfoTypesOutputReference
	_jsii_.Get(
		j,
		"selectedInfoTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) SelectedInfoTypesInput() *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsSelectedInfoTypes {
	var returns *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsSelectedInfoTypes
	_jsii_.Get(
		j,
		"selectedInfoTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference{}

	_jsii_.Create(
		"google-beta.googleDataLossPreventionDeidentifyTemplate.GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference_Override(g GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataLossPreventionDeidentifyTemplate.GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) PutAllInfoTypes(value *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllInfoTypes) {
	if err := g.validatePutAllInfoTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAllInfoTypes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) PutAllText(value *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllText) {
	if err := g.validatePutAllTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAllText",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) PutRedactionColor(value *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsRedactionColor) {
	if err := g.validatePutRedactionColorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRedactionColor",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) PutSelectedInfoTypes(value *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsSelectedInfoTypes) {
	if err := g.validatePutSelectedInfoTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSelectedInfoTypes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) ResetAllInfoTypes() {
	_jsii_.InvokeVoid(
		g,
		"resetAllInfoTypes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) ResetAllText() {
	_jsii_.InvokeVoid(
		g,
		"resetAllText",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) ResetRedactionColor() {
	_jsii_.InvokeVoid(
		g,
		"resetRedactionColor",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) ResetSelectedInfoTypes() {
	_jsii_.InvokeVoid(
		g,
		"resetSelectedInfoTypes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

