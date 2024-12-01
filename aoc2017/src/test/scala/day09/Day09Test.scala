package day09

import org.scalatest.funsuite.AnyFunSuite
import org.scalatest.matchers.should.*
import org.scalatest.prop.TableDrivenPropertyChecks.*

import day09.Day09Solution

class Day09SolutionTests extends AnyFunSuite with Matchers:
    test("partOne scores streams correctly") {
        val testCases = Table(
            ("input", "expected"),
            ("{}", "1"),
            ("{{{}}}", "6"),
            ("{{},{}}", "5"),
            ("{{{},{},{{}}}}", "16"),
            ("{<a>,<a>,<a>,<a>}", "1"),
            ("{{<ab>},{<ab>},{<ab>},{<ab>}}", "9"),
            ("{{<!!>},{<!!>},{<!!>},{<!!>}}", "9"),
            ("{{<a!>},{<a!>},{<a!>},{<ab>}}", "3"),
        )

        forAll(testCases) { (input: String, expected: String) => 
            val solution = Day09Solution()

            val result = solution.partOne(input)

            result should equal (Right(expected))    
        }
    }

    test("partTwo counts garbage correctly") {
        val testCases = Table(
            ("input", "expected"),
            ("<>", "0"),
            ("<random characters>", "17"),
            ("<<<<>", "3"),
            ("<{!>}>", "2"),
            ("<!!>", "0"),
            ("<!!!>>", "0"),
            ("<{o\"i!a,<{i<a>", "10"),
        )

        forAll(testCases) { (input: String, expected: String) => 
            val solution = Day09Solution()

            val result = solution.partTwo(input)

            result should equal (Right(expected))    
        }
    }
        