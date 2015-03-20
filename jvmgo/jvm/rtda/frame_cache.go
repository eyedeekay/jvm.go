package rtda

import rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"

type FrameCache struct {
	thread       *Thread
	cachedFrames []*Frame
	frameCount   uint
	maxFrame     uint
}

func newFrameCache(thread *Thread, maxFrame uint) *FrameCache {
	return &FrameCache{
		thread:       thread,
		maxFrame:     maxFrame,
		cachedFrames: make([]*Frame, maxFrame),
	}
}

func (self *FrameCache) borrowFrame(method *rtc.Method) *Frame {
	if self.frameCount > 0 {
		for i, frame := range self.cachedFrames {
			if frame != nil &&
				frame.maxLocals > method.MaxLocals() &&
				frame.maxStack > method.MaxStack() {

				self.frameCount--
				self.cachedFrames[i] = nil
				frame.reset(method)
				return frame
			}
		}
	}
	return newFrame(self.thread, method)
}

func (self *FrameCache) returnFrame(frame *Frame) {
	for i, cachedFrame := range self.cachedFrames {
		if cachedFrame == nil {
			self.cachedFrames[i] = frame
			self.frameCount++
			return
		}
		// if frame.maxLocals > cachedFrame.maxLocals {
		// 	cachedFrame.maxLocals = frame.maxLocals
		// 	cachedFrame.localVars = frame.localVars
		// }
		// if frame.maxStack > cachedFrame.maxStack {
		// 	cachedFrame.maxStack = frame.maxStack
		// 	cachedFrame.operandStack = frame.operandStack
		// }
	}
}