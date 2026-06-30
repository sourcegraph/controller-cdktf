package budgetsbudget

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/budgetsbudget/internal"
)

type BudgetsBudgetCostTypesOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IncludeCredit() any
	SetIncludeCredit(val any)
	IncludeCreditInput() any
	IncludeDiscount() any
	SetIncludeDiscount(val any)
	IncludeDiscountInput() any
	IncludeOtherSubscription() any
	SetIncludeOtherSubscription(val any)
	IncludeOtherSubscriptionInput() any
	IncludeRecurring() any
	SetIncludeRecurring(val any)
	IncludeRecurringInput() any
	IncludeRefund() any
	SetIncludeRefund(val any)
	IncludeRefundInput() any
	IncludeSubscription() any
	SetIncludeSubscription(val any)
	IncludeSubscriptionInput() any
	IncludeSupport() any
	SetIncludeSupport(val any)
	IncludeSupportInput() any
	IncludeTax() any
	SetIncludeTax(val any)
	IncludeTaxInput() any
	IncludeUpfront() any
	SetIncludeUpfront(val any)
	IncludeUpfrontInput() any
	InternalValue() *BudgetsBudgetCostTypes
	SetInternalValue(val *BudgetsBudgetCostTypes)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseAmortized() any
	SetUseAmortized(val any)
	UseAmortizedInput() any
	UseBlended() any
	SetUseBlended(val any)
	UseBlendedInput() any
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
	ResetIncludeCredit()
	ResetIncludeDiscount()
	ResetIncludeOtherSubscription()
	ResetIncludeRecurring()
	ResetIncludeRefund()
	ResetIncludeSubscription()
	ResetIncludeSupport()
	ResetIncludeTax()
	ResetIncludeUpfront()
	ResetUseAmortized()
	ResetUseBlended()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BudgetsBudgetCostTypesOutputReference
type jsiiProxy_BudgetsBudgetCostTypesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeCredit() any {
	var returns any
	_jsii_.Get(
		j,
		"includeCredit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeCreditInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeCreditInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeDiscount() any {
	var returns any
	_jsii_.Get(
		j,
		"includeDiscount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeDiscountInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeDiscountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeOtherSubscription() any {
	var returns any
	_jsii_.Get(
		j,
		"includeOtherSubscription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeOtherSubscriptionInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeOtherSubscriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeRecurring() any {
	var returns any
	_jsii_.Get(
		j,
		"includeRecurring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeRecurringInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeRecurringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeRefund() any {
	var returns any
	_jsii_.Get(
		j,
		"includeRefund",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeRefundInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeRefundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeSubscription() any {
	var returns any
	_jsii_.Get(
		j,
		"includeSubscription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeSubscriptionInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeSubscriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeSupport() any {
	var returns any
	_jsii_.Get(
		j,
		"includeSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeSupportInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeTax() any {
	var returns any
	_jsii_.Get(
		j,
		"includeTax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeTaxInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeTaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeUpfront() any {
	var returns any
	_jsii_.Get(
		j,
		"includeUpfront",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeUpfrontInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeUpfrontInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) InternalValue() *BudgetsBudgetCostTypes {
	var returns *BudgetsBudgetCostTypes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) UseAmortized() any {
	var returns any
	_jsii_.Get(
		j,
		"useAmortized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) UseAmortizedInput() any {
	var returns any
	_jsii_.Get(
		j,
		"useAmortizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) UseBlended() any {
	var returns any
	_jsii_.Get(
		j,
		"useBlended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) UseBlendedInput() any {
	var returns any
	_jsii_.Get(
		j,
		"useBlendedInput",
		&returns,
	)
	return returns
}

func NewBudgetsBudgetCostTypesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BudgetsBudgetCostTypesOutputReference {
	_init_.Initialize()

	if err := validateNewBudgetsBudgetCostTypesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BudgetsBudgetCostTypesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudgetCostTypesOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBudgetsBudgetCostTypesOutputReference_Override(b BudgetsBudgetCostTypesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudgetCostTypesOutputReference",
		[]any{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetIncludeCredit(val any) {
	if err := j.validateSetIncludeCreditParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeCredit",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetIncludeDiscount(val any) {
	if err := j.validateSetIncludeDiscountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeDiscount",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetIncludeOtherSubscription(val any) {
	if err := j.validateSetIncludeOtherSubscriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeOtherSubscription",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetIncludeRecurring(val any) {
	if err := j.validateSetIncludeRecurringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeRecurring",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetIncludeRefund(val any) {
	if err := j.validateSetIncludeRefundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeRefund",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetIncludeSubscription(val any) {
	if err := j.validateSetIncludeSubscriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSubscription",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetIncludeSupport(val any) {
	if err := j.validateSetIncludeSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSupport",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetIncludeTax(val any) {
	if err := j.validateSetIncludeTaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeTax",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetIncludeUpfront(val any) {
	if err := j.validateSetIncludeUpfrontParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeUpfront",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetInternalValue(val *BudgetsBudgetCostTypes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetUseAmortized(val any) {
	if err := j.validateSetUseAmortizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAmortized",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetUseBlended(val any) {
	if err := j.validateSetUseBlendedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useBlended",
		val,
	)
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ResetIncludeCredit() {
	_jsii_.InvokeVoid(
		b,
		"resetIncludeCredit",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ResetIncludeDiscount() {
	_jsii_.InvokeVoid(
		b,
		"resetIncludeDiscount",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ResetIncludeOtherSubscription() {
	_jsii_.InvokeVoid(
		b,
		"resetIncludeOtherSubscription",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ResetIncludeRecurring() {
	_jsii_.InvokeVoid(
		b,
		"resetIncludeRecurring",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ResetIncludeRefund() {
	_jsii_.InvokeVoid(
		b,
		"resetIncludeRefund",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ResetIncludeSubscription() {
	_jsii_.InvokeVoid(
		b,
		"resetIncludeSubscription",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ResetIncludeSupport() {
	_jsii_.InvokeVoid(
		b,
		"resetIncludeSupport",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ResetIncludeTax() {
	_jsii_.InvokeVoid(
		b,
		"resetIncludeTax",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ResetIncludeUpfront() {
	_jsii_.InvokeVoid(
		b,
		"resetIncludeUpfront",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ResetUseAmortized() {
	_jsii_.InvokeVoid(
		b,
		"resetUseAmortized",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ResetUseBlended() {
	_jsii_.InvokeVoid(
		b,
		"resetUseBlended",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		b,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
