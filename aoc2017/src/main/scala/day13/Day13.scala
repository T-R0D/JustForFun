package day13

import solution.Solution

class Day13Solution extends Solution:
    override def partOne(input: String): Either[String, String] =
        val scanners = parseScanners(input)

        val deepestScanner = scanners.keys.max
        val depths = findDepthsCaughtByScanners(scanners, deepestScanner, 0)
        val totalSeverity = depths.foldLeft(0) { (acc, depth) =>
            acc + (depth * scanners.getOrElse(depth, 0))    
        }

        Right(totalSeverity.toString)

    override def partTwo(input: String): Either[String, String] =
        val scanners = parseScanners(input)

        val deepestScanner = scanners.keys.max
        val targetDelay = findMinimumDelayToPassUnscathed(scanners, deepestScanner)

        targetDelay match
            case Some(delay) => Right(delay.toString)
            case None => Left("ideal delay was not found")

    def parseScanners(input: String): Map[Int, Int] =
        (
            for
                line <- input.split("\\n")
            yield
                val numbers = line.split(": ").map(_.toInt)
                numbers(0) -> numbers(1)
        ).toMap
    
    def findMinimumDelayToPassUnscathed(scanners: Map[Int, Int], lastScannerDepth: Int): Option[Int] =
        findMinimumDelayToPassUnscathedInternal(scanners, lastScannerDepth, 0)
    
    @scala.annotation.tailrec
    final def findMinimumDelayToPassUnscathedInternal(scanners: Map[Int, Int], lastScannerDepth: Int, startTime: Int): Option[Int] =
        val caughtDepths = findDepthsCaughtByScanners(scanners, lastScannerDepth, startTime)
        if caughtDepths.size == 0 then
            Some(startTime)
        else
            findMinimumDelayToPassUnscathedInternal(scanners, lastScannerDepth, startTime + 1)

    def findDepthsCaughtByScanners(scanners: Map[Int, Int], lastScannerDepth: Int, startTime: Int): Seq[Int] =
        (0 to lastScannerDepth).foldLeft(Seq.empty: Seq[Int]) { (acc, depth) =>
            scanners.get(depth) match
                case None => acc
                case Some(range) => {
                    if scanPosition(range, startTime + depth) != 0 then
                        acc
                    else
                        acc :+ depth
                }
        }
    
    def scanPosition(range: Int, t: Int): Int =
        val cycleLength = {
            if range > 1 then
                (2 * range) - 2
            else
                1
        }
        val mod = t % cycleLength
        {
            if mod >= range then
                val workBack = mod - range
                range - (workBack + 2)
            else
                mod
        }
