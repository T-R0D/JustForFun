import 'package:aoc2021/src/solver/solver.dart' show Solver;

class Day11Solver implements Solver {
  List<List<int>> _octopusMap = [];

  @override
  void consumeRawInput(String rawInput) {
//     rawInput = '''11111
// 19991
// 19191
// 19991
// 11111''';
//     rawInput = '''5483143223
// 2745854711
// 5264556173
// 6141336146
// 6357385478
// 4167524645
// 2176841721
// 6882881134
// 4846848554
// 5283751526''';

    _octopusMap = rawInput
        .split('\n')
        .map((line) => line.split('').map(int.parse).toList())
        .toList();
  }

  @override
  String solvePart1() {
    final octopusMap = [
      for (var row in _octopusMap) [for (var energyLevel in row) energyLevel]
    ];
    var totalFlashes = 0;
    for (var t = 0; t < 100; t++) {
      totalFlashes += _executeOneTimeStep(octopusMap);
    }
    return totalFlashes.toString();
  }

  @override
  String solvePart2() {
    final octopusMap = [
      for (var row in _octopusMap) [for (var energyLevel in row) energyLevel]
    ];
    final totalOctopuses = octopusMap.length * octopusMap[0].length;
    for (var t = 1;; t++) {
      final flashes = _executeOneTimeStep(octopusMap);
      if (flashes == totalOctopuses) {
        return t.toString();
      }
    }
  }
}

int _executeOneTimeStep(List<List<int>> octopusMap) {
  final flashMap = [
    for (var row in octopusMap) [for (var _ in row) false]
  ];
  var flashes = 0;

  for (var i = 0; i < octopusMap.length; i++) {
    for (var j = 0; j < octopusMap[0].length; j++) {
      flashes += _visitOctopus(octopusMap, flashMap, i, j);
    }
  }

  return flashes;
}

int _visitOctopus(List<List<int>> octopusMap, List<List<bool>> flashMap,
    final int i, final int j) {
  if (flashMap[i][j]) {
    return 0;
  }

  octopusMap[i][j]++;

  if (octopusMap[i][j] <= 9) {
    return 0;
  }

  var flashes = 1;
  flashMap[i][j] = true;
  octopusMap[i][j] = 0;

  for (var iOffset = -1; iOffset <= 1; iOffset++) {
    final neighborI = i + iOffset;
    if (neighborI < 0 || octopusMap.length <= neighborI) {
      continue;
    }

    for (var jOffset = -1; jOffset <= 1; jOffset++) {
      if (iOffset == 0 && jOffset == 0) {
        continue;
      }
      final neighborJ = j + jOffset;
      if (neighborJ < 0 || octopusMap[0].length <= neighborJ) {
        continue;
      }

      flashes += _visitOctopus(octopusMap, flashMap, neighborI, neighborJ);
    }
  }
  return flashes;
}

String _octopusMapToString(List<List<int>> octopusMap) {
  final buffer = StringBuffer();
  for (var i = 0; i < octopusMap.length; i++) {
    for (var j = 0; j < octopusMap[0].length; j++) {
      final energy = octopusMap[i][j];
      buffer.write('$energy${energy == 0 ? '!' : ' '}');
    }
    buffer.write('\n');
  }
  return buffer.toString();
}
