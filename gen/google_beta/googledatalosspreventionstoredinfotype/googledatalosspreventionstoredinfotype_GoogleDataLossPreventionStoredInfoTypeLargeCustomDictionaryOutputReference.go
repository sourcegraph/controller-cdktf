package googledatalosspreventionstoredinfotype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatalosspreventionstoredinfotype/internal"
)

type GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference interface {
	cdktf.ComplexObject
	BigQueryField() GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference
	BigQueryFieldInput() *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField
	CloudStorageFileSet() GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetOutputReference
	CloudStorageFileSetInput() *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet
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
	InternalValue() *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionary
	SetInternalValue(val *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionary)
	OutputPath() GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPathOutputReference
	OutputPathInput() *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPath
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
	PutBigQueryField(value *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField)
	PutCloudStorageFileSet(value *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet)
	PutOutputPath(value *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPath)
	ResetBigQueryField()
	ResetCloudStorageFileSet()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference
type jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) BigQueryField() GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference {
	var returns GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference
	_jsii_.Get(
		j,
		"bigQueryField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) BigQueryFieldInput() *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField {
	var returns *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField
	_jsii_.Get(
		j,
		"bigQueryFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) CloudStorageFileSet() GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetOutputReference {
	var returns GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetOutputReference
	_jsii_.Get(
		j,
		"cloudStorageFileSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) CloudStorageFileSetInput() *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet {
	var returns *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet
	_jsii_.Get(
		j,
		"cloudStorageFileSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) InternalValue() *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionary {
	var returns *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionary
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) OutputPath() GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPathOutputReference {
	var returns GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPathOutputReference
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) OutputPathInput() *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPath {
	var returns *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPath
	_jsii_.Get(
		j,
		"outputPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference{}

	_jsii_.Create(
		"google-beta.googleDataLossPreventionStoredInfoType.GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference_Override(g GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataLossPreventionStoredInfoType.GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference)SetInternalValue(val *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionary) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) PutBigQueryField(value *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField) {
	if err := g.validatePutBigQueryFieldParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBigQueryField",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) PutCloudStorageFileSet(value *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet) {
	if err := g.validatePutCloudStorageFileSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCloudStorageFileSet",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) PutOutputPath(value *GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPath) {
	if err := g.validatePutOutputPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOutputPath",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) ResetBigQueryField() {
	_jsii_.InvokeVoid(
		g,
		"resetBigQueryField",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) ResetCloudStorageFileSet() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudStorageFileSet",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

