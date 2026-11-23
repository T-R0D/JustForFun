import { Option } from "../monad"

const DEFAULT_UNDERLYING_SIZE = 100

export class Dequeue<T> {
  #data: (T | null)[]
  #start: number
  #stop: number
  #size: number

  constructor(initialElements?: Iterable<T>) {
    this.#data =
      initialElements !== undefined
        ? Array.from(initialElements)
        : Array(DEFAULT_UNDERLYING_SIZE).fill(null)
    this.#start = 0
    this.#stop = initialElements !== undefined ? this.#data.length : 0
    this.#size = initialElements !== undefined ? this.#data.length : 0
  }

  push(x: T) {
    if (this.size() < this.#data.length) {
      this.#data[this.#stop % this.#data.length] = x
      this.#stop = (this.#stop + 1) % this.#data.length
      this.#size += 1
      return
    }

    const newData = Array.from(this.iter())
    this.#data = newData

    this.#data.push(x)
    this.#start = 0
    this.#stop = this.#data.length
    this.#size = this.#data.length
  }

  pushFront(x: T) {
    if (this.size() < this.#data.length) {
      const newStart = (this.#data.length + this.#start - 1) % this.#data.length
      this.#data[newStart] = x
      this.#start = newStart
      this.#size += 1
      return
    }

    // It might be likely that this deque will be filling values from the
    // front, so let's allocate some extra space up front to reduce
    // re-allocations of underlying memory.
    const oldSize = this.size()
    const newData = [...Array(this.size()).fill(null), x, ...this.iter()]
    this.#data = newData

    // The push is just to make sure we resize the underlying memory allocation if necessary.
    this.#data.push()
    this.#start = oldSize
    this.#stop = oldSize + oldSize + 1
    this.#size += 1
  }

  pop(): Option<T> {
    if (this.size() == 0) {
      return Option.none()
    }

    const x = this.#data[this.#start] as T
    this.#data[this.#start] = null
    this.#start = (this.#start + this.#data.length + 1) % this.#data.length
    this.#size -= 1
    return Option.some(x)
  }

  popBack(): Option<T> {
    if (this.size() == 0) {
      return Option.none()
    }

    this.#stop = (this.#stop + this.#data.length - 1) % this.#data.length
    const x = this.#data[this.#stop] as T
    this.#data[this.#stop] = null
    this.#size -= 1
    return Option.some(x)
  }

  peek(): Option<T> {
    if (this.size() === 0) {
      return Option.none()
    }

    const x = this.#data[this.#start] as T
    return Option.some(x)
  }

  peekBack(): Option<T> {
    if (this.size() === 0) {
      return Option.none()
    }

    const x = this.#data[
      (this.#stop + this.#data.length - 1) % this.#data.length
    ] as T
    return Option.some(x)
  }

  size(): number {
    return this.#size
  }

  /** Iterates the dequeue in left to right order. */
  *iter(): Iterable<T> {
    for (let i = 0; i < this.size(); i += 1) {
      yield this.#data[
        (this.#start + this.#data.length + i) % this.#data.length
      ] as T
    }
  }
}
