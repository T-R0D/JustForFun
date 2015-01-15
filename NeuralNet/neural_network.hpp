#ifndef _NEURAL_NETWORK_HPP_
#define _NEURAL_NETWORK_HPP_ 1

#include "neuron.hpp"

class NeuralNetwork {
 public:
  typedef std::vector<Neuron> NeuronLayer;

  /**
   * Basic constructor.
   */
  NeuralNetwork(
    const std::vector<unsigned>& topology,
    const double learning_rate,
    const double momentum);

  /**
   * Apply some inputs to the neural network
   */
  void
  ApplyInputs(const std::vector<double>& inputs);

  /**
   * Returns the current outputs for the network.
   */
  std::vector<double>
  CollectOutputs() const;

  /**
   *
   */
  void
  TrainOnInput(
    const std::vector<std::vector<double>>& input_set,
    const std::vector<std::vector<double>>& target_set);

  /**
   *
   */
  std::string
  toString() const; 

 private:
  /**
   * Propogates some inputs across the network to obtain an output.
   */
  void
  FeedForward();


  /**
   * Back-propogates error across the network
   */
  void
  BackPropogateError(const std::vector<double>& targets);

 private:
  std::vector<NeuronLayer> network_layers_;
  double learning_rate_; // aka: eta
  double momentum_; // aka: alpha
};

#endif //_NEURAL_NETWORK_HPP_