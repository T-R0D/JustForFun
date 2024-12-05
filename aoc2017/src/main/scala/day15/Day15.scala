package day15

import solution.Solution

class Day15Solution extends Solution:
    override def partOne(input: String): Either[String, String] = {
        val startingValues = parseStartingValues(input)

        val lower16Matches = countLower16MatchesInGeneration(startingValues, 40_000_000)

        Right(lower16Matches.toString)
    }

    override def partTwo(input: String): Either[String, String] = ???

    def parseStartingValues(input: String): (Long, Long) = {
        val startingValues = {
            for {
                line <- input.split("\\n")
            } yield {
                line.split("starts with ").takeRight(1).map(_.toLong).toSeq(0)   
            }
        }
        (startingValues(0), startingValues(1))
    }

    def countLower16MatchesInGeneration(startingValues: (Long, Long), n: Int): Int = {
        val multiplicationFactors: (Long, Long) = (16807, 48271)
        val mod: Long = 2147483647
        val lower16Mask: Long = (1 << 16) - 1
        val (_, lower16MatchesFound) = (0 until n).foldLeft((startingValues, 0)) { (acc, _) =>
            val (previousValues, lower16MatchesFound) = acc
            
            val a = (previousValues._1 * multiplicationFactors._1) % mod
            val b = (previousValues._2 * multiplicationFactors._2) % mod

            val nextValues = (a, b)
            {
                if (a & lower16Mask) == (b & lower16Mask) then
                    (nextValues, lower16MatchesFound + 1)
                else
                    (nextValues, lower16MatchesFound)
            }
        }

        lower16MatchesFound
    }
