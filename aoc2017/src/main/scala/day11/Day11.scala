package day11

import solution.Solution

class Day11Solution extends Solution:
    override def partOne(input: String): Either[String, String] =
        for 
            path <- parseHexTilePath(input)
        yield
            tracePathToFindShortestPathDistance(path).toString

    override def partTwo(input: String): Either[String, String] = ???

    def parseHexTilePath(input: String): Either[String, Seq[HexDirection]] =
        val results = input.split(",").map(HexDirection.fromString)
        results.foldLeft(Right(Seq.empty): Either[String, Seq[HexDirection]]) { (acc, result) =>
            for
                steps <- acc
                nextStep <- result
            yield
                steps :+ nextStep
        }

    enum HexDirection(di: Int, dj: Int):
        case North extends HexDirection(-2, 0)
        case NorthEast extends HexDirection(-1, 1)
        case SouthEast extends HexDirection(1, 1)
        case South extends HexDirection(2, 0)
        case SouthWest extends HexDirection(1, -1)
        case NorthWest extends HexDirection(-1, -1)

        def toDelta(): (Int, Int) = (di, dj)

    object HexDirection:
        def fromString(str: String): Either[String, HexDirection] =
            str match
                case "n" => Right(HexDirection.North)
                case "ne" => Right(HexDirection.NorthEast)
                case "se" => Right(HexDirection.SouthEast)
                case "s" => Right(HexDirection.South)
                case "sw" => Right(HexDirection.SouthWest)
                case "nw" => Right(HexDirection.NorthWest)
                case _ => Left(s"'$str' is not a hex tile direction")

    def tracePathToFindShortestPathDistance(path: Seq[HexDirection]): Int =
        val resultingLocation = path.foldLeft((0, 0)) { (currentLocation, nextStep) =>
            val delta = nextStep.toDelta()
            (currentLocation._1 + delta._1, currentLocation._2 + delta._2)
        }
        hexagonalManhattanDistance((0, 0), resultingLocation)

    def hexagonalManhattanDistance(a: (Int, Int), b: (Int, Int)): Int =
        val verticalDiff = Math.abs(a._1 - b._2)
        val horizontalDiff = Math.abs(a._2 - b._2)
        val horizontalSteps = horizontalDiff
        val remainingVerticalDistance = Math.max(1, Math.abs(verticalDiff - horizontalSteps))
        val straightVerticalSteps = remainingVerticalDistance / 2

        println(s"$a -> $b = ${horizontalSteps + straightVerticalSteps}")
        println(s"\tremaining vertical distance: $remainingVerticalDistance")
        println(s"\t$horizontalSteps $straightVerticalSteps")

        horizontalSteps + straightVerticalSteps

