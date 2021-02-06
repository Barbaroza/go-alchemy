package coroutines

import "fmt"

type Any interface{}

type EvalFunc func(Any, Any) (Any, Any)

func BuildLazyEvaluator(evalFunc EvalFunc, state Any, lastState Any) func() Any {
	retChan := make(chan Any)

	processFunc := func() {
		tempState := state
		tempLastState := lastState
		for {
			ret, lastRet := evalFunc(tempState, tempLastState)
			tempState = ret
			tempLastState = lastRet
			retChan <- ret
		}
	}
	go processFunc()
	retFunc := func() Any {
		return <-retChan
	}

	return retFunc
}

func LazyEvaluatorDemo() {
	evalFunc := func(state Any, lastState Any) (ret Any, lastRet Any) {
		lastRet = state.(int)
		ret = state.(int) + lastState.(int)
		return ret, lastRet
	}

	evaluator := BuildLazyEvaluator(evalFunc, 1, 0)

	for i := 0; i < 10; i++ {
		fmt.Printf("%d : %d\n", i, evaluator())
	}
}
