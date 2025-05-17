package leetcode.easy;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

class TwoSumTest {
    @Test
    void testTwoSum() {
        TwoSum solution = new TwoSum();
        
        int[] nums = {2, 7, 11, 15};
        int target = 9;
        int[] expected = {0, 1};
        
        assertArrayEquals(expected, solution.twoSum(nums, target));
    }
}