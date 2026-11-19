package day24

import scala.util.{Try}

import solution.Solution

class Day24Solution extends Solution:
  override def partOne(input: String): Either[String, String] =
    val sortedSegments = parseInput(input)
    val graph = buildGraph(sortedSegments)
    val strongestStrength = findStrongestBridgeStrength(sortedSegments, graph)
    Right(strongestStrength.toString)

  override def partTwo(input: String): Either[String, String] =
    val sortedSegments = parseInput(input)
    val graph = buildGraph(sortedSegments)
    val strongestStrength =
      findStrongestLongestBridgeStrength(sortedSegments, graph)
    Right(strongestStrength.toString)

  case class Segment(start: Int, end: Int)

  enum Orientation:
    case Forward, Reversed, Nil

  def parseInput(input: String): Seq[Segment] =
    input.linesIterator
      .map(line => {
        val parts = line.split("/").map(_.toInt)
        Segment(parts(0), parts(1))
      })
      .toSeq
      .sortBy(_.start)

  def buildGraph(
      segments: Seq[Segment]
  ): Seq[Seq[(Orientation, Orientation)]] =
    for src <- segments
    yield
      for dst <- segments
      yield
        if src == dst then (Orientation.Nil, Orientation.Nil)
        else
          val nextOrientationIfForward = if src.end == dst.start then
            Orientation.Forward
          else if src.end == dst.end then Orientation.Reversed
          else Orientation.Nil

          val nextOrientationIfReversed = if src.start == dst.start then
            Orientation.Forward
          else if src.start == dst.end then Orientation.Reversed
          else Orientation.Nil

          (nextOrientationIfForward, nextOrientationIfReversed)

  def findStrongestBridgeStrength(
      sortedSegments: Seq[Segment],
      graph: Seq[Seq[(Orientation, Orientation)]]
  ): Int =
    val viableStartSegments =
      sortedSegments.zipWithIndex
        .filter(s => s._1.start == 0 || s._1.end == 0)
        .map(_._2)

    val solutions = for
      srcIndex <- viableStartSegments
      src = sortedSegments(srcIndex)
      orientation =
        if src.start == 0 then Orientation.Forward else Orientation.Reversed
    yield doFindStrongestBridgeStrength(
      sortedSegments,
      graph,
      srcIndex,
      orientation,
      0,
      Set()
    )

    solutions.map(_._1).max

  def doFindStrongestBridgeStrength(
      sortedSegments: Seq[Segment],
      graph: Seq[Seq[(Orientation, Orientation)]],
      nextSrc: Int,
      orientation: Orientation,
      accStrength: Int,
      accUsed: Set[Int]
  ): (Int, Set[Int]) =
    val src = sortedSegments(nextSrc)
    val strength = accStrength + src.start + src.end
    val used = accUsed + nextSrc

    val nextSegmentIndicesAndOrientations = graph(nextSrc)
      .map(options =>
        orientation match
          case Orientation.Forward  => options._1
          case Orientation.Reversed => options._2
          case Orientation.Nil      => throw new Error("uh oh")
      )
      .zipWithIndex
      .filter(x =>
        x._2 != nextSrc && x._1 != Orientation.Nil && !used.contains(x._2)
      )

    nextSegmentIndicesAndOrientations match
      case Seq() =>
        (strength, used)

      case _ =>
        val solutions =
          for (nextOrientation, nextSegmentIndex) <-
              nextSegmentIndicesAndOrientations
          yield doFindStrongestBridgeStrength(
            sortedSegments,
            graph,
            nextSegmentIndex,
            nextOrientation,
            strength,
            used
          )

        solutions.reduce((s1, s2) =>
          if s1._1 > s2._1 then s1
          else s2
        )

  def findStrongestLongestBridgeStrength(
      sortedSegments: Seq[Segment],
      graph: Seq[Seq[(Orientation, Orientation)]]
  ): Int =
    val viableStartSegments =
      sortedSegments.zipWithIndex
        .filter(s => s._1.start == 0 || s._1.end == 0)
        .map(_._2)

    val solutions = for
      srcIndex <- viableStartSegments
      src = sortedSegments(srcIndex)
      orientation =
        if src.start == 0 then Orientation.Forward else Orientation.Reversed
    yield doFindStrongestLongestBridgeStrength(
      sortedSegments,
      graph,
      srcIndex,
      orientation,
      0,
      Set()
    )

    solutions
      .reduce((s1, s2) =>
        if s1._2.size > s2._2.size then s1
        else if s1._2.size < s2._2.size then s2
        else if s1._1 > s2._1 then s1
        else s2
      )
      ._1

  def doFindStrongestLongestBridgeStrength(
      sortedSegments: Seq[Segment],
      graph: Seq[Seq[(Orientation, Orientation)]],
      nextSrc: Int,
      orientation: Orientation,
      accStrength: Int,
      accUsed: Set[Int]
  ): (Int, Set[Int]) =
    val src = sortedSegments(nextSrc)
    val strength = accStrength + src.start + src.end
    val used = accUsed + nextSrc

    val nextSegmentIndicesAndOrientations = graph(nextSrc)
      .map(options =>
        orientation match
          case Orientation.Forward  => options._1
          case Orientation.Reversed => options._2
          case Orientation.Nil      => throw new Error("uh oh")
      )
      .zipWithIndex
      .filter(x =>
        x._2 != nextSrc && x._1 != Orientation.Nil && !used.contains(x._2)
      )

    nextSegmentIndicesAndOrientations match
      case Seq() =>
        (strength, used)

      case _ =>
        val solutions =
          for (nextOrientation, nextSegmentIndex) <-
              nextSegmentIndicesAndOrientations
          yield doFindStrongestLongestBridgeStrength(
            sortedSegments,
            graph,
            nextSegmentIndex,
            nextOrientation,
            strength,
            used
          )

        solutions.reduce((s1, s2) =>
          if s1._2.size > s2._2.size then s1
          else if s1._2.size < s2._2.size then s2
          else if s1._1 > s2._1 then s1
          else s2
        )
