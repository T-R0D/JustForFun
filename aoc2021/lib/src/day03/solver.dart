import 'package:aoc2021/src/binary/binary.dart';
import 'package:aoc2021/src/solver/solver.dart' show Solver;

class Day03Solver implements Solver {
  List<String> _reportReadings = [];

  @override
  void consumeRawInput(String rawInput) {
    _reportReadings = rawInput.split('\n').toList();
  }

  @override
  String solvePart1() {
    final onBitCounts = _countOnBits(_reportReadings);
    final rates =
        _findPowerConsumptionRates(onBitCounts, _reportReadings.length);
    return '${rates.gammaRate * rates.epsilonRate}';
  }

  @override
  String solvePart2() {
    final lifeSupportRatings = _LifeSupportRatings(0, 0);
    lifeSupportRatings.o2GeneratorRating = _filterToFindRate(_reportReadings,
        (numOnBits, majorityCount) => numOnBits >= majorityCount ? '1' : '0');
    lifeSupportRatings.co2scrubberRating = _filterToFindRate(_reportReadings,
        (numOnBits, majorityCount) => numOnBits < majorityCount ? '1' : '0');

    return '${lifeSupportRatings.o2GeneratorRating * lifeSupportRatings.co2scrubberRating}';
  }

  List<int> _countOnBits(List<String> readings) {
    final readingLength = readings[0].length;
    final onBitCounts = List.filled(readingLength, 0);
    for (var reading in readings) {
      for (var i = 0; i < readingLength; i++) {
        if (reading[i] == '1') {
          onBitCounts[i]++;
        }
      }
    }
    return onBitCounts;
  }

  PowerConsumptionRates _findPowerConsumptionRates(
      List<int> onBitCounts, int numReadings) {
    var rates = PowerConsumptionRates(0, 0);
    final majority = (numReadings / 2) + 1;
    for (var i = 0; i < onBitCounts.length; i++) {
      rates.gammaRate <<= 1;
      rates.epsilonRate <<= 1;
      if (onBitCounts[i] >= majority) {
        rates.gammaRate |= 1;
      } else {
        rates.epsilonRate |= 1;
      }
    }
    return rates;
  }

  int _filterToFindRate(
      List<String> readings,
      String Function(int numOnBits, int majorityCount)
          determineDesiredBitValue) {
    final readingLength = readings[0].length;
    var remainingReadings = readings;
    for (var i = 0; i < readingLength && remainingReadings.length > 1; i++) {
      final onBitCounts = _countOnBits(remainingReadings);
      final majority = remainingReadings.length ~/ 2 +
          (remainingReadings.length & 1 == 1 ? 1 : 0);
      final desiredBitValue =
          determineDesiredBitValue(onBitCounts[i], majority);
      remainingReadings = remainingReadings
          .where((reading) => reading[i] == desiredBitValue)
          .toList();
    }

    if (remainingReadings.length != 1) {
      throw StateError('remainingReadings was not reduced to a single value');
    }

    return strToBin(remainingReadings[0]);
  }
}

class PowerConsumptionRates {
  int gammaRate;
  int epsilonRate;

  PowerConsumptionRates(this.gammaRate, this.epsilonRate);
}

class _LifeSupportRatings {
  int o2GeneratorRating;
  int co2scrubberRating;

  _LifeSupportRatings(this.o2GeneratorRating, this.co2scrubberRating);
}
