import 'dart:math';

import 'package:aoc2021/src/solver/solver.dart' show Solver;
import 'package:aoc2021/src/space/point.dart';

class Day17Solver implements Solver {
  int _targetX0 = 0;
  int _targetX1 = 0;
  int _targetY0 = 0;
  int _targetY1 = 0;

  @override
  void consumeRawInput(String rawInput) {
    final values = rawInput
        .replaceAll('target area: x=', '')
        .replaceAll(' y=', '')
        .replaceAll('..', ',')
        .split(',')
        .map(int.parse)
        .toList();

    _targetX0 = values[0];
    _targetX1 = values[1];
    // The y values are "reversed" in the input to go from least to greatest.
    // "Reverse" them here to have y0 be the first y bound checked.
    _targetY0 = values[3];
    _targetY1 = values[2];
  }

  @override
  String solvePart1() {
    final candidateXVelocities =
        _findCandidateXVelocities(_targetX0, _targetX1);
    final candidates = _findCandidateVelocityPairs(
        candidateXVelocities, _targetX0, _targetX1, _targetY0, _targetY1);
    final highestInitialYVelocity =
        candidates.keys.map((pair) => pair.y).reduce((a, b) => max(a, b));
    final apogee = _findVerticalDisplacement(
        highestInitialYVelocity, highestInitialYVelocity + 1);
    return apogee.toString();
  }

  @override
  String solvePart2() {
    final candidateXVelocities =
        _findCandidateXVelocities(_targetX0, _targetX1);
    final candidates = _findCandidateVelocityPairs(
        candidateXVelocities, _targetX0, _targetX1, _targetY0, _targetY1);
    return candidates.keys.length.toString();
  }
}

Map<int, int> _findCandidateXVelocities(int targetX0, int targetX1) {
  final canidateXVelocities = <int, int>{};

  final maxFeasibleInitialVelocity = targetX1;
  var minFeasibleInitialVelocity = 0;
  while (true) {
    final t = minFeasibleInitialVelocity + 1;
    if (_findHorizontalDisplacement(minFeasibleInitialVelocity, t) >=
        targetX0) {
      break;
    }
    minFeasibleInitialVelocity++;
  }

  X_INCREMENT:
  for (var x = minFeasibleInitialVelocity;
      x <= maxFeasibleInitialVelocity;
      x++) {
    for (var t = 1; true; t++) {
      final displacement = _findHorizontalDisplacement(x, t);
      if (targetX0 <= displacement && displacement <= targetX1) {
        canidateXVelocities[x] = t;
        continue X_INCREMENT;
      }
      if (targetX1 < displacement) {
        continue X_INCREMENT;
      }
    }
  }

  return canidateXVelocities;
}

Map<Point, List<int>> _findCandidateVelocityPairs(
    Map<int, int> candidateXVelocities,
    int targetX0,
    int targetX1,
    int targetY0,
    int targetY1) {
  final candidateVelocityPairs = <Point, List<int>>{};

  final minFeasibleInitialVelocity = targetY1;
  final maxFeasibleInitialVelocity = -targetY1 - 1;

  for (var y = minFeasibleInitialVelocity;
      y <= maxFeasibleInitialVelocity;
      y++) {
    X_INCREMENT:
    for (var entry in candidateXVelocities.entries) {
      final x = entry.key;
      final startingT = entry.value;
      final velocityPair = Point(x, y);
      for (var t = startingT; true; t++) {
        final xDisplacement = _findHorizontalDisplacement(x, t);
        final yDisplacement = _findVerticalDisplacement(y, t);
        if ((targetX0 <= xDisplacement && xDisplacement <= targetX1) &&
            (targetY0 >= yDisplacement && yDisplacement >= targetY1)) {
          candidateVelocityPairs.update(
              velocityPair, (timeSteps) => [...timeSteps, t],
              ifAbsent: () => [t]);
        }
        if (xDisplacement > targetX1 || yDisplacement < targetY1) {
          continue X_INCREMENT;
        }
      }
    }
  }

  return candidateVelocityPairs;
}

int _findHorizontalDisplacement(int initialVelocity, int timeSteps) {
  if (initialVelocity < 0) {
    throw Error();
  }
  if (timeSteps < 0) {
    throw Error();
  }

  var t = min(timeSteps, initialVelocity + 1);

  return findDisplacement(initialVelocity, t);
}

int _findVerticalDisplacement(int initialVelocity, int timeSteps) {
  if (timeSteps < 0) {
    throw Error();
  }

  return findDisplacement(initialVelocity, timeSteps);
}

int findDisplacement(int v, int t) {
  if (t & 0x01 == 0) {
    final numerator = t * ((2 * v) - t + 1);
    return numerator ~/ 2;
  }

  final oddStepDisplacement = v - (t ~/ 2);
  final numerator = (t - 1) * ((2 * v) - t + 1);
  return (numerator ~/ 2) + oddStepDisplacement;
}
