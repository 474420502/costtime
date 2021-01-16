package costtime

var defaultCost *CostTime = New()

// CostLog 计算消耗的时间
func CostLog(name string, run func()) {
	defaultCost.CostLog(name, run)
}

// Cost 里面计算消耗时间
func Cost(run func()) {
	defaultCost.Cost(run)
}

// SetLogCondition 设置输出cost条件
func SetLogCondition(cond ConditionFunc) {
	defaultCost.SetLogCondition(cond)
}

// SetEventCost 设置输出cost事件. 可以做邮件通知. 钉釘等办公通知
func SetEventCost(event EventFunc) {
	defaultCost.SetEeventCost(event)
}
