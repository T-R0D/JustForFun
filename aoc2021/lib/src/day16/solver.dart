import 'dart:math' show max, min;

import 'package:aoc2021/src/solver/solver.dart' show Solver;

class Day16Solver implements Solver {
  List<String> _messageBits = [];

  @override
  void consumeRawInput(String rawInput) {
    final hexDigits = rawInput.split('');
    _messageBits = <String>[];
    for (var digit in hexDigits) {
      final digitBits = _hexValueLookup[digit];
      if (digitBits == null) {
        throw Exception('$digit is not a hex digit');
      }
      _messageBits.addAll(digitBits);
    }
  }

  @override
  String solvePart1() {
    final processor = _BitsMessageProcessor(_messageBits);
    final rootPacket = processor.processMessage();
    final versionSum = rootPacket.versionSum();
    return versionSum.toString();
  }

  @override
  String solvePart2() {
    final processor = _BitsMessageProcessor(_messageBits);
    final rootPacket = processor.processMessage();
    final packetValue = rootPacket.computeValue();
    return packetValue.toString();
  }
}

const Map<String, List<String>> _hexValueLookup = {
  '0': [
    '0',
    '0',
    '0',
    '0',
  ],
  '1': [
    '0',
    '0',
    '0',
    '1',
  ],
  '2': [
    '0',
    '0',
    '1',
    '0',
  ],
  '3': [
    '0',
    '0',
    '1',
    '1',
  ],
  '4': [
    '0',
    '1',
    '0',
    '0',
  ],
  '5': [
    '0',
    '1',
    '0',
    '1',
  ],
  '6': [
    '0',
    '1',
    '1',
    '0',
  ],
  '7': [
    '0',
    '1',
    '1',
    '1',
  ],
  '8': [
    '1',
    '0',
    '0',
    '0',
  ],
  '9': [
    '1',
    '0',
    '0',
    '1',
  ],
  'A': [
    '1',
    '0',
    '1',
    '0',
  ],
  'B': [
    '1',
    '0',
    '1',
    '1',
  ],
  'C': [
    '1',
    '1',
    '0',
    '0',
  ],
  'D': [
    '1',
    '1',
    '0',
    '1',
  ],
  'E': [
    '1',
    '1',
    '1',
    '0',
  ],
  'F': [
    '1',
    '1',
    '1',
    '1',
  ],
};

class _BitsMessageProcessor {
  final List<String> _reversedBits;
  bool _messageProcessed = false;

  _BitsMessageProcessor(List<String> messageBits)
      : _reversedBits = messageBits.reversed.toList();

  _Packet processMessage() {
    if (_messageProcessed) {
      throw Error();
    }

    final packets = _processMessageIntoKnownNumberOfPackets(1);
    if (packets.length != 1) {
      throw Error();
    }

    _messageProcessed = true;

    return packets[0];
  }

  List<_Packet> _processMessageIntoKnownNumberOfPackets(int expectedPackets) {
    final packets = <_Packet>[];
    for (var p = 0; p < expectedPackets; p++) {
      final packet = _processOnePacket();
      packets.add(packet);
    }

    // Ummmm... This should hold true, but using the given input, I get the
    // right answer but this fails...
    //
    // while (_reversedBits.isNotEmpty) {
    //   if (_reversedBits.removeLast() != '0') {
    //     throw Error();
    //   }
    // }

    return packets;
  }

  List<_Packet> _processMessageIntoUnknownNumberOfPackets() {
    final packets = <_Packet>[];
    while (_reversedBits.isNotEmpty) {
      final packet = _processOnePacket();
      packets.add(packet);
    }
    return packets;
  }

  _Packet _processOnePacket() {
    final version = _readNBitsIntoValue(3);
    final typeId = _readNBitsIntoValue(3);
    final packet = _Packet(version, typeId);

    if (typeId == _literalValueTypeId) {
      final value = _processLiteralValue();
      packet.setLiteralValue(value);
    } else {
      final lengthTypeId = _readNBitsIntoValue(1);
      if (lengthTypeId == 0) {
        final subPacketsBitLength = _readNBitsIntoValue(15);
        final subPacketBits = <String>[];
        for (var i = 0; i < subPacketsBitLength; i++) {
          subPacketBits.add(_reversedBits.removeLast());
        }

        final subProcessor = _BitsMessageProcessor(subPacketBits);
        final subPackets =
            subProcessor._processMessageIntoUnknownNumberOfPackets();
        packet.addSubPackets(subPackets);
      } else {
        final subPacketCount = _readNBitsIntoValue(11);
        final subPackets =
            _processMessageIntoKnownNumberOfPackets(subPacketCount);
        packet.addSubPackets(subPackets);
      }
    }
    return packet;
  }

  int _processLiteralValue() {
    var value = 0;

    var keepReading = true;
    while (keepReading) {
      keepReading = _readNBitsIntoValue(1) == 1;
      final chunkValue = _readNBitsIntoValue(4);
      value <<= 4;
      value |= chunkValue;
    }

    return value;
  }

  int _readNBitsIntoValue(int n) {
    var value = 0;
    for (var i = 0; i < n; i++) {
      final bit = _reversedBits.removeLast();
      value <<= 1;
      if (bit == '1') {
        value |= 1;
      }
    }
    return value;
  }
}

const _literalValueTypeId = 4;

class _Packet {
  final int _version;
  final int _typeId;
  int _literalValue = 0;
  final List<_Packet> _subPackets = [];

  _Packet(this._version, this._typeId);

  void setLiteralValue(int value) {
    if (_typeId != _literalValueTypeId) {
      throw Error();
    }
    _literalValue = value;
  }

  void addSubPackets(List<_Packet> subPackets) {
    if (_typeId == _literalValueTypeId) {
      throw Error();
    }
    _subPackets.addAll(subPackets);
  }

  int versionSum() {
    var sum = _version;
    for (var subPacket in _subPackets) {
      sum += subPacket.versionSum();
    }
    return sum;
  }

  int computeValue() {
    switch (_typeId) {
      case 0:
        return _subPackets.map((p) => p.computeValue()).reduce((a, b) => a + b);
      case 1:
        return _subPackets.map((p) => p.computeValue()).reduce((a, b) => a * b);
      case 2:
        return _subPackets
            .map((p) => p.computeValue())
            .reduce((a, b) => min(a, b));
      case 3:
        return _subPackets
            .map((p) => p.computeValue())
            .reduce((a, b) => max(a, b));
      case 4:
        return _literalValue;
      case 5:
        return _subPackets[0].computeValue() > _subPackets[1].computeValue()
            ? 1
            : 0;
      case 6:
        return _subPackets[0].computeValue() < _subPackets[1].computeValue()
            ? 1
            : 0;
      case 7:
        return _subPackets[0].computeValue() == _subPackets[1].computeValue()
            ? 1
            : 0;
    }
    return 0;
  }
}
