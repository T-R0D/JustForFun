package day20

import scala.util.Try

import solution.Solution

class Day20Solution extends Solution {
    override def partOne(input: String): Either[String, String] = {
        val results = parseParticles(input)
        results match
            case x @ Left(errors) => Left(errors.mkString("; ")) 
            case Right(particles) => {
                val id = mostHomeboundParticle(particles)

                Right(id.toString)
            }
    }

    override def partTwo(input: String): Either[String, String] = ???

    case class Particle(position: Array[Int], velocity: Array[Int], acceleration: Array[Int])

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
                Particle(position = vectors(0), velocity = vectors(1), acceleration = vectors(0))
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
    
    def mostHomeboundParticle(particles: Seq[Particle]): Int = {
        val (id, _) = particles.zipWithIndex.foldLeft((0, particles.head.acceleration)) { (acc, particleWithId) =>
            val (homboundestId, homeboundestAcceleration) = acc
            val (particle, id) = particleWithId
            val accelerationMagnitude = particle.acceleration.reduce((a, b) => Math.abs(a) + Math.abs(b))
            val homeboundestAccelerationMagnitude = homeboundestAcceleration.reduce((a, b) => Math.abs(a) + Math.abs(b))
            if accelerationMagnitude < homeboundestAccelerationMagnitude then {
                (id, particle.acceleration)
            } else {
                acc
            }

        }

        id
    }
}
