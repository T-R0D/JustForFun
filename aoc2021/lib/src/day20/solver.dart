import 'package:aoc2021/src/solver/solver.dart' show Solver;
import 'package:aoc2021/src/space/point.dart';

class Day20Solver implements Solver {
  List<String> _enhancementAlgorithm = [];
  List<List<String>> _initialImageRegion = [];

  @override
  void consumeRawInput(String rawInput) {
    final parts = rawInput.split('\n\n');

    _enhancementAlgorithm = parts[0].split('');

    _initialImageRegion =
        parts[1].split('\n').map((line) => line.split('')).toList();
  }

  @override
  String solvePart1() {
    final image = _Image(_enhancementAlgorithm, _initialImageRegion);
    for (var i = 0; i < 2; i++) {
      image.enhance();
    }
    return image.litPixelCount.toString();
  }

  @override
  String solvePart2() {
    final image = _Image(_enhancementAlgorithm, _initialImageRegion);
    for (var i = 0; i < 50; i++) {
      image.enhance();
    }
    return image.litPixelCount.toString();
  }
}

class _Bounds {
  final int i0;
  final int i1;
  final int j0;
  final int j1;

  _Bounds(this.i0, this.i1, this.j0, this.j1);
}

class _Image {
  final List<String> _enhancementAlgorithm;
  Set<Point> _litPixels = {};
  _Bounds _bounds = _Bounds(1, 0, 1, 0);
  String _backgroud = '.';

  _Image(this._enhancementAlgorithm, List<List<String>> initialImageSample) {
    for (var i = 0; i < initialImageSample.length; i++) {
      for (var j = 0; j < initialImageSample[i].length; j++) {
        final pixelValue = initialImageSample[i][j];
        if (pixelValue == '#') {
          _litPixels.add(Point(i, j));
        }
      }
    }

    _bounds = _Bounds(
        0, initialImageSample.length - 1, 0, initialImageSample[0].length - 1);
    _backgroud = '.';
  }

  void enhance() {
    final imageSection = [
      Point(-1, -1),
      Point(-1, 0),
      Point(-1, 1),
      Point(0, -1),
      Point(0, 0),
      Point(0, 1),
      Point(1, -1),
      Point(1, 0),
      Point(1, 1),
    ];
    final bounds = _findBounds();
    final newLitPixels = <Point>{};
    final newBounds =
        _Bounds(bounds.i0 - 1, bounds.i1 + 1, bounds.j0 - 1, bounds.j1 + 1);

    for (var i = newBounds.i0; i <= newBounds.i1; i++) {
      for (var j = newBounds.j0; j <= newBounds.j1; j++) {
        var enhancementIndex = 0;
        for (var z = 0; z < imageSection.length; z++) {
          final p = imageSection[z];
          final pixel = Point(p.x + i, p.y + j);

          enhancementIndex <<= 1;
          if (pixel.x <= newBounds.i0 ||
              newBounds.i1 <= pixel.x ||
              pixel.y <= newBounds.j0 ||
              newBounds.j1 <= pixel.y) {
            if (_backgroud == '#') {
              enhancementIndex |= 1;
            }
          } else if (_litPixels.contains(pixel)) {
            enhancementIndex |= 1;
          }
        }

        final value = _enhancementAlgorithm[enhancementIndex];
        if (value == '#') {
          newLitPixels.add(Point(i, j));
        }
      }
    }

    _litPixels = newLitPixels;
    _backgroud = _backgroud == '#'
        ? _enhancementAlgorithm[0x1FF]
        : _enhancementAlgorithm[0];
    _bounds = newBounds;
  }

  int get litPixelCount => _litPixels.length;

  @override
  String toString() {
    final buffer = StringBuffer();
    final bounds = _findBounds();

    for (var i = bounds.i0; i <= bounds.i1; i++) {
      for (var j = bounds.j0; j <= bounds.j1; j++) {
        if (_litPixels.contains(Point(i, j))) {
          buffer.write('#');
        } else {
          buffer.write(' ');
        }
      }
      buffer.write('\n');
    }

    return buffer.toString();
  }

  _Bounds _findBounds() {
    return _bounds;
    // // TODO: Replace with MAX_INT.
    // int i0 = 99999;
    // int i1 = 0;
    // int j0 = 99999;
    // int j1 = 0;

    // for (var litPixel in _litPixels) {
    //   if (litPixel.x < i0) {
    //     i0 = litPixel.x;
    //   }
    //   if (litPixel.x > i1) {
    //     i1 = litPixel.x;
    //   }
    //   if (litPixel.y < j0) {
    //     j0 = litPixel.y;
    //   }
    //   if (litPixel.y > j1) {
    //     j1 = litPixel.y;
    //   }
    // }

    // return _Bounds(i0, i1, j0, j1);
  }
}
