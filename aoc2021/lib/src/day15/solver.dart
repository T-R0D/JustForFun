import 'package:aoc2021/src/solver/solver.dart' show Solver;
import 'package:aoc2021/src/space/point.dart' show Point;
import 'package:collection/collection.dart' show PriorityQueue;

const _maxInt = 9007199254740991; // 2^53 - 1 ()in case we compile to JS).

class Day15Solver implements Solver {
  List<List<int>> _riskGrid = [];

  @override
  void consumeRawInput(String rawInput) {
    _riskGrid = rawInput
        .split('\n')
        .map((row) => row.split('').map(int.parse).toList())
        .toList();
  }

  @override
  String solvePart1() {
    final riskGraph = riskGridToGraph(_riskGrid);
    final src = Point(0, 0);
    final dest = Point(_riskGrid.length - 1, _riskGrid[0].length - 1);
    final minTotalRisk = findRiskOfLowestRiskPath(riskGraph, src, dest);
    return minTotalRisk.toString();
  }

  @override
  String solvePart2() {
    final riskGraph = riskGridTo5xGraph(_riskGrid);
    final src = Point(0, 0);
    final dest =
        Point((_riskGrid.length * 5) - 1, (_riskGrid[0].length * 5) - 1);
    final minTotalRisk = findRiskOfLowestRiskPath(riskGraph, src, dest);
    return minTotalRisk.toString();
  }
}

class _PointRiskPair {
  final Point point;
  final int risk;

  _PointRiskPair(this.point, this.risk);
}

Map<Point, Map<Point, int>> riskGridToGraph(List<List<int>> grid) {
  final graph = <Point, Map<Point, int>>{};
  for (var i = 0; i < grid.length; i++) {
    for (var j = 0; j < grid[i].length; j++) {
      final successor = Point(i, j);
      final riskLevel = grid[i][j];

      for (var iOffset = -1; iOffset <= 1; iOffset++) {
        final predecessorI = i + iOffset;

        if (predecessorI < 0 || grid.length <= predecessorI) {
          continue;
        }

        for (var jOffset = -1; jOffset <= 1; jOffset++) {
          final predecessorJ = j + jOffset;

          if (predecessorJ < 0 || grid[i].length <= predecessorJ) {
            continue;
          }

          if (predecessorI == i && predecessorJ == j) {
            continue;
          }

          if (iOffset.abs() == 1 && jOffset.abs() == 1) {
            continue;
          }

          final predecessor = Point(predecessorI, predecessorJ);
          graph.update(predecessor, (edges) {
            edges[successor] = riskLevel;
            return edges;
          }, ifAbsent: () => {successor: riskLevel});
        }
      }
    }
  }
  return graph;
}

// Can probably speed this up by computing the additioinalRisk by computing
// it on the fly, rather than precomputing a graph.
int findRiskOfLowestRiskPath(
    Map<Point, Map<Point, int>> riskGraph, Point src, Point dest) {
  final cumulativeRiskToPoint = <Point, int>{src: 0};
  final predecessors = <Point, Point>{};
  final queue =
      PriorityQueue<_PointRiskPair>((a, b) => a.risk.compareTo(b.risk));

  for (var point in riskGraph.keys) {
    if (point != src) {
      cumulativeRiskToPoint[point] = _maxInt;
    }
  }

  queue.add(_PointRiskPair(src, 0));

  while (queue.isNotEmpty) {
    final pointRiskPair = queue.removeFirst();
    final point = pointRiskPair.point;
    final cumulativeRisk = pointRiskPair.risk;

    if (cumulativeRisk < (cumulativeRiskToPoint[point] ?? _maxInt)) {
      continue;
    } else if (point == dest) {
      break;
    }

    final neighbors = riskGraph[point];
    if (neighbors == null) {
      continue;
    }
    for (var entry in neighbors.entries) {
      final neighbor = entry.key;
      final additionalRisk = entry.value;
      final candidateCumulativeRisk = cumulativeRisk + additionalRisk;
      if (candidateCumulativeRisk <
          (cumulativeRiskToPoint[neighbor] ?? _maxInt)) {
        cumulativeRiskToPoint[neighbor] = candidateCumulativeRisk;
        predecessors[neighbor] = point;
        queue.add(_PointRiskPair(neighbor, candidateCumulativeRisk));
      }
    }
  }

  return cumulativeRiskToPoint[dest] ?? _maxInt;
}

Map<Point, Map<Point, int>> riskGridTo5xGraph(List<List<int>> grid) {
  final gridRowCount = grid.length;
  final gridColCount = grid[0].length;
  final caveRowCount = gridRowCount * 5;
  final caveColCount = gridColCount * 5;

  final graph = <Point, Map<Point, int>>{};
  for (var i = 0; i < caveRowCount; i++) {
    for (var j = 0; j < caveColCount; j++) {
      final successor = Point(i, j);

      final gridI = i % gridRowCount;
      final gridJ = j % gridColCount;
      final riskModifier = (i ~/ gridRowCount) + (j ~/ gridColCount);
      var riskLevel = grid[gridI][gridJ] + riskModifier;
      if (riskLevel > 9) {
        riskLevel = (riskLevel % 9);
        if (riskLevel == 0) {
          riskLevel = 1;
        }
      }

      for (var iOffset = -1; iOffset <= 1; iOffset++) {
        final predecessorI = i + iOffset;

        if (predecessorI < 0 || caveRowCount <= predecessorI) {
          continue;
        }

        for (var jOffset = -1; jOffset <= 1; jOffset++) {
          final predecessorJ = j + jOffset;

          if (predecessorJ < 0 || caveColCount <= predecessorJ) {
            continue;
          }

          if (predecessorI == i && predecessorJ == j) {
            continue;
          }

          if (iOffset.abs() == 1 && jOffset.abs() == 1) {
            continue;
          }

          final predecessor = Point(predecessorI, predecessorJ);
          graph.update(predecessor, (edges) {
            edges[successor] = riskLevel;
            return edges;
          }, ifAbsent: () => {successor: riskLevel});
        }
      }
    }
  }

  return graph;
}
