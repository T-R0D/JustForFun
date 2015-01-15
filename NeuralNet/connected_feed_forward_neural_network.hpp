#ifndef _CONNECTED_FEED_FORWARD_NEURAL_NETWORK_HPP_
#define _CONNECTED_FEED_FORWARD_NEURAL_NETWORK_HPP_ 1

#include "neuron.hpp"

class ConnectedFeedForwardNeuralNetwork {
 public:
  typedef std::vector<Neuron> NeuronLayer; // very convenient

  /**
   * Basic constructor.
   * Topology should be the sizes of each layer, in order, with the first
   * layer being the dimensionality of an input to the network and the last
   * being the dimensionality of the output of the network. There must be
   * at least 1 hidden layer, and topology sizes should not include bias
   * neurons (they are added automatically).
   */
  ConnectedFeedForwardNeuralNetwork(
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
   * Applies each vector of inputs in the given set and then back-propagates
   * the errors relative to each target vector. The sets input_set and
   * target_set must be the same size, and each example should be of
   * appropriate dimensions.
   */
  void
  TrainOnInput(
    const std::vector<std::vector<double>>& input_set,
    const std::vector<std::vector<double>>& target_set);

  /**
   * Produces a human readable string that contains the weights and outputs of
   * all neurons in the network.
   */
  std::string
  toString() const; 

 private:
  /**
   * Propagates some inputs across the network to obtain an output.
   */
  void
  FeedForward();


  /**
   * Back-propagates error across the network.
   */
  void
  BackPropogateError(const std::vector<double>& targets);

 private:
  std::vector<NeuronLayer> network_layers_;
  double learning_rate_; // aka: eta
  double momentum_; // aka: alpha
};

#endif //_CONNECTED_FEED_FORWARD_NEURAL_NETWORK_HPP_