type Enumerate<
  N extends number,
  Acc extends number[] = [],
> = Acc["length"] extends N
  ? Acc[number]
  : Enumerate<N, [...Acc, Acc["length"]]>

export type IntRangeInclusive<Start extends number, End extends number> =
  | Exclude<Enumerate<End>, Enumerate<Start>>
  | End
