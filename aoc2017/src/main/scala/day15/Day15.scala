package day15

import solution.Solution

class Day15Solution extends Solution:
    override def partOne(input: String): Either[String, String] = {
        val seeds = parseSeedValues(input)

        val lower16Matches = countLower16MatchesInGeneration(seeds, 40_000_000)

        Right(lower16Matches.toString)
    }

    override def partTwo(input: String): Either[String, String] = ???

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

    def countLower16MatchesInGeneration(seeds: (Long, Long), n: Int): Int = {
        val (seedA, seedB) = seeds
        val (multiplicationFactorA, multiplicationFactorB) = (16807l, 48271l)
        val mod = 2147483647l

        val (generatorA, generatorB) = (
            Generator.newGenerator(seed = seedA, multiplicationFactor = multiplicationFactorA, mod = mod, multipleRequirement = 1),
            Generator.newGenerator(seed = seedB, multiplicationFactor = multiplicationFactorB, mod = mod, multipleRequirement = 1),
        )

        val lower16Mask: Long = (1 << 16) - 1
        val (_, _, matchesFound) = (0 until n).foldLeft((generatorA, generatorB, 0)) { (acc, _) =>
            val (generatorA, generatorB, matchesFound) = acc
            
            val (valA, nextGeneratorA) = generatorA.next()
            val (valB, nextGeneratorB) = generatorB.next()

            val nextValues = (valA, valB)
            val newMatchesFound = {
                if (valA & lower16Mask) == (valB & lower16Mask) then
                    matchesFound + 1
                else
                    matchesFound
            }

            (nextGeneratorA, nextGeneratorB, newMatchesFound)
        }

        matchesFound
    }

    class Generator private (
        value: Long,
        sequence: LazyList[Long],
        multipleRequirement: Long,
    ) {
        def hasNext(): Boolean = {
            true
        }

        def next(): (Long, Generator) = {
            val advancedSequence = sequence.drop(1).dropWhile(x => x % multipleRequirement != 0)
            val nextPublishedValue = advancedSequence.head

            (nextPublishedValue, Generator(nextPublishedValue, advancedSequence, multipleRequirement))
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
            def f(previous: Long): Long = {
                previous * multiplicationFactor % mod
            }

            f
        }
    }
