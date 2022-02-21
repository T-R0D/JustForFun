import 'dart:collection';

import 'package:aoc2021/src/solver/solver.dart' show Solver;

class Day12Solver implements Solver {
  final Map<String, Set<String>> _caveMap = {};

  @override
  void consumeRawInput(String rawInput) {
    for (var line in rawInput.split('\n')) {
      final parts = line.split('-');
      _caveMap.update(parts[0], (existingEdges) => {...existingEdges, parts[1]},
          ifAbsent: () => {parts[1]});
      _caveMap.update(parts[1], (existingEdges) => {...existingEdges, parts[0]},
          ifAbsent: () => {parts[0]});
    }
  }

  @override
  String solvePart1() {
    final foundPaths = _findAllPathsThroughCaveSystem(_caveMap);
    return foundPaths.length.toString();
  }

  @override
  String solvePart2() {
    final foundPaths = _findAllLazyPathsThroughCaveSystem(_caveMap);
    return foundPaths.length.toString();
  }
}

const _caveEntrance = 'start';
const _caveExit = 'end';

class _PathSearch {
  final Set<String> exploredSmallCaves = {};
  final List<String> path = [];

  _PathSearch();

  _PathSearch.fromPrevious(_PathSearch previous) {
    exploredSmallCaves.addAll(previous.exploredSmallCaves);
    path.addAll(previous.path);
  }

  _PathSearch operator +(String caveIdentifier) {
    final newPathSearch = _PathSearch.fromPrevious(this);
    if (_isSmallCave(caveIdentifier)) {
      newPathSearch.exploredSmallCaves.add(caveIdentifier);
    }
    newPathSearch.path.add(caveIdentifier);
    return newPathSearch;
  }
}

bool _isSmallCave(String caveIdentifier) {
  return caveIdentifier == caveIdentifier.toLowerCase();
}

List<_PathSearch> _findAllPathsThroughCaveSystem(
    Map<String, Set<String>> caveMap) {
  final completePaths = <_PathSearch>[];
  final frontier = Queue<_PathSearch>.of([_PathSearch() + _caveEntrance]);
  while (frontier.isNotEmpty) {
    final pathSoFar = frontier.removeFirst();

    if (pathSoFar.path.last == _caveExit) {
      completePaths.add(pathSoFar);
      continue;
    }

    final lastLocation = pathSoFar.path.last;
    final unreducedPossibilities = caveMap[pathSoFar.path.last];
    if (unreducedPossibilities == null) {
      throw Exception(
          'expected there to be a set of possible paths from $lastLocation');
    }

    final nextPossibilities =
        unreducedPossibilities.difference(pathSoFar.exploredSmallCaves);

    for (var possibility in nextPossibilities) {
      frontier.addLast(pathSoFar + possibility);
    }
  }
  return completePaths;
}

class _LazyPathSearch {
  final Set<String> exploredSmallCaves = {};
  final List<String> path = [];
  bool hasExploredSmallCaveTwice = false;

  _LazyPathSearch();

  _LazyPathSearch.fromPrevious(_LazyPathSearch previous) {
    exploredSmallCaves.addAll(previous.exploredSmallCaves);
    path.addAll(previous.path);
    hasExploredSmallCaveTwice = previous.hasExploredSmallCaveTwice;
  }

  _LazyPathSearch operator +(String caveIdentifier) {
    final newPathSearch = _LazyPathSearch.fromPrevious(this);

    newPathSearch.hasExploredSmallCaveTwice =
        newPathSearch.hasExploredSmallCaveTwice ||
            newPathSearch.exploredSmallCaves.contains(caveIdentifier);

    if (_isSmallCave(caveIdentifier)) {
      newPathSearch.exploredSmallCaves.add(caveIdentifier);
    }

    newPathSearch.path.add(caveIdentifier);
    return newPathSearch;
  }
}

List<_LazyPathSearch> _findAllLazyPathsThroughCaveSystem(
    Map<String, Set<String>> caveMap) {
  final completePaths = <_LazyPathSearch>[];
  final frontier =
      Queue<_LazyPathSearch>.of([_LazyPathSearch() + _caveEntrance]);
  while (frontier.isNotEmpty) {
    final pathSoFar = frontier.removeFirst();

    if (pathSoFar.path.last == _caveExit) {
      completePaths.add(pathSoFar);
      continue;
    }

    final lastLocation = pathSoFar.path.last;
    final unreducedPossibilities = caveMap[pathSoFar.path.last];
    if (unreducedPossibilities == null) {
      throw Exception(
          'expected there to be a set of possible paths from $lastLocation');
    }

    var nextPossibilities = unreducedPossibilities.difference({_caveEntrance});
    if (pathSoFar.hasExploredSmallCaveTwice) {
      nextPossibilities =
          nextPossibilities.difference(pathSoFar.exploredSmallCaves);
    }

    for (var possibility in nextPossibilities) {
      frontier.addLast(pathSoFar + possibility);
    }
  }
  return completePaths;
}
