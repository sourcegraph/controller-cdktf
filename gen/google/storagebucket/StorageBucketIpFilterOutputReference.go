package storagebucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/storagebucket/internal"
)

type StorageBucketIpFilterOutputReference interface {
	cdktf.ComplexObject
	AllowAllServiceAgentAccess() any
	SetAllowAllServiceAgentAccess(val any)
	AllowAllServiceAgentAccessInput() any
	AllowCrossOrgVpcs() any
	SetAllowCrossOrgVpcs(val any)
	AllowCrossOrgVpcsInput() any
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
	// Experimental.
	Fqn() *string
	InternalValue() *StorageBucketIpFilter
	SetInternalValue(val *StorageBucketIpFilter)
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	PublicNetworkSource() StorageBucketIpFilterPublicNetworkSourceOutputReference
	PublicNetworkSourceInput() *StorageBucketIpFilterPublicNetworkSource
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpcNetworkSources() StorageBucketIpFilterVpcNetworkSourcesList
	VpcNetworkSourcesInput() any
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
	PutPublicNetworkSource(value *StorageBucketIpFilterPublicNetworkSource)
	PutVpcNetworkSources(value any)
	ResetAllowAllServiceAgentAccess()
	ResetAllowCrossOrgVpcs()
	ResetPublicNetworkSource()
	ResetVpcNetworkSources()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageBucketIpFilterOutputReference
type jsiiProxy_StorageBucketIpFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageBucketIpFilterOutputReference) AllowAllServiceAgentAccess() any {
	var returns any
	_jsii_.Get(
		j,
		"allowAllServiceAgentAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketIpFilterOutputReference) AllowAllServiceAgentAccessInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowAllServiceAgentAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketIpFilterOutputReference) AllowCrossOrgVpcs() any {
	var returns any
	_jsii_.Get(
		j,
		"allowCrossOrgVpcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketIpFilterOutputReference) AllowCrossOrgVpcsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowCrossOrgVpcsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketIpFilterOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketIpFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketIpFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketIpFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketIpFilterOutputReference) InternalValue() *StorageBucketIpFilter {
	var returns *StorageBucketIpFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketIpFilterOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketIpFilterOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketIpFilterOutputReference) PublicNetworkSource() StorageBucketIpFilterPublicNetworkSourceOutputReference {
	var returns StorageBucketIpFilterPublicNetworkSourceOutputReference
	_jsii_.Get(
		j,
		"publicNetworkSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketIpFilterOutputReference) PublicNetworkSourceInput() *StorageBucketIpFilterPublicNetworkSource {
	var returns *StorageBucketIpFilterPublicNetworkSource
	_jsii_.Get(
		j,
		"publicNetworkSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketIpFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketIpFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketIpFilterOutputReference) VpcNetworkSources() StorageBucketIpFilterVpcNetworkSourcesList {
	var returns StorageBucketIpFilterVpcNetworkSourcesList
	_jsii_.Get(
		j,
		"vpcNetworkSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBucketIpFilterOutputReference) VpcNetworkSourcesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"vpcNetworkSourcesInput",
		&returns,
	)
	return returns
}

func NewStorageBucketIpFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageBucketIpFilterOutputReference {
	_init_.Initialize()

	if err := validateNewStorageBucketIpFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageBucketIpFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.storageBucket.StorageBucketIpFilterOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageBucketIpFilterOutputReference_Override(s StorageBucketIpFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.storageBucket.StorageBucketIpFilterOutputReference",
		[]any{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageBucketIpFilterOutputReference) SetAllowAllServiceAgentAccess(val any) {
	if err := j.validateSetAllowAllServiceAgentAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAllServiceAgentAccess",
		val,
	)
}

func (j *jsiiProxy_StorageBucketIpFilterOutputReference) SetAllowCrossOrgVpcs(val any) {
	if err := j.validateSetAllowCrossOrgVpcsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowCrossOrgVpcs",
		val,
	)
}

func (j *jsiiProxy_StorageBucketIpFilterOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageBucketIpFilterOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageBucketIpFilterOutputReference) SetInternalValue(val *StorageBucketIpFilter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageBucketIpFilterOutputReference) SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_StorageBucketIpFilterOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageBucketIpFilterOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageBucketIpFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketIpFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (s *jsiiProxy_StorageBucketIpFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageBucketIpFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StorageBucketIpFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StorageBucketIpFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StorageBucketIpFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StorageBucketIpFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StorageBucketIpFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StorageBucketIpFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StorageBucketIpFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBucketIpFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageBucketIpFilterOutputReference) PutPublicNetworkSource(value *StorageBucketIpFilterPublicNetworkSource) {
	if err := s.validatePutPublicNetworkSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPublicNetworkSource",
		[]any{value},
	)
}

func (s *jsiiProxy_StorageBucketIpFilterOutputReference) PutVpcNetworkSources(value any) {
	if err := s.validatePutVpcNetworkSourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVpcNetworkSources",
		[]any{value},
	)
}

func (s *jsiiProxy_StorageBucketIpFilterOutputReference) ResetAllowAllServiceAgentAccess() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowAllServiceAgentAccess",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucketIpFilterOutputReference) ResetAllowCrossOrgVpcs() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowCrossOrgVpcs",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucketIpFilterOutputReference) ResetPublicNetworkSource() {
	_jsii_.InvokeVoid(
		s,
		"resetPublicNetworkSource",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucketIpFilterOutputReference) ResetVpcNetworkSources() {
	_jsii_.InvokeVoid(
		s,
		"resetVpcNetworkSources",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBucketIpFilterOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (s *jsiiProxy_StorageBucketIpFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
