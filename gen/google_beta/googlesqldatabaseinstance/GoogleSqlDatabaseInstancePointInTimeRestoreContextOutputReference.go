package googlesqldatabaseinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlesqldatabaseinstance/internal"
)

type GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference interface {
	cdktf.ComplexObject
	AllocatedIpRange() *string
	SetAllocatedIpRange(val *string)
	AllocatedIpRangeInput() *string
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
	Datasource() *string
	SetDatasource(val *string)
	DatasourceInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleSqlDatabaseInstancePointInTimeRestoreContext
	SetInternalValue(val *GoogleSqlDatabaseInstancePointInTimeRestoreContext)
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
	ResetAllocatedIpRange()
	ResetPointInTime()
	ResetPreferredZone()
	ResetRegion()
	ResetTargetInstance()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference
type jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) AllocatedIpRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocatedIpRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) AllocatedIpRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocatedIpRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) Datasource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) DatasourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) InternalValue() *GoogleSqlDatabaseInstancePointInTimeRestoreContext {
	var returns *GoogleSqlDatabaseInstancePointInTimeRestoreContext
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) PointInTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pointInTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) PointInTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pointInTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) PreferredZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) PreferredZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) TargetInstance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) TargetInstanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference_Override(g GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference)SetAllocatedIpRange(val *string) {
	if err := j.validateSetAllocatedIpRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocatedIpRange",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference)SetDatasource(val *string) {
	if err := j.validateSetDatasourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasource",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference)SetInternalValue(val *GoogleSqlDatabaseInstancePointInTimeRestoreContext) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference)SetPointInTime(val *string) {
	if err := j.validateSetPointInTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pointInTime",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference)SetPreferredZone(val *string) {
	if err := j.validateSetPreferredZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredZone",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference)SetTargetInstance(val *string) {
	if err := j.validateSetTargetInstanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetInstance",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) ResetAllocatedIpRange() {
	_jsii_.InvokeVoid(
		g,
		"resetAllocatedIpRange",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) ResetPointInTime() {
	_jsii_.InvokeVoid(
		g,
		"resetPointInTime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) ResetPreferredZone() {
	_jsii_.InvokeVoid(
		g,
		"resetPreferredZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) ResetTargetInstance() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetInstance",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstancePointInTimeRestoreContextOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

