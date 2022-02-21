int strToBin(String str) {
  int ret = 0;
  for (var i = 0; i < str.length; i++) {
    final bit = str[i];
    if (!['0', '1'].contains(bit)) {
      throw StateError('$bit is not a proper bit value');
    }

    ret <<= 1;
    if (bit == '1') {
      ret |= 1;
    }
  }
  return ret;
}
