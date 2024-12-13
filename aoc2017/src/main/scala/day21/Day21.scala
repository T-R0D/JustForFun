package day21

import solution.Solution

class Day21Solution extends Solution {
    override def partOne(input: String): Either[String, String] = {
        val ruleBook = parseRuleBook(input)

        val expandedRuleBook = expandRuleBook(ruleBook)

        val upscaledImage = (0 until 5).foldLeft(initialImage) { (acc, _) =>
            upscaleImage(expandedRuleBook, acc)
        }

        val nOnPixels = countOnPixels(upscaledImage)
    
        Right(nOnPixels.toString)
    }

    override def partTwo(input: String): Either[String, String] = {
        val ruleBook = parseRuleBook(input)

        val expandedRuleBook = expandRuleBook(ruleBook)

        val upscaledImage = (0 until 18).foldLeft(initialImage) { (acc, _) =>
            upscaleImage(expandedRuleBook, acc)
        }

        val nOnPixels = countOnPixels(upscaledImage)
    
        Right(nOnPixels.toString)
    }

    val initialImage = ".#...####"

    def parseRuleBook(input: String): Map[String, String] = {
        var ruleBook = (
            for {
                line <- input.split("\\n+")
                parts = line.split("\\s+=>\\s+")
            } yield {
                parts(0).replace("/", "") -> parts(1).replace("/", "")
            }
        ).toMap

        ruleBook
    }

    def expandRuleBook(ruleBook: Map[String, String]): Map[String, String] = {
        (
            for {
                (originalImage, upscaledImage) <- ruleBook
                k <- 0 until 4
            } yield {
                val rotatedImage = (0 until k).foldLeft(originalImage) { (acc, _) => rotateSquareImageClockwise(acc) }
                Seq(
                    rotatedImage,
                    flipSquareImageAcrossHorizontalAxis(rotatedImage),
                    flipSquareImageAcrossVerticalAxis(rotatedImage),
                ).map(_ -> upscaledImage)
            }
        ).flatten.toMap
    }

    def rotateSquareImageClockwise(image: String): String = {
        val n = Math.sqrt(image.size).toInt
        (
            for {
                i <- 0 until n
                j <- 0 until n
            } yield {
                image(((n - j - 1) * n) + i)
            }
        ).mkString
    }
    
    def flipSquareImageAcrossHorizontalAxis(image: String): String = {
        val n = Math.sqrt(image.size).toInt
        (
            for {
                i <- 0 until n
                j <- 0 until n
            } yield {
                image(((n - i - 1) * n) + j)
            }
        ).mkString
    }
        
    def flipSquareImageAcrossVerticalAxis(image: String): String = {
        val n = Math.sqrt(image.size).toInt
        (
            for {
                i <- 0 until n
                j <- 0 until n
            } yield {
                image((i * n) + (n - j - 1))
            }
        ).mkString
    }

    def upscaleImage(ruleBook: Map[String, String], image: String): String = {
        val n = Math.sqrt(image.size).toInt

        val (upscaledTiles, newImageSize, newTileSize) = {
            if (n & 0x01) == 0 then {
                imageToTiles(ruleBook, image, n, 2)
            } else {
                imageToTiles(ruleBook, image, n, 3)
            }
        }

        stitchTiles(upscaledTiles, newImageSize, newTileSize)
    }

    def imageToTiles(
        ruleBook: Map[String, String],
        image: String,
        imageSize: Int,
        tileSize: Int
    ): (Seq[String], Int, Int) = {
        val tiles = (
            for {
                i <- 0 until imageSize by tileSize
                j <- 0 until imageSize by tileSize
            } yield {
                (
                    for {
                        y <- 0 until tileSize
                        x <- 0 until tileSize
                    } yield {
                        image(((i + y) * imageSize) + j + x)
                    }
                ).mkString
            }
        )

        val newTileSize = tileSize + 1

        val newImageSize = (imageSize / tileSize) * newTileSize

        val upscaledTiles = tiles.map { (tile) =>
            ruleBook.get(tile) match {
                case None => throw Exception(s"tile '$tile' could not be found in rule book")
                case Some(value) => value
            }
        }

        (upscaledTiles, newImageSize, newTileSize)
    }

    def stitchTiles(tiles: Seq[String], newImageSize: Int, tileSize: Int): String = {
        val nTilesPerSide = Math.sqrt(tiles.size).toInt
        (
            for {
                i <- 0 until newImageSize
                j <- 0 until newImageSize
            } yield {
                val tileI = i / tileSize
                val y = i % tileSize
                val tileJ = j / tileSize
                val x = j % tileSize
                val tile = tiles((tileI * nTilesPerSide) + tileJ)
                tile((y * tileSize) + x)
            }
        ).mkString
    }

    def countOnPixels(image: String): Int = {
        image.foldLeft(0) { (acc, pixel) =>
            if pixel == '#' then {
                acc + 1
            } else {
                acc
            }
        }
    }
}