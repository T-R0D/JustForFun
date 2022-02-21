import 'package:aoc2021/src/solver/solver.dart' show Solver;
import 'package:aoc2021/src/space/point.dart';

class Day19Solver implements Solver {
  List<List<Point3D>> _scannedBeacons = [];

  @override
  void consumeRawInput(String rawInput) {
    final scannerLists = rawInput.split('\n\n');
    _scannedBeacons = scannerLists
        .map((scannerList) => scannerList
            .split('\n')
            .sublist(1)
            .map((line) => line.split(',').map(int.parse).toList())
            .map((list) => Point3D(list[0], list[1], list[2]))
            .toList())
        .toList();
  }

  @override
  String solvePart1() {
    final reorientedBeacons =
        _determineBeaconsRelativeToScanner0(_scannedBeacons);
    return reorientedBeacons.length.toString();
  }

  @override
  String solvePart2() {
    final translations =
        _determineTranslationsRelativeToScanner0(_scannedBeacons);

    var maxSeparation = 0;
    for (var i = 0; i < translations.length; i++) {
      final scannerPos1 = translations[i];
      for (var j = i + 1; j < translations.length; j++) {
        final scannerPos2 = translations[j];
        final separation = (scannerPos1.x - scannerPos2.x).abs() +
            (scannerPos1.y - scannerPos2.y).abs() +
            (scannerPos1.z - scannerPos2.z).abs();
        if (separation > maxSeparation) {
          maxSeparation = separation;
        }
      }
    }

    return maxSeparation.toString();
  }
}

Iterable<Point3D> _determineBeaconsRelativeToScanner0(
    List<List<Point3D>> scannedBeacons) {
  final beaconsRelativeToScanner0 = scannedBeacons[0].toSet();
  final beaconsFromOtherScanners = [...scannedBeacons.sublist(1)];

  while (beaconsFromOtherScanners.isNotEmpty) {
    final determinedBeacons = beaconsRelativeToScanner0.toList();

    TESTING_SCANNERS:
    for (var scanner = 0;
        scanner < beaconsFromOtherScanners.length;
        scanner++) {
      final beaconsFromOtherScanner = beaconsFromOtherScanners[scanner];
      for (var rotationId = 0; rotationId < _nRotationIds; rotationId++) {
        final rotatedBeacons =
            _rotateListOfPoints(beaconsFromOtherScanner, rotationId);
        final translation =
            _tryFindTranslation(determinedBeacons, rotatedBeacons, 12);
        if (translation != null) {
          beaconsFromOtherScanners.removeAt(scanner);
          final translatedBeacons =
              _translateListOfPoints(rotatedBeacons, translation);
          beaconsRelativeToScanner0.addAll(translatedBeacons);
          break TESTING_SCANNERS;
        }
      }
    }
  }

  return beaconsRelativeToScanner0;
}

List<Point3D> _determineTranslationsRelativeToScanner0(
    List<List<Point3D>> scannedBeacons) {
  final translations = <Point3D>[Point3D(0, 0, 0)];
  final beaconsRelativeToScanner0 = scannedBeacons[0].toSet();
  final beaconsFromOtherScanners = [...scannedBeacons.sublist(1)];

  while (beaconsFromOtherScanners.isNotEmpty) {
    final determinedBeacons = beaconsRelativeToScanner0.toList();

    TESTING_SCANNERS:
    for (var scanner = 0;
        scanner < beaconsFromOtherScanners.length;
        scanner++) {
      final beaconsFromOtherScanner = beaconsFromOtherScanners[scanner];
      for (var rotationId = 0; rotationId < _nRotationIds; rotationId++) {
        final rotatedBeacons =
            _rotateListOfPoints(beaconsFromOtherScanner, rotationId);
        final translation =
            _tryFindTranslation(determinedBeacons, rotatedBeacons, 12);
        if (translation != null) {
          translations.add(translation);
          beaconsFromOtherScanners.removeAt(scanner);
          final translatedBeacons =
              _translateListOfPoints(rotatedBeacons, translation);
          beaconsRelativeToScanner0.addAll(translatedBeacons);
          break TESTING_SCANNERS;
        }
      }
    }
  }

  return translations;
}

Iterable<Point3D> _rotateListOfPoints(List<Point3D> list, int rotationId) {
  return list.map((p) => _rotatePoint3D(p, rotationId)).toList();
}

Iterable<Point3D> _translateListOfPoints(
    Iterable<Point3D> list, Point3D translation) {
  return list.map((p) =>
      Point3D(p.x + translation.x, p.y + translation.y, p.z + translation.z));
}

const _nRotationIds = 24;

Point3D _rotatePoint3D(Point3D p, int rotationId) {
  switch (rotationId) {
    case 0:
      return Point3D(p.x, p.y, p.z);
    case 1:
      return Point3D(-p.x, -p.y, p.z);
    case 2:
      return Point3D(-p.x, p.y, -p.z);
    case 3:
      return Point3D(p.x, -p.y, -p.z);
    case 4:
      return Point3D(-p.x, p.z, p.y);
    case 5:
      return Point3D(p.x, -p.z, p.y);
    case 6:
      return Point3D(p.x, p.z, -p.y);
    case 7:
      return Point3D(-p.x, -p.z, -p.y);
    case 8:
      return Point3D(-p.y, p.x, p.z);
    case 9:
      return Point3D(p.y, -p.x, p.z);
    case 10:
      return Point3D(p.y, p.x, -p.z);
    case 11:
      return Point3D(-p.y, -p.x, -p.z);
    case 12:
      return Point3D(p.y, p.z, p.x);
    case 13:
      return Point3D(-p.y, -p.z, p.x);
    case 14:
      return Point3D(-p.y, p.z, -p.x);
    case 15:
      return Point3D(p.y, -p.z, -p.x);
    case 16:
      return Point3D(p.z, p.x, p.y);
    case 17:
      return Point3D(-p.z, -p.x, p.y);
    case 18:
      return Point3D(-p.z, p.x, -p.y);
    case 19:
      return Point3D(p.z, -p.x, -p.y);
    case 20:
      return Point3D(-p.z, p.y, p.x);
    case 21:
      return Point3D(p.z, -p.y, p.x);
    case 22:
      return Point3D(p.z, p.y, -p.x);
    case 23:
      return Point3D(-p.z, -p.y, -p.x);
  }
  throw Error();
}

Point3D? _tryFindTranslation(Iterable<Point3D> determinedBeacons,
    Iterable<Point3D> otherBeacons, int requiredOverlapCount) {
  final diffs = <Point3D, int>{};

  for (var determinedBeacon in determinedBeacons) {
    for (var otherBeacon in otherBeacons) {
      final diff = Point3D(
          determinedBeacon.x - otherBeacon.x,
          determinedBeacon.y - otherBeacon.y,
          determinedBeacon.z - otherBeacon.z);
      diffs.update(diff, (value) => value + 1, ifAbsent: () => 1);
    }
  }

  for (var entry in diffs.entries) {
    final translation = entry.key;
    final count = entry.value;
    if (count >= requiredOverlapCount) {
      return translation;
    }
  }

  return null;
}
