import java.io.*;

class Result {

    /*
     * Complete the 'staircase' function below.
     *
     * The function accepts INTEGER n as parameter.
     */

    public static void staircase(int n) {
        // Write your code here
        String space = " ";
        String hash = "#";
        for (int i = 0; i < n; i++) {
            System.out.printf(space.repeat(n - i - 1));
            System.out.printf(hash.repeat(i + 1));
            System.out.println();
        }
    }

}

public class Staircase {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));

        int n = Integer.parseInt(bufferedReader.readLine().trim());

        Result.staircase(n);

        bufferedReader.close();
    }
}
