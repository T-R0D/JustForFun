package day11

import org.scalatest.funsuite.AnyFunSuite
import org.scalatest.matchers.should.*
import org.scalatest.prop.TableDrivenPropertyChecks.*

import day11.Day11Solution

class Day11SolutionTests extends AnyFunSuite with Matchers:
    test("partOne finds 'manhattan' distance to end of path") {
        val testCases = Table(
            ("input", "expected"),
            ("ne,ne,sw,sw", "0"),
            ("ne,ne,s,s", "2"),
            ("se,sw,se,sw,sw", "3"),
            ("se,ne", "2"),
            ("n", "1"),
            ("ne,ne", "2"),
            ("sw,nw", "2"),
            ("sw,nw,n", "3"),
        )

        forAll(testCases) { (input, expected) =>
            val solution = Day11Solution()

            val result = solution.partOne(input)

            result should equal (Right(expected))
        }
    }
        