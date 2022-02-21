import 'package:aoc2021/src/solver/solver.dart' show Solver;
import 'package:aoc2021/src/space/point.dart';

class Day02Solver implements Solver {
  List<_Instruction> instructions = [];

  @override
  void consumeRawInput(String rawInput) {
    instructions = [];

    final lines = rawInput.split('\n');
    for (var i = 0; i < lines.length; i += 1) {
      final line = lines[i];
      final parts = line.split(' ');
      if (parts.length != 2) {
        throw StateError(
            'Unexpected number of parts on line $i: ${parts.length}');
      }
      _Direction? direction;
      switch (parts[0]) {
        case 'down':
          direction = _Direction.down;
          break;
        case 'forward':
          direction = _Direction.forward;
          break;
        case 'up':
          direction = _Direction.up;
          break;
        default:
          throw StateError('Unexpected direction: ${parts[0]}');
      }

      final magnitude = int.parse(parts[1]);

      instructions.add(_Instruction(direction, magnitude));
    }
  }

  @override
  String solvePart1() {
    var currentLocation = Point(0, 0);

    for (var instruction in instructions) {
      switch (instruction.direction) {
        case _Direction.down:
          currentLocation.y += instruction.magnitude;
          break;
        case _Direction.forward:
          currentLocation.x += instruction.magnitude;
          break;
        case _Direction.up:
          currentLocation.y -= instruction.magnitude;
          break;
      }
    }

    return '${currentLocation.x * currentLocation.y}';
  }

  @override
  String solvePart2() {
    var currentLocation = Point3D(0, 0, 0);

    for (var instruction in instructions) {
      switch (instruction.direction) {
        case _Direction.down:
          currentLocation.z += instruction.magnitude;
          break;
        case _Direction.forward:
          currentLocation.x += instruction.magnitude;
          currentLocation.y += (instruction.magnitude * currentLocation.z);
          break;
        case _Direction.up:
          currentLocation.z -= instruction.magnitude;
          break;
      }
    }

    return '${currentLocation.x * currentLocation.y}';
  }
}

enum _Direction {
  down,
  forward,
  up,
}

class _Instruction {
  _Direction direction;
  int magnitude;

  _Instruction(this.direction, this.magnitude);
}
