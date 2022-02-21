class Point {
  int x, y;

  Point(this.x, this.y);

  @override
  String toString() {
    return '($x, $y)';
  }

  @override
  bool operator ==(Object other) {
    return other is Point && x == other.x && y == other.y;
  }

  @override
  int get hashCode => toString().hashCode;
}

class Point3D {
  int x, y, z;

  Point3D(this.x, this.y, this.z);

  @override
  String toString() {
    return '($x, $y, $z)';
  }

  @override
  bool operator ==(Object other) {
    return other is Point3D && x == other.x && y == other.y && z == other.z;
  }

  @override
  int get hashCode => toString().hashCode;
}
