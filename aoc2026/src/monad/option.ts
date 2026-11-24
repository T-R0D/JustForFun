import { Result } from "./result"

type Some<T> = {
  readonly kind: "some"
  value: T
}

type None = {
  readonly kind: "none"
}

/** An `Option` known to hold a value. */
type SomeOption<T> = OptionImpl<T> & Some<T>

/** An `Option` known to hold nothing. */
type NoneOption<T> = OptionImpl<T> & None

export type Option<T> = SomeOption<T> | NoneOption<T>

export class OptionImpl<T> {
  /** Construct a filled Option. */
  static some<T>(value: T): SomeOption<T> {
    return new OptionImpl("some", value) as SomeOption<T>
  }

  /** Construct an empty Option. */
  static none(): NoneOption<never> {
    return new OptionImpl("none", undefined) as NoneOption<never>
  }

  readonly kind: "some" | "none"
  readonly #value: T | undefined

  private constructor(kind: "some" | "none", value: T | undefined) {
    this.kind = kind
    this.#value = value
  }

  isSome(): this is { readonly kind: "some" } {
    return this.kind == "some"
  }

  isNone(): this is { readonly kind: "none" } {
    return this.kind == "none"
  }

  value(this: SomeOption<T>): T {
    return this.#value as T
  }

  unwrap(this: Option<T>): T {
    if (!this.isSome()) {
      throw new Error("option could not be unwrapped")
    }

    return this.value()
  }

  unwrapOrElse<U>(this: Option<T>, f: () => U): T | U {
    if (!this.isSome()) {
      return f()
    }

    return this.value()
  }

  map<T2>(this: Option<T>, f: (t: T) => T2): Option<T2> {
    if (this.isNone()) {
      return OptionImpl.none()
    }

    return OptionImpl.some(f(this.value()))
  }

  flatMap<T2>(this: Option<T>, f: (t: T) => Option<T2>): Option<T2> {
    if (this.isNone()) {
      return OptionImpl.none()
    }

    return f(this.value())
  }

  toResult<E>(this: Option<T>, ifNone: () => E): Result<T, E> {
    if (this.isNone()) {
      return Result.err(ifNone())
    }

    return Result.ok(this.value())
  }
}

export const Option = OptionImpl
