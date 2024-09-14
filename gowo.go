// Package gowo provides text owofication capabilities.
//
// To perform a basic owofication with the default vars (replacements, actions,
// emoticons) and probabilities, [Owoify] and [OwoifyTokens] functions can be
// used:
//
//	fmt.Println(gowo.Owoify("Lorem ipsum dolor sit amet"))
//
// For advanced usage with custom vars and probabilities, [Owoifier] struct can
// be used.
package gowo

// Owoify performs text owofication:
//   - Replaces symbols and substrings with replacements
//   - Adds s-stutters
//   - Adds actions *blushes*
//   - Adds emoticons ≧◉◡◉≦
//
// The text is split into tokens by the whitespace (`" "`) character.
//
// Owoify uses the default [Owoifier] to owoify text.
func Owoify(text string) string {
	return New().Owoify(text)
}

// Owoify performs tokens owofication:
//   - Replaces symbols and substrings with replacements
//   - Adds s-stutters
//   - Adds actions *blushes*
//   - Adds emoticons ≧◉◡◉≦
//
// Owoify uses the default [Owoifier] to owoify tokens.
func OwoifyTokens(tokens []string) []string {
	return New().OwoifyTokens(tokens)
}
