import java.io.*;
import java.util.*;
import java.util.stream.*;
import static java.util.stream.Collectors.toList;

class Result {

    /*
     * Complete the 'insertionSort2' function below.
     *
     * The function accepts following parameters:
     *  1. INTEGER n
     *  2. INTEGER_ARRAY arr
     */

    public static void insertionSort2(int n, List<Integer> arr) {
    // Write your code here
        for (int i = 1; i < arr.size(); i++) {
            for (int j = 0; j < i; j++) {
                if (arr.get(i) < arr.get(j)) {
                    int value = arr.get(i);
                    arr.subList(i, i + 1).clear();
                    arr.add(j, value);
                    break;
                }
            }
            
            for (Integer value : arr) {
                System.out.print(value + " ");
            }
            System.out.println();
        }
    }

}

public class InsertionSort2 {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));

        int n = Integer.parseInt(bufferedReader.readLine().trim());

        List<Integer> arr = Stream.of(bufferedReader.readLine().replaceAll("\\s+$", "").split(" "))
            .map(Integer::parseInt)
            .collect(toList());

        Result.insertionSort2(n, arr);

        bufferedReader.close();
    }
}
