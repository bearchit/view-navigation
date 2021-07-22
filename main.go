package main

import "strings"

func main() {
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Navigation
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Navigation struct {
	screens stack
}

func (nav *Navigation) PushScreen(screen Screen) {
	nav.screens.Push(screen)
}

func (nav *Navigation) PopScreen() Screen {
	return nav.screens.Pop().(Screen)
}

func (nav *Navigation) Stack() []Screen {
	screens := make([]Screen, len(nav.screens.items))
	for i, x := range nav.screens.items {
		screens[i] = x.(Screen)
	}
	return screens
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Screen
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Screen struct {
	id    string
	views stack
}

func NewScreen(id string) Screen {
	return Screen{
		id: id,
	}
}

func (s Screen) ID() string {
	return s.id
}

func (s Screen) String() string {
	sb := strings.Builder{}
	sb.WriteString(s.ID())

	if len(s.views.items) > 0 {
		sb.WriteString("[")
		for _, x := range s.views.items {
			sb.WriteString(x.ID())
		}
		sb.WriteString("]")
	}

	return s.ID()
}

func (s *Screen) PushView(view View) {
	s.views.Push(view)
}

func (s *Screen) PopView() View {
	return s.views.Pop().(View)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// View
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type View struct {
	id string
}

func NewView(id string) View {
	return View{
		id: id,
	}
}

func (v View) ID() string {
	return v.id
}

func (v View) String() string {
	return v.ID()
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Stack
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type stack struct {
	items []stackItem
}

type stackItem interface {
	ID() string
}

func (s *stack) Push(v stackItem) stackItem {
	s.items = append(s.items, v)
	return v
}

func (s *stack) Pop() stackItem {
	if len(s.items) > 0 {
		lastIdx := len(s.items) - 1
		v := s.items[lastIdx]
		s.items = s.items[0:lastIdx]
		return v
	}

	return nil
}
