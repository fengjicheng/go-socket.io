package gcexcel

import "time"

func RelativeDays(duration time.Duration) (uint, uint) {
	// 1. 构造基准时间：2000-01-01 00:00:00
	// 注意：指定时区（这里使用本地时区，与当前时间时区一致，避免时区偏差）
	// 若需使用 UTC 时区，将 time.Local 改为 time.UTC 即可
	baseTime := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.Local)

	// 2. 获取当前系统时间
	start := time.Now()
	//start = time.Date(2026, 1, 29, 0, 0, 0, 0, time.Local)
	// 3. 计算两种类型的天数差
	one := calculateFullDays(start, baseTime)

	if duration == 0 {
		return one, 0
	}

	two := calculateFullDays(start.Add(duration), baseTime)
	return one, two
}

/**
 * 计算两个时间的完整天数差（忽略时分秒，取整，返回整数）
 * @param current 当前时间
 * @param base 基准时间（2000-01-01）
 * @return 完整天数（int64）
 */
func calculateFullDays(current, base time.Time) uint {
	// 步骤1：将两个时间都「截断到日期」（忽略时分秒、毫秒等，只保留年-月-日）
	currentTruncated := time.Date(current.Year(), current.Month(), current.Day(), 0, 0, 0, 0, current.Location())
	baseTruncated := time.Date(base.Year(), base.Month(), base.Day(), 0, 0, 0, 0, base.Location())

	// 步骤2：计算截断后的时间差，转换为完整天数（取整）
	duration := currentTruncated.Sub(baseTruncated)

	// 转换为int64类型的完整天数（无小数）
	return uint(int64(duration.Hours() / 24.0))
}
