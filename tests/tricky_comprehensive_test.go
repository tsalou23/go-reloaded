package tests

import (
	"go-reloaded/internal/processor"
	"testing"
)

func TestTrickyComprehensive(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// ARTICLES
		{"Articles_1", "I am a orange.", "I am an orange."},
		{"Articles_2", "I am A orange.", "I am An orange."},
		{"Articles_3", "I am A ORANGE.", "I am AN ORANGE."},
		{"Articles_4", "I am AN ORANGE.", "I am AN ORANGE."},
		{"Articles_5", "I am AN phone.", "I am A phone."},
		{"Articles_6", "I am a phone.", "I am a phone."},
		{"Articles_7", "I am A phone.", "I am A phone."},

		// CASE COMMANDS
		{"Case_1", "this is a (cap) apple and an (cap) banana.", "this is An apple and An banana."},
		{"Case_2", "make these words (up,2) louder please.", "make THESE WORDS louder please."},
		{"Case_3", "quietly (low,3) YELLING AFTER NOW.", "quietly YELLING AFTER NOW."},
		{"Case_4", "check john doe (cap,2) now.", "check John Doe now."},

		// HEX / BIN
		{"Hex_1", "the number 1E (hex) should become 30.", "the number 30 should become 30."},
		{"Hex_2", "the number 1E (low) (hex) should become 30.", "the number 30 should become 30."},
		{"Bin_1", "the number 1010 (bin) should become 10.", "the number 10 should become 10."},
		{"Hex_Invalid", "but 1G (hex) stays (hex) same because invalid.", "but 1G stays same because invalid."},

		// QUOTES
		{"Quotes_1", "' this is a quoted text , with punctuation ! '", "'this is a quoted text, with punctuation!'"},
		{"Quotes_2", "' mixed CASE inside (low,3) HERE , and (cap) there . '", "'mixed case inside HERE, And there.'"},

		// PUNCTUATION
		{"Punct_1", "this is amazing!!!", "this is amazing!!!"},
		{"Punct_2", "wait ... are you sure?!", "wait... are you sure?!"},
		{"Punct_3", "yes , i am ; absolutely .", "yes, i am; absolutely."},
		{"Punct_4", "i love apples , oranges ; bananas : and grapes !", "i love apples, oranges; bananas: and grapes!"},

		// MIXED ARTICLES + COMMANDS
		{"Mixed_1", "asdf a (cap) orange is better than a (cap) apple.", "asdf An orange is better than An apple."},
		{"Mixed_2", "a (low) ORANGE tastes like AN (low) apple.", "an ORANGE tastes like an apple."},
		{"Mixed_3", "an orange (cap,2) and a fruit salad (up,3).", "An Orange and A FRUIT SALAD."},
		{"Mixed_4", "AN ORANGE (cap) and AN BANANA (up).", "AN ORANGE and A BANANA."},

		// MULTILINE
		{"Multi_1", "this is a test.\na new line starts here.\na (cap) orange\nbut a (up) apple", "this is a test.\na new line starts here.\nAn orange\nbut An apple"},

		// QUOTES + COMMANDS
		{"QuoteCmd_1", "' I am a (cap) optimist , but a (up) realist . '", "'I am An optimist, but A realist.'"},
		{"QuoteCmd_2", "' a (cap) apple a day keeps a (low) doctor away . '", "'An apple a day keeps a doctor away.'"},

		// HEX/BIN + CASE
		{"HexCase_1", "this hex number is 2A (hex), and this binary 1111 (bin).", "this hex number is 42, and this binary 15."},
		{"HexCase_2", "a (cap) 2A (hex) banana and a 1111 (bin) orange.", "A 42 banana and a 15 orange."},

		// MIXED PUNCTUATION
		{"MixedPunct", "a orange?! a phone! an apple... an ORANGE?!", "an orange?! a phone! an apple... an ORANGE?!"},

		// EDGE CASES
		{"Edge_1", "a (cap) (up) orange and an (low) (cap) phone. ???", "An orange and An phone.???"},
		{"Edge_2", "a a a an an a (cap,2) (low,3) orange.", "an a an an an an orange."},
		{"Edge_3", "a , a (cap) orange . a : an (low) apple !", "a, An orange. a: an apple!"},

		// ADVANCED MIXES
		{"Advanced_1", "Behold 1f4 (hex) warriors and 101001 (bin) enemies marching ... slowly ,but surely !!", "Behold 500 warriors and 41 enemies marching... slowly, but surely!!"},
		{"Advanced_2", "He whispered ' the END is NEAR ' (low, 3) ,or maybe not (up) ?", "He whispered 'the END is near', or maybe NOT?"},
		{"Advanced_3", "A hour ago, a elephant walked into a hotel (cap, 5) unexpectedly.", "An hour ago, an Elephant Walked Into A Hotel unexpectedly."},
		{"Advanced_4", "THIS (low, 4) LINE HAS (up, 2) mixed COMMANDS (low, 5) to test priority.", "THIS (low, 4) line has mixed commands to test priority."},
		{"Advanced_5", "Three dots ... or maybe !? punctuation should test spacing ,and emotion !!", "Three dots... or maybe!? punctuation should test spacing, and emotion!!"},
		{"Advanced_6", "He claimed ' i am invincible (up, 2) and immortal ' (cap, 4) ,yet fell to a arrow.", "He claimed 'i AM INVINCIBLE And Immortal', yet fell to an arrow."},
		{"Advanced_7", "He claimed'i AM Invincible And Immortal ', yet fell to an arrow.", "He claimed'i AM Invincible And Immortal', yet fell to an arrow."},
	}

	pipeline := processor.NewPipeline()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pipeline.Process(tt.input)
			if result != tt.expected {
				t.Errorf("Test %s failed:\nInput:    %q\nExpected: %q\nGot:      %q", tt.name, tt.input, tt.expected, result)
			}
		})
	}
}