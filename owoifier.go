package gowo

// Owoifier is a struct that performs text owofication.
type Owoifier struct {
	actions      []string          // Actions slice.
	emoticons    []string          // Emoticons slice.
	replacements map[string]string // Word replacements.

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
		replacements:             Replacements,
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
		replacements:             Replacements,
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
		replacements:             replacements,
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
		replacements:             replacements,
		probabilityOfAction:      probabilityOfAction,
		probabilityOfEmoticon:    probabilityOfEmoticon,
		probabilityOfReplacement: probabilityOfReplacement,
		probabilityOfStutter:     probabilityOfStutter,
	}
}
