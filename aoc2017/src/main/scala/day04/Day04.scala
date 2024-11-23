package day04

import solution.Solution

class Day04Solution extends Solution:
    override def partOne(input: String): Either[String, String] =
        val candidates = parsePassphrases(input)

        val nValidPassphrases = candidates.filter(isValidPassphrase(_)).size

        Right(nValidPassphrases.toString)

    override def partTwo(input: String): Either[String, String] =
        val candidates = parsePassphrases(input)

        val nValidPassphrases =candidates.filter(isExtraValidPassphrase(_)).size

        Right(nValidPassphrases.toString)
        
    def parsePassphrases(input: String): Seq[Seq[String]] =
        input.split("\\n+").toSeq.map { line =>
            line.split("\\s+").toSeq
        }

    def isValidPassphrase(candidate: Seq[String]): Boolean =
        val counts = candidate.groupBy(identity).mapValues(_.size)
        return counts.values.forall(_ == 1)

    def isExtraValidPassphrase(candidate: Seq[String]): Boolean =
        val letterCounts = candidate.map{ word =>
            word.groupBy(identity).mapValues(_.size).toMap
        }.toSeq

        (
            for
                (countsA, i) <- letterCounts.zipWithIndex
                (countsB, j) <- letterCounts.zipWithIndex
                    if (i != j)
            yield
                val mostUniqueLetters = Math.max(countsA.size, countsB.size)
                countsA.toSet.intersect(countsB.toSet).size < mostUniqueLetters
        ).forall(identity)

        
