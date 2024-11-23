package day04

import org.scalatest.funsuite.AnyFunSuite
import org.scalatest.matchers.should.*

import day04.Day04Solution

class Day04SolutionTests extends AnyFunSuite with Matchers:
    test("partOne counts valid passphrases") {
        val input = """aa bb cc dd ee
                      |aa bb cc dd aa
                      |aa bb cc dd aaa
                      |""".stripMargin
        val expected = "2"
        val solution = Day04Solution()

        val result = solution.partOne(input)

        result should equal (Right(expected))
    }

    test("partTwo counts extra-valid passphrases") {
        val input = """abcde fghij
                      |abcde xyz ecdab
                      |a ab abc abd abf abj
                      |iiii oiii ooii oooi oooo
                      |oiii ioii iioi iiio
                      |""".stripMargin
        val expected = "3"
        val solution = Day04Solution()

        val result = solution.partTwo(input)

        result should equal (Right(expected))
    }
