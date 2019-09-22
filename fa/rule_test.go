package fa

import "testing"

func TestNewRuleBook(t *testing.T) {
	rb := NewRuleBook([]*Rule{
		NewRule(0, Charset{'a', 'u'}, 1),
		NewRule(1, Charset{'a', 'd'}, 2),
		NewRule(2, Charset{'w', 'w'}, 3),
		NewRule(3, Charset{'u', 'z'}, 4),
	})
	if len(rb.Rules) != 6 {
		t.Errorf("NewRuleBook() length expected %d. got %d", 6, len(rb.Rules))
	}
}
