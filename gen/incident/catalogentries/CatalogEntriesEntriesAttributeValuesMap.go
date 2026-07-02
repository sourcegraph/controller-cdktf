package catalogentries

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/incident/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/incident/catalogentries/internal"
)

type CatalogEntriesEntriesAttributeValuesMap interface {
	cdktf.ComplexMap
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	Get(key *string) CatalogEntriesEntriesAttributeValuesOutputReference
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

// The jsii proxy struct for CatalogEntriesEntriesAttributeValuesMap
type jsiiProxy_CatalogEntriesEntriesAttributeValuesMap struct {
	internal.Type__cdktfComplexMap
}

func (j *jsiiProxy_CatalogEntriesEntriesAttributeValuesMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogEntriesEntriesAttributeValuesMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogEntriesEntriesAttributeValuesMap) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogEntriesEntriesAttributeValuesMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogEntriesEntriesAttributeValuesMap) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCatalogEntriesEntriesAttributeValuesMap(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CatalogEntriesEntriesAttributeValuesMap {
	_init_.Initialize()

	if err := validateNewCatalogEntriesEntriesAttributeValuesMapParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CatalogEntriesEntriesAttributeValuesMap{}

	_jsii_.Create(
		"@cdktf/provider-incident.catalogEntries.CatalogEntriesEntriesAttributeValuesMap",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCatalogEntriesEntriesAttributeValuesMap_Override(c CatalogEntriesEntriesAttributeValuesMap, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-incident.catalogEntries.CatalogEntriesEntriesAttributeValuesMap",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CatalogEntriesEntriesAttributeValuesMap) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CatalogEntriesEntriesAttributeValuesMap) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CatalogEntriesEntriesAttributeValuesMap) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CatalogEntriesEntriesAttributeValuesMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogEntriesEntriesAttributeValuesMap) Get(key *string) CatalogEntriesEntriesAttributeValuesOutputReference {
	if err := c.validateGetParameters(key); err != nil {
		panic(err)
	}
	var returns CatalogEntriesEntriesAttributeValuesOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]any{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogEntriesEntriesAttributeValuesMap) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogEntriesEntriesAttributeValuesMap) Resolve(_context cdktf.IResolveContext) any {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		c,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogEntriesEntriesAttributeValuesMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
