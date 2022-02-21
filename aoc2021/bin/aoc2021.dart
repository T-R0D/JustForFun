import 'dart:io';

import 'package:aoc2021/src/day01/solver.dart' show Day01Solver;
import 'package:aoc2021/src/day02/solver.dart' show Day02Solver;
import 'package:aoc2021/src/day03/solver.dart' show Day03Solver;
import 'package:aoc2021/src/day04/solver.dart' show Day04Solver;
import 'package:aoc2021/src/day05/solver.dart' show Day05Solver;
import 'package:aoc2021/src/day06/solver.dart' show Day06Solver;
import 'package:aoc2021/src/day07/solver.dart' show Day07Solver;
import 'package:aoc2021/src/day08/solver.dart' show Day08Solver;
import 'package:aoc2021/src/day09/solver.dart' show Day09Solver;
import 'package:aoc2021/src/day10/solver.dart' show Day10Solver;
import 'package:aoc2021/src/day11/solver.dart' show Day11Solver;
import 'package:aoc2021/src/day12/solver.dart' show Day12Solver;
import 'package:aoc2021/src/day13/solver.dart' show Day13Solver;
import 'package:aoc2021/src/day14/solver.dart' show Day14Solver;
import 'package:aoc2021/src/day15/solver.dart' show Day15Solver;
import 'package:aoc2021/src/day16/solver.dart' show Day16Solver;
import 'package:aoc2021/src/day17/solver.dart' show Day17Solver;
import 'package:aoc2021/src/day18/solver.dart' show Day18Solver;
import 'package:aoc2021/src/day19/solver.dart' show Day19Solver;
import 'package:aoc2021/src/day20/solver.dart' show Day20Solver;
import 'package:aoc2021/src/day21/solver.dart' show Day21Solver;
import 'package:aoc2021/src/day22/solver.dart' show Day22Solver;
import 'package:aoc2021/src/day23/solver.dart' show Day23Solver;
import 'package:aoc2021/src/day24/solver.dart' show Day24Solver;
import 'package:aoc2021/src/day25/solver.dart' show Day25Solver;
import 'package:aoc2021/src/solver/solver.dart' show Solver;
import 'package:args/args.dart';

const help = 'help';
const dayArgName = 'day';
const partArgName = 'part';

void main(List<String> arguments) {
  final parser = ArgParser()
    ..addFlag(help,
        abbr: 'h', help: 'Print this help message', defaultsTo: false)
    ..addOption(dayArgName,
        help: 'The day to process',
        valueHelp: '1-25',
        allowed: List<String>.generate(25, (i) => '${i + 1}'))
    ..addOption(partArgName,
        help: 'The part of the day to process',
        valueHelp: '1-2',
        allowed: ['1', '2']);
  final argResults = parser.parse(arguments);

  if (argResults[help] as bool) {
    print(parser.usage);
    return;
  }

  final dayArg = argResults[dayArgName];
  final partArg = argResults[partArgName];

  if (dayArg == null) {
    print('Argument for "day" is required.');
    return;
  } else if (partArg == null) {
    print('Argument for "part" is required.');
    return;
  }

  final day = int.parse(dayArg);
  final part = int.parse(partArg);

  print('Solving day $dayArg part $part...');

  final inputFilePath =
      './lib/src/day${day.toString().padLeft(2, '0')}/input.txt';
  Solver? solver;
  switch (day) {
    case 1:
      solver = Day01Solver();
      break;
    case 2:
      solver = Day02Solver();
      break;
    case 3:
      solver = Day03Solver();
      break;
    case 4:
      solver = Day04Solver();
      break;
    case 5:
      solver = Day05Solver();
      break;
    case 6:
      solver = Day06Solver();
      break;
    case 7:
      solver = Day07Solver();
      break;
    case 8:
      solver = Day08Solver();
      break;
    case 9:
      solver = Day09Solver();
      break;
    case 10:
      solver = Day10Solver();
      break;
    case 11:
      solver = Day11Solver();
      break;
    case 12:
      solver = Day12Solver();
      break;
    case 13:
      solver = Day13Solver();
      break;
    case 14:
      solver = Day14Solver();
      break;
    case 15:
      solver = Day15Solver();
      break;
    case 16:
      solver = Day16Solver();
      break;
    case 17:
      solver = Day17Solver();
      break;
    case 18:
      solver = Day18Solver();
      break;
    case 19:
      solver = Day19Solver();
      break;
    case 20:
      solver = Day20Solver();
      break;
    case 21:
      solver = Day21Solver();
      break;
    case 22:
      solver = Day22Solver();
      break;
    case 23:
      solver = Day23Solver();
      break;
    case 24:
      solver = Day24Solver();
      break;
    case 25:
      solver = Day25Solver();
      break;
  }

  if (solver == null) {
    print('Unable to find a relevant solver.');
    return;
  }

  final rawInput = _readRawInput(inputFilePath);

  solver.consumeRawInput(rawInput);

  String? result;
  if (part == 1) {
    result = solver.solvePart1();
  } else {
    result = solver.solvePart2();
  }

  print(result);
}

String _readRawInput(String inputFilePath) {
  return File(inputFilePath).readAsStringSync();
}
