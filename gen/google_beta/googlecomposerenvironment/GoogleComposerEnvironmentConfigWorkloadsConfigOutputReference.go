package googlecomposerenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomposerenvironment/internal"
)

type GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference interface {
	cdktf.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComposerEnvironmentConfigWorkloadsConfig
	SetInternalValue(val *GoogleComposerEnvironmentConfigWorkloadsConfig)
	Scheduler() GoogleComposerEnvironmentConfigWorkloadsConfigSchedulerOutputReference
	SchedulerInput() *GoogleComposerEnvironmentConfigWorkloadsConfigScheduler
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Triggerer() GoogleComposerEnvironmentConfigWorkloadsConfigTriggererOutputReference
	TriggererInput() *GoogleComposerEnvironmentConfigWorkloadsConfigTriggerer
	WebServer() GoogleComposerEnvironmentConfigWorkloadsConfigWebServerOutputReference
	WebServerInput() *GoogleComposerEnvironmentConfigWorkloadsConfigWebServer
	Worker() GoogleComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference
	WorkerInput() *GoogleComposerEnvironmentConfigWorkloadsConfigWorker
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
	PutScheduler(value *GoogleComposerEnvironmentConfigWorkloadsConfigScheduler)
	PutTriggerer(value *GoogleComposerEnvironmentConfigWorkloadsConfigTriggerer)
	PutWebServer(value *GoogleComposerEnvironmentConfigWorkloadsConfigWebServer)
	PutWorker(value *GoogleComposerEnvironmentConfigWorkloadsConfigWorker)
	ResetScheduler()
	ResetTriggerer()
	ResetWebServer()
	ResetWorker()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference
type jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) InternalValue() *GoogleComposerEnvironmentConfigWorkloadsConfig {
	var returns *GoogleComposerEnvironmentConfigWorkloadsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) Scheduler() GoogleComposerEnvironmentConfigWorkloadsConfigSchedulerOutputReference {
	var returns GoogleComposerEnvironmentConfigWorkloadsConfigSchedulerOutputReference
	_jsii_.Get(
		j,
		"scheduler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) SchedulerInput() *GoogleComposerEnvironmentConfigWorkloadsConfigScheduler {
	var returns *GoogleComposerEnvironmentConfigWorkloadsConfigScheduler
	_jsii_.Get(
		j,
		"schedulerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) Triggerer() GoogleComposerEnvironmentConfigWorkloadsConfigTriggererOutputReference {
	var returns GoogleComposerEnvironmentConfigWorkloadsConfigTriggererOutputReference
	_jsii_.Get(
		j,
		"triggerer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) TriggererInput() *GoogleComposerEnvironmentConfigWorkloadsConfigTriggerer {
	var returns *GoogleComposerEnvironmentConfigWorkloadsConfigTriggerer
	_jsii_.Get(
		j,
		"triggererInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) WebServer() GoogleComposerEnvironmentConfigWorkloadsConfigWebServerOutputReference {
	var returns GoogleComposerEnvironmentConfigWorkloadsConfigWebServerOutputReference
	_jsii_.Get(
		j,
		"webServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) WebServerInput() *GoogleComposerEnvironmentConfigWorkloadsConfigWebServer {
	var returns *GoogleComposerEnvironmentConfigWorkloadsConfigWebServer
	_jsii_.Get(
		j,
		"webServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) Worker() GoogleComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference {
	var returns GoogleComposerEnvironmentConfigWorkloadsConfigWorkerOutputReference
	_jsii_.Get(
		j,
		"worker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) WorkerInput() *GoogleComposerEnvironmentConfigWorkloadsConfigWorker {
	var returns *GoogleComposerEnvironmentConfigWorkloadsConfigWorker
	_jsii_.Get(
		j,
		"workerInput",
		&returns,
	)
	return returns
}


func NewGoogleComposerEnvironmentConfigWorkloadsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComposerEnvironmentConfigWorkloadsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googleComposerEnvironment.GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComposerEnvironmentConfigWorkloadsConfigOutputReference_Override(g GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComposerEnvironment.GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference)SetInternalValue(val *GoogleComposerEnvironmentConfigWorkloadsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) PutScheduler(value *GoogleComposerEnvironmentConfigWorkloadsConfigScheduler) {
	if err := g.validatePutSchedulerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScheduler",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) PutTriggerer(value *GoogleComposerEnvironmentConfigWorkloadsConfigTriggerer) {
	if err := g.validatePutTriggererParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTriggerer",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) PutWebServer(value *GoogleComposerEnvironmentConfigWorkloadsConfigWebServer) {
	if err := g.validatePutWebServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWebServer",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) PutWorker(value *GoogleComposerEnvironmentConfigWorkloadsConfigWorker) {
	if err := g.validatePutWorkerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWorker",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) ResetScheduler() {
	_jsii_.InvokeVoid(
		g,
		"resetScheduler",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) ResetTriggerer() {
	_jsii_.InvokeVoid(
		g,
		"resetTriggerer",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) ResetWebServer() {
	_jsii_.InvokeVoid(
		g,
		"resetWebServer",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) ResetWorker() {
	_jsii_.InvokeVoid(
		g,
		"resetWorker",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigWorkloadsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

