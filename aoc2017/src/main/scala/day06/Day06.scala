package day06

import solution.Solution
import scala.collection.immutable.HashMap

class Day06Solution extends Solution:
    override def partOne(input: String): Either[String, String] =
        val initialMemoryLayout = parseMemoryBanks(input)

        val iterations = cyclesTilRepeat(
            initialMemoryLayout,
            Set[Seq[Int]](),
            0,
        )

        Right(iterations.toString)

    override def partTwo(input: String): Either[String, String] =
        val initialMemoryLayout = parseMemoryBanks(input)

        val length = cycleLength(
            initialMemoryLayout,
            Map[Seq[Int], Int](),
            0,
        )

        Right(length.toString)

    def parseMemoryBanks(input: String): Seq[Int] =
        input.split("\\s+").map(_.toInt).toSeq

    @scala.annotation.tailrec
    final def cyclesTilRepeat(memory: Seq[Int], previousLayouts: Set[Seq[Int]], redistributions: Int): Int =
        if previousLayouts.contains(memory) then
            redistributions
        else
            val newMemory = simulateRedistribution(memory)
            cyclesTilRepeat(newMemory, previousLayouts + memory, redistributions + 1)

    @scala.annotation.tailrec
    final def cycleLength(memory: Seq[Int], previousLayouts: Map[Seq[Int], Int], redistributions: Int): Int =
        if previousLayouts.contains(memory) then
            redistributions - previousLayouts.getOrElse(memory, 0)
        else
            val newMemory = simulateRedistribution(memory)
            cycleLength(newMemory, previousLayouts.updated(memory, redistributions), redistributions + 1)

    def simulateRedistribution(memory: Seq[Int]): Seq[Int] =
        val (value, j) = memory.zipWithIndex.maxBy(_._1)
        val evenlyDistributed = value / memory.size
        val remainder = value % memory.size

        for
            (x, i) <- memory.zipWithIndex
        yield
            if (i < j && remainder > 0 && j + remainder >= memory.size && i <= (j + remainder) % memory.size) then
                x + evenlyDistributed + 1
            else if (i == j) then
                evenlyDistributed
            else if (i > j && remainder > 0 && i <= j + remainder) then
                x + evenlyDistributed + 1
            else
                x + evenlyDistributed
