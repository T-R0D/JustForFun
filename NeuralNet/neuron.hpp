#ifndef _NEURON_HPP_
#define _NEURON_HPP_ 1

#include <functional>
#include <vector>
#include <string>

typedef struct {
  double weight;
  double delta_weight;
} InputConnection;

class Neuron {
 public:
  typedef std::vector<Neuron> NeuronLayer;

  /**
   * A basic constructor.
   */
  Neuron(
    const unsigned index_in_layer,
    const unsigned num_inputs,
    const std::function<double(double)>& activation_function,
    const std::function<double(double)>& activation_function_derivative);

  /**
   * A constructor intended for input neurons only.
   */
  Neuron(const unsigned index_in_layer);

  /**
   *
   */
  unsigned
  index() const;

  /**
   * Returns the most recently computed dot product of the inputs and
   * their weights.
   */
  double
  input_sum() const;

  /**
   * Returns the most recently computed output.
   */
  double
  output() const;

  /**
   * Returns the most recently computed delta
   * (delta_{i} = g'(x) * sum_{j}{w_{i,j} * delta_{j}})
   */
  double
  delta() const;

  /**
   *
   */
  void
  SetInputWeights(const std::vector<double>& weights);

  /**
   *
   */
  std::string
  toString() const;

  /**
   * Applies the given inputs and produces the output.
   */
  void
  ApplyInputs(const std::vector<double>& inputs);

  /**
   *
   */
  void
  ComputeDelta(const NeuronLayer& subsequent_layer);

  /**
   *
   */
  void
  ComputeOutputDelta(const double target);

  /**
   *
   */
  void
  UpdateInputWeights(
    const NeuronLayer& previous_layer,
    const double eta,
    const double alpha);

  /**
   *
   */
  static double
  GenerateRandomWeight();

 private:
  unsigned index_in_layer_;
  double input_sum_;
  double output_;
  double delta_;
  std::vector<double> input_weights_;
  std::vector<double> previous_weight_adjustments_;
  std::function<double(double)> activation_function_;
  std::function<double(double)> activation_function_derivative_;
};

#endif //_NEURON_HPP_