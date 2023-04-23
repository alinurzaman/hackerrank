import java.io.*;
import java.util.*;
import java.util.stream.*;
import static java.util.stream.Collectors.toList;

class Result {

    /*
     * Complete the 'plusMinus' function below.
     *
     * The function accepts INTEGER_ARRAY arr as parameter.
     */

    public static void plusMinus(List<Integer> arr) {
        // Write your code here
        Integer pos = 0;
        Integer neg = 0;
        Integer zero = 0;
        for (Integer value : arr) {
            if (value > 0)
                pos++;
            else if (value < 0)
                neg++;
            else
                zero++;
        }
        Integer total = arr.size();
        System.out.printf("%.6f\n", pos.doubleValue() / total.doubleValue());
        System.out.printf("%.6f\n", neg.doubleValue() / total.doubleValue());
        System.out.printf("%.6f\n", zero.doubleValue() / total.doubleValue());
    }

}

public class PlusMinus {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));

        int n = Integer.parseInt(bufferedReader.readLine().trim());

        List<Integer> arr = Stream.of(bufferedReader.readLine().replaceAll("\\s+$", "").split(" "))
                .map(Integer::parseInt)
                .collect(toList());

        Result.plusMinus(arr);

        bufferedReader.close();
    }
}
