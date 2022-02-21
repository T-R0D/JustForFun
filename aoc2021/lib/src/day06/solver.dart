import 'package:aoc2021/src/solver/solver.dart' show Solver;

class Day06Solver implements Solver {
  static const _simulatedDays = 80;
  static const _extendedSimulatedDays = 256;
  static const _reproductionCycleDays = 6;
  static const _immatureFishReproductionCycleDays = 8;

  List<int> _lanternFishTimers = [];

  @override
  void consumeRawInput(String rawInput) {
    _lanternFishTimers =
        rawInput.split(',').map((listValue) => int.parse(listValue)).toList();
  }

  @override
  String solvePart1() {
    final totalFish =
        _simulateFishReproduction(_lanternFishTimers, _simulatedDays);

    return '$totalFish';
  }

  @override
  String solvePart2() {
    final totalFish =
        _simulateFishReproduction(_lanternFishTimers, _extendedSimulatedDays);

    return '$totalFish';
  }

  int _simulateFishReproduction(
      List<int> initialFishTimers, int daysToSimulate) {
    final fishCounts = List.filled(9, 0);

    for (var startingTimerValue in _lanternFishTimers) {
      fishCounts[startingTimerValue]++;
    }

    for (var day = 0; day < daysToSimulate; day++) {
      final reproducingFish = fishCounts[0];
      for (var i = 1; i < fishCounts.length; i++) {
        fishCounts[i - 1] = fishCounts[i];
      }
      fishCounts[_reproductionCycleDays] += reproducingFish;
      fishCounts[_immatureFishReproductionCycleDays] = reproducingFish;
    }

    var totalFish = 0;
    for (var count in fishCounts) {
      totalFish += count;
    }
    return totalFish;
  }
}
