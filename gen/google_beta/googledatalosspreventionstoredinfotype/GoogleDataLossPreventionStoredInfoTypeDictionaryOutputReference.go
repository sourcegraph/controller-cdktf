package googledatalosspreventionstoredinfotype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatalosspreventionstoredinfotype/internal"
)

type GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference interface {
	cdktf.ComplexObject
	CloudStoragePath() GoogleDataLossPreventionStoredInfoTypeDictionaryCloudStoragePathOutputReference
	CloudStoragePathInput() *GoogleDataLossPreventionStoredInfoTypeDictionaryCloudStoragePath
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
	InternalValue() *GoogleDataLossPreventionStoredInfoTypeDictionary
	SetInternalValue(val *GoogleDataLossPreventionStoredInfoTypeDictionary)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WordList() GoogleDataLossPreventionStoredInfoTypeDictionaryWordListStructOutputReference
	WordListInput() *GoogleDataLossPreventionStoredInfoTypeDictionaryWordListStruct
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
	PutCloudStoragePath(value *GoogleDataLossPreventionStoredInfoTypeDictionaryCloudStoragePath)
	PutWordList(value *GoogleDataLossPreventionStoredInfoTypeDictionaryWordListStruct)
	ResetCloudStoragePath()
	ResetWordList()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference
type jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) CloudStoragePath() GoogleDataLossPreventionStoredInfoTypeDictionaryCloudStoragePathOutputReference {
	var returns GoogleDataLossPreventionStoredInfoTypeDictionaryCloudStoragePathOutputReference
	_jsii_.Get(
		j,
		"cloudStoragePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) CloudStoragePathInput() *GoogleDataLossPreventionStoredInfoTypeDictionaryCloudStoragePath {
	var returns *GoogleDataLossPreventionStoredInfoTypeDictionaryCloudStoragePath
	_jsii_.Get(
		j,
		"cloudStoragePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) InternalValue() *GoogleDataLossPreventionStoredInfoTypeDictionary {
	var returns *GoogleDataLossPreventionStoredInfoTypeDictionary
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) WordList() GoogleDataLossPreventionStoredInfoTypeDictionaryWordListStructOutputReference {
	var returns GoogleDataLossPreventionStoredInfoTypeDictionaryWordListStructOutputReference
	_jsii_.Get(
		j,
		"wordList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) WordListInput() *GoogleDataLossPreventionStoredInfoTypeDictionaryWordListStruct {
	var returns *GoogleDataLossPreventionStoredInfoTypeDictionaryWordListStruct
	_jsii_.Get(
		j,
		"wordListInput",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionStoredInfoTypeDictionaryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference{}

	_jsii_.Create(
		"google-beta.googleDataLossPreventionStoredInfoType.GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference_Override(g GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataLossPreventionStoredInfoType.GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference)SetInternalValue(val *GoogleDataLossPreventionStoredInfoTypeDictionary) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) PutCloudStoragePath(value *GoogleDataLossPreventionStoredInfoTypeDictionaryCloudStoragePath) {
	if err := g.validatePutCloudStoragePathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCloudStoragePath",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) PutWordList(value *GoogleDataLossPreventionStoredInfoTypeDictionaryWordListStruct) {
	if err := g.validatePutWordListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWordList",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) ResetCloudStoragePath() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudStoragePath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) ResetWordList() {
	_jsii_.InvokeVoid(
		g,
		"resetWordList",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionStoredInfoTypeDictionaryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

