package pipeline

import (
	"net/http"
)
type RequestPipeline func(*ComponentContext)

var emptyPipeline RequestPipeline = func(*ComponentContext) {
	/*todo*/
}

func CreatePipeline(components ...MiddlewareComponent) RequestPipeline {
	f := emptyPipeline

	for i := len(components) - 1; i >= 0; i--{
		currentComponent := components[i]
		nextFunc := f
		f = func(context *ComponentContext){
			if (context.error == nil){
				currentComponent.PorcessRequest(context, nextFunc)
			}
		}
		currentComponent.init()
	}

	return f
}

func (pl RequestPipeline) PorcessRequest(req *http.Request,
		resp http.ResponseWriter) error {
	ctx := ComponentContext {
		Request: req,
		ResponseWriter: resp,
	}
	pl(&ctx)

	return ctx.error

}

