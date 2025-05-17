package leetcode.easy;

import static org.junit.jupiter.api.Assertions.assertEquals;
import org.junit.jupiter.api.Test;

public class JewelsAndStonesTest {
   @Test
   public void testNumJewelsInStones() {
      JewelsAndStones solution = new JewelsAndStones();

      String jewels = "aA";
      String stones = "aAAbbbb";

      int expected = 3;

      assertEquals(expected, solution.numJewelsInStones(jewels, stones));
   }
}
