import 'dart:math';

import 'package:aoc2021/src/solver/solver.dart' show Solver;
import 'package:aoc2021/src/space/point.dart';
import 'package:collection/collection.dart';

class Day23Solver implements Solver {
  final Map<Point, String> _initialLayout = {};
  final Set<Point> _floorSpaces = {};
  final Map<Point, String> _initialAmphipodPositions = {};

  @override
  void consumeRawInput(String rawInput) {
    final lines = rawInput.split('\n');
    for (var i = 0; i < lines.length; i++) {
      final line = lines[i];
      final cells = line.split('');
      for (var j = 0; j < cells.length; j++) {
        final cell = cells[j];
        final p = Point(i, j);
        _initialLayout[p] = cell;
        if (!['#', ' '].contains(cell)) {
          _floorSpaces.add(p);
          if (['A', 'B', 'C', 'D'].contains(cell)) {
            _initialAmphipodPositions[p] = cell;
          }
        }
      }
    }
  }

  @override
  String solvePart1() {
    final cost = findMostEfficientReorganizationCost(_initialAmphipodPositions);
    return cost.toString();
  }

  @override
  String solvePart2() {
    final initialAmphipodPositions = {..._initialAmphipodPositions};
    for (var j = 3; j <= 9; j += 2) {
      final p1 = Point(3, j);
      final p2 = Point(4, j);
      final p3 = Point(5, j);

      initialAmphipodPositions[p3] = initialAmphipodPositions[p1] ?? '?';
      initialAmphipodPositions[p1] = _foldedPositions[p1] ?? '?';
      initialAmphipodPositions[p2] = _foldedPositions[p2] ?? '?';
    }
    final cost =
        findMostEfficientReorganizationCostFull(initialAmphipodPositions);
    return cost.toString();
  }
}

int findMostEfficientReorganizationCost(
    Map<Point, String> initialAmphipodPositions) {
  final frontier = PriorityQueue<_State>((a, b) => a.cost.compareTo(b.cost));
  frontier.add(_State(initialAmphipodPositions, 0));
  final seen = <String>{};

  while (frontier.isNotEmpty) {
    final stateUnderConsideration = frontier.removeFirst();

    if (stateUnderConsideration.isFinished) {
      return stateUnderConsideration.cost;
    }

    final nextStr = stateUnderConsideration.toString();
    if (seen.contains(nextStr)) {
      continue;
    }

    final nextStates = stateUnderConsideration.generateNextStates();
    frontier.addAll(nextStates);

    seen.add(nextStr);
  }

  return -1;
}

final _foldedPositions = <Point, String>{
  Point(3, 3): 'D',
  Point(4, 3): 'D',
  Point(3, 5): 'C',
  Point(4, 5): 'B',
  Point(3, 7): 'B',
  Point(4, 7): 'A',
  Point(3, 9): 'A',
  Point(4, 9): 'C',
};

int findMostEfficientReorganizationCostFull(
    Map<Point, String> initialAmphipodPositions) {
  final frontier = PriorityQueue<_State2>((a, b) => a.cost.compareTo(b.cost));
  frontier.add(_State2(initialAmphipodPositions, 0));
  final seen = <String>{};

  while (frontier.isNotEmpty) {
    final stateUnderConsideration = frontier.removeFirst();

    if (stateUnderConsideration.isFinished) {
      return stateUnderConsideration.cost;
    }

    final nextStr = stateUnderConsideration.toString();
    if (seen.contains(nextStr)) {
      continue;
    }

    final nextStates = stateUnderConsideration.generateNextStates();
    frontier.addAll(nextStates);

    seen.add(nextStr);
  }

  return -1;
}

const _hallwayI = 1;

bool _isInHallway(Point p) {
  return p.x == _hallwayI;
}

final _waitingHallwaySpots = <Point>{
  Point(1, 1),
  Point(1, 2),
  Point(1, 4),
  Point(1, 6),
  Point(1, 8),
  Point(1, 10),
  Point(1, 11),
};

