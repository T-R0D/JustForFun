import { Result } from "./result"

type Some<T> = {
  kind: "some"
  value: T
}

type None = {
  kind: "none"
}

/** An `Option` known to hold a value. */
export type SomeOption<T> = OptionImpl<T> & { readonly contents: Some<T> }

/** An `Option` known to hold nothing. */
export type NoneOption<T> = OptionImpl<T> & { readonly contents: None }

export type Option<T> = SomeOption<T> | NoneOption<T>

type O<T> = Some<T> | None

export class OptionImpl<T> {
  /** Construct a filled Option. */
  static some<T>(value: T): SomeOption<T> {
    return new OptionImpl({ kind: "some", value }) as SomeOption<T>
  }

  /** Construct an empty Option. */
  static none(): NoneOption<never> {
    return new OptionImpl({ kind: "none" }) as NoneOption<never>
  }

  readonly contents: O<T>

  private constructor(contents: O<T>) {
    this.contents = contents
  }

  isSome(): this is { readonly contents: Some<T> } {
    return this.contents.kind == "some"
  }

  isNone(): this is { readonly contents: None } {
    return this.contents.kind == "none"
  }

  value(this: { contents: Some<T> }): T {
    return this.contents.value
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
