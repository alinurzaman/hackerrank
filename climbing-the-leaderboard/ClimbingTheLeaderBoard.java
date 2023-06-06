import java.io.*;
import java.util.*;
import java.util.stream.*;
import static java.util.stream.Collectors.joining;
import static java.util.stream.Collectors.toList;

class Result {

    /*
     * Complete the 'climbingLeaderboard' function below.
     *
     * The function is expected to return an INTEGER_ARRAY.
     * The function accepts following parameters:
     * 1. INTEGER_ARRAY ranked
     * 2. INTEGER_ARRAY player
     */

    public static List<Integer> climbingLeaderboard(List<Integer> ranked, List<Integer> player) {
        // Write your code here
        List<Integer> rank = new ArrayList<>();
        List<Integer> listWithoutDuplicates = ranked.stream()
                .distinct()
                .collect(Collectors.toList());
        int index = listWithoutDuplicates.size() - 1;
        for (int i = 0; i < player.size(); i++) {
            for (int j = index; j >= 0; j--) {
                if (player.get(i) >= listWithoutDuplicates.get(0)) {
                    rank.add(1);
                    break;
                } else if (player.get(i) < listWithoutDuplicates.get(j)) {
                    rank.add(j + 2);
                    index = j;
                    break;
                } else if (player.get(i) == listWithoutDuplicates.get(j)) {
                    rank.add(j + 1);
                    index = j;
                    break;
                }
            }
        }

        return rank;
    }

}

public class ClimbingTheLeaderBoard {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        int rankedCount = Integer.parseInt(bufferedReader.readLine().trim());

        List<Integer> ranked = Stream.of(bufferedReader.readLine().replaceAll("\\s+$", "").split(" "))
                .map(Integer::parseInt)
                .collect(toList());

        int playerCount = Integer.parseInt(bufferedReader.readLine().trim());

        List<Integer> player = Stream.of(bufferedReader.readLine().replaceAll("\\s+$", "").split(" "))
                .map(Integer::parseInt)
                .collect(toList());

        List<Integer> result = Result.climbingLeaderboard(ranked, player);

        bufferedWriter.write(
                result.stream()
                        .map(Object::toString)
                        .collect(joining("\n"))
                        + "\n");

        bufferedReader.close();
        bufferedWriter.close();
    }
}
