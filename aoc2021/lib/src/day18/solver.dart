import 'package:aoc2021/src/solver/solver.dart' show Solver;

class Day18Solver implements Solver {
  List<String> _inputSpecs = [];

  @override
  void consumeRawInput(String rawInput) {
    _inputSpecs = rawInput.split('\n');
  }

  @override
  String solvePart1() {
    final startingNumbers = [
      for (var spec in _inputSpecs) _specToSnailfishNumber(spec),
    ];
    final result = startingNumbers.reduce((value, element) => value + element);
    return result.magnitude.toString();
  }

  @override
  String solvePart2() {
    var maxMagnitude = 0;
    for (var i = 0; i < _inputSpecs.length; i++) {
      for (var j = 0; j < _inputSpecs.length; j++) {
        if (i == j) {
          continue;
        }

        final numbers = [
      for (var spec in _inputSpecs) _specToSnailfishNumber(spec),
    ];
      final x = numbers[i];
        final y = numbers[j];

        final candidateMagnitude = (x + y).magnitude;
        if (candidateMagnitude > maxMagnitude) {
          maxMagnitude = candidateMagnitude;
        }
      }
    }

    return maxMagnitude.toString();
  }
}

class _SnailfishNumber {
  static const _noValue = -1;

  int value = _noValue;
  _SnailfishNumber? parent;
  _SnailfishNumber? left;
  _SnailfishNumber? right;

  _SnailfishNumber operator +(_SnailfishNumber that) {
    final newNumber = _SnailfishNumber();
    newNumber.left = this;
    newNumber.left!.parent = newNumber;
    newNumber.right = that;
    newNumber.right!.parent = newNumber;
    newNumber.reduce();
    return newNumber;
  }

  void reduce() {
    do {
      if (_explodeSomething()) {
        continue;
      }
      if (_splitSomething()) {
        continue;
      }
      break;
    } while (true);
  }

  int get magnitude {
    return _magnitudeHelper(this);
  }

  bool _explodeSomething() {
    final explodingPair = _explodeSearch(this, 0);
    if (explodingPair == null) {
      return false;
    }

    explodingPair.value = 0;
    final leftValue = explodingPair.left!.value;
    explodingPair.left = null;
    final rightValue = explodingPair.right!.value;
    explodingPair.right = null;

    final inOrderPredecessor = _findInOrderPredecessor(explodingPair);
    if (inOrderPredecessor != null) {
      inOrderPredecessor.value += leftValue;
    }

    final inOrderSuccessor = _findInOrderSuccessor(explodingPair);
    if (inOrderSuccessor != null) {
      inOrderSuccessor.value += rightValue;
    }

    return true;
  }

  _SnailfishNumber? _explodeSearch(_SnailfishNumber current, int depth) {
    if (current.value != _noValue) {
      return null;
    }

    if (depth == 4 &&
        (current.left != null &&
            current.left!.value != _noValue &&
            current.right != null &&
            current.right!.value != _noValue)) {
      return current;
    }

    final left = current.left;
    if (left == null) {
      throw Error();
    }
    final leftResult = _explodeSearch(left, depth + 1);
    if (leftResult != null) {
      return leftResult;
    }

    final right = current.right;
    if (right == null) {
      throw Error();
    }
    return _explodeSearch(right, depth + 1);
  }

  _SnailfishNumber? _findInOrderPredecessor(_SnailfishNumber x) {
    var previous = x;
    var current = x.parent;
    while (current != null && identical(previous, current.left)) {
      previous = current;
      current = current.parent;
    }

    if (current == null) {
      return null;
    }

    current = current.left;
    while (current!.right != null) {
      current = current.right;
    }

    return current;
  }

  _SnailfishNumber? _findInOrderSuccessor(_SnailfishNumber x) {
    var previous = x;
    var current = x.parent;
    while (current != null && identical(previous, current.right)) {
      previous = current;
      current = current.parent;
    }

    if (current == null) {
      return null;
    }

    current = current.right;
    while (current!.left != null) {
      current = current.left;
    }

    return current;
  }

  bool _splitSomething() {
    final splitResult = _splitSearch(this);
    if (splitResult == null) {
      return false;
    }

    final leftValue = splitResult.value ~/ 2;
    var rightValue = splitResult.value ~/ 2;
    if (splitResult.value & 0x01 == 1) {
      rightValue += 1;
    }

    splitResult.left = _SnailfishNumber();
    splitResult.left!.value = leftValue;
    splitResult.left!.parent = splitResult;
    splitResult.right = _SnailfishNumber();
    splitResult.right!.value = rightValue;
    splitResult.right!.parent = splitResult;
    splitResult.value = _noValue;

    return true;
  }

  _SnailfishNumber? _splitSearch(_SnailfishNumber? x) {
    if (x == null) {
      return null;
    }

    if (x.value != _noValue && x.value > 9) {
      return x;
    }

    final leftResult = _splitSearch(x.left);
    if (leftResult != null) {
      return leftResult;
    }

    return _splitSearch(x.right);
  }

  int _magnitudeHelper(_SnailfishNumber current) {
    if (current.value != _noValue) {
      return current.value;
    }

    final left = current.left;
    if (left == null) {
      throw Error();
    }
    var sum = 3 * _magnitudeHelper(left);

    final right = current.right;
    if (right == null) {
      throw Error();
    }
    sum += 2 * _magnitudeHelper(right);

    return sum;
  }

  @override
  String toString() {
    final buffer = StringBuffer();
    _toStringHelper(buffer, this);
    return buffer.toString();
  }

  void _toStringHelper(StringBuffer buffer, _SnailfishNumber current) {
    if (current.value != _noValue) {
      buffer.write(current.value);
      return;
    }

    buffer.write('[');

    final left = current.left;
    if (left == null) {
      throw Error();
    }
    _toStringHelper(buffer, left);

    buffer.write(',');

    final right = current.right;
    if (right == null) {
      throw Error();
    }
    _toStringHelper(buffer, right);

    buffer.write(']');
  }
}

_SnailfishNumber _specToSnailfishNumber(String spec) {
  final root = _SnailfishNumber();
  _SnailfishNumber? current = root;
  for (var char in spec.split('')) {
    if (char == '[') {
      current!.left = _SnailfishNumber();
      current.left!.parent = current;
      current = current.left;
    } else if (char == ']') {
      current = current!.parent;
    } else if (char == ',') {
      current!.right = _SnailfishNumber();
      current.right!.parent = current;
      current = current.right;
    } else if (int.tryParse(char) != null) {
      current!.value = int.parse(char);
      current = current.parent;
    } else {
      throw Error();
    }
  }

  return root;
}
