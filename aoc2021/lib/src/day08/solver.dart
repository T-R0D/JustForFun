import 'package:aoc2021/src/solver/solver.dart' show Solver;

class Day08Solver implements Solver {
  List<_NoteEntry> notes = [];

  @override
  void consumeRawInput(String rawInput) {
//     rawInput =
//         // 'acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf';
//         '''be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
// edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
// fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
// fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
// aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
// fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
// dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
// bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
// egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
// gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce''';

    final lines = rawInput.split('\n');
    notes = [];
    for (var line in lines) {
      final lineParts = line.split(' | ');
      final noteEntry =
          _NoteEntry(lineParts[0].split(' '), lineParts[1].split(' '));
      notes.add(noteEntry);
    }
  }

  @override
  String solvePart1() {
    const uniqueNumbersOfSegments = <int>{
      2,
      4,
      3,
      7,
    };
    var uniqueNumberOfSegmentDigitCount = 0;
    for (var note in notes) {
      for (var displayDigit in note.outputValues) {
        if (uniqueNumbersOfSegments.contains(displayDigit.length)) {
          uniqueNumberOfSegmentDigitCount++;
        }
      }
    }
    return '$uniqueNumberOfSegmentDigitCount';
  }

  @override
  String solvePart2() {
    var outputSum = 0;

    for (var note in notes) {
      final solver = _SevenSegmentSolver();
      solver.solveScrambling(note.observedSignalPatterns);

      var outputValue = 0;
      for (var litSegmentString in note.outputValues) {
        outputValue = (outputValue * 10) + solver.decodeDigit(litSegmentString);
      }

      outputSum += outputValue;
    }

    return outputSum.toString();
  }
}

class _NoteEntry {
  final List<String> observedSignalPatterns;
  final List<String> outputValues;

  _NoteEntry(this.observedSignalPatterns, this.outputValues);
}

enum Segment {
  a,
  b,
  c,
  d,
  e,
  f,
  g,
}

class _SevenSegmentSolver {
  static const stringToSegment = <String, Segment>{
    'a': Segment.a,
    'b': Segment.b,
    'c': Segment.c,
    'd': Segment.d,
    'e': Segment.e,
    'f': Segment.f,
    'g': Segment.g,
  };

  static const _correctMappings = {
    0: {
      Segment.a,
      Segment.b,
      Segment.c,
      Segment.e,
      Segment.f,
      Segment.g,
    },
    1: {
      Segment.c,
      Segment.f,
    },
    2: {
      Segment.a,
      Segment.c,
      Segment.d,
      Segment.e,
      Segment.g,
    },
    3: {
      Segment.a,
      Segment.c,
      Segment.d,
      Segment.f,
      Segment.g,
    },
    4: {
      Segment.b,
      Segment.c,
      Segment.d,
      Segment.f,
    },
    5: {
      Segment.a,
      Segment.b,
      Segment.d,
      Segment.f,
      Segment.g,
    },
    6: {
      Segment.a,
      Segment.b,
      Segment.d,
      Segment.e,
      Segment.f,
      Segment.g,
    },
    7: {
      Segment.a,
      Segment.c,
      Segment.f,
    },
    8: {
      Segment.a,
      Segment.b,
      Segment.c,
      Segment.d,
      Segment.e,
      Segment.f,
      Segment.g,
    },
    9: {
      Segment.a,
      Segment.b,
      Segment.c,
      Segment.d,
      Segment.f,
      Segment.g,
    },
  };

  final Map<Segment, Set<Segment>> _wireToSegment = {
    for (var i in Segment.values) i: {for (var j in Segment.values) j}
  };

  final Map<Set<Segment>, int> _litSegmentsToDigitValue = {};

  bool get isSolved {
    for (var entry in _wireToSegment.entries) {
      if (entry.value.length != 1) {
        return false;
      }
    }
    return true;
  }

