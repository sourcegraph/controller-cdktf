package storagetransferjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/storagetransferjob/internal"
)

type StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference interface {
	cdktf.ComplexObject
	AuthMethod() *string
	SetAuthMethod(val *string)
	AuthMethodInput() *string
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
	InternalValue() *StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3Metadata
	SetInternalValue(val *StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3Metadata)
	ListApi() *string
	SetListApi(val *string)
	ListApiInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	RequestModel() *string
	SetRequestModel(val *string)
	RequestModelInput() *string
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
	ResetAuthMethod()
	ResetListApi()
	ResetProtocol()
	ResetRequestModel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference
type jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) AuthMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) AuthMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) InternalValue() *StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3Metadata {
	var returns *StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3Metadata
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) ListApi() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) ListApiInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) RequestModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) RequestModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference {
	_init_.Initialize()

	if err := validateNewStorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.storageTransferJob.StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference_Override(s StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.storageTransferJob.StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference)SetAuthMethod(val *string) {
	if err := j.validateSetAuthMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authMethod",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference)SetInternalValue(val *StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3Metadata) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference)SetListApi(val *string) {
	if err := j.validateSetListApiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listApi",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference)SetRequestModel(val *string) {
	if err := j.validateSetRequestModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestModel",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) ResetAuthMethod() {
	_jsii_.InvokeVoid(
		s,
		"resetAuthMethod",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) ResetListApi() {
	_jsii_.InvokeVoid(
		s,
		"resetListApi",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		s,
		"resetProtocol",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) ResetRequestModel() {
	_jsii_.InvokeVoid(
		s,
		"resetRequestModel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecAwsS3CompatibleDataSourceS3MetadataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

