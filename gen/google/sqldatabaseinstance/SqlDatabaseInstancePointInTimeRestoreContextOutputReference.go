package sqldatabaseinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/sqldatabaseinstance/internal"
)

type SqlDatabaseInstancePointInTimeRestoreContextOutputReference interface {
	cdktf.ComplexObject
	AllocatedIpRange() *string
	SetAllocatedIpRange(val *string)
	AllocatedIpRangeInput() *string
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
	Datasource() *string
	SetDatasource(val *string)
	DatasourceInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *SqlDatabaseInstancePointInTimeRestoreContext
	SetInternalValue(val *SqlDatabaseInstancePointInTimeRestoreContext)
	PointInTime() *string
	SetPointInTime(val *string)
	PointInTimeInput() *string
	PreferredZone() *string
	SetPreferredZone(val *string)
	PreferredZoneInput() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	TargetInstance() *string
	SetTargetInstance(val *string)
	TargetInstanceInput() *string
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
	ResetAllocatedIpRange()
	ResetPointInTime()
	ResetPreferredZone()
	ResetRegion()
	ResetTargetInstance()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SqlDatabaseInstancePointInTimeRestoreContextOutputReference
type jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) AllocatedIpRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocatedIpRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) AllocatedIpRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocatedIpRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) Datasource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) DatasourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) InternalValue() *SqlDatabaseInstancePointInTimeRestoreContext {
	var returns *SqlDatabaseInstancePointInTimeRestoreContext
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) PointInTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pointInTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) PointInTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pointInTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) PreferredZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) PreferredZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) TargetInstance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) TargetInstanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewSqlDatabaseInstancePointInTimeRestoreContextOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SqlDatabaseInstancePointInTimeRestoreContextOutputReference {
	_init_.Initialize()

	if err := validateNewSqlDatabaseInstancePointInTimeRestoreContextOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.sqlDatabaseInstance.SqlDatabaseInstancePointInTimeRestoreContextOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSqlDatabaseInstancePointInTimeRestoreContextOutputReference_Override(s SqlDatabaseInstancePointInTimeRestoreContextOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.sqlDatabaseInstance.SqlDatabaseInstancePointInTimeRestoreContextOutputReference",
		[]any{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) SetAllocatedIpRange(val *string) {
	if err := j.validateSetAllocatedIpRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocatedIpRange",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) SetDatasource(val *string) {
	if err := j.validateSetDatasourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasource",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) SetInternalValue(val *SqlDatabaseInstancePointInTimeRestoreContext) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) SetPointInTime(val *string) {
	if err := j.validateSetPointInTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pointInTime",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) SetPreferredZone(val *string) {
	if err := j.validateSetPreferredZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredZone",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) SetTargetInstance(val *string) {
	if err := j.validateSetTargetInstanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetInstance",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (s *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) ResetAllocatedIpRange() {
	_jsii_.InvokeVoid(
		s,
		"resetAllocatedIpRange",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) ResetPointInTime() {
	_jsii_.InvokeVoid(
		s,
		"resetPointInTime",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) ResetPreferredZone() {
	_jsii_.InvokeVoid(
		s,
		"resetPreferredZone",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		s,
		"resetRegion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) ResetTargetInstance() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetInstance",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (s *jsiiProxy_SqlDatabaseInstancePointInTimeRestoreContextOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
