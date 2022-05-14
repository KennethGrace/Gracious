package util

import (
	"github.com/Art-of-the-Living/gracious"
)

// A Sensor is any structure which can produce a slice of base.QualitativeSignal values. The Sensor interface is
// implemented by the FunctionalSensor
type Sensor interface {
	Evoke() gracious.QualitativeSignal
}

// A FunctionalSensor produces phenomenal experience as direct sensory
// observations. A FunctionalSensor component processes incoming raw-data into
// distributed signals that the system can begin to work with. Unlike other
// evocation a FunctionalSensor component doesn't have any QualitativeSignal
// inputs, instead Sensors produces a QualitativeSignal based on the output of an
// externally implemented function. This function can be set using SetProcessor
// and should be set before any calls to Evoke. Otherwise, Evoke will return an
// empty signal.
type FunctionalSensor struct {
	Id        string
	processor func() gracious.QualitativeSignal
}

// NewFunctionalSensor creates a new FunctionalSensor
func NewFunctionalSensor(name string) *FunctionalSensor {
	fs := FunctionalSensor{Id: "FunctionalSensor:" + name}
	return &fs
}

// GetName returns the ID of this FunctionalSensor
func (n *FunctionalSensor) GetName() string {
	return n.Id
}

// SetProcessor sets the function that should be run when this FunctionalSensor component is evoked.
func (n *FunctionalSensor) SetProcessor(processor func() gracious.QualitativeSignal) {
	n.processor = processor
}

// Evoke returns the result of the specified processor function. If no processor is defined, then an empty slice of
// QualitativeSignal values will be returned.
func (n *FunctionalSensor) Evoke() gracious.QualitativeSignal {
	if n.processor != nil {
		return n.processor()
	} else {
		return gracious.NewQualitativeSignal(n.Id)
	}
}
