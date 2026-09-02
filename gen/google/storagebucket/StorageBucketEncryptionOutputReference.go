package storagebucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/storagebucket/internal"
)

type StorageBucketEncryptionOutputReference interface {
	cdktn.ComplexObject
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
	CustomerManagedEncryptionEnforcementConfig() StorageBucketEncryptionCustomerManagedEncryptionEnforcementConfigOutputReference
	CustomerManagedEncryptionEnforcementConfigInput() *StorageBucketEncryptionCustomerManagedEncryptionEnforcementConfig
	CustomerSuppliedEncryptionEnforcementConfig() StorageBucketEncryptionCustomerSuppliedEncryptionEnforcementConfigOutputReference
	CustomerSuppliedEncryptionEnforcementConfigInput() *StorageBucketEncryptionCustomerSuppliedEncryptionEnforcementConfig
	DefaultKmsKeyName() *string
	SetDefaultKmsKeyName(val *string)
	DefaultKmsKeyNameInput() *string
	// Experimental.
	Fqn() *string
	GoogleManagedEncryptionEnforcementConfig() StorageBucketEncryptionGoogleManagedEncryptionEnforcementConfigOutputReference
	GoogleManagedEncryptionEnforcementConfigInput() *StorageBucketEncryptionGoogleManagedEncryptionEnforcementConfig
	InternalValue() *StorageBucketEncryption
	SetInternalValue(val *StorageBucketEncryption)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutCustomerManagedEncryptionEnforcementConfig(value *StorageBucketEncryptionCustomerManagedEncryptionEnforcementConfig)
	PutCustomerSuppliedEncryptionEnforcementConfig(value *StorageBucketEncryptionCustomerSuppliedEncryptionEnforcementConfig)
	PutGoogleManagedEncryptionEnforcementConfig(value *StorageBucketEncryptionGoogleManagedEncryptionEnforcementConfig)
	ResetCustomerManagedEncryptionEnforcementConfig()
	ResetCustomerSuppliedEncryptionEnforcementConfig()
	ResetDefaultKmsKeyName()
	ResetGoogleManagedEncryptionEnforcementConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageBucketEncryptionOutputReference
type jsiiProxy_StorageBucketEncryptionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_StorageBucketEncryptionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketEncryptionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketEncryptionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketEncryptionOutputReference) CustomerManagedEncryptionEnforcementConfig() StorageBucketEncryptionCustomerManagedEncryptionEnforcementConfigOutputReference {
	var returns StorageBucketEncryptionCustomerManagedEncryptionEnforcementConfigOutputReference
	_jsii_.Get(
		j,
		"customerManagedEncryptionEnforcementConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketEncryptionOutputReference) CustomerManagedEncryptionEnforcementConfigInput() *StorageBucketEncryptionCustomerManagedEncryptionEnforcementConfig {
	var returns *StorageBucketEncryptionCustomerManagedEncryptionEnforcementConfig
	_jsii_.Get(
		j,
		"customerManagedEncryptionEnforcementConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketEncryptionOutputReference) CustomerSuppliedEncryptionEnforcementConfig() StorageBucketEncryptionCustomerSuppliedEncryptionEnforcementConfigOutputReference {
	var returns StorageBucketEncryptionCustomerSuppliedEncryptionEnforcementConfigOutputReference
	_jsii_.Get(
		j,
		"customerSuppliedEncryptionEnforcementConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketEncryptionOutputReference) CustomerSuppliedEncryptionEnforcementConfigInput() *StorageBucketEncryptionCustomerSuppliedEncryptionEnforcementConfig {
	var returns *StorageBucketEncryptionCustomerSuppliedEncryptionEnforcementConfig
	_jsii_.Get(
		j,
		"customerSuppliedEncryptionEnforcementConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketEncryptionOutputReference) DefaultKmsKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultKmsKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketEncryptionOutputReference) DefaultKmsKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultKmsKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketEncryptionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketEncryptionOutputReference) GoogleManagedEncryptionEnforcementConfig() StorageBucketEncryptionGoogleManagedEncryptionEnforcementConfigOutputReference {
	var returns StorageBucketEncryptionGoogleManagedEncryptionEnforcementConfigOutputReference
	_jsii_.Get(
		j,
		"googleManagedEncryptionEnforcementConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketEncryptionOutputReference) GoogleManagedEncryptionEnforcementConfigInput() *StorageBucketEncryptionGoogleManagedEncryptionEnforcementConfig {
	var returns *StorageBucketEncryptionGoogleManagedEncryptionEnforcementConfig
	_jsii_.Get(
		j,
		"googleManagedEncryptionEnforcementConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketEncryptionOutputReference) InternalValue() *StorageBucketEncryption {
	var returns *StorageBucketEncryption
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketEncryptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketEncryptionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageBucketEncryptionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) StorageBucketEncryptionOutputReference {
	_init_.Initialize()

	if err := validateNewStorageBucketEncryptionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageBucketEncryptionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.storageBucket.StorageBucketEncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageBucketEncryptionOutputReference_Override(s StorageBucketEncryptionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.storageBucket.StorageBucketEncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageBucketEncryptionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageBucketEncryptionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageBucketEncryptionOutputReference)SetDefaultKmsKeyName(val *string) {
	if err := j.validateSetDefaultKmsKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultKmsKeyName",
		val,
	)
}

func (j *jsiiProxy_StorageBucketEncryptionOutputReference)SetInternalValue(val *StorageBucketEncryption) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageBucketEncryptionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageBucketEncryptionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageBucketEncryptionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketEncryptionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketEncryptionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketEncryptionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketEncryptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketEncryptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketEncryptionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketEncryptionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketEncryptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketEncryptionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketEncryptionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketEncryptionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketEncryptionOutputReference) PutCustomerManagedEncryptionEnforcementConfig(value *StorageBucketEncryptionCustomerManagedEncryptionEnforcementConfig) {
	if err := s.validatePutCustomerManagedEncryptionEnforcementConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCustomerManagedEncryptionEnforcementConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageBucketEncryptionOutputReference) PutCustomerSuppliedEncryptionEnforcementConfig(value *StorageBucketEncryptionCustomerSuppliedEncryptionEnforcementConfig) {
	if err := s.validatePutCustomerSuppliedEncryptionEnforcementConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCustomerSuppliedEncryptionEnforcementConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageBucketEncryptionOutputReference) PutGoogleManagedEncryptionEnforcementConfig(value *StorageBucketEncryptionGoogleManagedEncryptionEnforcementConfig) {
	if err := s.validatePutGoogleManagedEncryptionEnforcementConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGoogleManagedEncryptionEnforcementConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageBucketEncryptionOutputReference) ResetCustomerManagedEncryptionEnforcementConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomerManagedEncryptionEnforcementConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucketEncryptionOutputReference) ResetCustomerSuppliedEncryptionEnforcementConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomerSuppliedEncryptionEnforcementConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucketEncryptionOutputReference) ResetDefaultKmsKeyName() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultKmsKeyName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucketEncryptionOutputReference) ResetGoogleManagedEncryptionEnforcementConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetGoogleManagedEncryptionEnforcementConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucketEncryptionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketEncryptionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

