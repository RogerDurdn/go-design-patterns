package main

import "fmt"

/*
In this scenario we focus on force the client
to use the builder, this approach makes necessary
to one function a parameter which is the same builder
with this the client should use the builder in order
to execute and action
*/

func main() {
	// this is the base use case
	SendEmail(func(b *EmailBuilder) {
		b.To("some@mail.com").
			From("tom@mail.com").
			Subject("hello").
			Body("dam")
	})
	// can be used on this way too, but is not fluent
	eb := EmailBuilder{}
	eb.To("some").From("todo")
	SendEmail(func(builder *EmailBuilder) {
		*builder = eb // we have to copy values between objects
	})
}

/*
A normal builder but in this case the type is private
*/
type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	b.email.to = to
	return b
}
func (b *EmailBuilder) From(from string) *EmailBuilder {
	b.email.from = from
	return b
}
func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subject = subject
	return b
}
func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}

/*
This is the parameter part, here we are creating a function and force the client
to use another function that requires a builder to send and email
*/
// sendMailImp is private
func sendMailImp(email *email) {
	fmt.Println(email)
}

// build is a private func type
type build func(*EmailBuilder)

// SendEmail expose the functionality but only through a build private function
func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendMailImp(&builder.email)
}
