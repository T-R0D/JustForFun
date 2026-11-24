import { Option } from "../monad"

export class PriorityQueue<T> {
  #data: T[]
  #isHigherPriority: (a: T, b: T) => boolean

  constructor(
    isHigherPriority: (a: T, b: T) => boolean,
    initialElements?: IterableIterator<T>,
  ) {
    this.#data =
      initialElements !== undefined ? Array.from(initialElements) : []
    this.#isHigherPriority = isHigherPriority
  }

  push(x: T) {
    this.#data.push(x)
    this.#bubbleUp()
  }

  pop(): Option<T> {
    if (this.size() == 0) {
      return Option.none()
    }

    const x = this.#data[0] as T
    const newTop = this.#data.pop() as T
    if (this.#data.length > 0) {
      this.#data[0] = newTop
      this.#siftDown()
    }

    return Option.some(x)
  }

  peek(): Option<T> {
    if (this.size() == 0) {
      return Option.none()
    }

    return Option.some(this.#data[0] as T)
  }

  size(): number {
    return this.#data.length
  }

  #bubbleUp() {
    if (this.size() == 0) {
      return
    }

    let i = this.#data.length - 1
    while (i > 0) {
      const p = Math.trunc(i / 2)
      const child = this.#data[i] as T
      const parent = this.#data[p] as T

      if (!this.#isHigherPriority(child, parent)) {
        return
      }

      this.#data[p] = child
      this.#data[i] = parent
      i = p
    }
  }

  #siftDown() {
    let i = 0
    siftLoop: while (i < this.#data.length) {
      let highPriority = i

      const left = i * 2 + 1
      if (
        left < this.#data.length &&
        this.#isHigherPriority(
          this.#data[left] as T,
          this.#data[highPriority] as T,
        )
      ) {
        highPriority = left
      }

      const right = i * 2 + 2
      if (
        right < this.#data.length &&
        this.#isHigherPriority(
          this.#data[right] as T,
          this.#data[highPriority] as T,
        )
      ) {
        highPriority = right
      }

      if (highPriority === i) {
        break siftLoop
      }

      ;[this.#data[i], this.#data[highPriority]] = [
        this.#data[highPriority] as T,
        this.#data[i] as T,
      ]
      i = highPriority
    }
  }
}
