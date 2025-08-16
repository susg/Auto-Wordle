package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/susg/autowordle/internal/config"
	"github.com/susg/autowordle/internal/models"
)

func TestRulesCheckerImpl_AreRulesSatisfied(t *testing.T) {
	cfg := config.GetConfig()
	wi := models.NewWordleInfo(5, cfg)
	wi.Update([]string{"ag", "pb", "pb", "ly", "eb"})
	rc := &RulesCheckerImpl{}
	assert.False(t, rc.AreRulesSatisfied(*wi, "plain"))
	assert.False(t, rc.AreRulesSatisfied(*wi, "light"))
	assert.False(t, rc.AreRulesSatisfied(*wi, "about"))
	assert.False(t, rc.AreRulesSatisfied(*wi, "abblb"))
	assert.False(t, rc.AreRulesSatisfied(*wi, "apples"))
	assert.True(t, rc.AreRulesSatisfied(*wi, "aloft"))
}
