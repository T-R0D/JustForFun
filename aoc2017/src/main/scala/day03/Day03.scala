package day03

import solution.Solution

class Day03Solution extends Solution:
    override def partOne(input: String): Either[String, String] =
        val target = parseTargetLocation(input)

        val distanceFromOrigin = distanceFromOriginOnSpiral(target)

        scala.Right(distanceFromOrigin.toString)

    override def partTwo(input: String): Either[String, String] =
        val target = parseTargetLocation(input)

        val firstSumGreaterThanTarget = sumOnSpiralGreaterThanTarget(target)

        scala.Right(firstSumGreaterThanTarget.toString)

    def parseTargetLocation(input: String): Int =
        input.toInt
        
    def distanceFromOriginOnSpiral(target: Int): Int =
        if target == 1 then
            0
        else
            val ring = ringNumber(target)
            val position = positionOnRing(target, ring)
            val additive = distanceFromOriginAdditive(position, ring)
            ring + additive
    
    def ringNumber(target: Int): Int =
        LazyList
            .from(0)
            .dropWhile(!isInRing(target, _))
            .head

    def isInRing(target: Int, ring: Int): Boolean =
        val doubleRing = 2 * ring
        ((doubleRing - 1) * (doubleRing - 1)) + 1 <= target &&
            target <= ((doubleRing + 1) * (doubleRing + 1))

    def positionOnRing(target: Int, ringNumber: Int): Int =
        val previousRing = ringNumber - 1
        val doubleRing = 2 * previousRing
        val previousRingMax = (doubleRing + 1) * (doubleRing + 1)
        target - previousRingMax

    def distanceFromOriginAdditive(positionOnRing: Int, ringNumber: Int) = 
        (1 to positionOnRing)
            .scanLeft((ringNumber, -1)) {
                case ((previousAdditive, direction), position) =>
                    val additive = previousAdditive + direction
                    val nextDirection = 
                        if additive == 0 then
                            1
                        else if additive == ringNumber then
                            -1
                        else
                            direction
                    (additive, nextDirection)
            }
            .map(_._1)
            .takeRight(1)
            .head

    def sumOnSpiralGreaterThanTarget(target: Int): Int =
        sumOnSpiralHelper(target, (0, 0), Right(2), 0, Map())

    @scala.annotation.tailrec
    final def sumOnSpiralHelper(
        target: Int,
        position: (Int, Int),
        direction: Direction,
        ring: Int,
        discovered: Map[(Int, Int), Int]
    ): Int =
        val currentValue = if (position == (0, 0)) then
            1
        else
            val (i, j) = position
            (
                for
                    di <- -1 to 1
                    dj <- -1 to 1
                        if (di != 0 || dj != 0)
                yield
                    discovered.getOrElse((i + di, j + dj), 0)
            ).sum()

        if (currentValue > target) then
            currentValue
        else
            val (nextDirection, nextPos, nextRing) = direction.next(position, ring)
            sumOnSpiralHelper(
                target,
                nextPos,
                nextDirection,
                nextRing,
                discovered.updated(position, currentValue)
            )

    sealed trait Direction:
        def next(position: (Int, Int), ring: Int): (Direction, (Int, Int), Int) =
            val (i, j) = position

            this match
                case Up(remaining) =>
                    if remaining > 1 then
                        (Up(remaining - 1), (i - 1, j), ring)
                    else
                        (Left(2 * ring), (i, j -1), ring)
                case Left(remaining) =>
                    if remaining > 1 then
                        (Left(remaining - 1), (i, j - 1), ring)
                    else
                        (Down(2 * ring), (i + 1, j), ring)
                case Down(remaining) =>  
                    if remaining > 1 then
                        (Down(remaining - 1), (i + 1, j), ring)
                    else
                        (Right(2 * ring + 1), (i, j + 1), ring)
                case Right(remaining) => 
                    if remaining > 1 then
                        (Right(remaining - 1), (i, j + 1), ring)
                    else
                        (Up(2 * (ring+1) - 1), (i - 1, j), ring + 1)
                    

    case class Up(val remaining: Int) extends Direction
    case class Down(val remaining: Int) extends Direction
    case class Left(val remaining: Int) extends Direction
    case class Right(val remaining: Int) extends Direction
        