import { Dequeue } from "../queue"
import { range } from "../utils"

type ZipOptions<Ts extends Iterable<unknown>[]> =
  | {
      mode: "shortest"
    }
  | {
      mode: "longest"
      fills: NoInfer<Zipped<Ts>>
    }

type Zipped<Ts extends Iterable<unknown>[]> = {
  [K in keyof Ts]: Ts[K] extends Iterable<infer U> ? U : never
}

export function* zip<Ts extends Iterable<unknown>[]>(
  iterables: [...Ts],
  options: ZipOptions<Ts> = { mode: "shortest" },
): IteratorObject<Zipped<Ts>> {
  if (iterables.length === 0) {
    return
  }

  const iterators = iterables.map((iterable) => iterable[Symbol.iterator]())

  let collectTuple: (iterators: Iterator<unknown>[]) => [unknown[], boolean]
  if (options.mode === "shortest") {
    collectTuple = collectTupleShortest
  } else {
    collectTuple = (iterators: Iterator<unknown>[]): [unknown[], boolean] =>
      collectTupleLongest(iterators, options.fills)
  }

  try {
    while (true) {
      const [tuple, ok] = collectTuple(iterators)
      if (!ok) {
        return
      }
      yield tuple as Zipped<Ts>
    }
  } finally {
    iterators.forEach((iterator) => iterator.return?.())
  }
}

function collectTupleShortest(
  iterators: Iterator<unknown>[],
): [unknown[], boolean] {
  const tuple: unknown[] = []
  for (const iterator of iterators) {
    const next = iterator.next()
    if (next.done) {
      return [tuple, false]
    }
    tuple.push(next.value)
  }
  return [tuple, true]
}

function collectTupleLongest(
  iterators: Iterator<unknown>[],
  fills: unknown[],
): [unknown[], boolean] {
  const tuple: unknown[] = []
  let allDone = true
  for (let i = 0; i < iterators.length; i += 1) {
    const next = iterators[i]!.next()
    if (next.done) {
      tuple.push(fills[i])
      continue
    }
    allDone = false
    tuple.push(next.value)
  }
  return [tuple, !allDone]
}

export function* chunks<T>(
  iterable: Iterable<T>,
  n: number,
): IteratorObject<T[]> {
  const iterator = iterable[Symbol.iterator]()
  while (true) {
    const chunk: T[] = []
    for (const _ of range({ stop: n })) {
      const next = iterator.next()
      if (next.done) {
        if (chunk.length > 0) {
          yield chunk
        }
        return
      }
      chunk.push(next.value)
    }
    yield chunk
  }
}

export function* enumerate<T>(
  iterable: Iterable<T>,
): IteratorObject<[number, T], undefined> {
  let i = 0
  for (const x of iterable) {
    yield [i, x]
    i += 1
  }
}

export function* windows<T>(
  iterable: Iterable<T>,
  size: number,
  shiftBy: number = 1,
): IteratorObject<T[]> {
  const iterator = iterable[Symbol.iterator]()
  const window = new Dequeue<T>()
  for (const _ of range({ stop: size })) {
    const next = iterator.next()
    if (next.done) {
      break
    }
    window.push(next.value)
  }

  if (window.size() === 0) {
    return
  }

  yield Array.from(window)

  while (true) {
    for (const _ of range({ stop: shiftBy })) {
      window.pop()
    }

    for (const _ of range({ stop: shiftBy })) {
      const next = iterator.next()
      if (next.done) {
        if (window.size() > 0) {
          yield Array.from(window)
        }
        return
      }
      window.push(next.value)
    }

    yield Array.from(window)
  }
}
