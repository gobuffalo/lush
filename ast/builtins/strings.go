package builtins

import "strings"

type Strings struct{}

func (Strings) Compare(a, b string) int {
	return strings.Compare(a, b)
}

func (Strings) Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}
func (Strings) ContainsAny(s, chars string) bool {
	return strings.ContainsAny(s, chars)
}
func (Strings) ContainsRune(s string, r rune) bool {
	return strings.ContainsRune(s, r)
}
func (Strings) Count(s, substr string) int {
	return strings.Count(s, substr)
}
func (Strings) EqualFold(s, t string) bool {
	return strings.EqualFold(s, t)
}
func (Strings) Fields(s string) []string {
	return strings.Fields(s)
}
func (Strings) HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}
func (Strings) HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}
func (Strings) Index(s, substr string) int {
	return strings.Index(s, substr)
}
func (Strings) IndexAny(s, chars string) int {
	return strings.IndexAny(s, chars)
}
func (Strings) IndexByte(s string, c byte) int {
	return strings.IndexByte(s, c)
}
func (Strings) IndexRune(s string, r rune) int {
	return strings.IndexRune(s, r)
}
func (Strings) Join(a []string, sep string) string {
	return strings.Join(a, sep)
}
func (Strings) LastIndex(s, substr string) int {
	return strings.LastIndex(s, substr)
}
func (Strings) LastIndexAny(s, chars string) int {
	return strings.LastIndexAny(s, chars)
}
func (Strings) LastIndexByte(s string, c byte) int {
	return strings.LastIndexByte(s, c)
}
func (Strings) Repeat(s string, count int) string {
	return strings.Repeat(s, count)
}
func (Strings) Replace(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}
func (Strings) ReplaceAll(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}
func (Strings) Split(s, sep string) []string {
	return strings.Split(s, sep)
}
func (Strings) SplitAfter(s, sep string) []string {
	return strings.SplitAfter(s, sep)
}
func (Strings) SplitAfterN(s, sep string, n int) []string {
	return strings.SplitAfterN(s, sep, n)
}
func (Strings) SplitN(s, sep string, n int) []string {
	return strings.SplitN(s, sep, n)
}
func (Strings) Title(s string) string {
	return strings.Title(s)
}
func (Strings) ToLower(s string) string {
	return strings.ToLower(s)
}
func (Strings) ToTitle(s string) string {
	return strings.ToTitle(s)
}
func (Strings) ToUpper(s string) string {
	return strings.ToUpper(s)
}
func (Strings) Trim(s string, cutset string) string {
	return strings.Trim(s, cutset)
}
func (Strings) TrimLeft(s string, cutset string) string {
	return strings.TrimLeft(s, cutset)
}
func (Strings) TrimPrefix(s, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}
func (Strings) TrimRight(s string, cutset string) string {
	return strings.TrimRight(s, cutset)
}
func (Strings) TrimSpace(s string) string {
	return strings.TrimSpace(s)
}
func (Strings) TrimSuffix(s, suffix string) string {
	return strings.TrimSuffix(s, suffix)
}
