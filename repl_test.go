package main

import (
    "testing"
)

func TestCleanInput(t *testing.T) {

    cases := []struct {
        input string
        expected []string
    }{
        {
            input: "  hello  world  ",
            expected: []string{"hello", "world"},
        },
        {
            input: "  this  and  that  ",
            expected: []string{"this", "and", "that"},
        },
    }

    for _, c := range cases {
        actual := cleanInput(c.input)
        if len(actual) != len(c.expected) {
            t.Errorf("expected length %d does not match actual length %d", len(actual), len(c.expected))
        }

        for i := range actual {
            word := actual[i]
            expectedWord := c.expected[i]

            if word != expectedWord {
                t.Errorf("word: %s does not match expected word: %s", word, expectedWord)
            }
        }
    }

}
