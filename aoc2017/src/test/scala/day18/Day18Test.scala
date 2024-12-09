package day18

import org.scalatest.funsuite.AnyFunSuite
import org.scalatest.matchers.should.*

import day18.Day18Solution

class Day18SolutionTests extends AnyFunSuite with Matchers {
  test("partOne determines the first recovered frequency") {
    val input = """|set a 1
                   |add a 2
                   |mul a a
                   |mod a 5
                   |snd a
                   |set a 0
                   |rcv a
                   |jgz a -1
                   |set a 1
                   |jgz a -2""".stripMargin
    val expected = "4"
    val solution = Day18Solution()

    val result = solution.partOne(input)

    result should equal(Right(expected))
  }

  test("partTwo counts the sends made by program 0") {
    val input = """|snd 1
                   |snd 2
                   |snd p
                   |rcv a
                   |rcv b
                   |rcv c
                   |rcv d""".stripMargin
    val expected = "3"
    val solution = Day18Solution()

    val result = solution.partTwo(input)

    result should equal(Right(expected))
  }
}
