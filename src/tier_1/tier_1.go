package tier1

import fmt_tier_1 "goLibBasics/src/tier_1/fmt"

type Tier_1 struct {
	fmt_tier_1.LocalTier1
}

func (t *Tier_1) Tier1() {
	t.Fmt_all_packages()
}
