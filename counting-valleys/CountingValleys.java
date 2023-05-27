import java.io.*;

class Result {

    /*
     * Complete the 'countingValleys' function below.
     *
     * The function is expected to return an INTEGER.
     * The function accepts following parameters:
     * 1. INTEGER steps
     * 2. STRING path
     */

    public static int countingValleys(int steps, String path) {
        // Write your code here
        int totalValley = 0;
        int level = 0;
        for (int i = 0; i < steps; i++) {
            if (path.charAt(i) == 'U') {
                level++;
                if (level == 0)
                    totalValley++;
            } else if (path.charAt(i) == 'D')
                level--;
        }
        return totalValley;
    }

}

public class CountingValleys {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        int steps = Integer.parseInt(bufferedReader.readLine().trim());

        String path = bufferedReader.readLine();

        int result = Result.countingValleys(steps, path);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedReader.close();
        bufferedWriter.close();
    }
}
