package day20

import scala.util.Try

import solution.Solution

class Day20Solution extends Solution {
    override def partOne(input: String): Either[String, String] = {
        val results = parseParticles(input)
        results match
            case x @ Left(errors) => Left(errors.mkString("; ")) 
            case Right(particles) => {
                val id = mostHomeboundParticleId(particles)

                Right(id.toString)
            }
    }

    override def partTwo(input: String): Either[String, String] = {
        val results = parseParticles(input)
        results match
            case x @ Left(errors) => Left(errors.mkString("; ")) 
            case Right(particles) => {
                val nRemainingParticles = simulateToEliminate(particles)

                Right(nRemainingParticles.toString)
            }
    }

    case class Particle(position: Array[Int], velocity: Array[Int], acceleration: Array[Int]) {
        def next(): Particle = {
            val newVelocity = velocity.zip(acceleration).map((v, a) => v + a)
            val newPosition = position.zip(newVelocity).map((p, v) => p + v)
            Particle(newPosition, newVelocity, acceleration)
        }

        def distanceFromOrigin(): Int = {
            position.fold(0) { (acc, p) => acc + Math.abs(p) }
        }

        def positionTuple(): (Int, Int, Int) = {
            (position(0), position(1), position(2))
        }

        override
        def toString(): String = {
            s"Particle(${position.mkString(",")}, ${velocity.mkString(",")}, ${acceleration.mkString(",")})"
        }
    }

    def parseParticles(input: String): Either[Seq[String], Seq[Particle]] = {
        val prefixes = Seq("p=<", "v=<", "a=<")
        val results = input.split("\\n+").map { (line) =>
            Try {
                val vectors = line.split(">,\\s+").zipWithIndex.map{ (quantityStr, i) =>
                    quantityStr
                        .replace(">", "")
                        .replace(prefixes(i), "")
                        .split(",")
                        .map(_.toInt)
                        .toArray
                }.toSeq
                Particle(position = vectors(0), velocity = vectors(1), acceleration = vectors(2))
            }.toEither.left.map(_.toString)
        }

        val failures = results.collect({
            case Left(x) => x
        }).toSeq

        if !failures.isEmpty then {
            Left(failures)
        } else {
            Right(results.collect({
                case Right(x) => x
            }))
        }
    }
    
    def mostHomeboundParticleId(particles: Seq[Particle]): Int = {
        val particlesWithIds = particles
            .zipWithIndex
            .sortBy((p, _) => p.distanceFromOrigin())

        mostHomeboundParticleIdInternal(particlesWithIds, 0)
    }

    @scala.annotation.tailrec
    final def mostHomeboundParticleIdInternal(particles: Seq[(Particle, Int)], nTimesSameHead: Int): Int = {
        val currentRankings = particles.map(_._2)
        val nextParticles = particles
            .map((p, id) => (p.next(), id))
            .sortBy((p, _) => p.distanceFromOrigin())

        val nextTimesSameHead = {
            if nextParticles.head._2 == particles.head._2 then {
                nTimesSameHead + 1
            } else {
                0
            }
        }

        // Arbitrary definition of "the long term".
        if nextTimesSameHead == 1000 then {
            nextParticles.head._2
        } else {
            mostHomeboundParticleIdInternal(nextParticles, nextTimesSameHead)
        }
    }

    def simulateToEliminate(particles: Seq[Particle]): Int = {
        simulateToEliminateInternal(particles, 0)
    }

    @scala.annotation.tailrec
    final def simulateToEliminateInternal(particles: Seq[Particle], nTimesWithoutCollisions: Int): Int = {
        val nextParticles = particles
            .map(_.next())
            .groupBy(_.positionTuple())
            .filter((p, colocatedParticles) => colocatedParticles.size <= 1)
            .toSeq
            .flatMap((k, v) => v)

        val nextTimesWithoutCollistions = {
            if particles.size == nextParticles.size then {
                nTimesWithoutCollisions + 1
            } else {
                0
            }
        }

        // Arbitrary definition of "the long term".
        if nTimesWithoutCollisions == 1000 then {
            nextParticles.size
        } else {
            simulateToEliminateInternal(nextParticles, nextTimesWithoutCollistions)
        }
    }
}
