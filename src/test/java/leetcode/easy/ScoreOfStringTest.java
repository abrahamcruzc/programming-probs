package leetcode.easy;

import static org.junit.jupiter.api.Assertions.assertEquals;
import org.junit.jupiter.api.Test;

public class ScoreOfStringTest {
    @Test
    void testScoreOfString() {
        ScoreOfString solution = new ScoreOfString();

        String s = "hello";

        int expected = 13;

        assertEquals(expected, solution.scoreOfString(s));
    }
}
