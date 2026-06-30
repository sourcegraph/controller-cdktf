package s3bucketlifecycleconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/s3bucketlifecycleconfiguration/internal"
)

type S3BucketLifecycleConfigurationRuleOutputReference interface {
	cdktf.ComplexObject
	AbortIncompleteMultipartUpload() S3BucketLifecycleConfigurationRuleAbortIncompleteMultipartUploadOutputReference
	AbortIncompleteMultipartUploadInput() *S3BucketLifecycleConfigurationRuleAbortIncompleteMultipartUpload
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
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
	Expiration() S3BucketLifecycleConfigurationRuleExpirationOutputReference
	ExpirationInput() *S3BucketLifecycleConfigurationRuleExpiration
	Filter() S3BucketLifecycleConfigurationRuleFilterOutputReference
	FilterInput() *S3BucketLifecycleConfigurationRuleFilter
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() any
	SetInternalValue(val any)
	NoncurrentVersionExpiration() S3BucketLifecycleConfigurationRuleNoncurrentVersionExpirationOutputReference
	NoncurrentVersionExpirationInput() *S3BucketLifecycleConfigurationRuleNoncurrentVersionExpiration
	NoncurrentVersionTransition() S3BucketLifecycleConfigurationRuleNoncurrentVersionTransitionList
	NoncurrentVersionTransitionInput() any
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Transition() S3BucketLifecycleConfigurationRuleTransitionList
	TransitionInput() any
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	PutAbortIncompleteMultipartUpload(value *S3BucketLifecycleConfigurationRuleAbortIncompleteMultipartUpload)
	PutExpiration(value *S3BucketLifecycleConfigurationRuleExpiration)
	PutFilter(value *S3BucketLifecycleConfigurationRuleFilter)
	PutNoncurrentVersionExpiration(value *S3BucketLifecycleConfigurationRuleNoncurrentVersionExpiration)
	PutNoncurrentVersionTransition(value any)
	PutTransition(value any)
	ResetAbortIncompleteMultipartUpload()
	ResetExpiration()
	ResetFilter()
	ResetNoncurrentVersionExpiration()
	ResetNoncurrentVersionTransition()
	ResetPrefix()
	ResetTransition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3BucketLifecycleConfigurationRuleOutputReference
type jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) AbortIncompleteMultipartUpload() S3BucketLifecycleConfigurationRuleAbortIncompleteMultipartUploadOutputReference {
	var returns S3BucketLifecycleConfigurationRuleAbortIncompleteMultipartUploadOutputReference
	_jsii_.Get(
		j,
		"abortIncompleteMultipartUpload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) AbortIncompleteMultipartUploadInput() *S3BucketLifecycleConfigurationRuleAbortIncompleteMultipartUpload {
	var returns *S3BucketLifecycleConfigurationRuleAbortIncompleteMultipartUpload
	_jsii_.Get(
		j,
		"abortIncompleteMultipartUploadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) Expiration() S3BucketLifecycleConfigurationRuleExpirationOutputReference {
	var returns S3BucketLifecycleConfigurationRuleExpirationOutputReference
	_jsii_.Get(
		j,
		"expiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) ExpirationInput() *S3BucketLifecycleConfigurationRuleExpiration {
	var returns *S3BucketLifecycleConfigurationRuleExpiration
	_jsii_.Get(
		j,
		"expirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) Filter() S3BucketLifecycleConfigurationRuleFilterOutputReference {
	var returns S3BucketLifecycleConfigurationRuleFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) FilterInput() *S3BucketLifecycleConfigurationRuleFilter {
	var returns *S3BucketLifecycleConfigurationRuleFilter
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) NoncurrentVersionExpiration() S3BucketLifecycleConfigurationRuleNoncurrentVersionExpirationOutputReference {
	var returns S3BucketLifecycleConfigurationRuleNoncurrentVersionExpirationOutputReference
	_jsii_.Get(
		j,
		"noncurrentVersionExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) NoncurrentVersionExpirationInput() *S3BucketLifecycleConfigurationRuleNoncurrentVersionExpiration {
	var returns *S3BucketLifecycleConfigurationRuleNoncurrentVersionExpiration
	_jsii_.Get(
		j,
		"noncurrentVersionExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) NoncurrentVersionTransition() S3BucketLifecycleConfigurationRuleNoncurrentVersionTransitionList {
	var returns S3BucketLifecycleConfigurationRuleNoncurrentVersionTransitionList
	_jsii_.Get(
		j,
		"noncurrentVersionTransition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) NoncurrentVersionTransitionInput() any {
	var returns any
	_jsii_.Get(
		j,
		"noncurrentVersionTransitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) Transition() S3BucketLifecycleConfigurationRuleTransitionList {
	var returns S3BucketLifecycleConfigurationRuleTransitionList
	_jsii_.Get(
		j,
		"transition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) TransitionInput() any {
	var returns any
	_jsii_.Get(
		j,
		"transitionInput",
		&returns,
	)
	return returns
}

func NewS3BucketLifecycleConfigurationRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) S3BucketLifecycleConfigurationRuleOutputReference {
	_init_.Initialize()

	if err := validateNewS3BucketLifecycleConfigurationRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.s3BucketLifecycleConfiguration.S3BucketLifecycleConfigurationRuleOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewS3BucketLifecycleConfigurationRuleOutputReference_Override(s S3BucketLifecycleConfigurationRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.s3BucketLifecycleConfiguration.S3BucketLifecycleConfigurationRuleOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) PutAbortIncompleteMultipartUpload(value *S3BucketLifecycleConfigurationRuleAbortIncompleteMultipartUpload) {
	if err := s.validatePutAbortIncompleteMultipartUploadParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAbortIncompleteMultipartUpload",
		[]any{value},
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) PutExpiration(value *S3BucketLifecycleConfigurationRuleExpiration) {
	if err := s.validatePutExpirationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putExpiration",
		[]any{value},
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) PutFilter(value *S3BucketLifecycleConfigurationRuleFilter) {
	if err := s.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFilter",
		[]any{value},
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) PutNoncurrentVersionExpiration(value *S3BucketLifecycleConfigurationRuleNoncurrentVersionExpiration) {
	if err := s.validatePutNoncurrentVersionExpirationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNoncurrentVersionExpiration",
		[]any{value},
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) PutNoncurrentVersionTransition(value any) {
	if err := s.validatePutNoncurrentVersionTransitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNoncurrentVersionTransition",
		[]any{value},
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) PutTransition(value any) {
	if err := s.validatePutTransitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTransition",
		[]any{value},
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) ResetAbortIncompleteMultipartUpload() {
	_jsii_.InvokeVoid(
		s,
		"resetAbortIncompleteMultipartUpload",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) ResetExpiration() {
	_jsii_.InvokeVoid(
		s,
		"resetExpiration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		s,
		"resetFilter",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) ResetNoncurrentVersionExpiration() {
	_jsii_.InvokeVoid(
		s,
		"resetNoncurrentVersionExpiration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) ResetNoncurrentVersionTransition() {
	_jsii_.InvokeVoid(
		s,
		"resetNoncurrentVersionTransition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetPrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) ResetTransition() {
	_jsii_.InvokeVoid(
		s,
		"resetTransition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		s,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketLifecycleConfigurationRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
