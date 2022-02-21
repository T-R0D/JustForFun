import 'package:aoc2021/src/solver/solver.dart';
import 'package:aoc2021/src/space/point.dart';

class Day04Solver implements Solver {
  List<int> _calledNumbers = [];
  List<_BingoBoard> _boards = [];

  @override
  void consumeRawInput(String rawInput) {
    final parts = rawInput.split('\n\n');
    _calledNumbers = parts[0]
        .split(',')
        .map((calledNumberStr) => int.parse(calledNumberStr))
        .toList();
    _boards =
        parts.sublist(1).map((boardSpec) => _BingoBoard(boardSpec)).toList();
  }

  @override
  String solvePart1() {
    var winningBoardNdx = -1;
    var winningNumber = -1;
    CALL_NUMBERS:
    for (var calledNumber in _calledNumbers) {
      for (var i = 0; i < _boards.length; i++) {
        final board = _boards[i];
        final isWin = board.markNumber(calledNumber);
        if (isWin) {
          winningBoardNdx = i;
          winningNumber = calledNumber;
          break CALL_NUMBERS;
        }
      }
    }
    if (winningBoardNdx == -1) {
      throw StateError('No winning board found');
    }

    return '${_boards[winningBoardNdx].scoreBoard(winningNumber)}';
  }

  @override
  String solvePart2() {
    Set<int> remainingBoards =
        Set.from(List.generate(_boards.length, (index) => index));
    List<int> winningBoardOrder = [];
    int lastWinningNumber = -1;

    CALLING_NUMBERS:
    for (var calledNumber in _calledNumbers) {
      for (var i = 0; i < _boards.length; i++) {
        if (remainingBoards.contains(i)) {
          final board = _boards[i];
          final isWin = board.markNumber(calledNumber);
          if (isWin) {
            remainingBoards.remove(i);
            winningBoardOrder.add(i);
          }
        }

        if (remainingBoards.isEmpty) {
          lastWinningNumber = calledNumber;
          break CALLING_NUMBERS;
        }
      }
    }

    if (lastWinningNumber == -1) {
      throw StateError('Not all boards were winners after calling all numbers');
    }

    final lastWinningBoardNdx = winningBoardOrder[winningBoardOrder.length - 1];
    final loserBoard = _boards[lastWinningBoardNdx];
    return '${loserBoard.scoreBoard(lastWinningNumber)}';
  }
}

class _BingoBoard {
  int _boardSize = 0;
  List<List<_Spot>> _spots = [];
  final Map<int, Point> _spotLocations = {};

  _BingoBoard(String boardSpec) {
    final rows = boardSpec.split('\n');
    _boardSize = rows.length;
    _spots = List.generate(
        _boardSize, (index) => List.generate(_boardSize, (index) => _Spot(0)));
    for (var i = 0; i < rows.length; i++) {
      final row = rows[i].trimLeft().split(RegExp(r'\s+'));
      final spotValues =
          row.map((spotValueStr) => int.parse(spotValueStr)).toList();
      for (var j = 0; j < spotValues.length; j++) {
        final spotValue = spotValues[j];
        _spots[i][j].value = spotValue;
        _spotLocations[spotValue] = Point(i, j);
      }
    }
  }

  bool markNumber(int number) {
    final spotCoordinates = _spotLocations[number];
    if (spotCoordinates == null) {
      return false;
    }

    final i = spotCoordinates.x;
    final j = spotCoordinates.y;

    _spots[i][j].marked = true;

    return _checkForWin(i, j);
  }

  int scoreBoard(int justCalledNumber) {
    int unmarkedSum = 0;
    for (var i = 0; i < _boardSize; i++) {
      for (var j = 0; j < _boardSize; j++) {
        final spot = _spots[i][j];
        if (!spot.marked) {
          unmarkedSum += spot.value;
        }
      }
    }
    return unmarkedSum * justCalledNumber;
  }

  @override
  String toString() {
    final buffer = StringBuffer();
    for (var i = 0; i < _boardSize; i++) {
      for (var j = 0; j < _boardSize; j++) {
        final spot = _spots[i][j];
        buffer.write(
            '${spot.value.toString().padLeft(2)} ${spot.marked ? 'X' : ' '} ');
      }
      buffer.write('\n');
    }
    return buffer.toString();
  }

  bool _checkForWin(int i, int j) {
    var hasRowWin = true;
    for (var j2 = 0; j2 < _boardSize; j2++) {
      if (!_spots[i][j2].marked) {
        hasRowWin = false;
        break;
      }
    }
    if (hasRowWin) {
      return hasRowWin;
    }

    var hasColWin = true;
    for (var i2 = 0; i2 < _boardSize; i2++) {
      if (!_spots[i2][j].marked) {
        hasColWin = false;
        break;
      }
    }
    return hasColWin;
  }
}

class _Spot {
  int value = 0;
  bool marked = false;

  _Spot(this.value);
}
