package leetcode.easy;

import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;
import org.junit.jupiter.api.Test;

public class WordsContainingCharacterTest {
    @Test
    public void testFindWordsContaining() {
        WordsContainingCharacter solution = new WordsContainingCharacter();

        String[] words = { "abc", "bcd", "aaaa", "cbc" };
        char x = 'a';

        List<Integer> expected = Arrays.asList(0, 2);
        List<Integer> result = solution.findWordsContaining(words, x);

        assertEquals(expected, result);
    }
}
