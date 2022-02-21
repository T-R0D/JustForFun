import 'dart:collection';
import 'dart:math';

import 'package:aoc2021/src/solver/solver.dart' show Solver;

// TODO: Come back, re-analyze, and really explain what's going on here.

class Day24Solver implements Solver {
  List<_ProgramInstruction> _programInstructions = [];

  @override
  void consumeRawInput(String rawInput) {
    final lines = rawInput.split('\n');
    _programInstructions = lines.map((line) {
      final parts = line.split(' ');
      final instruction = toInstruction(parts[0]);
      final dest = toRegister(parts[1]);

      if (parts.length == 2) {
        return _ProgramInstruction(instruction, dest);
      }

      final srcValLiteral = int.tryParse(parts[2]);
      if (srcValLiteral != null) {
        return _ProgramInstruction(instruction, dest,
            isLiteral: true, literalValue: srcValLiteral);
      }

      final srcRegister = toRegister(parts[2]);
      return _ProgramInstruction(instruction, dest, src: srcRegister);
    }).toList();
  }

  @override
  String solvePart1() {
    final modelNumber = _findHighestValitModelNumber(_programInstructions);
    return modelNumber.toString();
  }

  @override
  String solvePart2() {
    final modelNumber = _findLowestValitModelNumber(_programInstructions);
    return modelNumber.toString();
  }
}

int _findHighestValitModelNumber(List<_ProgramInstruction> monadProgram) {
  final relevantInstructionValues =
      _findRelevantInstructionValues(monadProgram);
  final stack = Queue<_IndexValuePair>();
  final digits = List.generate(relevantInstructionValues.length, (_) => 0);

  for (var i = 0; i < relevantInstructionValues.length; i++) {
    final value = relevantInstructionValues[i];
    if (value > 0) {
      stack.addLast(_IndexValuePair(i, value));
    } else {
      final complement = stack.removeLast();
      final complementIndex = complement.index;
      final complementValue = complement.value;

      final diff = complementValue + value;

      digits[complementIndex] = min(9, 9 - diff);
      digits[i] = min(9, 9 + diff);
    }
  }

  var modelNumber = 0;
  for (var digit in digits) {
    modelNumber *= 10;
    modelNumber += digit;
  }

  return modelNumber;
}

int _findLowestValitModelNumber(List<_ProgramInstruction> monadProgram) {
  final relevantInstructionValues =
      _findRelevantInstructionValues(monadProgram);
  final stack = Queue<_IndexValuePair>();
  final digits = List.generate(relevantInstructionValues.length, (_) => 0);

  for (var i = 0; i < relevantInstructionValues.length; i++) {
    final value = relevantInstructionValues[i];
    if (value > 0) {
      stack.addLast(_IndexValuePair(i, value));
    } else {
      final complement = stack.removeLast();
      final complementIndex = complement.index;
      final complementValue = complement.value;

      final diff = complementValue + value;

      digits[complementIndex] = max(1, 1 - diff);
      digits[i] = max(1, 1 + diff);
    }
  }

  var modelNumber = 0;
  for (var digit in digits) {
    modelNumber *= 10;
    modelNumber += digit;
  }

  return modelNumber;
}

List<int> _findRelevantInstructionValues(
    List<_ProgramInstruction> monadProgram) {
  final relevantInstructionValues = <int>[];

  for (var i = 0; i < monadProgram.length; i += 18) {
    final checkOffsetValue = monadProgram[i + 5].literalValue;
    if (checkOffsetValue == 0) {
      throw Error();
    }

    final pushOffsetValue = monadProgram[i + 15].literalValue;
    if (pushOffsetValue == 0) {
      throw Error();
    }

    if (checkOffsetValue < 0) {
      relevantInstructionValues.add(checkOffsetValue);
    } else {
      relevantInstructionValues.add(pushOffsetValue);
    }
  }

  return relevantInstructionValues;
}

class _IndexValuePair {
  final int index;
  final int value;

  const _IndexValuePair(this.index, this.value);
}

