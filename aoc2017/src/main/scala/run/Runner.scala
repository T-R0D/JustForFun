package run

import scala.io.{Source}
import scala.util.{Failure, Success, Try}

import args.Config
import day01.{Day01Solution}
import day02.{Day02Solution}
import day03.{Day03Solution}
import day04.{Day04Solution}
import day05.{Day05Solution}
import day06.{Day06Solution}
import day07.{Day07Solution}
import day08.{Day08Solution}
import day09.{Day09Solution}
import day10.{Day10Solution}
import day11.{Day11Solution}
import day12.{Day12Solution}
import day13.{Day13Solution}
import day14.{Day14Solution}
import day15.{Day15Solution}
import day16.{Day16Solution}
import day17.{Day17Solution}
import day18.{Day18Solution}
import day19.{Day19Solution}
import day20.{Day20Solution}
import day21.{Day21Solution}
import day22.{Day22Solution}
import day23.{Day23Solution}
import day24.{Day24Solution}
import day25.{Day25Solution}
import solution.{Solution}

class Runner:
    def run(config: Config): Unit =
        val result = Try {
            val input = Source.fromFile(config.input).mkString.stripLineEnd

            val solution: Solution = config.day match
                case 1 => Day01Solution()
                case 2 => Day02Solution()
                case 3 => Day03Solution()
                case 4 => Day04Solution()
                case 5 => Day05Solution()
                case 6 => Day06Solution()
                case 7 => Day07Solution()
                case 8 => Day08Solution()
                case 9 => Day09Solution()
                case 10 => Day10Solution()
                case 11 => Day11Solution()
                case 12 => Day12Solution()
                case 13 => Day13Solution()
                case 14 => Day14Solution()
                case 15 => Day15Solution()
                case 16 => Day16Solution()
                case 17 => Day17Solution()
                case 18 => Day18Solution()
                case 19 => Day19Solution()
                case 20 => Day20Solution()
                case 21 => Day21Solution()
                case 22 => Day22Solution()
                case 23 => Day23Solution()
                case 24 => Day24Solution()
                case 25 => Day25Solution()
            
            val solve = config.part match
                case 1 => solution.partOne
                case 2 => solution.partTwo

            val start = System.nanoTime()
            val answer = solve(input)
            val stop = System.nanoTime()

            (answer, stop - start)
        }
        
        result match
            case Success((Right(answer), executionTimeNanos)) =>
                println(s"""$answer
                           |(${executionTimeNanos / 1000}us)
                         """.stripMargin)
            case Success((Left(msg), executionTimeNanos)) =>
                println(s"Failure: $msg (${executionTimeNanos}ns)")
            case Failure(exc) =>
                println(s"Failure: ${exc.toString}")
