// executorsnetworking
package executorsnetworking

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/executorsnetworking/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/executorsnetworking/internal"
)

type Executorsnetworking interface {
	cdktf.TerraformModule
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IpCidrOutput() *string
	Nat() *bool
	SetNat(val *bool)
	NatIpOutput() *string
	NatMinPortsPerVm() *float64
	SetNatMinPortsPerVm(val *float64)
	NetworkIdOutput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Providers() *[]interface{}
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	// Experimental.
	SkipAssetCreationFromLocalModules() *bool
	// Experimental.
	Source() *string
	SubnetIdOutput() *string
	// Experimental.
	Version() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	AddProvider(provider interface{})
	// Experimental.
	GetString(output *string) *string
	// Experimental.
	InterpolationForOutput(moduleOutput *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Executorsnetworking
type jsiiProxy_Executorsnetworking struct {
	internal.Type__cdktfTerraformModule
}

func (j *jsiiProxy_Executorsnetworking) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsnetworking) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsnetworking) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsnetworking) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsnetworking) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsnetworking) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsnetworking) IpCidrOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipCidrOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsnetworking) Nat() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"nat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsnetworking) NatIpOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natIpOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsnetworking) NatMinPortsPerVm() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"natMinPortsPerVm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsnetworking) NetworkIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsnetworking) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsnetworking) Providers() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"providers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsnetworking) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsnetworking) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsnetworking) SkipAssetCreationFromLocalModules() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"skipAssetCreationFromLocalModules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsnetworking) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsnetworking) SubnetIdOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsnetworking) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func NewExecutorsnetworking(scope constructs.Construct, id *string, options *ExecutorsnetworkingOptions) Executorsnetworking {
	_init_.Initialize()

	if err := validateNewExecutorsnetworkingParameters(scope, id, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Executorsnetworking{}

	_jsii_.Create(
		"executorsnetworking.Executorsnetworking",
		[]interface{}{scope, id, options},
		&j,
	)

	return &j
}

func NewExecutorsnetworking_Override(e Executorsnetworking, scope constructs.Construct, id *string, options *ExecutorsnetworkingOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"executorsnetworking.Executorsnetworking",
		[]interface{}{scope, id, options},
		e,
	)
}

func (j *jsiiProxy_Executorsnetworking)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Executorsnetworking)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Executorsnetworking)SetNat(val *bool) {
	_jsii_.Set(
		j,
		"nat",
		val,
	)
}

func (j *jsiiProxy_Executorsnetworking)SetNatMinPortsPerVm(val *float64) {
	_jsii_.Set(
		j,
		"natMinPortsPerVm",
		val,
	)
}

func (j *jsiiProxy_Executorsnetworking)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Executorsnetworking_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExecutorsnetworking_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"executorsnetworking.Executorsnetworking",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Executorsnetworking) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_Executorsnetworking) AddProvider(provider interface{}) {
	if err := e.validateAddProviderParameters(provider); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addProvider",
		[]interface{}{provider},
	)
}

func (e *jsiiProxy_Executorsnetworking) GetString(output *string) *string {
	if err := e.validateGetStringParameters(output); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Executorsnetworking) InterpolationForOutput(moduleOutput *string) cdktf.IResolvable {
	if err := e.validateInterpolationForOutputParameters(moduleOutput); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForOutput",
		[]interface{}{moduleOutput},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Executorsnetworking) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_Executorsnetworking) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Executorsnetworking) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Executorsnetworking) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Executorsnetworking) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Executorsnetworking) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

