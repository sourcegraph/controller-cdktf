package netappbackupvault

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/netappbackupvault/internal"
)

type NetappBackupVaultBackupRetentionPolicyOutputReference interface {
	cdktf.ComplexObject
	BackupMinimumEnforcedRetentionDays() *float64
	SetBackupMinimumEnforcedRetentionDays(val *float64)
	BackupMinimumEnforcedRetentionDaysInput() *float64
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
	DailyBackupImmutable() any
	SetDailyBackupImmutable(val any)
	DailyBackupImmutableInput() any
	// Experimental.
	Fqn() *string
	InternalValue() *NetappBackupVaultBackupRetentionPolicy
	SetInternalValue(val *NetappBackupVaultBackupRetentionPolicy)
	ManualBackupImmutable() any
	SetManualBackupImmutable(val any)
	ManualBackupImmutableInput() any
	MonthlyBackupImmutable() any
	SetMonthlyBackupImmutable(val any)
	MonthlyBackupImmutableInput() any
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeeklyBackupImmutable() any
	SetWeeklyBackupImmutable(val any)
	WeeklyBackupImmutableInput() any
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
	ResetDailyBackupImmutable()
	ResetManualBackupImmutable()
	ResetMonthlyBackupImmutable()
	ResetWeeklyBackupImmutable()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetappBackupVaultBackupRetentionPolicyOutputReference
type jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) BackupMinimumEnforcedRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupMinimumEnforcedRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) BackupMinimumEnforcedRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupMinimumEnforcedRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) DailyBackupImmutable() any {
	var returns any
	_jsii_.Get(
		j,
		"dailyBackupImmutable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) DailyBackupImmutableInput() any {
	var returns any
	_jsii_.Get(
		j,
		"dailyBackupImmutableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) InternalValue() *NetappBackupVaultBackupRetentionPolicy {
	var returns *NetappBackupVaultBackupRetentionPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) ManualBackupImmutable() any {
	var returns any
	_jsii_.Get(
		j,
		"manualBackupImmutable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) ManualBackupImmutableInput() any {
	var returns any
	_jsii_.Get(
		j,
		"manualBackupImmutableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) MonthlyBackupImmutable() any {
	var returns any
	_jsii_.Get(
		j,
		"monthlyBackupImmutable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) MonthlyBackupImmutableInput() any {
	var returns any
	_jsii_.Get(
		j,
		"monthlyBackupImmutableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) WeeklyBackupImmutable() any {
	var returns any
	_jsii_.Get(
		j,
		"weeklyBackupImmutable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) WeeklyBackupImmutableInput() any {
	var returns any
	_jsii_.Get(
		j,
		"weeklyBackupImmutableInput",
		&returns,
	)
	return returns
}

func NewNetappBackupVaultBackupRetentionPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetappBackupVaultBackupRetentionPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewNetappBackupVaultBackupRetentionPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.netappBackupVault.NetappBackupVaultBackupRetentionPolicyOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetappBackupVaultBackupRetentionPolicyOutputReference_Override(n NetappBackupVaultBackupRetentionPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.netappBackupVault.NetappBackupVaultBackupRetentionPolicyOutputReference",
		[]any{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) SetBackupMinimumEnforcedRetentionDays(val *float64) {
	if err := j.validateSetBackupMinimumEnforcedRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupMinimumEnforcedRetentionDays",
		val,
	)
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) SetDailyBackupImmutable(val any) {
	if err := j.validateSetDailyBackupImmutableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dailyBackupImmutable",
		val,
	)
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) SetInternalValue(val *NetappBackupVaultBackupRetentionPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) SetManualBackupImmutable(val any) {
	if err := j.validateSetManualBackupImmutableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manualBackupImmutable",
		val,
	)
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) SetMonthlyBackupImmutable(val any) {
	if err := j.validateSetMonthlyBackupImmutableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monthlyBackupImmutable",
		val,
	)
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) SetWeeklyBackupImmutable(val any) {
	if err := j.validateSetWeeklyBackupImmutableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weeklyBackupImmutable",
		val,
	)
}

func (n *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) ResetDailyBackupImmutable() {
	_jsii_.InvokeVoid(
		n,
		"resetDailyBackupImmutable",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) ResetManualBackupImmutable() {
	_jsii_.InvokeVoid(
		n,
		"resetManualBackupImmutable",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) ResetMonthlyBackupImmutable() {
	_jsii_.InvokeVoid(
		n,
		"resetMonthlyBackupImmutable",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) ResetWeeklyBackupImmutable() {
	_jsii_.InvokeVoid(
		n,
		"resetWeeklyBackupImmutable",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		n,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappBackupVaultBackupRetentionPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
