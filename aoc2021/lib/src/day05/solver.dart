import 'package:aoc2021/src/solver/solver.dart' show Solver;
import 'package:aoc2021/src/space/line.dart';

class Day05Solver implements Solver {
  List<Line45Degree> _ventLines = [];

  @override
  void consumeRawInput(String rawInput) {
    final List<Line45Degree> ventLines = [];
    final inputLines = rawInput.split('\n');
    for (var inputLine in inputLines) {
      final inputLineParts = inputLine.split(' -> ');
      final startParts = inputLineParts[0].split(',');
      final endParts = inputLineParts[1].split(',');
      final x1 = int.parse(startParts[0]);
      final y1 = int.parse(startParts[1]);
      final x2 = int.parse(endParts[0]);
      final y2 = int.parse(endParts[1]);
      ventLines.add(Line45Degree(x1, y1, x2, y2));
    }
    _ventLines = ventLines;
  }

  @override
  String solvePart1() {
    final coveredPoints = <String, int>{};
    List<Line45Degree> axisParallelLines = _ventLines
        .where((line) => line.isHorizontal || line.isVertical)
        .toList();

    for (var line in axisParallelLines) {
      final pointsCoveredByLine = line.getCoveredPoints();
      for (var point in pointsCoveredByLine) {
        coveredPoints.update(point.toString(), (value) => value + 1,
            ifAbsent: () => 1);
      }
    }

    coveredPoints.removeWhere((_, value) => value < 2);

    return '${coveredPoints.length}';
  }

  @override
  String solvePart2() {
    final coveredPoints = <String, int>{};
    for (var line in _ventLines) {
      final pointsCoveredByLine = line.getCoveredPoints();
      for (var point in pointsCoveredByLine) {
        coveredPoints.update(point.toString(), (value) => value + 1,
            ifAbsent: () => 1);
      }
    }

    coveredPoints.removeWhere((_, value) => value < 2);

    return '${coveredPoints.length}';
  }
}