  void solveScrambling(List<String> hintStrings) {
    if (isSolved) {
      return;
    }

    final hints = hintStrings
        .map((hintString) =>
            hintString.split('').map(_convertStringToSegment).toSet())
        .toList();
    hints.sort((a, b) => a.length.compareTo(b.length));

    _determineMappingsByFrequency(hints);

    _determineMappingsByNumberOfSegmentsUsed(hints);

    if (!isSolved) {
      throw Exception('Mapping not solved after applying hints!');
    }

    _configureTranslations();
  }

  int decodeDigit(String pattern) {
    final litSegments = pattern.split('').map(_convertStringToSegment).toSet();
    for (var entry in _litSegmentsToDigitValue.entries) {
      final targetLitSegments = entry.key;
      if (targetLitSegments.length == litSegments.length &&
          targetLitSegments.containsAll(litSegments)) {
        return entry.value;
      }
    }
    throw Exception('unable to convert segments $litSegments to digit value');
  }

  Segment _convertStringToSegment(String char) {
    final maybeSegment = stringToSegment[char];
    if (maybeSegment == null) {
      throw StateError('$char is not a segment');
    }
    return maybeSegment;
  }

  void _determineMappingsByFrequency(List<Set<Segment>> hints) {
    final segmentCounts = <Segment, int>{};

    for (var hint in hints) {
      for (var segment in hint) {
        segmentCounts.update(segment, (value) => value + 1, ifAbsent: () => 1);
      }
    }

    for (var entry in segmentCounts.entries) {
      final segment = entry.key;
      final count = entry.value;
      switch (count) {
        case 4:
          _finalizeSingleMapping(Segment.e, segment);
          break;
        case 6:
          _finalizeSingleMapping(Segment.b, segment);
          break;
        case 9:
          _finalizeSingleMapping(Segment.f, segment);
          break;
      }
    }
  }

  void _finalizeSingleMapping(Segment src, Segment dest) {
    final litSegment = <Segment>{dest};
    _wireToSegment[src] = litSegment;
    for (var entry in _wireToSegment.entries) {
      final srcSegment = entry.key;
      if (srcSegment == src) {
        continue;
      }
      entry.value.remove(dest);
    }
  }

  void _determineMappingsByNumberOfSegmentsUsed(List<Set<Segment>> hints) {
    for (var hint in hints) {
      switch (hint.length) {
        case 2:
          _eliminatePossibilitiesForKnownLengthDigit(1, hint);
          break;
        case 3:
          _eliminatePossibilitiesForKnownLengthDigit(7, hint);
          break;
        case 4:
          _eliminatePossibilitiesForKnownLengthDigit(4, hint);
          break;
      }
    }
  }

  _eliminatePossibilitiesForKnownLengthDigit(int digit, Set<Segment> hint) {
    final correctMapping = _correctMappings[digit];
    if (correctMapping == null) {
      throw Exception('Mapping for $digit does not exist');
    }

    for (var srcSegment in Segment.values) {
      final currentPossibilities = _wireToSegment[srcSegment];
      if (currentPossibilities == null) {
        throw Exception('No current possibilities for src $srcSegment');
      }

      if (correctMapping.contains(srcSegment)) {
        _wireToSegment[srcSegment] = currentPossibilities.intersection(hint);
      } else {
        _wireToSegment[srcSegment] = currentPossibilities.difference(hint);
      }
    }
  }

  void _configureTranslations() {
    for (var correctMapping in _correctMappings.entries) {
      final digitValue = correctMapping.key;
      final srcSegments = correctMapping.value;

      final litSegments = <Segment>{};
      for (var srcSegment in srcSegments) {
        final litSegment = _wireToSegment[srcSegment];
        if (litSegment == null || litSegment.length != 1) {
          throw Exception('Invalid mapping for src=$srcSegment to $litSegment');
        }
        litSegments.addAll(litSegment);
      }
      _litSegmentsToDigitValue[litSegments] = digitValue;
    }
  }
}