const _amphipodMovementCosts = <String, int>{
  'A': 1,
  'B': 10,
  'C': 100,
  'D': 1000,
};

int _manhattanDistance(Point p1, Point p2) {
  return (p1.x - p2.x).abs() + (p1.y - p2.y).abs();
}

class _State {
  static final _roomSpots = <String, List<Point>>{
    'A': [Point(2, 3), Point(3, 3)],
    'B': [Point(2, 5), Point(3, 5)],
    'C': [Point(2, 7), Point(3, 7)],
    'D': [Point(2, 9), Point(3, 9)],
  };

  final Map<Point, String> _amphipodPositions;
  final int _costToReach;

  _State(this._amphipodPositions, this._costToReach);

  int get cost => _costToReach;

  @override
  String toString() {
    final buffer = StringBuffer();
    for (var j = 1; j <= 11; j++) {
      final p = Point(_hallwayI, j);
      final occupant = _amphipodPositions[p] ?? '.';
      buffer.write(occupant);
    }
    buffer.write('\n');

    for (var i = 2; i <= 3; i++) {
      buffer.write('  ');
      for (var j = 3; j <= 9; j += 2) {
        final p = Point(i, j);
        final occupant = _amphipodPositions[p] ?? '.';
        buffer.write(occupant);
        buffer.write(' ');
      }
      buffer.write('\n');
    }

    return buffer.toString();
  }

  bool get isFinished {
    for (var entry in _roomSpots.entries) {
      final amphipodType = entry.key;
      final roomPoints = entry.value;
      for (var roomPoint in roomPoints) {
        final occupant = _amphipodPositions[roomPoint];
        if (occupant == null || occupant != amphipodType) {
          return false;
        }
      }
    }
    return true;
  }

  Iterable<_State> generateNextStates() {
    final nextStates = <_State>[];

    for (var amphipodEntry in _amphipodPositions.entries) {
      final currentPosition = amphipodEntry.key;
      final kind = amphipodEntry.value;

      if (_isInHallway(currentPosition)) {
        if (!_roomContainsIncorrectAmphipods(kind)) {
          final deepestAvailableRoomSlot = _getDeepestAvailableRoomSlot(kind);
          if (deepestAvailableRoomSlot != null &&
              _isReachableFromHall(deepestAvailableRoomSlot, currentPosition)) {
            final newState = _createNeighboringState(
                currentPosition, deepestAvailableRoomSlot, kind);
            nextStates.add(newState);
          }
        }
      } else if (!_isInCorrectRoom(kind, currentPosition) ||
          _isBlocking(currentPosition, kind)) {
        final availableHallwaySpots =
            _waitingHallwaySpots.difference(_amphipodPositions.keys.toSet());
        for (var availableHallwaySpot in availableHallwaySpots) {
          if (_isReachableFromRoom(availableHallwaySpot, currentPosition)) {
            final newState = _createNeighboringState(
                currentPosition, availableHallwaySpot, kind);
            nextStates.add(newState);
          }
        }
      }
    }

    return nextStates;
  }

  bool _isInCorrectRoom(String kind, Point position) {
    final validPositions = _roomSpots[kind];
    if (validPositions == null) {
      throw Error();
    }
    return validPositions.contains(position);
  }

  bool _roomContainsIncorrectAmphipods(String amphipodType) {
    final roomSpots = _roomSpots[amphipodType];
    if (roomSpots == null) {
      return false;
    }

    for (var entry in _amphipodPositions.entries) {
      if (entry.value != amphipodType && roomSpots.contains(entry.key)) {
        return true;
      }
    }

    return false;
  }

  Point? _getDeepestAvailableRoomSlot(String kind) {
    final roomSpots = _roomSpots[kind];
    if (roomSpots == null) {
      return null;
    }

    final availableSpots =
        roomSpots.toSet().difference(_amphipodPositions.keys.toSet()).toList();
    if (availableSpots.isEmpty) {
      return null;
    }
    final deepestRoom = availableSpots
        .reduce((value, element) => value.x > element.x ? value : element);
    return deepestRoom;
  }

