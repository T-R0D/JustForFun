import { Option } from "../monad"

export class Stack<T> {
  readonly #data: T[]

  constructor(initialElements?: Iterable<T>) {
    this.#data =
      initialElements !== undefined ? Array.from(initialElements) : []
  }

  push(x: T) {
    this.#data.push(x)
  }

  pop(): Option<T> {
    const previousSize = this.#data.length
    const x = this.#data.pop()
    return previousSize > 0 ? Option.some(x as T) : Option.none()
  }

  peek(): Option<T> {
    const x = this.#data.at(-1)
    return this.#data.length > 0 ? Option.some(x as T) : Option.none()
  }

  size(): number {
    return this.#data.length
  }

  /** Iterate the elements in reverse order (the order they will leave the stack). */
  *[Symbol.iterator](): Iterable<T> {
    for (let i = this.#data.length - 1; i >= 0; i -= 1) {
      yield this.#data[i] as T
    }
  }
}
