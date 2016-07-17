import java.math.BigInteger;

/**
 * Created by Terence
 * on 8/16/2014.
 */
public class Problem16Solver {
    public static void main(String[] args) {
        BigInteger sum = new BigInteger("0");
        BigInteger workingNum = new BigInteger("2");

        workingNum = workingNum.pow(1000);

        while (workingNum.compareTo(BigInteger.valueOf(0)) > 0) {
            sum = sum.add(workingNum.mod(BigInteger.valueOf(10)));
            workingNum = workingNum.divide(BigInteger.valueOf(10));
            System.out.println("Working with: " + workingNum);
            System.out.println("Current sum: " + sum);
        }
    }
}