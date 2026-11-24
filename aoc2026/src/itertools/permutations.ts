import { Stack } from "../queue"

export function* combinations<T>(elements: T[]): IterableIterator<T[]> {
  const stop = (1 << elements.length) - 1
  for (let s = 0; s <= stop; s += 1) {
    const combination = []
    for (let i = 0; i < elements.length; i += 1) {
      if (((s >> i) & 1) === 1) {
        combination.push(elements[i] as T)
      }
    }
    yield combination
  }
}

export function* permutations<T>(elements: T[]): IterableIterator<T[]> {
  for (const combination of combinations(elements)) {
    for (const permutation of permute(combination)) {
      yield permutation
    }
  }
}

export function* permute<T>(elements: T[]): IterableIterator<T[]> {
  const stack = new Stack<number[]>([[]].values())
  generate: while (stack.size() > 0) {
    const membersOption = stack.pop()
    if (membersOption.isNone()) {
      continue
    }

    const members = membersOption.value()

    if (members.length === elements.length) {
      const permutation: T[] = []
      for (let i = 0; i < members.length; i += 1) {
        permutation.push(elements[members[i] as number] as T)
      }
      yield permutation

      continue generate
    }

    extend: for (let i = 0; i < elements.length; i += 1) {
      if (members.includes(i)) {
        continue extend
      }

      const clone = [...members].concat(i)
      if (clone.length < elements.length) {
        stack.push(clone)
        continue extend
      }

      const permutation: T[] = []
      for (let i = 0; i < clone.length; i += 1) {
        permutation.push(elements[clone[i] as number] as T)
      }
      yield permutation
    }
  }
}
