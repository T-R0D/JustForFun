export function* range({
  start = 0,
  step = 1,
  stop,
}: {
  start?: number
  step?: number
  stop: number
}): Generator<number> {
  for (let i = start; i < stop; i += step) {
    yield i
  }
}
