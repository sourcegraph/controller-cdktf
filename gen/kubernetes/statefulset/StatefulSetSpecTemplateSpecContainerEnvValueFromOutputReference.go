package statefulset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/kubernetes/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/kubernetes/statefulset/internal"
)

type StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference interface {
	cdktf.ComplexObject
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
	ConfigMapKeyRef() StatefulSetSpecTemplateSpecContainerEnvValueFromConfigMapKeyRefOutputReference
	ConfigMapKeyRefInput() *StatefulSetSpecTemplateSpecContainerEnvValueFromConfigMapKeyRef
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FieldRef() StatefulSetSpecTemplateSpecContainerEnvValueFromFieldRefOutputReference
	FieldRefInput() *StatefulSetSpecTemplateSpecContainerEnvValueFromFieldRef
	// Experimental.
	Fqn() *string
	InternalValue() *StatefulSetSpecTemplateSpecContainerEnvValueFrom
	SetInternalValue(val *StatefulSetSpecTemplateSpecContainerEnvValueFrom)
	ResourceFieldRef() StatefulSetSpecTemplateSpecContainerEnvValueFromResourceFieldRefOutputReference
	ResourceFieldRefInput() *StatefulSetSpecTemplateSpecContainerEnvValueFromResourceFieldRef
	SecretKeyRef() StatefulSetSpecTemplateSpecContainerEnvValueFromSecretKeyRefOutputReference
	SecretKeyRefInput() *StatefulSetSpecTemplateSpecContainerEnvValueFromSecretKeyRef
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
	PutConfigMapKeyRef(value *StatefulSetSpecTemplateSpecContainerEnvValueFromConfigMapKeyRef)
	PutFieldRef(value *StatefulSetSpecTemplateSpecContainerEnvValueFromFieldRef)
	PutResourceFieldRef(value *StatefulSetSpecTemplateSpecContainerEnvValueFromResourceFieldRef)
	PutSecretKeyRef(value *StatefulSetSpecTemplateSpecContainerEnvValueFromSecretKeyRef)
	ResetConfigMapKeyRef()
	ResetFieldRef()
	ResetResourceFieldRef()
	ResetSecretKeyRef()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference
type jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) ConfigMapKeyRef() StatefulSetSpecTemplateSpecContainerEnvValueFromConfigMapKeyRefOutputReference {
	var returns StatefulSetSpecTemplateSpecContainerEnvValueFromConfigMapKeyRefOutputReference
	_jsii_.Get(
		j,
		"configMapKeyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) ConfigMapKeyRefInput() *StatefulSetSpecTemplateSpecContainerEnvValueFromConfigMapKeyRef {
	var returns *StatefulSetSpecTemplateSpecContainerEnvValueFromConfigMapKeyRef
	_jsii_.Get(
		j,
		"configMapKeyRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) FieldRef() StatefulSetSpecTemplateSpecContainerEnvValueFromFieldRefOutputReference {
	var returns StatefulSetSpecTemplateSpecContainerEnvValueFromFieldRefOutputReference
	_jsii_.Get(
		j,
		"fieldRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) FieldRefInput() *StatefulSetSpecTemplateSpecContainerEnvValueFromFieldRef {
	var returns *StatefulSetSpecTemplateSpecContainerEnvValueFromFieldRef
	_jsii_.Get(
		j,
		"fieldRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) InternalValue() *StatefulSetSpecTemplateSpecContainerEnvValueFrom {
	var returns *StatefulSetSpecTemplateSpecContainerEnvValueFrom
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) ResourceFieldRef() StatefulSetSpecTemplateSpecContainerEnvValueFromResourceFieldRefOutputReference {
	var returns StatefulSetSpecTemplateSpecContainerEnvValueFromResourceFieldRefOutputReference
	_jsii_.Get(
		j,
		"resourceFieldRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) ResourceFieldRefInput() *StatefulSetSpecTemplateSpecContainerEnvValueFromResourceFieldRef {
	var returns *StatefulSetSpecTemplateSpecContainerEnvValueFromResourceFieldRef
	_jsii_.Get(
		j,
		"resourceFieldRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) SecretKeyRef() StatefulSetSpecTemplateSpecContainerEnvValueFromSecretKeyRefOutputReference {
	var returns StatefulSetSpecTemplateSpecContainerEnvValueFromSecretKeyRefOutputReference
	_jsii_.Get(
		j,
		"secretKeyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) SecretKeyRefInput() *StatefulSetSpecTemplateSpecContainerEnvValueFromSecretKeyRef {
	var returns *StatefulSetSpecTemplateSpecContainerEnvValueFromSecretKeyRef
	_jsii_.Get(
		j,
		"secretKeyRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewStatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference {
	_init_.Initialize()

	if err := validateNewStatefulSetSpecTemplateSpecContainerEnvValueFromOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.statefulSet.StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference_Override(s StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.statefulSet.StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference",
		[]any{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) SetInternalValue(val *StatefulSetSpecTemplateSpecContainerEnvValueFrom) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (s *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) PutConfigMapKeyRef(value *StatefulSetSpecTemplateSpecContainerEnvValueFromConfigMapKeyRef) {
	if err := s.validatePutConfigMapKeyRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putConfigMapKeyRef",
		[]any{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) PutFieldRef(value *StatefulSetSpecTemplateSpecContainerEnvValueFromFieldRef) {
	if err := s.validatePutFieldRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFieldRef",
		[]any{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) PutResourceFieldRef(value *StatefulSetSpecTemplateSpecContainerEnvValueFromResourceFieldRef) {
	if err := s.validatePutResourceFieldRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceFieldRef",
		[]any{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) PutSecretKeyRef(value *StatefulSetSpecTemplateSpecContainerEnvValueFromSecretKeyRef) {
	if err := s.validatePutSecretKeyRefParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSecretKeyRef",
		[]any{value},
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) ResetConfigMapKeyRef() {
	_jsii_.InvokeVoid(
		s,
		"resetConfigMapKeyRef",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) ResetFieldRef() {
	_jsii_.InvokeVoid(
		s,
		"resetFieldRef",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) ResetResourceFieldRef() {
	_jsii_.InvokeVoid(
		s,
		"resetResourceFieldRef",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) ResetSecretKeyRef() {
	_jsii_.InvokeVoid(
		s,
		"resetSecretKeyRef",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (s *jsiiProxy_StatefulSetSpecTemplateSpecContainerEnvValueFromOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
