import 'package:aoc2021/src/solver/solver.dart' show Solver;

class Day14Solver implements Solver {
  String _initialPolymer = '';
  final Map<String, String> _expansions = {};

  @override
  void consumeRawInput(String rawInput) {
    final parts = rawInput.split('\n\n');

    _initialPolymer = parts[0];

    for (var expansionRule in parts[1].split('\n')) {
      final ruleParts = expansionRule.split(' -> ');
      _expansions[ruleParts[0]] = ruleParts[1];
    }
  }

  @override
  String solvePart1() {
    final polymerChain = _PolymerChain.fromString(_initialPolymer, _expansions);
    for (var t = 0; t < 10; t++) {
      polymerChain.grow();
    }
    final frequencies = polymerChain.elementFrequencies;
    final sortedFrequencies = frequencies.entries.toList()
      ..sort((a, b) => a.value.compareTo(b.value));
    final nakedCounts = sortedFrequencies.map((entry) => entry.value).toList();
    return '${nakedCounts[nakedCounts.length - 1] - nakedCounts[0]}';
  }

  @override
  String solvePart2() {
    final polymerChain =
        _PolymerChainSummary.fromString(_initialPolymer, _expansions);
    for (var t = 0; t < 40; t++) {
      polymerChain.grow();
    }
    final frequencies = polymerChain.elementFrequencies;
    final sortedFrequencies = frequencies.entries.toList()
      ..sort((a, b) => a.value.compareTo(b.value));
    final nakedCounts = sortedFrequencies.map((entry) => entry.value).toList();
    return '${nakedCounts[nakedCounts.length - 1] - nakedCounts[0]}';
  }
}

class _PolymerChain {
  Map<String, String> _expansions = {};
  _PolymerNode? _head;

  _PolymerChain.fromString(
      String initialPolymer, Map<String, String> expansions) {
    _expansions = expansions;
    if (initialPolymer.length > 1) {
      var previous = _PolymerNode(initialPolymer.substring(0, 1), null);
      _head = previous;
      for (var element in initialPolymer.substring(1).split('')) {
        final newNode = _PolymerNode(element, null);
        previous.next = newNode;
        previous = newNode;
      }
    }
  }

  void grow() {
    var previous = _head;
    if (previous == null) {
      return;
    }
    var current = previous.next;
    for (;
        previous != null && current != null;
        previous = current, current = current.next) {
      final pair = previous.element + current.element;
      final insertedElement = _expansions[pair];
      if (insertedElement == null) {
        continue;
      }
      previous.next = _PolymerNode(insertedElement, current);
    }
  }

  Map<String, int> get elementFrequencies {
    final frequencies = <String, int>{};
    for (var cursor = _head; cursor != null; cursor = cursor.next) {
      frequencies.update(cursor.element, (value) => value + 1,
          ifAbsent: () => 1);
    }
    return frequencies;
  }

  @override
  String toString() {
    final buffer = StringBuffer();
    for (var current = _head; current != null; current = current.next) {
      buffer.write(current.element);
    }
    return buffer.toString();
  }
}

class _PolymerNode {
  final String element;
  _PolymerNode? next;

  _PolymerNode(this.element, this.next);
}

class _PolymerChainSummary {
  Map<String, int> _elementFrequencies = {};
  Map<String, int> _elementPairCounts = {};
  Map<String, String> _expansions = {};

  _PolymerChainSummary.fromString(String initialPolymer, this._expansions) {
    final elements = initialPolymer.split('');
    _elementFrequencies[elements[0]] = 1;
    for (var i = 1; i < elements.length; i++) {
      _elementFrequencies.update(elements[i], (frequency) => frequency + 1,
          ifAbsent: () => 1);
      final elementPair = elements[i - 1] + elements[i];
      _elementPairCounts.update(elementPair, (count) => count + 1,
          ifAbsent: () => 1);
    }
  }

  void grow() {
    final newPairCounts = <String, int>{};
    for (var entry in _elementPairCounts.entries) {
      final pair = entry.key;
      final count = entry.value;
      final insertedElement = _expansions[pair];
      if (insertedElement == null) {
        newPairCounts[pair] = count;
        continue;
      }
      _elementFrequencies.update(
          insertedElement, (frequency) => frequency + count,
          ifAbsent: () => 1);
      final pairParts = pair.split('');
      final firstNewPair = pairParts[0] + insertedElement;
      final secondNewPair = insertedElement + pairParts[1];
      newPairCounts.update(firstNewPair, (value) => value + count,
          ifAbsent: () => count);
      newPairCounts.update(secondNewPair, (value) => value + count,
          ifAbsent: () => count);
    }
    _elementPairCounts = newPairCounts;
  }

  Map<String, int> get elementFrequencies => _elementFrequencies
      .map((element, frequency) => MapEntry(element, frequency));

  @override
  String toString() {
    return _elementPairCounts.toString();
  }
}
