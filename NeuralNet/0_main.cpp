#include "neural_network.hpp"

#include <vector>

typedef struct {
  std::vector<double> inputs;
  std::vector<double> targets;
} TestCase;

std::vector<TestCase> test_cases = {
  {{0.0, 0.0}, {0.0}},
  {{0.0, 1.0}, {1.0}},
  {{1.0, 0.0}, {1.0}},
  {{1.0, 1.0}, {0.0}}
};

TestCase
generateXorTestCase();

bool
PerformRoundOfTraining(NeuralNetwork& test_network);

void
ViewNetworkAfterInputs();

int main(int argc, char** argv) {
  NeuralNetwork test_network(
    {2, 4, 1},
    0.2, // eta
    0.2  // alpha
  );

  // puts(test_network.toString().c_str());
  // test_network.TrainOnInput({{1.0, 0.0}}, {{1.0}});
  // puts(test_network.toString().c_str());

// puts(test_network.toString().c_str());
// char z;
// scanf("%c", &z);

  char y;
  while (true) {
    int correct = 0;
    int num_trainings = 100;
    for (unsigned i = 0; i < num_trainings; i++) {
      if (PerformRoundOfTraining(test_network)) {
        correct++;
      }
    }
    puts(test_network.toString().c_str());
    printf(
      "Percent correct: %i%%\n",
      (int) (correct * 100.0 / (double) num_trainings)
    );
    scanf("%c", &y);
  }

  return 0;
}

TestCase
generateXorTestCase() {
  unsigned input1 = rand() & 0x01;
  unsigned input2 = rand() & 0x01;
  unsigned result = input1 ^ input2;

  return {{(double) input1, (double) input2}, {(double) result}};
}

bool
PerformRoundOfTraining(NeuralNetwork& test_network) {
  for (TestCase& test_case : test_cases) {
    test_network.TrainOnInput({test_case.inputs}, {test_case.targets});
  }

  TestCase trial_case = generateXorTestCase();
  test_network.ApplyInputs({trial_case.inputs});
  std::vector<double> output = test_network.CollectOutputs();

// printf(
//   "inputs: %f %f\n"
//   "target: %f | %f :output\n",
//   trial_case.inputs[0], trial_case.inputs[1],
//   trial_case.targets[0], output[0]
// );

  bool result;
  double out = output[0];
  double target = trial_case.targets[0];
  if ((target == 1.0 && out >= 0.8) ||
      (target == 0.0 && out <= 0.2)) {
    result = true;
  } else {
    result = false;
  }

  return (output[0] >= 0.5 ? 1 : 0) == (unsigned) trial_case.targets[0];
}

void
ViewNetworkAfterInputs() {
  NeuralNetwork test_network(
    {2, 4, 1},
    0.2, // eta
    0.5  // alpha
  );

  puts("\n=================================== Before inputs:");
  puts(test_network.toString().c_str());
  test_network.ApplyInputs({0.0, 0.0});
  puts("\n=================================== 0, 0:");
  puts(test_network.toString().c_str());
  test_network.ApplyInputs({0.0, 1.0});
  puts("\n=================================== 0, 1:");
  puts(test_network.toString().c_str());
  test_network.ApplyInputs({1.0, 0.0});
  puts("\n=================================== 1, 0:");
  puts(test_network.toString().c_str());
  test_network.ApplyInputs({1.0, 1.0});
  puts("\n=================================== 1, 1:");
  puts(test_network.toString().c_str());
}