package actions

func (as *ActionSuite) Test_Search() {
	res := as.JSON("/api/v1/search").Post(`{"city": "this"}`)
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "message")
}
