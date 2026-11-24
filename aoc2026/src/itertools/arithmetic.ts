export function sum(numbers: IterableIterator<number>): number {
  return Array.from(numbers).reduce((acc, x) => acc + x, 0)
}

export function product(numbers: IterableIterator<number>): number {
  return Array.from(numbers).reduce((acc, x) => acc * x, 1)
}