  bool _isBlocking(Point currentLocation, String kind) {
    final roomSpots = _roomSpots[kind];
    if (roomSpots == null) {
      throw Error();
    }

    final deeperSpots =
        roomSpots.where((roomSpot) => roomSpot.x > currentLocation.x);
    for (var spot in deeperSpots) {
      final occupant = _amphipodPositions[spot];
      if (occupant == null || occupant != kind) {
        return true;
      }
    }
    return false;
  }

  bool _isReachableFromHall(Point to, Point from) {
    final pointsInPath = <Point>[];
    if (from.y < to.y) {
      for (var j = from.y + 1; j <= to.y; j++) {
        pointsInPath.add(Point(_hallwayI, j));
      }
    } else {
      for (var j = from.y - 1; j >= to.y; j--) {
        pointsInPath.add(Point(_hallwayI, j));
      }
    }
    for (var i = _hallwayI + 1; i <= to.x; i++) {
      pointsInPath.add(Point(i, to.y));
    }
    return pointsInPath
        .every((point) => !_amphipodPositions.containsKey(point));
  }

  bool _isReachableFromRoom(Point to, Point from) {
    final pointsInPath = <Point>[];
    for (var i = from.x - 1; i > _hallwayI; i--) {
      pointsInPath.add(Point(i, from.y));
    }
    for (var j = min(to.y, from.y); j <= max(to.y, from.y); j++) {
      pointsInPath.add(Point(_hallwayI, j));
    }
    return pointsInPath
        .every((point) => !_amphipodPositions.containsKey(point));
  }

  _State _createNeighboringState(
      Point oldPosition, Point newPosition, String amphipodType) {
    final newPositions = {..._amphipodPositions};
    newPositions.remove(oldPosition);
    newPositions[newPosition] = amphipodType;
    final costToMove = _manhattanDistance(oldPosition, newPosition) *
        (_amphipodMovementCosts[amphipodType] ?? 1);
    return _State(newPositions, cost + costToMove);
  }
}

class _State2 {
  static final _roomSpots2 = <String, List<Point>>{
    'A': [Point(2, 3), Point(3, 3), Point(4, 3), Point(5, 3)],
    'B': [Point(2, 5), Point(3, 5), Point(4, 5), Point(5, 5)],
    'C': [Point(2, 7), Point(3, 7), Point(4, 7), Point(5, 7)],
    'D': [Point(2, 9), Point(3, 9), Point(4, 9), Point(5, 9)],
  };

  final Map<Point, String> _amphipodPositions;
  final int _costToReach;

  _State2(this._amphipodPositions, this._costToReach);

  int get cost => _costToReach;

  @override
  String toString() {
    final buffer = StringBuffer();
    for (var j = 1; j <= 11; j++) {
      final p = Point(_hallwayI, j);
      final occupant = _amphipodPositions[p] ?? '.';
      buffer.write(occupant);
    }
    buffer.write('\n');

    for (var i = 2; i <= 5; i++) {
      buffer.write('  ');
      for (var j = 3; j <= 9; j += 2) {
        final p = Point(i, j);
        final occupant = _amphipodPositions[p] ?? '.';
        buffer.write(occupant);
        buffer.write(' ');
      }
      buffer.write('\n');
    }

    return buffer.toString();
  }

  bool get isFinished {
    for (var entry in _roomSpots2.entries) {
      final amphipodType = entry.key;
      final roomPoints = entry.value;
      for (var roomPoint in roomPoints) {
        final occupant = _amphipodPositions[roomPoint];
        if (occupant == null || occupant != amphipodType) {
          return false;
        }
      }
    }
    return true;
  }

