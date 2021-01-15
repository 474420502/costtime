package costtime

import "time"

// CondGTE 创建cost>=t时间输出日志的条件
func CondGTE(t time.Duration) ConditionFunc {
	return func(cost time.Duration) bool {
		if cost >= t {
			return true
		}
		return false
	}
}

// CondGT 创建cost>t时间输出日志的条件
func CondGT(t time.Duration) ConditionFunc {
	return func(cost time.Duration) bool {
		if cost > t {
			return true
		}
		return false
	}
}

// CondLTE 创建cost<=t时间输出日志的条件
func CondLTE(t time.Duration) ConditionFunc {
	return func(cost time.Duration) bool {
		if cost >= t {
			return true
		}
		return false
	}
}

// CondLT 创建cost<t时间输出日志的条件
func CondLT(t time.Duration) ConditionFunc {
	return func(cost time.Duration) bool {
		if cost < t {
			return true
		}
		return false
	}
}

// CondRange 创建t1<=cost<=t2时间输出日志的条件
func CondRange(t1, t2 time.Duration) ConditionFunc {
	return func(cost time.Duration) bool {
		if t1 <= cost && cost <= t2 {
			return true
		}
		return false
	}
}
