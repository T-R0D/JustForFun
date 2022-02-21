import 'dart:math';

import 'package:aoc2021/src/solver/solver.dart' show Solver;
import 'package:aoc2021/src/space/point.dart' show Point3D;

class Day22Solver implements Solver {
  List<_Instruction> _instructions = [];

  @override
  void consumeRawInput(String rawInput) {
    final lines = rawInput.split('\n');
    _instructions = lines.map((line) {
      final actionAndRange = line.split(' ');
      final action =
          actionAndRange[0] == 'on' ? _Action.turnOn : _Action.turnOff;

      final ranges = actionAndRange[1]
          .replaceAll('x=', '')
          .replaceAll('y=', '')
          .replaceAll('z=', '')
          .split(',');
      final rangeValues = ranges.map((rangeStr) {
        return rangeStr.split('..').map(int.parse).toList();
      }).toList();

      return _Instruction(
        action,
        rangeValues[0][0],
        rangeValues[0][1],
        rangeValues[1][0],
        rangeValues[1][1],
        rangeValues[2][0],
        rangeValues[2][1],
      );
    }).toList();
  }

  @override
  String solvePart1() {
    final activeCubes = _executeInitializationInstructions(_instructions);
    return activeCubes.length.toString();
  }

  @override
  String solvePart2() {
    final activeCubes = _executeRebootInstructios(_instructions);
    return activeCubes.toString();
  }
}

enum _Action {
  turnOff,
  turnOn,
}

class _Instruction {
  final _Action action;
  final int x0;
  final int x1;
  final int y0;
  final int y1;
  final int z0;
  final int z1;

  _Instruction(
    this.action,
    this.x0,
    this.x1,
    this.y0,
    this.y1,
    this.z0,
    this.z1,
  );
}

Set<Point3D> _executeInitializationInstructions(
    List<_Instruction> instructions) {
  final activeCubes = <Point3D>{};
  for (var instruction in instructions) {
    for (var x = max(instruction.x0, -50); x <= min(instruction.x1, 50); x++) {
      for (var y = max(instruction.y0, -50);
          y <= min(instruction.y1, 50);
          y++) {
        for (var z = max(instruction.z0, -50);
            z <= min(instruction.z1, 50);
            z++) {
          final cube = Point3D(x, y, z);
          if (instruction.action == _Action.turnOn) {
            activeCubes.add(cube);
          } else {
            activeCubes.remove(cube);
          }
        }
      }
    }
  }
  return activeCubes;
}

int _executeRebootInstructios(List<_Instruction> instructions) {
  List<_AddCuboidInstruction> distilledInstructions = [];
  for (var instruction in instructions) {
    final cuboid = _Cuboid(
      instruction.x0,
      instruction.x1,
      instruction.y0,
      instruction.y1,
      instruction.z0,
      instruction.z1,
    );

    final intersections = <_AddCuboidInstruction>[];
    for (var distilledInstruction in distilledInstructions) {
      final intersection = distilledInstruction.cuboid.intersection(cuboid);
      if (intersection != null) {
        intersections.add(
            _AddCuboidInstruction(!distilledInstruction.add, intersection));
      }
    }
    distilledInstructions.addAll(intersections);

    if (instruction.action == _Action.turnOn) {
      distilledInstructions.add(_AddCuboidInstruction(true, cuboid));
    }
  }

  int activeCubeCount = 0;
  for (var distilledInstruction in distilledInstructions) {
    if (distilledInstruction.add) {
      activeCubeCount += distilledInstruction.cuboid.volume;
    } else {
      activeCubeCount -= distilledInstruction.cuboid.volume;
    }
  }
  return activeCubeCount;
}

class _Cuboid {
  final List<int> rangeEndPoints;

  _Cuboid(
    int x0,
    int x1,
    int y0,
    int y1,
    int z0,
    int z1,
  ) : rangeEndPoints = [x0, x1, y0, y1, z0, z1];

  _Cuboid? intersection(_Cuboid other) {
    final intersectionRangeEndpoints = [];
    for (var i = 1; i < rangeEndPoints.length; i += 2) {
      final startEndpoint =
          max(rangeEndPoints[i - 1], other.rangeEndPoints[i - 1]);
      final endEndpoint = min(rangeEndPoints[i], other.rangeEndPoints[i]);
      if (startEndpoint > endEndpoint) {
        return null;
      }
      intersectionRangeEndpoints.addAll([startEndpoint, endEndpoint]);
    }

    return _Cuboid(
      intersectionRangeEndpoints[0],
      intersectionRangeEndpoints[1],
      intersectionRangeEndpoints[2],
      intersectionRangeEndpoints[3],
      intersectionRangeEndpoints[4],
      intersectionRangeEndpoints[5],
    );
  }

  int get volume {
    int v = 1;
    for (var i = 1; i < rangeEndPoints.length; i += 2) {
      v *= (rangeEndPoints[i] - rangeEndPoints[i - 1] + 1);
    }
    return v;
  }
}

class _AddCuboidInstruction {
  final bool add;
  final _Cuboid cuboid;

  _AddCuboidInstruction(this.add, this.cuboid);
}
