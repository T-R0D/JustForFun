#include <cassert>

#include "neural_network.hpp"

void testNeuronApplyInputs() {
  Neuron test_neuron(
    1,
    2,
    [] (const double x) -> double {return x;},
    [] (const double x) -> double {return 1.0;}
  );
  test_neuron.SetInputWeights({1.0, 1.0, 1.0});

  test_neuron.ApplyInputs({1.0, 1.0});

  assert(test_neuron.output() == 3.0);
}

int main() {
  testNeuronApplyInputs();
}