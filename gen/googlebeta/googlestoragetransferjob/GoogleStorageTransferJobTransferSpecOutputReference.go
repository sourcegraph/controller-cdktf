package googlestoragetransferjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlestoragetransferjob/internal"
)

type GoogleStorageTransferJobTransferSpecOutputReference interface {
	cdktf.ComplexObject
	AwsS3DataSource() GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference
	AwsS3DataSourceInput() *GoogleStorageTransferJobTransferSpecAwsS3DataSource
	AzureBlobStorageDataSource() GoogleStorageTransferJobTransferSpecAzureBlobStorageDataSourceOutputReference
	AzureBlobStorageDataSourceInput() *GoogleStorageTransferJobTransferSpecAzureBlobStorageDataSource
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
	GcsDataSink() GoogleStorageTransferJobTransferSpecGcsDataSinkOutputReference
	GcsDataSinkInput() *GoogleStorageTransferJobTransferSpecGcsDataSink
	GcsDataSource() GoogleStorageTransferJobTransferSpecGcsDataSourceOutputReference
	GcsDataSourceInput() *GoogleStorageTransferJobTransferSpecGcsDataSource
	HttpDataSource() GoogleStorageTransferJobTransferSpecHttpDataSourceOutputReference
	HttpDataSourceInput() *GoogleStorageTransferJobTransferSpecHttpDataSource
	InternalValue() *GoogleStorageTransferJobTransferSpec
	SetInternalValue(val *GoogleStorageTransferJobTransferSpec)
	ObjectConditions() GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference
	ObjectConditionsInput() *GoogleStorageTransferJobTransferSpecObjectConditions
	PosixDataSink() GoogleStorageTransferJobTransferSpecPosixDataSinkOutputReference
	PosixDataSinkInput() *GoogleStorageTransferJobTransferSpecPosixDataSink
	PosixDataSource() GoogleStorageTransferJobTransferSpecPosixDataSourceOutputReference
	PosixDataSourceInput() *GoogleStorageTransferJobTransferSpecPosixDataSource
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransferOptions() GoogleStorageTransferJobTransferSpecTransferOptionsOutputReference
	TransferOptionsInput() *GoogleStorageTransferJobTransferSpecTransferOptions
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
	PutAwsS3DataSource(value *GoogleStorageTransferJobTransferSpecAwsS3DataSource)
	PutAzureBlobStorageDataSource(value *GoogleStorageTransferJobTransferSpecAzureBlobStorageDataSource)
	PutGcsDataSink(value *GoogleStorageTransferJobTransferSpecGcsDataSink)
	PutGcsDataSource(value *GoogleStorageTransferJobTransferSpecGcsDataSource)
	PutHttpDataSource(value *GoogleStorageTransferJobTransferSpecHttpDataSource)
	PutObjectConditions(value *GoogleStorageTransferJobTransferSpecObjectConditions)
	PutPosixDataSink(value *GoogleStorageTransferJobTransferSpecPosixDataSink)
	PutPosixDataSource(value *GoogleStorageTransferJobTransferSpecPosixDataSource)
	PutTransferOptions(value *GoogleStorageTransferJobTransferSpecTransferOptions)
	ResetAwsS3DataSource()
	ResetAzureBlobStorageDataSource()
	ResetGcsDataSink()
	ResetGcsDataSource()
	ResetHttpDataSource()
	ResetObjectConditions()
	ResetPosixDataSink()
	ResetPosixDataSource()
	ResetTransferOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleStorageTransferJobTransferSpecOutputReference
type jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) AwsS3DataSource() GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference {
	var returns GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference
	_jsii_.Get(
		j,
		"awsS3DataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) AwsS3DataSourceInput() *GoogleStorageTransferJobTransferSpecAwsS3DataSource {
	var returns *GoogleStorageTransferJobTransferSpecAwsS3DataSource
	_jsii_.Get(
		j,
		"awsS3DataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) AzureBlobStorageDataSource() GoogleStorageTransferJobTransferSpecAzureBlobStorageDataSourceOutputReference {
	var returns GoogleStorageTransferJobTransferSpecAzureBlobStorageDataSourceOutputReference
	_jsii_.Get(
		j,
		"azureBlobStorageDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) AzureBlobStorageDataSourceInput() *GoogleStorageTransferJobTransferSpecAzureBlobStorageDataSource {
	var returns *GoogleStorageTransferJobTransferSpecAzureBlobStorageDataSource
	_jsii_.Get(
		j,
		"azureBlobStorageDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) GcsDataSink() GoogleStorageTransferJobTransferSpecGcsDataSinkOutputReference {
	var returns GoogleStorageTransferJobTransferSpecGcsDataSinkOutputReference
	_jsii_.Get(
		j,
		"gcsDataSink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) GcsDataSinkInput() *GoogleStorageTransferJobTransferSpecGcsDataSink {
	var returns *GoogleStorageTransferJobTransferSpecGcsDataSink
	_jsii_.Get(
		j,
		"gcsDataSinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) GcsDataSource() GoogleStorageTransferJobTransferSpecGcsDataSourceOutputReference {
	var returns GoogleStorageTransferJobTransferSpecGcsDataSourceOutputReference
	_jsii_.Get(
		j,
		"gcsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) GcsDataSourceInput() *GoogleStorageTransferJobTransferSpecGcsDataSource {
	var returns *GoogleStorageTransferJobTransferSpecGcsDataSource
	_jsii_.Get(
		j,
		"gcsDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) HttpDataSource() GoogleStorageTransferJobTransferSpecHttpDataSourceOutputReference {
	var returns GoogleStorageTransferJobTransferSpecHttpDataSourceOutputReference
	_jsii_.Get(
		j,
		"httpDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) HttpDataSourceInput() *GoogleStorageTransferJobTransferSpecHttpDataSource {
	var returns *GoogleStorageTransferJobTransferSpecHttpDataSource
	_jsii_.Get(
		j,
		"httpDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) InternalValue() *GoogleStorageTransferJobTransferSpec {
	var returns *GoogleStorageTransferJobTransferSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) ObjectConditions() GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference {
	var returns GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference
	_jsii_.Get(
		j,
		"objectConditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) ObjectConditionsInput() *GoogleStorageTransferJobTransferSpecObjectConditions {
	var returns *GoogleStorageTransferJobTransferSpecObjectConditions
	_jsii_.Get(
		j,
		"objectConditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) PosixDataSink() GoogleStorageTransferJobTransferSpecPosixDataSinkOutputReference {
	var returns GoogleStorageTransferJobTransferSpecPosixDataSinkOutputReference
	_jsii_.Get(
		j,
		"posixDataSink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) PosixDataSinkInput() *GoogleStorageTransferJobTransferSpecPosixDataSink {
	var returns *GoogleStorageTransferJobTransferSpecPosixDataSink
	_jsii_.Get(
		j,
		"posixDataSinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) PosixDataSource() GoogleStorageTransferJobTransferSpecPosixDataSourceOutputReference {
	var returns GoogleStorageTransferJobTransferSpecPosixDataSourceOutputReference
	_jsii_.Get(
		j,
		"posixDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) PosixDataSourceInput() *GoogleStorageTransferJobTransferSpecPosixDataSource {
	var returns *GoogleStorageTransferJobTransferSpecPosixDataSource
	_jsii_.Get(
		j,
		"posixDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) TransferOptions() GoogleStorageTransferJobTransferSpecTransferOptionsOutputReference {
	var returns GoogleStorageTransferJobTransferSpecTransferOptionsOutputReference
	_jsii_.Get(
		j,
		"transferOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) TransferOptionsInput() *GoogleStorageTransferJobTransferSpecTransferOptions {
	var returns *GoogleStorageTransferJobTransferSpecTransferOptions
	_jsii_.Get(
		j,
		"transferOptionsInput",
		&returns,
	)
	return returns
}


func NewGoogleStorageTransferJobTransferSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleStorageTransferJobTransferSpecOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleStorageTransferJobTransferSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleStorageTransferJob.GoogleStorageTransferJobTransferSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleStorageTransferJobTransferSpecOutputReference_Override(g GoogleStorageTransferJobTransferSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleStorageTransferJob.GoogleStorageTransferJobTransferSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference)SetInternalValue(val *GoogleStorageTransferJobTransferSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) PutAwsS3DataSource(value *GoogleStorageTransferJobTransferSpecAwsS3DataSource) {
	if err := g.validatePutAwsS3DataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAwsS3DataSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) PutAzureBlobStorageDataSource(value *GoogleStorageTransferJobTransferSpecAzureBlobStorageDataSource) {
	if err := g.validatePutAzureBlobStorageDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAzureBlobStorageDataSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) PutGcsDataSink(value *GoogleStorageTransferJobTransferSpecGcsDataSink) {
	if err := g.validatePutGcsDataSinkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGcsDataSink",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) PutGcsDataSource(value *GoogleStorageTransferJobTransferSpecGcsDataSource) {
	if err := g.validatePutGcsDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGcsDataSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) PutHttpDataSource(value *GoogleStorageTransferJobTransferSpecHttpDataSource) {
	if err := g.validatePutHttpDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHttpDataSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) PutObjectConditions(value *GoogleStorageTransferJobTransferSpecObjectConditions) {
	if err := g.validatePutObjectConditionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putObjectConditions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) PutPosixDataSink(value *GoogleStorageTransferJobTransferSpecPosixDataSink) {
	if err := g.validatePutPosixDataSinkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPosixDataSink",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) PutPosixDataSource(value *GoogleStorageTransferJobTransferSpecPosixDataSource) {
	if err := g.validatePutPosixDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPosixDataSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) PutTransferOptions(value *GoogleStorageTransferJobTransferSpecTransferOptions) {
	if err := g.validatePutTransferOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTransferOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) ResetAwsS3DataSource() {
	_jsii_.InvokeVoid(
		g,
		"resetAwsS3DataSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) ResetAzureBlobStorageDataSource() {
	_jsii_.InvokeVoid(
		g,
		"resetAzureBlobStorageDataSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) ResetGcsDataSink() {
	_jsii_.InvokeVoid(
		g,
		"resetGcsDataSink",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) ResetGcsDataSource() {
	_jsii_.InvokeVoid(
		g,
		"resetGcsDataSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) ResetHttpDataSource() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpDataSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) ResetObjectConditions() {
	_jsii_.InvokeVoid(
		g,
		"resetObjectConditions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) ResetPosixDataSink() {
	_jsii_.InvokeVoid(
		g,
		"resetPosixDataSink",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) ResetPosixDataSource() {
	_jsii_.InvokeVoid(
		g,
		"resetPosixDataSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) ResetTransferOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetTransferOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

