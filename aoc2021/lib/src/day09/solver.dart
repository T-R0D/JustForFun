import 'dart:collection';

import 'package:aoc2021/src/solver/solver.dart' show Solver;
import 'package:aoc2021/src/space/point.dart';

class Day09Solver implements Solver {
  List<List<int>> _elevationMap = [];
  int _rows = 0;
  int _cols = 0;

  @override
  void consumeRawInput(String rawInput) {
    final rows = rawInput.split('\n');
    final elevationMap = <List<int>>[];
    for (var i = 0; i < rows.length; i++) {
      final row = rows[i];
      final rowValues = <int>[];
      for (var j = 0; j < row.length; j++) {
        rowValues.add(int.parse(row[j]));
      }
      elevationMap.add(rowValues);
    }
    _elevationMap = elevationMap;
    _rows = elevationMap.length;
    _cols = elevationMap[0].length;
  }

  @override
  String solvePart1() {
    var riskScoreSum = 0;
    for (var i = 0; i < _rows; i++) {
      for (var j = 0; j < _cols; j++) {
        if (_isLowPoint(i, j)) {
          final target = _elevationMap[i][j];
          riskScoreSum += _computeRiskScore(target);
        }
      }
    }

    return riskScoreSum.toString();
  }

  @override
  String solvePart2() {
    final basins = <Set<Point>>[];
    for (var i = 0; i < _rows; i++) {
      for (var j = 0; j < _cols; j++) {
        if (_isLowPoint(i, j)) {
          final basin = _mapBasin(i, j);
          basins.add(basin);
        }
      }
    }

    basins.sort((a, b) => b.length.compareTo(a.length));
    var topBasinSizeProduct = 1;
    for (var basin in basins.sublist(0, 3)) {
      topBasinSizeProduct *= basin.length;
    }

    return topBasinSizeProduct.toString();
  }

  bool _isLowPoint(int i, int j) {
    final targetElevation = _elevationMap[i][j];

    for (var iOffset = -1; iOffset <= 1; iOffset++) {
      final adjacentI = i + iOffset;
      if (adjacentI < 0 || _rows <= adjacentI) {
        continue;
      }
      for (var jOffset = -1; jOffset <= 1; jOffset++) {
        final adjacentJ = j + jOffset;
        if (adjacentJ < 0 || _cols <= adjacentJ) {
          continue;
        }
        if (iOffset == 0 && jOffset == 0 || iOffset == jOffset) {
          continue;
        }

        final adjacentElevation = _elevationMap[adjacentI][adjacentJ];
        if (adjacentElevation <= targetElevation) {
          return false;
        }
      }
    }

    return true;
  }

  int _computeRiskScore(int elevation) {
    return elevation + 1;
  }

  Set<Point> _mapBasin(int i, int j) {
    final explored = <String>{};
    final frontier = Queue<Point>.of([Point(i, j)]);
    final basin = <Point>{};

    while (frontier.isNotEmpty) {
      final next = frontier.removeFirst();

      final nextString = next.toString();

      if (explored.contains(nextString)) {
        continue;
      }

      explored.add(next.toString());

      if (_elevationMap[next.x][next.y] == 9) {
        continue;
      }

      basin.add(next);

      for (var iOffset = -1; iOffset <= 1; iOffset++) {
        final adjacentI = next.x + iOffset;
        if (adjacentI < 0 || _rows <= adjacentI) {
          continue;
        }
        for (var jOffset = -1; jOffset <= 1; jOffset++) {
          final adjacentJ = next.y + jOffset;
          if (adjacentJ < 0 || _cols <= adjacentJ) {
            continue;
          }
          if (iOffset == 0 && jOffset == 0 || iOffset.abs() == jOffset.abs()) {
            continue;
          }

          final neighbor = Point(adjacentI, adjacentJ);
          frontier.add(neighbor);
        }
      }
    }

    return basin;
  }
}
