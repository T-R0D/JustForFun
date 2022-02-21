import 'dart:math';

import 'package:aoc2021/src/solver/solver.dart' show Solver;

class Day21Solver implements Solver {
  List<int> _startingPositions = [];

  @override
  void consumeRawInput(String rawInput) {
    final lines = rawInput.split('\n');

    _startingPositions = lines.map((line) {
      final parts = line.split('position: ');
      return int.parse(parts[1]);
    }).toList();
  }

  @override
  String solvePart1() {
    final die = _DeterministicDie();
    final result =
        _playDiracDiceGame(die, _startingPositions[0], _startingPositions[1]);
    return (result.loserScore * result.diceRollCount).toString();
  }

  @override
  String solvePart2() {
    final result =
        _quantumDiracDiceGameDP(_startingPositions[0], _startingPositions[1]);
    return max(result.player1Wins, result.player2Wins).toString();
  }
}

class _DiracDiceResult {
  final int winningPlayer;
  final int winnerScore;
  final int loserScore;
  final int diceRollCount;

  _DiracDiceResult(this.winningPlayer, this.winnerScore, this.loserScore,
      this.diceRollCount);

  @override
  String toString() {
    return 'Winner: $winningPlayer\nWinning Score: $winnerScore\nLosing Score: $loserScore\nDice Rolls: $diceRollCount\n';
  }
}

abstract class _Die {
  int roll();

  int get rollCount;
}

class _DeterministicDie implements _Die {
  int _nextValue = 1;
  int _rollCount = 0;

  @override
  int roll() {
    final result = _nextValue;
    _nextValue = (_nextValue % 100) + 1;
    _rollCount++;
    return result;
  }

  @override
  int get rollCount => _rollCount;
}

_DiracDiceResult _playDiracDiceGame(
    _Die die, int player1Start, int player2Start) {
  List<int> playerPositions = [player1Start - 1, player2Start - 1];
  List<int> scores = [0, 0];

  for (var turnIndicator = 0;; turnIndicator = (turnIndicator + 1) % 2) {
    final movement = (die.roll() + die.roll() + die.roll());
    final newPosition = playerPositions[turnIndicator] =
        (playerPositions[turnIndicator] + movement) % 10;
    final score = scores[turnIndicator] + newPosition + 1;
    scores[turnIndicator] = score;

    if (score >= 1000) {
      return _DiracDiceResult(turnIndicator + 1, score,
          scores[(turnIndicator + 1) % 2], die.rollCount);
    }
  }
}

class _QuantumDiracDiceGameResult {
  int player1Wins;
  int player2Wins;

  _QuantumDiracDiceGameResult(this.player1Wins, this.player2Wins);
}

const _roll3OutcomeFrequency = <int, int>{
  3: 1,
  4: 3,
  5: 6,
  6: 7,
  7: 6,
  8: 3,
  9: 1,
};

class _DPTableEntryKey {
  final int playerAScore;
  final int playerBScore;
  final int playerAPosition;
  final int playerBPosition;

  _DPTableEntryKey(
    this.playerAScore,
    this.playerBScore,
    this.playerAPosition,
    this.playerBPosition,
  );

  @override
  String toString() {
    return [
      playerAScore,
      playerBScore,
      playerAPosition,
      playerBPosition,
    ].toString();
  }

  @override
  bool operator ==(Object other) {
    return other is _DPTableEntryKey && (toString() == other.toString());
  }

  @override
  int get hashCode => toString().hashCode;
}

_QuantumDiracDiceGameResult _quantumDiracDiceGameDP(
    int player1Start, int player2Start) {
  final memo = <_DPTableEntryKey, int>{
    _DPTableEntryKey(0, 0, player1Start - 1, player2Start - 1): 1,
  };

  for (var playerAScore = 0; playerAScore < 21; playerAScore++) {
    for (var playerBScore = 0; playerBScore < 21; playerBScore++) {
      for (var playerAPosition = 0; playerAPosition < 10; playerAPosition++) {
        for (var playerBPosition = 0; playerBPosition < 10; playerBPosition++) {
          final waysToGetHere = memo[_DPTableEntryKey(
            playerAScore,
            playerBScore,
            playerAPosition,
            playerBPosition,
          )];

          if (waysToGetHere == null) {
            continue;
          }

          for (var entry in _roll3OutcomeFrequency.entries) {
            final roll = entry.key;
            final rollFrequency = entry.value;
            final newPosition = (playerAPosition + roll) % 10;
            final newScore = playerAScore + newPosition + 1;

            if (newScore >= 21) {
              final key = _DPTableEntryKey(
                newScore,
                playerBScore,
                newPosition,
                playerBPosition,
              );
              final additionalOutcomes = waysToGetHere * rollFrequency;
              memo.update(key, (value) => value + additionalOutcomes,
                  ifAbsent: () => additionalOutcomes);
            } else {
              for (var entryB in _roll3OutcomeFrequency.entries) {
                final rollB = entryB.key;
                final rollFrequencyB =
                    entryB.value;
                final newPositionB = (playerBPosition + rollB) % 10;
                final newScoreB = playerBScore + newPositionB + 1;
                final additionalOutcomes =
                    waysToGetHere * rollFrequency * rollFrequencyB;
                final key = _DPTableEntryKey(
                  newScore,
                  newScoreB,
                  newPosition,
                  newPositionB,
                );
                memo.update(key, (value) => value + additionalOutcomes,
                    ifAbsent: () => additionalOutcomes);
              }
            }
          }
        }
      }
    }
  }

  final result = _QuantumDiracDiceGameResult(0, 0);
  for (var winnerScore = 21; winnerScore <= 30; winnerScore++) {
    for (var loserScore = 0; loserScore < 21; loserScore++) {
      for (var playerAPosition = 0; playerAPosition < 10; playerAPosition++) {
        for (var playerBPosition = 0; playerBPosition < 10; playerBPosition++) {
          final keyA = _DPTableEntryKey(
            winnerScore,
            loserScore,
            playerAPosition,
            playerBPosition,
          );
          final outcomesA = memo[keyA];
          if (outcomesA != null) {
            result.player1Wins += outcomesA;
          }

          final keyB = _DPTableEntryKey(
            loserScore,
            winnerScore,
            playerAPosition,
            playerBPosition,
          );
          final outcomesB = memo[keyB];
          if (outcomesB != null) {
            result.player2Wins += outcomesB;
          }
        }
      }
    }
  }

  return result;
}
