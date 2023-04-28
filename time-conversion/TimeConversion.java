import java.io.*;

class Result {

    /*
     * Complete the 'timeConversion' function below.
     *
     * The function is expected to return a STRING.
     * The function accepts STRING s as parameter.
     */

    public static String timeConversion(String s) {
        // Write your code here
        String hours = s.substring(0, 2);
        String minutesAndSeconds = s.substring(2, 8);
        String meridiem = s.substring(s.length() - 2);

        String finalHours = hours;
        if (meridiem.equals("PM") && !hours.equals("12"))
            finalHours = String.valueOf(Integer.parseInt(hours) + 12);
        else if (meridiem.equals("AM") && hours.equals("12"))
            finalHours = "00";

        return finalHours + minutesAndSeconds;
    }

}

public class TimeConversion {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        String s = bufferedReader.readLine();

        String result = Result.timeConversion(s);

        bufferedWriter.write(result);
        bufferedWriter.newLine();

        bufferedReader.close();
        bufferedWriter.close();
    }
}
