package googledatalosspreventiondeidentifytemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatalosspreventiondeidentifytemplate/internal"
)

type GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference interface {
	cdktf.ComplexObject
	BooleanValue() interface{}
	SetBooleanValue(val interface{})
	BooleanValueInput() interface{}
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
	DateValue() GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValueOutputReference
	DateValueInput() *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue
	DayOfWeekValue() *string
	SetDayOfWeekValue(val *string)
	DayOfWeekValueInput() *string
	FloatValue() *float64
	SetFloatValue(val *float64)
	FloatValueInput() *float64
	// Experimental.
	Fqn() *string
	IntegerValue() *string
	SetIntegerValue(val *string)
	IntegerValueInput() *string
	InternalValue() *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound
	SetInternalValue(val *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound)
	StringValue() *string
	SetStringValue(val *string)
	StringValueInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimestampValue() *string
	SetTimestampValue(val *string)
	TimestampValueInput() *string
	TimeValue() GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValueOutputReference
	TimeValueInput() *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue
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
	PutDateValue(value *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue)
	PutTimeValue(value *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue)
	ResetBooleanValue()
	ResetDateValue()
	ResetDayOfWeekValue()
	ResetFloatValue()
	ResetIntegerValue()
	ResetStringValue()
	ResetTimestampValue()
	ResetTimeValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference
type jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) BooleanValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"booleanValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) BooleanValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"booleanValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) DateValue() GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValueOutputReference {
	var returns GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValueOutputReference
	_jsii_.Get(
		j,
		"dateValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) DateValueInput() *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue {
	var returns *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue
	_jsii_.Get(
		j,
		"dateValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) DayOfWeekValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayOfWeekValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) DayOfWeekValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayOfWeekValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) FloatValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"floatValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) FloatValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"floatValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) IntegerValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integerValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) IntegerValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integerValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) InternalValue() *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound {
	var returns *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) StringValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) StringValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) TimestampValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) TimestampValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) TimeValue() GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValueOutputReference {
	var returns GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValueOutputReference
	_jsii_.Get(
		j,
		"timeValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) TimeValueInput() *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue {
	var returns *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue
	_jsii_.Get(
		j,
		"timeValueInput",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference{}

	_jsii_.Create(
		"google-beta.googleDataLossPreventionDeidentifyTemplate.GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference_Override(g GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataLossPreventionDeidentifyTemplate.GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference)SetBooleanValue(val interface{}) {
	if err := j.validateSetBooleanValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"booleanValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference)SetDayOfWeekValue(val *string) {
	if err := j.validateSetDayOfWeekValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dayOfWeekValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference)SetFloatValue(val *float64) {
	if err := j.validateSetFloatValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"floatValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference)SetIntegerValue(val *string) {
	if err := j.validateSetIntegerValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integerValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference)SetInternalValue(val *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBound) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference)SetStringValue(val *string) {
	if err := j.validateSetStringValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stringValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference)SetTimestampValue(val *string) {
	if err := j.validateSetTimestampValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampValue",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) PutDateValue(value *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundDateValue) {
	if err := g.validatePutDateValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDateValue",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) PutTimeValue(value *GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundTimeValue) {
	if err := g.validatePutTimeValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeValue",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) ResetBooleanValue() {
	_jsii_.InvokeVoid(
		g,
		"resetBooleanValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) ResetDateValue() {
	_jsii_.InvokeVoid(
		g,
		"resetDateValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) ResetDayOfWeekValue() {
	_jsii_.InvokeVoid(
		g,
		"resetDayOfWeekValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) ResetFloatValue() {
	_jsii_.InvokeVoid(
		g,
		"resetFloatValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) ResetIntegerValue() {
	_jsii_.InvokeVoid(
		g,
		"resetIntegerValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) ResetStringValue() {
	_jsii_.InvokeVoid(
		g,
		"resetStringValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) ResetTimestampValue() {
	_jsii_.InvokeVoid(
		g,
		"resetTimestampValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) ResetTimeValue() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfigUpperBoundOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

