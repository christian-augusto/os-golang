package operational_system

import "sync"

var fe *processQueue
var ftr *processQueue
var fu *processQueue
var fu2 *processQueue
var fu3 *processQueue
var osStatus string
var asyncProcess sync.WaitGroup
