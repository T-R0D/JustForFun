package generate

import java.nio.file.{Files, Path, Paths}
import scala.util.{Failure, Success, Try}

import args.Config
import java.nio.file.attribute.BasicFileAttributes
import java.nio.file.{FileVisitResult, SimpleFileVisitor}
import java.io.IOException

class Generator:
    def generate(config: Config): Unit =
        if (config.delete) then
            clean(config.genTargetDir)
        else
            createSolutionFiles(config.genTargetDir).fold(identity, x => ())
            

    def createSolutionFiles(targetDir: String): Try[Unit] =
        (1 to 25).foldLeft(Try(())) { (acc, day) =>
            for
                _ <- acc

                solutionPath = Paths.get(targetDir, "/src/main/scala/", dayDirStr(day))
                result <- createDir(solutionPath)
                solutionFilePath = Paths.get(solutionPath.toString, dayFileStr(day))
                _ <- createFile(solutionFilePath, solutionFileContents(day))
            
                testPath = Paths.get(targetDir, "/src/test/scala/", dayDirStr(day))
                _ <- createDir(testPath)
                testFilePath = Paths.get(testPath.toString, dayTestFileStr(day))
                _ <- createFile(testFilePath, testFileContents(day))
            yield
                ()
        }

    def clean(targetDir: String): Unit =
        for 
            day <- 1 to 25
            projectPath <- Seq("/src/main/scala/", "/src/test/scala/")
        do
            val dayStr = dayDirStr(day)
            val path = Paths.get(targetDir, projectPath, dayStr)
            Files.walkFileTree(
                path,
                new SimpleFileVisitor[Path] {
                    override def visitFile(file: Path, attrs: BasicFileAttributes): FileVisitResult = 
                        Files.delete(file)
                        FileVisitResult.CONTINUE

                    override def postVisitDirectory(dir: Path, exc: IOException): FileVisitResult = 
                        Files.delete(dir)
                        FileVisitResult.CONTINUE
                }
            )

    def dayDirStr(day: Int): String = f"day$day%02d"

    def dayFileStr(day: Int): String = f"Day$day%02d.scala"

    def dayTestFileStr(day: Int): String = f"Day$day%02dTest.scala"

    def createDir(path: Path): Try[Unit] =
        Files.createDirectories(path)
        Success(())
        
    def createFile(path: Path, content: String): Try[Unit] =
        Files.writeString(path, content)
        Success(())

    def solutionFileContents(day: Int): String =
        f"""package day$day%02d
           |
           |import solution.Solution
           |
           |class Day$day%02dSolution extends Solution:
           |    override def partOne(input: String): Either[String, String] = ???
           |
           |    override def partTwo(input: String): Either[String, String] = ???
        """.stripMargin

    def testFileContents(day: Int): String = 
        f"""package Day$day%02d
           |
           |import org.scalatest.funsuite.AnyFunSuite
           |
           |class Day$day%02dSolutionTests extends AnyFunSuite
        """.stripMargin

        
