package steganography;

import ar.com.hjg.pngj.*;
import ar.com.hjg.pngj.chunks.ChunkCopyBehaviour;
import ar.com.hjg.pngj.chunks.PngChunkTextVar;

import java.io.File;

/**
 * Created by Terence
 * on 1/15/2015.
 */
public class ImageHider {
    public static final String TEST_FOLDER_PATH =
        "C:\\Users\\Terence\\Documents\\GitHub\\JustForFun\\Steganography\\TestImages\\";

    public static void main(String[] args) {
        String coverImage = TEST_FOLDER_PATH + "toor.png";
        String messageImage = TEST_FOLDER_PATH + "Mushroom2.PNG";
        String encodedFile = TEST_FOLDER_PATH + "test_encoding.png";
        String decodedFile = TEST_FOLDER_PATH + "test_decoding.png";

        hideImage(messageImage, coverImage, encodedFile);
        extractImage(encodedFile, decodedFile);

//        copyImageSorta(messageImage, encodedFile);
    }

    static void hideImage(String coverImageFileName, String messageImageFileName, String secretImageFileName) {

        PngReader coverImageReader = new PngReader(new File(coverImageFileName));
        PngReader messageImageReader = new PngReader(new File(messageImageFileName));
        System.out.println(coverImageFileName + ":\n" + coverImageReader.toString());
        System.out.println(messageImageFileName + ":\n" + messageImageReader.toString());

        ImageInfo steganoImageInfo = createCoverImageInfo(coverImageReader.imgInfo, 8, false);

        PngWriter secretFileWriter = new PngWriter(new File(secretImageFileName), steganoImageInfo, true);
        secretFileWriter.getMetadata().setText(PngChunkTextVar.KEY_Description, "I put a secret here!");

        int channels = coverImageReader.imgInfo.channels;
        for (int i = 0; i < coverImageReader.imgInfo.rows; ++i) {
            IImageLine coverLine = coverImageReader.readRow();
            IImageLine messageLine = messageImageReader.readRow();
            IImageLine steganoLine = ImageLineInt.getFactory().createImageLine(coverImageReader.imgInfo);

            int[] coverLineValues = ((ImageLineInt) coverLine).getScanline();
            int[] messageLineValues = ((ImageLineInt) messageLine).getScanline();
            int[] steganoLineValues = ((ImageLineInt) steganoLine).getScanline();

            for (int j = 0; j < coverImageReader.imgInfo.cols; ++j) {
                for (int channel = 0; channel < channels; ++channel) {
                    int elementIndex = (j * channels) + channel;

                    steganoLineValues[elementIndex] = computeHidingValue(
                        coverLineValues[elementIndex],
                        messageLineValues[elementIndex]
                    );
                }
            }
            secretFileWriter.writeRow(steganoLine);
        }

        coverImageReader.end();
        messageImageReader.end();
        secretFileWriter.end();
    }

    static int computeHidingValue(int coverValue, int messageValue) {
        int steganoValue = coverValue;
        steganoValue &= 0b11110000;
        int messageTemp = messageValue;
        steganoValue |= messageTemp >>> 4;

        return steganoValue;
    }

    static void extractImage(String encodedFileName, String decodedFileName) {
        PngReader encodedImageReader = new PngReader(new File(encodedFileName));
        PngWriter decodedImageWriter = new PngWriter(new File(decodedFileName), encodedImageReader.imgInfo, true);
        decodedImageWriter.getMetadata().setText(PngChunkTextVar.KEY_Comment, "This is a decoded message!");

        int channels = encodedImageReader.imgInfo.channels;
        for (int i = 0; i < encodedImageReader.imgInfo.rows; ++i) {
            IImageLine encodedLine = encodedImageReader.readRow();
            int[] encodedValues = ((ImageLineInt) encodedLine).getScanline();

            IImageLine decodedLine = ImageLineInt.getFactory().createImageLine(encodedImageReader.imgInfo);
            int[] decodedValues = ((ImageLineInt) decodedLine).getScanline();

            for (int j = 0; j < encodedImageReader.imgInfo.cols; ++j) {
                for (int channel = 0; channel < channels; ++channel) {
                    int elementIndex = (j * channels) + channel;
                    decodedValues[elementIndex] = extractMessageValue(encodedValues[elementIndex]);
                }
            }
            decodedImageWriter.writeRow(decodedLine);
        }

        encodedImageReader.end();
        decodedImageWriter.end();
    }

    static int extractMessageValue(int steganoValue) {
        /* assume that stegano value is 8 bits, there are 3 message bits. */
        return  steganoValue << 5;
    }

    static ImageInfo createCoverImageInfo(ImageInfo coverInfo, int bitDepth, boolean hasAlpha) {
        return new ImageInfo(coverInfo.cols, coverInfo.rows, bitDepth, hasAlpha);
    }

    static void exampleRedDecreaseGreenIncreaseCode() {
        String origFilename = TEST_FOLDER_PATH + "toor.PNG";
        String destFilename = TEST_FOLDER_PATH + "other_file.png";

        PngReader pngr = new PngReader(new File(origFilename));
        System.out.println(pngr.toString());
        int channels = pngr.imgInfo.channels;
        if (channels < 3 || pngr.imgInfo.bitDepth != 8)
            throw new RuntimeException("This method is for RGB8/RGBA8 images");
        PngWriter pngw = new PngWriter(new File(destFilename), pngr.imgInfo, true);
        pngw.copyChunksFrom(pngr.getChunksList(), ChunkCopyBehaviour.COPY_ALL_SAFE);
        pngw.getMetadata().setText(PngChunkTextVar.KEY_Description, "Decreased red and increased green");
        for (int row = 0; row < pngr.imgInfo.rows; row++) { // also: while(pngr.hasMoreRows())
            IImageLine l1 = pngr.readRow();
            int[] scanline = ((ImageLineInt) l1).getScanline(); // to save typing
            for (int j = 0; j < pngr.imgInfo.cols; j++) {
                scanline[j * channels] /= 2;
                scanline[j * channels + 1] = ImageLineHelper.clampTo_0_255(scanline[j * channels + 1] + 20);
            }
            pngw.writeRow(l1);
        }
        pngr.end(); // it's recommended to end the reader first, in case there are trailing chunks to read
        pngw.end();
    }

    static void copyImageSorta(String pictureFileName, String newPictureFileName) {
        PngReader pngReader = new PngReader(new File(pictureFileName));
        PngWriter pngWriter = new PngWriter(new File(newPictureFileName), pngReader.imgInfo, true);

        for (int i = 0; i < pngReader.imgInfo.rows; i++) {
            IImageLine originalLine = pngReader.readRow();
            int[] lineValues = ((ImageLineInt) originalLine).getScanline();
            IImageLine newLine = ImageLineInt.getFactory().createImageLine(pngReader.imgInfo);
            int[] newValues = ((ImageLineInt) newLine).getScanline();

            for (int j = 0; j < pngReader.imgInfo.cols; j++) {
                int elementIndex = j * pngReader.imgInfo.channels;

                for (int channel = 0; channel < pngReader.imgInfo.channels; ++ channel) {
                    newValues[elementIndex + channel] = lineValues[elementIndex + channel];
                }
            }

            pngWriter.writeRow(newLine);
        }

        pngReader.end();
        pngWriter.end();
    }
}
