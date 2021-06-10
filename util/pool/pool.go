package pool

import "github.com/dreamlu/gt/tool/gsync"

// Gsync 数量个线程池
var Gsync = gsync.NewPool(2000)
