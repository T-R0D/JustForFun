import 'package:aoc2021/src/solver/solver.dart';

class Day01Solver implements Solver {
  List<int> input = [];

  @override
  void consumeRawInput(String rawInput) {
    input =
        rawInput.split('\n').map((depthStr) => int.parse(depthStr)).toList();
  }

  @override
  String solvePart1() {
    var increases = 0;
    var recentReading = input[0];
    for (var depthReading in input.sublist(1)) {
      if (depthReading > recentReading) {
        increases += 1;
      }
      recentReading = depthReading;
    }

    return '$increases';
  }

  @override
  String solvePart2() {
    var increases = 0;

    var currentWindowSum = 0;
    for (var i = 0; i < 3; i += 1) {
      currentWindowSum += input[i];
    }

    for (var i = 0; i < input.length - 3; i += 1) {
      final oldestReading = input[i];
      final newReading = input[i + 3];
      final newWindowSum = currentWindowSum - oldestReading + newReading;

      if (newWindowSum > currentWindowSum) {
        increases += 1;
      }
      
      currentWindowSum = newWindowSum;
    }

    return '$increases';
  }
}
