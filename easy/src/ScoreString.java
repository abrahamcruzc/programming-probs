public class ScoreString {
    public int scoreOfString(String s) {
        if ( s == null || s.length() < 2) {
            return 0;
        }

        int score = 0;

        for (int i = 1; i < s.length(); i++) {
            int currentASCII = (int) s.charAt(i);
            int previousASCII = (int) s.charAt(i - 1);

            score += Math.abs(currentASCII - previousASCII);
        }   

        return score;
    }


    public static void main(String[] args) {
        ScoreString solution = new ScoreString();

        String test1 = "hello";
        int result1 = solution.scoreOfString(test1);
        System.out.println("Input: " + test1);
        System.out.println("Output: " + result1);
        System.out.println("Expected: 13");
        System.out.println("Test 1 " + (result1 == 13 ? "PASSED" : "FAILED"));
    }
} 