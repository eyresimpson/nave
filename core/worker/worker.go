package worker

import (
	"nave/tools/log"
	"nave/types/blueprint"
	"net/http"
)

// Run Work编排
// TODO： 设计上这部分应该是流水线的工作（Work分配），后续整理代码时注意
func Run(sid string, parent interface{}, w http.ResponseWriter, r *http.Request) {
	switch parent.(type) {
	case blueprint.Worker:
		parents := parent.(blueprint.Worker)
		// 这个循环的是需要执行节点和其兄弟节点
		for _, child := range parents.Children {
			if child.Sid == sid {
				if child.Type == 0 {
					RestfulWorker(child, w, r)
				} else {
					ConditionWorker(child)
				}
				// 如果有后续节点
				if child.Next != "-1" {
					Run(child.Next, parent, w, r)
				} else {
					break
				}
			}
		}
	case []blueprint.Worker:
		parents := parent.([]blueprint.Worker)
		// 这个循环的是需要执行节点和其兄弟节点
		for _, child := range parents {
			if child.Sid == sid {
				if child.Type == 0 {
					RestfulWorker(child, w, r)
				} else {
					ConditionWorker(child)
				}
				// 如果有后续节点
				if child.Next != "-1" {
					Run(child.Next, parent, w, r)
				} else {
					break
				}
			}
		}
	default:
		log.Warn("Received an unknown type")
	}

}
