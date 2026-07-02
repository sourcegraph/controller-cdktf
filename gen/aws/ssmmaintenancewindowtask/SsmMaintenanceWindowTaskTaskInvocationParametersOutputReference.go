package ssmmaintenancewindowtask

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/ssmmaintenancewindowtask/internal"
)

type SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference interface {
	cdktf.ComplexObject
	AutomationParameters() SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference
	AutomationParametersInput() *SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParameters
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
	InternalValue() *SsmMaintenanceWindowTaskTaskInvocationParameters
	SetInternalValue(val *SsmMaintenanceWindowTaskTaskInvocationParameters)
	LambdaParameters() SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference
	LambdaParametersInput() *SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParameters
	RunCommandParameters() SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference
	RunCommandParametersInput() *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters
	StepFunctionsParameters() SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference
	StepFunctionsParametersInput() *SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters
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
	PutAutomationParameters(value *SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParameters)
	PutLambdaParameters(value *SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParameters)
	PutRunCommandParameters(value *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters)
	PutStepFunctionsParameters(value *SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters)
	ResetAutomationParameters()
	ResetLambdaParameters()
	ResetRunCommandParameters()
	ResetStepFunctionsParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference
type jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) AutomationParameters() SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference {
	var returns SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParametersOutputReference
	_jsii_.Get(
		j,
		"automationParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) AutomationParametersInput() *SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParameters {
	var returns *SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParameters
	_jsii_.Get(
		j,
		"automationParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) InternalValue() *SsmMaintenanceWindowTaskTaskInvocationParameters {
	var returns *SsmMaintenanceWindowTaskTaskInvocationParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) LambdaParameters() SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference {
	var returns SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParametersOutputReference
	_jsii_.Get(
		j,
		"lambdaParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) LambdaParametersInput() *SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParameters {
	var returns *SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParameters
	_jsii_.Get(
		j,
		"lambdaParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) RunCommandParameters() SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference {
	var returns SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersOutputReference
	_jsii_.Get(
		j,
		"runCommandParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) RunCommandParametersInput() *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters {
	var returns *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters
	_jsii_.Get(
		j,
		"runCommandParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) StepFunctionsParameters() SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference {
	var returns SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersOutputReference
	_jsii_.Get(
		j,
		"stepFunctionsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) StepFunctionsParametersInput() *SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters {
	var returns *SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters
	_jsii_.Get(
		j,
		"stepFunctionsParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewSsmMaintenanceWindowTaskTaskInvocationParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference {
	_init_.Initialize()

	if err := validateNewSsmMaintenanceWindowTaskTaskInvocationParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSsmMaintenanceWindowTaskTaskInvocationParametersOutputReference_Override(s SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ssmMaintenanceWindowTask.SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference",
		[]any{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) SetInternalValue(val *SsmMaintenanceWindowTaskTaskInvocationParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) PutAutomationParameters(value *SsmMaintenanceWindowTaskTaskInvocationParametersAutomationParameters) {
	if err := s.validatePutAutomationParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAutomationParameters",
		[]any{value},
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) PutLambdaParameters(value *SsmMaintenanceWindowTaskTaskInvocationParametersLambdaParameters) {
	if err := s.validatePutLambdaParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLambdaParameters",
		[]any{value},
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) PutRunCommandParameters(value *SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParameters) {
	if err := s.validatePutRunCommandParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRunCommandParameters",
		[]any{value},
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) PutStepFunctionsParameters(value *SsmMaintenanceWindowTaskTaskInvocationParametersStepFunctionsParameters) {
	if err := s.validatePutStepFunctionsParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putStepFunctionsParameters",
		[]any{value},
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) ResetAutomationParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetAutomationParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) ResetLambdaParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetLambdaParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) ResetRunCommandParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetRunCommandParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) ResetStepFunctionsParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetStepFunctionsParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (s *jsiiProxy_SsmMaintenanceWindowTaskTaskInvocationParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
