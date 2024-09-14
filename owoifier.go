package gowo

import (
	"fmt"
	"math/rand"
	"strings"
)

// makeReplacer creates an instance of [strings.Replacer] with map m used as
// oldnew.
func makeReplacer(m map[string]string) *strings.Replacer {
	oldnew := make([]string, 2*len(m))
	i := 0

	for k, v := range m {
		oldnew[i] = k
		oldnew[i+1] = v
		i += 2
	}

	return strings.NewReplacer(oldnew...)
}

// Owoifier is a struct that performs text owofication.
type Owoifier struct {
	actions   []string          // Actions slice.
	emoticons []string          // Emoticons slice.
	replacer  *strings.Replacer // Token replacer.

	probabilityOfAction      float64 // Probability of action occurrance.
	probabilityOfEmoticon    float64 // Probability of emoticon occurrance.
	probabilityOfReplacement float64 // Probability of replacement occurrance.
	probabilityOfStutter     float64 // Probability of stutter occurrance.
}

// New returns a new instance of Owoifier with the default vars and probabilities.
func New() *Owoifier {
	return &Owoifier{
		actions:                  Actions,
		emoticons:                Emoticons,
		replacer:                 makeReplacer(Replacements),
		probabilityOfAction:      ProbabilityOfAction,
		probabilityOfEmoticon:    ProbabilityOfEmoticon,
		probabilityOfReplacement: ProbabilityOfReplacement,
		probabilityOfStutter:     ProbabilityOfStutter,
	}
}

// NewWithProbabilities returns a new instance of Owoifier with arbitrary probabilities.
func NewWithProbabilities(
	probabilityOfAction float64,
	probabilityOfEmoticon float64,
	probabilityOfReplacement float64,
	probabilityOfStutter float64,
) *Owoifier {
	return &Owoifier{
		actions:                  Actions,
		emoticons:                Emoticons,
		replacer:                 makeReplacer(Replacements),
		probabilityOfAction:      probabilityOfAction,
		probabilityOfEmoticon:    probabilityOfEmoticon,
		probabilityOfReplacement: probabilityOfReplacement,
		probabilityOfStutter:     probabilityOfStutter,
	}
}

// NewWithVars returns a new instance of Owoifier with arbitrary vars.
func NewWithVars(
	actions []string,
	emoticons []string,
	replacements map[string]string,
) *Owoifier {
	return &Owoifier{
		actions:                  actions,
		emoticons:                emoticons,
		replacer:                 makeReplacer(replacements),
		probabilityOfAction:      ProbabilityOfAction,
		probabilityOfEmoticon:    ProbabilityOfEmoticon,
		probabilityOfReplacement: ProbabilityOfReplacement,
		probabilityOfStutter:     ProbabilityOfStutter,
	}
}

// NewWithVarsProbabilities returns a new instance of Owoifier with arbitrary
// vars and probabilities.
func NewWithVarsProbabilities(
	actions []string,
	emoticons []string,
	replacements map[string]string,
	probabilityOfAction float64,
	probabilityOfEmoticon float64,
	probabilityOfReplacement float64,
	probabilityOfStutter float64,
) *Owoifier {
	return &Owoifier{
		actions:                  actions,
		emoticons:                emoticons,
		replacer:                 makeReplacer(replacements),
		probabilityOfAction:      probabilityOfAction,
		probabilityOfEmoticon:    probabilityOfEmoticon,
		probabilityOfReplacement: probabilityOfReplacement,
		probabilityOfStutter:     probabilityOfStutter,
	}
}

// Owoify performs text owofication:
//   - Replaces symbols and substrings with replacements
//   - Adds s-stutters
//   - Adds actions *blushes*
//   - Adds emoticons ≧◉◡◉≦
func (o Owoifier) Owoify(text string) string {
	tokens := strings.Split(text, " ")
	result := make([]string, 0, len(tokens))

	for _, token := range tokens {
		newToken := token

		if rand.Float64() < o.probabilityOfReplacement {
			newToken = o.replacer.Replace(token)
		}

		if rand.Float64() < o.probabilityOfStutter {
			newToken = fmt.Sprintf("%c-%s", newToken[0], newToken)
		}

		result = append(result, newToken)

		if rand.Float64() < o.probabilityOfAction {
			result = append(result, o.RandomAction())
		}

		if rand.Float64() < o.probabilityOfEmoticon {
			result = append(result, o.RandomEmoticon())
		}
	}

	return strings.Join(result, " ")
}

// RandomAction returns a random action of the Owoifier.
func (o Owoifier) RandomAction() string {
	i := rand.Intn(len(o.actions))
	return o.actions[i]
}

// RandomEmoticon returns a random emoticon of the Owoifier.
func (o Owoifier) RandomEmoticon() string {
	i := rand.Intn(len(o.emoticons))
	return o.emoticons[i]
}
