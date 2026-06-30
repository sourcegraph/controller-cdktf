package bigqueryconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/bigqueryconnection/internal"
)

type BigqueryConnectionCloudSpannerOutputReference interface {
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
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	DatabaseRole() *string
	SetDatabaseRole(val *string)
	DatabaseRoleInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *BigqueryConnectionCloudSpanner
	SetInternalValue(val *BigqueryConnectionCloudSpanner)
	MaxParallelism() *float64
	SetMaxParallelism(val *float64)
	MaxParallelismInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseDataBoost() any
	SetUseDataBoost(val any)
	UseDataBoostInput() any
	UseParallelism() any
	SetUseParallelism(val any)
	UseParallelismInput() any
	UseServerlessAnalytics() any
	SetUseServerlessAnalytics(val any)
	UseServerlessAnalyticsInput() any
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
	ResetDatabaseRole()
	ResetMaxParallelism()
	ResetUseDataBoost()
	ResetUseParallelism()
	ResetUseServerlessAnalytics()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BigqueryConnectionCloudSpannerOutputReference
type jsiiProxy_BigqueryConnectionCloudSpannerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) DatabaseRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) DatabaseRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) InternalValue() *BigqueryConnectionCloudSpanner {
	var returns *BigqueryConnectionCloudSpanner
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) MaxParallelism() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) MaxParallelismInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) UseDataBoost() any {
	var returns any
	_jsii_.Get(
		j,
		"useDataBoost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) UseDataBoostInput() any {
	var returns any
	_jsii_.Get(
		j,
		"useDataBoostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) UseParallelism() any {
	var returns any
	_jsii_.Get(
		j,
		"useParallelism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) UseParallelismInput() any {
	var returns any
	_jsii_.Get(
		j,
		"useParallelismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) UseServerlessAnalytics() any {
	var returns any
	_jsii_.Get(
		j,
		"useServerlessAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) UseServerlessAnalyticsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"useServerlessAnalyticsInput",
		&returns,
	)
	return returns
}

func NewBigqueryConnectionCloudSpannerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BigqueryConnectionCloudSpannerOutputReference {
	_init_.Initialize()

	if err := validateNewBigqueryConnectionCloudSpannerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BigqueryConnectionCloudSpannerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryConnection.BigqueryConnectionCloudSpannerOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBigqueryConnectionCloudSpannerOutputReference_Override(b BigqueryConnectionCloudSpannerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryConnection.BigqueryConnectionCloudSpannerOutputReference",
		[]any{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) SetDatabaseRole(val *string) {
	if err := j.validateSetDatabaseRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseRole",
		val,
	)
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) SetInternalValue(val *BigqueryConnectionCloudSpanner) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) SetMaxParallelism(val *float64) {
	if err := j.validateSetMaxParallelismParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxParallelism",
		val,
	)
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) SetUseDataBoost(val any) {
	if err := j.validateSetUseDataBoostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useDataBoost",
		val,
	)
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) SetUseParallelism(val any) {
	if err := j.validateSetUseParallelismParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useParallelism",
		val,
	)
}

func (j *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) SetUseServerlessAnalytics(val any) {
	if err := j.validateSetUseServerlessAnalyticsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useServerlessAnalytics",
		val,
	)
}

func (b *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (b *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) ResetDatabaseRole() {
	_jsii_.InvokeVoid(
		b,
		"resetDatabaseRole",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) ResetMaxParallelism() {
	_jsii_.InvokeVoid(
		b,
		"resetMaxParallelism",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) ResetUseDataBoost() {
	_jsii_.InvokeVoid(
		b,
		"resetUseDataBoost",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) ResetUseParallelism() {
	_jsii_.InvokeVoid(
		b,
		"resetUseParallelism",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) ResetUseServerlessAnalytics() {
	_jsii_.InvokeVoid(
		b,
		"resetUseServerlessAnalytics",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (b *jsiiProxy_BigqueryConnectionCloudSpannerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
