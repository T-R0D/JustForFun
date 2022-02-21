import 'dart:collection';

import 'package:aoc2021/src/solver/solver.dart' show Solver;

class Day10Solver implements Solver {
  List<List<String>> _lines = [];

  @override
  void consumeRawInput(String rawInput) {
    _lines = rawInput.split('\n').map((line) => line.split('')).toList();
  }

  @override
  String solvePart1() {
    var corruptionScoreSum = 0;
    for (var line in _lines) {
      corruptionScoreSum += _getCorruptionScoreForLine(line);
    }
    return corruptionScoreSum.toString();
  }

  @override
  String solvePart2() {
    final incompleteLines =
        _lines.where((line) => _getCorruptionScoreForLine(line) == 0);

    final autocompleteScores = <int>[];
    for (var line in incompleteLines) {
      final autocompleteScore = _getAutocompleteScore(line);
      autocompleteScores.add(autocompleteScore);
    }

    autocompleteScores.sort((a, b) => a.compareTo(b));
    final medianScore =
        autocompleteScores[autocompleteScores.length ~/ 2];
    return medianScore.toString();
  }
}

int _getCorruptionScoreForLine(List<String> line) {
  final stack = Queue<String>();

  for (var chunkDelimiter in line) {
    if (_matchingDelimiters.keys.contains(chunkDelimiter)) {
      stack.addLast(chunkDelimiter);
    } else if (_matchingDelimiters.values.contains(chunkDelimiter)) {
      final previousChunkDelimiter = stack.removeLast();
      if (!delimitersMatch(previousChunkDelimiter, chunkDelimiter)) {
        return _getCorruptionScore(chunkDelimiter);
      }
    }
  }

  return 0;
}

const _matchingDelimiters = <String, String>{
  '(': ')',
  '[': ']',
  '{': '}',
  '<': '>',
};

bool delimitersMatch(String openingDelimeter, String closingDelimiter) {
  return _matchingDelimiters[openingDelimeter] == closingDelimiter;
}

const _corruptionScores = <String, int>{
  ')': 3,
  ']': 57,
  '}': 1197,
  '>': 25137,
};

int _getCorruptionScore(String closingDelimiter) {
  return _corruptionScores[closingDelimiter] ?? 0;
}

const _completionScores = <String, int>{')': 1, ']': 2, '}': 3, '>': 4};

int _getAutocompleteScore(List<String> line) {
  final processingQueue = Queue<String>.of(line);
  var autocompleteScore = 0;
  while (processingQueue.isNotEmpty) {
    final delimeter = processingQueue.removeLast();
    if (_matchingDelimiters.keys.contains(delimeter)) {
      final closingDelimiter = _matchingDelimiters[delimeter];
      final score = _completionScores[closingDelimiter];
      if (score == null) {
        throw Exception('No score for delimeter: $closingDelimiter');
      }
      autocompleteScore = (autocompleteScore * 5) + score;
    } else {
      _trimCompletePairs(processingQueue, delimeter);
    }
  }

  return autocompleteScore;
}

void _trimCompletePairs(
    Queue<String> processingQueue, String startingDelimiter) {
  final stack = Queue<String>.of([startingDelimiter]);
  while (stack.isNotEmpty) {
    final nextDelimiter = processingQueue.removeLast();
    if (_matchingDelimiters.values.contains(nextDelimiter)) {
      stack.addFirst(nextDelimiter);
    } else if (_matchingDelimiters.keys.contains(nextDelimiter)) {
      final maybeMatchingClosingDelimiter = stack.removeFirst();
      if (_matchingDelimiters[nextDelimiter] != maybeMatchingClosingDelimiter) {
        throw Exception(
            'Non matching pair discovered: $nextDelimiter =/= $maybeMatchingClosingDelimiter');
      }
    }
  }
}
