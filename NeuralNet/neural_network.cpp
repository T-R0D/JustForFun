#ifndef _NEURAL_NETWORK_CPP_
#define _NEURAL_NETWORK_CPP_ 1

#include "neural_network.hpp"

#include <cassert>
#include <cmath>

NeuralNetwork::NeuralNetwork(
  const std::vector<unsigned>& topology,
  const double learning_rate,
  const double momentum)
    : learning_rate_(learning_rate),
      momentum_(momentum) {

  assert(topology.size() >= 3);

  NeuronLayer input_layer;
  unsigned input_dim = topology[0];
  for (unsigned j = 0; j < input_dim; ++j) {
    input_layer.push_back(Neuron(j));
  }
  network_layers_.push_back(input_layer);

  for (unsigned i = 1; i < topology.size(); ++i) {
    unsigned layer_size = topology[i];
    unsigned previous_layer_size = topology[i - 1];
    NeuronLayer new_layer;
    for (unsigned j = 0; j < layer_size; ++j) {
      new_layer.push_back(Neuron(
        j,
        previous_layer_size,
        [] (const double x) -> double {return tanh(x);},
        [] (const double x) -> double {return (1.0 - (x * x));}
        // [] (const double x) -> double {return x;},
        // [] (const double x) -> double {return 1.0;}
      ));
    }
    network_layers_.push_back(new_layer);
  }
}

void
NeuralNetwork::ApplyInputs(const std::vector<double>& inputs) {
  NeuronLayer& input_layer = network_layers_.front();
  unsigned input_dim = input_layer.size();

  assert(inputs.size() == input_dim);

  for (unsigned j = 0; j < input_dim; ++j) {
    input_layer[j].ApplyInputs({inputs[j]});
  }

  FeedForward();
}

std::vector<double>
NeuralNetwork::CollectOutputs() const {
  const NeuronLayer& output_layer = network_layers_.back();

  std::vector<double> outputs(output_layer.size());
  for (unsigned j = 0; j < output_layer.size(); ++j) {
    outputs[j] = output_layer[j].output();
  }

  return outputs;
}

void
NeuralNetwork::TrainOnInput(
  const std::vector<std::vector<double>>& input_set,
  const std::vector<std::vector<double>>& target_set) {

  assert(input_set.size() == target_set.size());
  for (unsigned i = 0; i < input_set.size(); ++i) {
    ApplyInputs(input_set[i]);
    BackPropogateError(target_set[i]);
  }
}

std::string
NeuralNetwork::toString() const {
  std::string to_string = "";
  char dummy[200];

  for (unsigned i = 0; i < network_layers_.size(); ++i) {
    sprintf(dummy, "Layer %u\n===============\n", i);
    to_string += dummy;

    for (const Neuron& neuron : network_layers_[i]) {
      to_string += neuron.toString();
    }

    to_string += "\n";
  }

  return to_string;
}

void
NeuralNetwork::FeedForward() {
  for (unsigned i = 1; i < network_layers_.size(); ++i) {
    NeuronLayer& previous_layer = network_layers_[i - 1];
    NeuronLayer& current_layer = network_layers_[i];

    std::vector<double> previous_layer_outputs(previous_layer.size());
    for (unsigned j = 0; j < previous_layer.size(); ++j) {
      previous_layer_outputs[j] = previous_layer[j].output();
    }

    for (Neuron& neuron : current_layer) {
      neuron.ApplyInputs(previous_layer_outputs);
    }
  }
}

void
NeuralNetwork::BackPropogateError(const std::vector<double>& targets) {
  NeuronLayer& output_layer = network_layers_.back();

  for (unsigned j = 0; j < output_layer.size(); ++j) {
    output_layer[j].ComputeOutputDelta(targets[j]);
  }

  for (unsigned i = network_layers_.size() - 2; i > 0; --i) {
    NeuronLayer& current_layer = network_layers_[i];
    NeuronLayer& subsequent_layer = network_layers_[i + 1];

    for (Neuron& neuron : current_layer) {
      neuron.ComputeDelta(subsequent_layer);
    }
  }

  for (unsigned i = 1; i < network_layers_.size(); ++i) {
    NeuronLayer& previous_layer = network_layers_[i - 1];
    NeuronLayer& current_layer = network_layers_[i];
    
    for (Neuron& neuron : current_layer) {
      neuron.UpdateInputWeights(
        previous_layer,
        learning_rate_,
        momentum_
      );
    }
  }
}

#endif //_NEURAL_NETWORK_CPP_