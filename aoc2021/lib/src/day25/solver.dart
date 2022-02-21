import 'package:aoc2021/src/solver/solver.dart' show Solver;

class Day25Solver implements Solver {
  List<List<String>> _initialGrid = [];

  @override
  void consumeRawInput(String rawInput) {
    final grid = <List<String>>[];
    for (var line in rawInput.split('\n')) {
      final row = <String>[];
      for (var cell in line.split('')) {
        row.add(cell);
      }
      grid.add(row);
    }
    _initialGrid = grid;
  }

  @override
  String solvePart1() {
    final steps = _simulateMovementUntilItStops(_initialGrid);
    return steps.toString();
  }

  @override
  String solvePart2() {
    return 'Merry Christmas!';
  }

  int _simulateMovementUntilItStops(List<List<String>> initialGrid) {
    var steps = 0;
    var movesMade = 0;
    var currentState = initialGrid;

    do {
      final pair = _simulateMovement(currentState);
      movesMade = pair.movesMade;
      currentState = pair.grid;
      steps++;
    } while (movesMade > 0);

    return steps;
  }
}

_GridMovesPair _simulateMovement(List<List<String>> start) {
  final intermediate = List<List<String>>.generate(start.length,
      (i) => List<String>.generate(start[i].length, (j) => start[i][j]));

  var movesMade = 0;

  final rowCount = start.length;
  for (var i = 0; i < rowCount; i++) {
    final columnCount = start[i].length;
    for (var j = 0; j < columnCount; j++) {
      final occupant = start[i][j];
      if (occupant == '.' || occupant == 'v') {
        continue;
      }

      final nextI = i;
      final nextJ = (j + 1) % columnCount;

      final nextOccupant = start[nextI][nextJ];
      if (nextOccupant == '.') {
        intermediate[i][j] = '.';
        intermediate[nextI][nextJ] = occupant;
        movesMade++;
      }
    }
  }

  final next = List<List<String>>.generate(
      intermediate.length,
      (i) => List<String>.generate(
          intermediate[i].length, (j) => intermediate[i][j]));

  for (var i = 0; i < rowCount; i++) {
    final columnCount = intermediate[i].length;
    for (var j = 0; j < columnCount; j++) {
      final occupant = intermediate[i][j];
      if (occupant == '.' || occupant == '>') {
        continue;
      }

      final nextI = (i + 1) % rowCount;
      final nextJ = j;

      final nextOccupant = intermediate[nextI][nextJ];
      if (nextOccupant == '.') {
        next[i][j] = '.';
        next[nextI][nextJ] = occupant;
        movesMade++;
      }
    }
  }

  return _GridMovesPair(next, movesMade);
}

class _GridMovesPair {
  final List<List<String>> grid;
  final int movesMade;

  const _GridMovesPair(this.grid, this.movesMade);
}

String _gridToString(List<List<String>> grid) {
  final buffer = StringBuffer();
  for (var row in grid) {
    for (var cell in row) {
      buffer.write(cell);
    }
    buffer.write('\n');
  }
  return buffer.toString();
}
