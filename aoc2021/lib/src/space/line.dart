import 'dart:math';

import 'package:aoc2021/src/space/point.dart';

/// Assumes that the provided points form a line along a slope that is a
/// multiple of 45 degrees.
class Line45Degree {
  final Point start;
  final Point end;

  Line45Degree(int x1, int y1, int x2, int y2)
      : start = Point(x1, y1),
        end = Point(x2, y2);

  bool get isHorizontal => start.y == end.y;

  bool get isVertical => start.x == end.x;

  List<Point> getCoveredPoints() {
    final coveredPoints = <Point>[];
    if (isVertical) {
      for (var y = min(start.y, end.y); y <= max(start.y, end.y); y++) {
        coveredPoints.add(Point(start.x, y));
      }
    } else if (isHorizontal) {
      for (var x = min(start.x, end.x); x <= max(start.x, end.x); x++) {
        coveredPoints.add(Point(x, start.y));
      }
    } else {
      // We can cover the points moving left to right and also moving up or
      // down depenging on the slope.
      var leftmostPoint = start;
      var verticalIncrement = start.y < end.y ? 1 : -1;
      if (end.x < start.x) {
        leftmostPoint = end;
        verticalIncrement = end.y < start.y ? 1 : -1;
      }

      final length = (end.x - start.x).abs();

      for (var i = 0; i <= length; i++) {
        coveredPoints.add(Point(
            leftmostPoint.x + i, leftmostPoint.y + (i * verticalIncrement)));
      }
    }

    return coveredPoints;
  }
}
