import 'package:aoc2021/src/solver/solver.dart' show Solver;
import 'package:aoc2021/src/space/point.dart';

class Day13Solver implements Solver {
  List<Point> _initialDots = [];
  List<_FoldInstruction> _foldInstructions = [];

  @override
  void consumeRawInput(String rawInput) {
//     rawInput = '''6,10
// 0,14
// 9,10
// 0,3
// 10,4
// 4,11
// 6,0
// 6,12
// 4,1
// 0,13
// 10,12
// 3,4
// 3,0
// 8,4
// 1,10
// 2,14
// 8,10
// 9,0

// fold along y=7
// fold along x=5''';

    final parts = rawInput.split('\n\n');

    _initialDots = parts[0].split('\n').map((pair) {
      final pairParts = pair.split(',').map(int.parse).toList();
      return Point(pairParts[0], pairParts[1]);
    }).toList();

    _foldInstructions = parts[1].split('\n').map((instruction) {
      final instructionParts =
          instruction.replaceAll('fold along ', '').split('=');
      return _FoldInstruction(
          instructionParts[0], int.parse(instructionParts[1]));
    }).toList();
  }

  @override
  String solvePart1() {
    final paper = _TransparencyPaper(_initialDots);
    paper.fold(_foldInstructions[0].axis, _foldInstructions[0].linePosition);
    return paper.dotCount.toString();
  }

  @override
  String solvePart2() {
    final paper = _TransparencyPaper(_initialDots);
    for (var instruction in _foldInstructions) {
      paper.fold(instruction.axis, instruction.linePosition);
    }
    return paper.toString();
  }
}

class _FoldInstruction {
  String axis = "";
  int linePosition = 0;

  _FoldInstruction(this.axis, this.linePosition);
}

class _TransparencyPaper {
  List<Point> _dots = [];

  _TransparencyPaper(this._dots);

  int get dotCount => _dots.length;

  fold(String axis, int linePosition) {
    if (axis == 'x') {
      _foldLeft(linePosition);
    } else if (axis == 'y') {
      _foldUp(linePosition);
    } else {
      throw Error();
    }
  }

  _foldLeft(int linePosition) {
    final newDots = <Point>[];
    final existingDots = <String>{};

    for (var dot in _dots) {
      var newX = dot.x;
      if (dot.x > linePosition) {
        newX = (2 * linePosition) - dot.x;
      }

      final newDot = Point(newX, dot.y);
      if (existingDots.contains(newDot.toString())) {
        continue;
      }
      existingDots.add(newDot.toString());
      newDots.add(newDot);
    }

    _dots = newDots;
  }

  _foldUp(int linePosition) {
    final newDots = <Point>[];
    final existingDots = <String>{};

    for (var dot in _dots) {
      var newY = dot.y;
      if (dot.y > linePosition) {
        newY = (2 * linePosition) - dot.y;
      }

      final newDot = Point(dot.x, newY);
      if (existingDots.contains(newDot.toString())) {
        continue;
      }
      existingDots.add(newDot.toString());
      newDots.add(newDot);
    }

    _dots = newDots;
  }

  @override
  String toString() {
    var maxX = 0;
    var maxY = 0;
    for (var dot in _dots) {
      if (dot.x > maxX) {
        maxX = dot.x;
      }
      if (dot.y > maxY) {
        maxY = dot.y;
      }
    }

    final grid = List<List<String>>.generate(
        maxY + 1, (i) => List<String>.generate(maxX + 1, (j) => ' '));

    for (var dot in _dots) {
      grid[dot.y][dot.x] = '#';
    }

    final buffer = StringBuffer();
    for (var row in grid) {
      for (var cell in row) {
        buffer.write(cell);
      }
      buffer.write('\n');
    }
    return buffer.toString();
  }
}
