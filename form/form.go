// Package form implements a builder for HTML forms utilizing Angular Material Design. The forms are constructed to
// follow any conventions used in the Universal Data Tools UI application.
package form

import (
	"github.com/jhallat/udt-devtools/input"
)

type Type int

const (
	Template Type = iota
	Reactive
)


type Builder interface {
	appendInput(input input.Input)
	Build() string
}


type ReactiveFormBuilder struct {
	inputs []input.Input
}

type TemplateFormBuilder struct {
	inputs []input.Input
}


func NewFormBuilder(formType Type) Builder {
	if formType == Reactive {
		return &ReactiveFormBuilder{
			inputs: []input.Input{},
		}
	}
	return &TemplateFormBuilder{
		inputs: []input.Input{},
	}
}

func (form *ReactiveFormBuilder) appendInput(input input.Input) {
	form.inputs = append(form.inputs, input)
}

func (form *ReactiveFormBuilder) Build() string {
	return ""
}

func (form *TemplateFormBuilder) appendInput(input input.Input) {
	form.inputs = append(form.inputs, input)
}

func (form *TemplateFormBuilder) Build() string {

	//<form class="example-form">
	//<mat-form-field class="example-full-width">
	//<mat-label>Favorite food</mat-label>
	//<input matInput placeholder="Ex. Pizza" value="Sushi">
	//</mat-form-field>
    //
	//<mat-form-field class="example-full-width">
	//<mat-label>Leave a comment</mat-label>
	//<textarea matInput placeholder="Ex. It makes me feel..."></textarea>
	//</mat-form-field>
	//</form>

	return ""
}