/* Wow! Would this finish in a year? */

int _findHighestValidModelNumberBruteForce(
    Iterable<_ProgramInstruction> monadProgram) {
  for (var modelNumber = 99999999999999;
      modelNumber >= 11111111111111;
      modelNumber--) {
    if (_containsAnyZeros(modelNumber)) {
      continue;
    }

    print('processing: $modelNumber');

    final alu = _ALU();
    alu.addInput(modelNumber.toString().split('').map(int.parse));
    for (var programInstruction in monadProgram) {
      if (programInstruction.isLiteral) {
        alu.op(programInstruction.instruction, programInstruction.dest,
            programInstruction.literalValue);
      } else {
        alu.opR(programInstruction.instruction, programInstruction.dest,
            programInstruction.src);
      }
    }

    if (alu.zRegister == 0) {
      return modelNumber;
    }
  }

  return -1;
}

bool _containsAnyZeros(int n) {
  return n.toString().split('').any((digit) => digit == '0');
}

class _ProgramInstruction {
  final _Instruction instruction;
  final _Register dest;
  final bool isLiteral;
  final int literalValue;
  final _Register src;

  _ProgramInstruction(this.instruction, this.dest,
      {this.isLiteral = false, this.literalValue = 0, this.src = _Register.w});
}

enum _Instruction { inp, add, mul, div, mod, eql }

_Instruction toInstruction(String instructionStr) {
  switch (instructionStr) {
    case 'inp':
      return _Instruction.inp;
    case 'add':
      return _Instruction.add;
    case 'mul':
      return _Instruction.mul;
    case 'div':
      return _Instruction.div;
    case 'mod':
      return _Instruction.mod;
    case 'eql':
      return _Instruction.eql;
    default:
      throw Error();
  }
}

enum _Register { w, x, y, z }

_Register toRegister(String registerStr) {
  switch (registerStr) {
    case 'w':
      return _Register.w;
    case 'x':
      return _Register.x;
    case 'y':
      return _Register.y;
    case 'z':
      return _Register.z;
    default:
      throw Error();
  }
}

class _ALU {
  final _registers = {
    _Register.w: 0,
    _Register.x: 0,
    _Register.y: 0,
    _Register.z: 0,
  };

  final Queue<int> _inputQueue = Queue();

  _ALU();

  void addInput(Iterable<int> newInput) {
    _inputQueue.addAll(newInput);
  }

  void opR(_Instruction instruction, _Register dest, _Register src) {
    final srcVal = _get(src);
    op(instruction, dest, srcVal);
  }

  void op(_Instruction instruction, _Register dest, int val) {
    switch (instruction) {
      case _Instruction.inp:
        _inp(dest);
        break;
      case _Instruction.add:
        _add(dest, val);
        break;
      case _Instruction.mul:
        _mul(dest, val);
        break;
      case _Instruction.div:
        _div(dest, val);
        break;
      case _Instruction.mod:
        _mod(dest, val);
        break;
      case _Instruction.eql:
        _eql(dest, val);
        break;
    }
  }

  int get zRegister => _get(_Register.z);

  void _inp(_Register reg) {
    final value = _inputQueue.removeFirst();
    _store(reg, value);
  }

  void _add(_Register reg, int value) {
    final first = _get(reg);
    _store(reg, first + value);
  }

  void _mul(_Register reg, int value) {
    final first = _get(reg);
    _store(reg, first * value);
  }

  void _div(_Register reg, int value) {
    final first = _get(reg);
    _store(reg, first ~/ value);
  }

  void _mod(_Register reg, int value) {
    final first = _get(reg);
    _store(reg, first % value);
  }

  void _eql(_Register reg, int value) {
    final first = _get(reg);
    final result = first == value ? 1 : 0;
    _store(reg, result);
  }

  int _get(_Register reg) {
    final value = _registers[reg];
    if (value == null) {
      throw Error();
    }
    return value;
  }

  void _store(_Register reg, int value) {
    _registers[reg] = value;
  }
}
