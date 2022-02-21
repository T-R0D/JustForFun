import 'package:aoc2021/src/solver/solver.dart' show Solver;

class Day07Solver implements Solver {
  List<int> _initialPositions = [];

  @override
  void consumeRawInput(String rawInput) {
    // rawInput = '16,1,2,0,4,2,7,1,2,14';
    _initialPositions = rawInput.split(',').map((x) => int.parse(x)).toList();
  }

  @override
  String solvePart1() {
    final positionCounts = <int, int>{};
    for (var initialPosition in _initialPositions) {
      positionCounts.update(initialPosition, (value) => value + 1,
          ifAbsent: () => 1);
    }

    // TODO: Clean up these hacky 99999s...
    var firstPosition = 999999;
    var lastPosition = 0;
    for (var initialPostion in positionCounts.keys) {
      if (initialPostion < firstPosition) {
        firstPosition = initialPostion;
      }
      if (initialPostion > lastPosition) {
        lastPosition = initialPostion;
      }
    }

    var bestFuelCost = lastPosition * _initialPositions.length;
    for (var i = firstPosition; i <= lastPosition; i++) {
      var fuelCost = 0;
      for (var entry in positionCounts.entries) {
        fuelCost += (i - entry.key).abs() * entry.value;
      }

      if (fuelCost < bestFuelCost) {
        bestFuelCost = fuelCost;
      }
    }

    return '$bestFuelCost';
  }

  @override
  String solvePart2() {
    final positionCounts = <int, int>{};
    for (var initialPosition in _initialPositions) {
      positionCounts.update(initialPosition, (value) => value + 1,
          ifAbsent: () => 1);
    }

    // TODO: Clean up these hacky 99999s...
    var firstPosition = 999999;
    var lastPosition = 0;
    for (var initialPostion in positionCounts.keys) {
      if (initialPostion < firstPosition) {
        firstPosition = initialPostion;
      }
      if (initialPostion > lastPosition) {
        lastPosition = initialPostion;
      }
    }

    // TODO: Clean up these hacky 99999s...
    var bestFuelCost = lastPosition * _initialPositions.length * 99999;
    for (var i = firstPosition; i <= lastPosition; i++) {
      var fuelCost = 0;
      for (var entry in positionCounts.entries) {
        var distanceToCandidatePosition = (i - entry.key).abs();
        fuelCost +=
            _sumNumbersZeroToN(distanceToCandidatePosition) * entry.value;
      }

      if (fuelCost < bestFuelCost) {
        bestFuelCost = fuelCost;
      }
    }

    return '$bestFuelCost';
  }

  int _sumNumbersZeroToN(int n) {
    if (n < 0) {
      throw StateError('Supply only numbers 0+');
    }
    if (n & 1 == 0) {
      return (n + 1) * (n ~/ 2);
    }
    return (n * (n ~/ 2)) + n;
  }
}
