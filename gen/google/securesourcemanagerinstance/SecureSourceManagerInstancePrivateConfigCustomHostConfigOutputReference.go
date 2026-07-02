package securesourcemanagerinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/securesourcemanagerinstance/internal"
)

type SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference interface {
	cdktf.ComplexObject
	Api() *string
	SetApi(val *string)
	ApiInput() *string
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
	GitHttp() *string
	SetGitHttp(val *string)
	GitHttpInput() *string
	GitSsh() *string
	SetGitSsh(val *string)
	GitSshInput() *string
	Html() *string
	SetHtml(val *string)
	HtmlInput() *string
	InternalValue() *SecureSourceManagerInstancePrivateConfigCustomHostConfig
	SetInternalValue(val *SecureSourceManagerInstancePrivateConfigCustomHostConfig)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference
type jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) Api() *string {
	var returns *string
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) ApiInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) GitHttp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitHttp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) GitHttpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitHttpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) GitSsh() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitSsh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) GitSshInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitSshInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) Html() *string {
	var returns *string
	_jsii_.Get(
		j,
		"html",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) HtmlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"htmlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) InternalValue() *SecureSourceManagerInstancePrivateConfigCustomHostConfig {
	var returns *SecureSourceManagerInstancePrivateConfigCustomHostConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewSecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.secureSourceManagerInstance.SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference_Override(s SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.secureSourceManagerInstance.SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) SetApi(val *string) {
	if err := j.validateSetApiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"api",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) SetGitHttp(val *string) {
	if err := j.validateSetGitHttpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitHttp",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) SetGitSsh(val *string) {
	if err := j.validateSetGitSshParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitSsh",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) SetHtml(val *string) {
	if err := j.validateSetHtmlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"html",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) SetInternalValue(val *SecureSourceManagerInstancePrivateConfigCustomHostConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (s *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (s *jsiiProxy_SecureSourceManagerInstancePrivateConfigCustomHostConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
