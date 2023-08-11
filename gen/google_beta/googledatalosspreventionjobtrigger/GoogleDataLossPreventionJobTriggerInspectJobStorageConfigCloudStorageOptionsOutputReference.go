package googledatalosspreventionjobtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatalosspreventionjobtrigger/internal"
)

type GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference interface {
	cdktf.ComplexObject
	BytesLimitPerFile() *float64
	SetBytesLimitPerFile(val *float64)
	BytesLimitPerFileInput() *float64
	BytesLimitPerFilePercent() *float64
	SetBytesLimitPerFilePercent(val *float64)
	BytesLimitPerFilePercentInput() *float64
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
	FileSet() GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetOutputReference
	FileSetInput() *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet
	FilesLimitPercent() *float64
	SetFilesLimitPercent(val *float64)
	FilesLimitPercentInput() *float64
	FileTypes() *[]*string
	SetFileTypes(val *[]*string)
	FileTypesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptions
	SetInternalValue(val *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptions)
	SampleMethod() *string
	SetSampleMethod(val *string)
	SampleMethodInput() *string
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
	PutFileSet(value *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet)
	ResetBytesLimitPerFile()
	ResetBytesLimitPerFilePercent()
	ResetFilesLimitPercent()
	ResetFileTypes()
	ResetSampleMethod()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference
type jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) BytesLimitPerFile() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bytesLimitPerFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) BytesLimitPerFileInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bytesLimitPerFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) BytesLimitPerFilePercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bytesLimitPerFilePercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) BytesLimitPerFilePercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bytesLimitPerFilePercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) FileSet() GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetOutputReference {
	var returns GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetOutputReference
	_jsii_.Get(
		j,
		"fileSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) FileSetInput() *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet
	_jsii_.Get(
		j,
		"fileSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) FilesLimitPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"filesLimitPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) FilesLimitPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"filesLimitPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) FileTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fileTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) FileTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fileTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) InternalValue() *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptions {
	var returns *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) SampleMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sampleMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) SampleMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sampleMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataLossPreventionJobTrigger.GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference_Override(g GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataLossPreventionJobTrigger.GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference)SetBytesLimitPerFile(val *float64) {
	if err := j.validateSetBytesLimitPerFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bytesLimitPerFile",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference)SetBytesLimitPerFilePercent(val *float64) {
	if err := j.validateSetBytesLimitPerFilePercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bytesLimitPerFilePercent",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference)SetFilesLimitPercent(val *float64) {
	if err := j.validateSetFilesLimitPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filesLimitPercent",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference)SetFileTypes(val *[]*string) {
	if err := j.validateSetFileTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileTypes",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference)SetInternalValue(val *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference)SetSampleMethod(val *string) {
	if err := j.validateSetSampleMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleMethod",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) PutFileSet(value *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet) {
	if err := g.validatePutFileSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFileSet",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) ResetBytesLimitPerFile() {
	_jsii_.InvokeVoid(
		g,
		"resetBytesLimitPerFile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) ResetBytesLimitPerFilePercent() {
	_jsii_.InvokeVoid(
		g,
		"resetBytesLimitPerFilePercent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) ResetFilesLimitPercent() {
	_jsii_.InvokeVoid(
		g,
		"resetFilesLimitPercent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) ResetFileTypes() {
	_jsii_.InvokeVoid(
		g,
		"resetFileTypes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) ResetSampleMethod() {
	_jsii_.InvokeVoid(
		g,
		"resetSampleMethod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

