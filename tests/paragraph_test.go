package tests

import (
	"go-reloaded/internal/processor"
	"testing"
)

func TestLargeParagraph(t *testing.T) {
	input := `A friend sent me 1E (hex) messages yesterday , and I replied 10 (bin) times ! 
He said ' thanks ' but then shouted HELLO (low) . 
It was A honest mistake , I guess . 
This is truly amazing (up, 2) experience (cap) . 
Later that night , I wrote ' I am happy ' in my notebook (up) . 
Then I realized it was just a dream , a illusion that felt real . 
We talked about a orange , a apple , and a umbrella — all while laughing (cap, 4) . 
Finally , before I slept , I whispered ' good night ' and turned off the lights . 
It was a peaceful moment ... but also a reminder of how amazing (up, 3) everything (cap) can be .`

	expected := `A friend sent me 30 messages yesterday, and I replied 2 times! 
He said 'thanks' but then shouted hello. 
It was An honest mistake, I guess. 
This is TRULY AMAZING Experience. 
Later that night, I wrote 'I am happy' in my NOTEBOOK. 
Then I realized it was just a dream, an illusion that felt real. 
We talked about an orange, an apple, and an umbrella — All While Laughing. 
Finally, before I slept, I whispered 'good night' and turned off the lights. 
It was a peaceful moment... but also a reminder OF HOW AMAZING Everything can be.`

	pipeline := processor.NewPipeline()
	result := pipeline.Process(input)

	if result != expected {
		t.Errorf("Paragraph test failed:\nInput:\n%s\n\nExpected:\n%s\n\nGot:\n%s", input, expected, result)
	}
}