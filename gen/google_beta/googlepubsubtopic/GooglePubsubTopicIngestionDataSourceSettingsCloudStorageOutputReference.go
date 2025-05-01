package googlepubsubtopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlepubsubtopic/internal"
)

type GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference interface {
	cdktf.ComplexObject
	AvroFormat() GooglePubsubTopicIngestionDataSourceSettingsCloudStorageAvroFormatOutputReference
	AvroFormatInput() *GooglePubsubTopicIngestionDataSourceSettingsCloudStorageAvroFormat
	Bucket() *string
	SetBucket(val *string)
	BucketInput() *string
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
	InternalValue() *GooglePubsubTopicIngestionDataSourceSettingsCloudStorage
	SetInternalValue(val *GooglePubsubTopicIngestionDataSourceSettingsCloudStorage)
	MatchGlob() *string
	SetMatchGlob(val *string)
	MatchGlobInput() *string
	MinimumObjectCreateTime() *string
	SetMinimumObjectCreateTime(val *string)
	MinimumObjectCreateTimeInput() *string
	PubsubAvroFormat() GooglePubsubTopicIngestionDataSourceSettingsCloudStoragePubsubAvroFormatOutputReference
	PubsubAvroFormatInput() *GooglePubsubTopicIngestionDataSourceSettingsCloudStoragePubsubAvroFormat
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TextFormat() GooglePubsubTopicIngestionDataSourceSettingsCloudStorageTextFormatOutputReference
	TextFormatInput() *GooglePubsubTopicIngestionDataSourceSettingsCloudStorageTextFormat
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
	PutAvroFormat(value *GooglePubsubTopicIngestionDataSourceSettingsCloudStorageAvroFormat)
	PutPubsubAvroFormat(value *GooglePubsubTopicIngestionDataSourceSettingsCloudStoragePubsubAvroFormat)
	PutTextFormat(value *GooglePubsubTopicIngestionDataSourceSettingsCloudStorageTextFormat)
	ResetAvroFormat()
	ResetMatchGlob()
	ResetMinimumObjectCreateTime()
	ResetPubsubAvroFormat()
	ResetTextFormat()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference
type jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) AvroFormat() GooglePubsubTopicIngestionDataSourceSettingsCloudStorageAvroFormatOutputReference {
	var returns GooglePubsubTopicIngestionDataSourceSettingsCloudStorageAvroFormatOutputReference
	_jsii_.Get(
		j,
		"avroFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) AvroFormatInput() *GooglePubsubTopicIngestionDataSourceSettingsCloudStorageAvroFormat {
	var returns *GooglePubsubTopicIngestionDataSourceSettingsCloudStorageAvroFormat
	_jsii_.Get(
		j,
		"avroFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) InternalValue() *GooglePubsubTopicIngestionDataSourceSettingsCloudStorage {
	var returns *GooglePubsubTopicIngestionDataSourceSettingsCloudStorage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) MatchGlob() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchGlob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) MatchGlobInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchGlobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) MinimumObjectCreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumObjectCreateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) MinimumObjectCreateTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumObjectCreateTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) PubsubAvroFormat() GooglePubsubTopicIngestionDataSourceSettingsCloudStoragePubsubAvroFormatOutputReference {
	var returns GooglePubsubTopicIngestionDataSourceSettingsCloudStoragePubsubAvroFormatOutputReference
	_jsii_.Get(
		j,
		"pubsubAvroFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) PubsubAvroFormatInput() *GooglePubsubTopicIngestionDataSourceSettingsCloudStoragePubsubAvroFormat {
	var returns *GooglePubsubTopicIngestionDataSourceSettingsCloudStoragePubsubAvroFormat
	_jsii_.Get(
		j,
		"pubsubAvroFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) TextFormat() GooglePubsubTopicIngestionDataSourceSettingsCloudStorageTextFormatOutputReference {
	var returns GooglePubsubTopicIngestionDataSourceSettingsCloudStorageTextFormatOutputReference
	_jsii_.Get(
		j,
		"textFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) TextFormatInput() *GooglePubsubTopicIngestionDataSourceSettingsCloudStorageTextFormat {
	var returns *GooglePubsubTopicIngestionDataSourceSettingsCloudStorageTextFormat
	_jsii_.Get(
		j,
		"textFormatInput",
		&returns,
	)
	return returns
}


func NewGooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference {
	_init_.Initialize()

	if err := validateNewGooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googlePubsubTopic.GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference_Override(g GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googlePubsubTopic.GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference)SetInternalValue(val *GooglePubsubTopicIngestionDataSourceSettingsCloudStorage) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference)SetMatchGlob(val *string) {
	if err := j.validateSetMatchGlobParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchGlob",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference)SetMinimumObjectCreateTime(val *string) {
	if err := j.validateSetMinimumObjectCreateTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumObjectCreateTime",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) PutAvroFormat(value *GooglePubsubTopicIngestionDataSourceSettingsCloudStorageAvroFormat) {
	if err := g.validatePutAvroFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAvroFormat",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) PutPubsubAvroFormat(value *GooglePubsubTopicIngestionDataSourceSettingsCloudStoragePubsubAvroFormat) {
	if err := g.validatePutPubsubAvroFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPubsubAvroFormat",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) PutTextFormat(value *GooglePubsubTopicIngestionDataSourceSettingsCloudStorageTextFormat) {
	if err := g.validatePutTextFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTextFormat",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) ResetAvroFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetAvroFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) ResetMatchGlob() {
	_jsii_.InvokeVoid(
		g,
		"resetMatchGlob",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) ResetMinimumObjectCreateTime() {
	_jsii_.InvokeVoid(
		g,
		"resetMinimumObjectCreateTime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) ResetPubsubAvroFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetPubsubAvroFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) ResetTextFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetTextFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GooglePubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

