package day15

import solution.Solution

class Day15Solution extends Solution:
    override def partOne(input: String): Either[String, String] = {
        val seeds = parseSeedValues(input)

        val lower16Matches = countLower16MatchesInGeneration(seeds, (1, 1), 40_000_000)

        Right(lower16Matches.toString)
    }

    override def partTwo(input: String): Either[String, String] = {
        val seeds = parseSeedValues(input)

        val lower16Matches = countLower16MatchesInGeneration(seeds, (4, 8), 5_000_000)

        Right(lower16Matches.toString)
    }

    def parseSeedValues(input: String): (Long, Long) = {
        val startingValues = {
            for {
                line <- input.split("\\n")
            } yield {
                line.split("starts with ").takeRight(1).map(_.toLong).toSeq(0)   
            }
        }
        (startingValues(0), startingValues(1))
    }

    def countLower16MatchesInGeneration(
        seeds: (Long, Long),
        multipleRequirements: (Long, Long),
        n: Int,
    ): Int = {
        val (seedA, seedB) = seeds
        val (multiplicationFactorA, multiplicationFactorB) = (16807l, 48271l)
        val mod = 2147483647l
        val (multipleRequirementA, multipleRequirementB) = multipleRequirements

        val (generatorA, generatorB) = (
            Generator.newGenerator(seedA, multiplicationFactorA, mod, multipleRequirementA),
            Generator.newGenerator(seedB, multiplicationFactorB, mod, multipleRequirementB),
        )

        val lower16Mask: Long = (1 << 16) - 1
        val matchesFound = (0 until n).foldLeft(0) { (acc, _) =>
            val valA = generatorA.next()
            val valB = generatorB.next()

            val nextValues = (valA, valB)
            {
                if (valA & lower16Mask) == (valB & lower16Mask) then
                    acc + 1
                else
                    acc
            }
        }

        matchesFound
    }

    class Generator private (
        value: Long,
        var sequence: LazyList[Long],
        multipleRequirement: Long,
    ) {
        def hasNext(): Boolean = {
            true
        }

        def next(): Long = {
            sequence = sequence.tail.dropWhile(x => x % multipleRequirement != 0)
            sequence.head
        }
    }

    object Generator {
        def newGenerator(
            seed: Long,
            multiplicationFactor: Long,
            mod: Long,
            multipleRequirement: Long,
        ): Generator = {
            val nextFn = newNextNumberFn(multiplicationFactor, mod)
            val sequence = LazyList.iterate(seed)(nextFn)

            Generator(sequence.head, sequence, multipleRequirement)
        }

        private def newNextNumberFn(multiplicationFactor: Long, mod: Long): (Long => Long) = {
            (previous: Long) => {
                previous * multiplicationFactor % mod
            }
        }
    }
