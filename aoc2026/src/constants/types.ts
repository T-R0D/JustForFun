import type { IntRangeInclusive } from "../utils"
import { ADVENT_DAYS } from "./values"

export type AdventDayNumber = IntRangeInclusive<1, typeof ADVENT_DAYS>

export type AdventPartNumber = 1 | 2
