export function sum(numbers: Iterable<number>): number {
  return Iterator.from(numbers).reduce((acc, x) => acc + x, 0)
}

export function product(numbers: Iterable<number>): number {
  return Iterator.from(numbers).reduce((acc, x) => acc * x, 1)
}

export function bigSum(numbers: Iterable<bigint>): bigint {
  return Iterator.from(numbers).reduce((acc, x) => acc + x, 0n)
}

export function bigProduct(numbers: Iterable<bigint>): bigint {
  return Iterator.from(numbers).reduce((acc, x) => acc * x, 1n)
}
