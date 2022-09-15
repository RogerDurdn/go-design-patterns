package main

import "fmt"

func main() {
	n1, n2 := NewNeuron(), NewNeuron()
	l1, l2 := NewNeuronLayer(2), NewNeuronLayer(5)

	// The Connect function can handle both, the origin struct and a struct with collections of that origin
	Connect(n1, n2)
	Connect(n1, l1)
	Connect(l2, l1)
	Connect(l1, l2)

	fmt.Println(n1)
	fmt.Println(l1)
	fmt.Println(l2)
}

// NeuronInterface provides a unified interface that is used to relate two different objects with something in common
type NeuronInterface interface {
	Collect() []*Neuron
}

func Connect(left, right NeuronInterface) {
	for _, nl := range left.Collect() {
		for _, nr := range right.Collect() {
			nl.ConnectTo(nr)
		}
	}
}

type Neuron struct {
	In, Out []*Neuron
}

func NewNeuron() *Neuron {
	return &Neuron{}
}

func (n *Neuron) ConnectTo(other *Neuron) {
	n.Out = append(n.Out, other)
	other.In = append(other.In, n)
}

func (n *Neuron) Collect() []*Neuron {
	return []*Neuron{n}
}

type NeuronLayer struct {
	Neurons []Neuron
}

func (l *NeuronLayer) Collect() []*Neuron {
	collection := make([]*Neuron, 0)
	for i := range l.Neurons {
		collection = append(collection, &l.Neurons[i])
	}
	return collection
}

func NewNeuronLayer(count int) *NeuronLayer {
	return &NeuronLayer{make([]Neuron, count)}
}
