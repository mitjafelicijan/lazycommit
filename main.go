package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const systemPrompt = `
	Given a set of Git diffs, generate a concise and informative commit message that summarizes the changes. The commit message should follow these guidelines:
	1. The first line (the 'subject line') should be a short, descriptive summary of the changes, no more than 50 characters long.
	2. If necessary, provide additional context or explanation in the body of the commit message, with each paragraph separated by a blank line.
	3. Use the Git diff output to determine the appropriate level of detail to include in the commit message. Focus on explaining the "why" behind the changes, not just the "what".
	4. Avoid unnecessary details or redundant information that can be inferred from the Git diff itself.
	5. Use proper grammar, capitalization, and punctuation throughout the commit message.`

func main() {
	openaiApiKey, found := os.LookupEnv("OPENAI_API_KEY")
	if !found {
		fmt.Fprintln(os.Stderr, "`OPENAI_API_KEY` environmental variable is missing.")
		fmt.Fprintln(os.Stderr, "You should put `OPENAI_API_KEY` into your bashrc.")
		os.Exit(1)
	}

	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: reading from stdin:", err)
		os.Exit(1)
	}

	commitMessage := strings.TrimSpace(string(bytes))
	if len(commitMessage) == 0 {
		fmt.Fprintln(os.Stderr, "Error: no message provided")
		os.Exit(1)
	}

	answer, err := openai(commitMessage, systemPrompt, openaiApiKey)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: LLM did not return any data:", err)
		os.Exit(1)
	}

	// Answer was given without any errors.
	// Relaying answer back to callee.
	fmt.Println(answer)
}