  Iterable<_State2> generateNextStates() {
    final nextStates = <_State2>[];

    for (var amphipodEntry in _amphipodPositions.entries) {
      final currentPosition = amphipodEntry.key;
      final kind = amphipodEntry.value;

      if (_isInHallway(currentPosition)) {
        if (!_roomContainsIncorrectAmphipods(kind)) {
          final deepestAvailableRoomSlot = _getDeepestAvailableRoomSlot(kind);
          if (deepestAvailableRoomSlot != null &&
              _isReachableFromHall(deepestAvailableRoomSlot, currentPosition)) {
            final newState = _createNeighboringState(
                currentPosition, deepestAvailableRoomSlot, kind);
            nextStates.add(newState);
          }
        }
      } else if (!_isInCorrectRoom(kind, currentPosition) ||
          _isBlocking(currentPosition, kind)) {
        final availableHallwaySpots =
            _waitingHallwaySpots.difference(_amphipodPositions.keys.toSet());
        for (var availableHallwaySpot in availableHallwaySpots) {
          if (_isReachableFromRoom(availableHallwaySpot, currentPosition)) {
            final newState = _createNeighboringState(
                currentPosition, availableHallwaySpot, kind);
            nextStates.add(newState);
          }
        }
      }
    }

    return nextStates;
  }

  bool _isInCorrectRoom(String kind, Point position) {
    final validPositions = _roomSpots2[kind];
    if (validPositions == null) {
      throw Error();
    }
    return validPositions.contains(position);
  }

  bool _roomContainsIncorrectAmphipods(String amphipodType) {
    final roomSpots = _roomSpots2[amphipodType];
    if (roomSpots == null) {
      return false;
    }

    for (var entry in _amphipodPositions.entries) {
      if (entry.value != amphipodType && roomSpots.contains(entry.key)) {
        return true;
      }
    }

    return false;
  }

  Point? _getDeepestAvailableRoomSlot(String kind) {
    final roomSpots = _roomSpots2[kind];
    if (roomSpots == null) {
      return null;
    }

    final availableSpots =
        roomSpots.toSet().difference(_amphipodPositions.keys.toSet()).toList();
    if (availableSpots.isEmpty) {
      return null;
    }
    final deepestRoom = availableSpots
        .reduce((value, element) => value.x > element.x ? value : element);
    return deepestRoom;
  }

  bool _isBlocking(Point currentLocation, String kind) {
    final roomSpots = _roomSpots2[kind];
    if (roomSpots == null) {
      throw Error();
    }

    final deeperSpots =
        roomSpots.where((roomSpot) => roomSpot.x > currentLocation.x);
    for (var spot in deeperSpots) {
      final occupant = _amphipodPositions[spot];
      if (occupant == null || occupant != kind) {
        return true;
      }
    }
    return false;
  }

  bool _isReachableFromHall(Point to, Point from) {
    final pointsInPath = <Point>[];
    if (from.y < to.y) {
      for (var j = from.y + 1; j <= to.y; j++) {
        pointsInPath.add(Point(_hallwayI, j));
      }
    } else {
      for (var j = from.y - 1; j >= to.y; j--) {
        pointsInPath.add(Point(_hallwayI, j));
      }
    }
    for (var i = _hallwayI + 1; i <= to.x; i++) {
      pointsInPath.add(Point(i, to.y));
    }
    return pointsInPath
        .every((point) => !_amphipodPositions.containsKey(point));
  }

  bool _isReachableFromRoom(Point to, Point from) {
    final pointsInPath = <Point>[];
    for (var i = from.x - 1; i > _hallwayI; i--) {
      pointsInPath.add(Point(i, from.y));
    }
    for (var j = min(to.y, from.y); j <= max(to.y, from.y); j++) {
      pointsInPath.add(Point(_hallwayI, j));
    }
    return pointsInPath
        .every((point) => !_amphipodPositions.containsKey(point));
  }

  _State2 _createNeighboringState(
      Point oldPosition, Point newPosition, String amphipodType) {
    final newPositions = {..._amphipodPositions};
    newPositions.remove(oldPosition);
    newPositions[newPosition] = amphipodType;
    final costToMove = _manhattanDistance(oldPosition, newPosition) *
        (_amphipodMovementCosts[amphipodType] ?? 1);
    return _State2(newPositions, cost + costToMove);
  }
}
