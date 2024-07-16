package googlepubsubsubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlepubsubsubscription/internal"
)

type GooglePubsubSubscriptionCloudStorageConfigOutputReference interface {
	cdktf.ComplexObject
	AvroConfig() GooglePubsubSubscriptionCloudStorageConfigAvroConfigOutputReference
	AvroConfigInput() *GooglePubsubSubscriptionCloudStorageConfigAvroConfig
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
	FilenameDatetimeFormat() *string
	SetFilenameDatetimeFormat(val *string)
	FilenameDatetimeFormatInput() *string
	FilenamePrefix() *string
	SetFilenamePrefix(val *string)
	FilenamePrefixInput() *string
	FilenameSuffix() *string
	SetFilenameSuffix(val *string)
	FilenameSuffixInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GooglePubsubSubscriptionCloudStorageConfig
	SetInternalValue(val *GooglePubsubSubscriptionCloudStorageConfig)
	MaxBytes() *float64
	SetMaxBytes(val *float64)
	MaxBytesInput() *float64
	MaxDuration() *string
	SetMaxDuration(val *string)
	MaxDurationInput() *string
	ServiceAccountEmail() *string
	SetServiceAccountEmail(val *string)
	ServiceAccountEmailInput() *string
	State() *string
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
	PutAvroConfig(value *GooglePubsubSubscriptionCloudStorageConfigAvroConfig)
	ResetAvroConfig()
	ResetFilenameDatetimeFormat()
	ResetFilenamePrefix()
	ResetFilenameSuffix()
	ResetMaxBytes()
	ResetMaxDuration()
	ResetServiceAccountEmail()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GooglePubsubSubscriptionCloudStorageConfigOutputReference
type jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) AvroConfig() GooglePubsubSubscriptionCloudStorageConfigAvroConfigOutputReference {
	var returns GooglePubsubSubscriptionCloudStorageConfigAvroConfigOutputReference
	_jsii_.Get(
		j,
		"avroConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) AvroConfigInput() *GooglePubsubSubscriptionCloudStorageConfigAvroConfig {
	var returns *GooglePubsubSubscriptionCloudStorageConfigAvroConfig
	_jsii_.Get(
		j,
		"avroConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) FilenameDatetimeFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filenameDatetimeFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) FilenameDatetimeFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filenameDatetimeFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) FilenamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filenamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) FilenamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filenamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) FilenameSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filenameSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) FilenameSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filenameSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) InternalValue() *GooglePubsubSubscriptionCloudStorageConfig {
	var returns *GooglePubsubSubscriptionCloudStorageConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) MaxBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) MaxBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) MaxDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) MaxDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) ServiceAccountEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) ServiceAccountEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGooglePubsubSubscriptionCloudStorageConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GooglePubsubSubscriptionCloudStorageConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGooglePubsubSubscriptionCloudStorageConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googlePubsubSubscription.GooglePubsubSubscriptionCloudStorageConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGooglePubsubSubscriptionCloudStorageConfigOutputReference_Override(g GooglePubsubSubscriptionCloudStorageConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googlePubsubSubscription.GooglePubsubSubscriptionCloudStorageConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference)SetFilenameDatetimeFormat(val *string) {
	if err := j.validateSetFilenameDatetimeFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filenameDatetimeFormat",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference)SetFilenamePrefix(val *string) {
	if err := j.validateSetFilenamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filenamePrefix",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference)SetFilenameSuffix(val *string) {
	if err := j.validateSetFilenameSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filenameSuffix",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference)SetInternalValue(val *GooglePubsubSubscriptionCloudStorageConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference)SetMaxBytes(val *float64) {
	if err := j.validateSetMaxBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBytes",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference)SetMaxDuration(val *string) {
	if err := j.validateSetMaxDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDuration",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference)SetServiceAccountEmail(val *string) {
	if err := j.validateSetServiceAccountEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountEmail",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) PutAvroConfig(value *GooglePubsubSubscriptionCloudStorageConfigAvroConfig) {
	if err := g.validatePutAvroConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAvroConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) ResetAvroConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAvroConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) ResetFilenameDatetimeFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetFilenameDatetimeFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) ResetFilenamePrefix() {
	_jsii_.InvokeVoid(
		g,
		"resetFilenamePrefix",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) ResetFilenameSuffix() {
	_jsii_.InvokeVoid(
		g,
		"resetFilenameSuffix",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) ResetMaxBytes() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxBytes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) ResetMaxDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) ResetServiceAccountEmail() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccountEmail",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GooglePubsubSubscriptionCloudStorageConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

