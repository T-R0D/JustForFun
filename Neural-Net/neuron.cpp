#ifndef _NEURON_CPP_
#define _NEURON_CPP_ 1

#include "neuron.hpp"

#include <cassert>

Neuron::Neuron(
  const unsigned index_in_layer)
    : index_in_layer_(index_in_layer),
      input_weights_(1),
      activation_function_([] (const double x) -> double {return x;}),
      activation_function_derivative_(
        [] (const double x) -> double {return 1.0;}) {

  // initialize the only input weight
  input_weights_[0] = 1.0;
}

Neuron::Neuron(
  const unsigned index_in_layer,
  const unsigned num_inputs,
  const std::function<double(double)>& activation_function,
  const std::function<double(double)>& activation_function_derivative)
    : index_in_layer_(index_in_layer),
      input_weights_(num_inputs + 1), // for a bias neuron
      previous_weight_adjustments_(num_inputs + 1),
      activation_function_(activation_function),
      activation_function_derivative_(activation_function_derivative) {

  // the only initialization necessary is random weight initialization
  for (double& input_weight : input_weights_) {
    input_weight = GenerateRandomWeight();
  }
}

unsigned
Neuron::index() const {
  return index_in_layer_;
}

double
Neuron::input_sum() const {
  return input_sum_;
}

double
Neuron::output() const {
  return output_;
}

double
Neuron::delta() const {
  return delta_;
}

void
Neuron::SetInputWeights(const std::vector<double>& weights) {
  assert(weights.size() == input_weights_.size());

  for (unsigned i = 0; i < weights.size(); ++i) {
    input_weights_[i] = weights[i];
  }
}

std::string
Neuron::toString() const {
  char dummy[200];
  std::string to_string = "";

  to_string += "Input Weights:";
  for (double weight : input_weights_) {
    sprintf(dummy, " %f", weight);
    to_string += (dummy);
  }
  to_string += "\n";

  sprintf(dummy, "Output:        %f\n", output_);
  to_string += (dummy);

  return to_string;
}

void
Neuron::ApplyInputs(const std::vector<double>& inputs) {
  if (input_weights_.size() > 1) { // a hack
    // if the neuron is not an input one, do the test
    assert(inputs.size() == input_weights_.size() - 1);
  }

  input_sum_ = 0.0;
  for (unsigned i = 0; i < inputs.size(); ++i) {
    input_sum_ += (inputs[i] * input_weights_[i]);
  }

  // add the bias term (1 * w_{i,j})
  input_sum_ += input_weights_.back();

  output_ = activation_function_(input_sum_);
}

void
Neuron::ComputeDelta(const NeuronLayer& subsequent_layer) {
  double weighted_delta_sum = 0.0;

  for (const Neuron& neuron : subsequent_layer) {
    weighted_delta_sum +=
      neuron.input_weights_[index_in_layer_] *
      neuron.delta();
  }

  // pretty sure a bias term is not needed here

  delta_ =
    activation_function_derivative_(output_) *  // input_sum_) *
    weighted_delta_sum;
}

void
Neuron::ComputeOutputDelta(const double target) {
  delta_ =
    activation_function_derivative_(output_) * // input_sum_) *
    (target - output_);
}

void
Neuron::UpdateInputWeights(
  const NeuronLayer& previous_layer,
  const double eta,
  const double alpha) {

  assert(previous_layer.size() == input_weights_.size() - 1);

  double weight_adjustment = 0.0;

  for (unsigned i = 0; i < previous_layer.size(); ++i) {
    weight_adjustment =
      (eta * previous_layer[i].output() * delta_) +
      (alpha * previous_weight_adjustments_[i]);

    input_weights_[i] += weight_adjustment;
    previous_weight_adjustments_[i] = weight_adjustment;
  }

  // the bias weight
  weight_adjustment =
    (eta * delta_) +
    (alpha * previous_weight_adjustments_.back());
  input_weights_.back() += weight_adjustment;
  previous_weight_adjustments_.back() = weight_adjustment;
}

double
Neuron::GenerateRandomWeight() {
  return 0.3 + ((double) (rand() % 5) / 10.0); // simple, easy on eyes for
                                               // debugging 
  // It is very necessary to have weights that differ so that back-propagation
  // can have somewhere to assign error to
}

#endif //_NEURON_CPP_