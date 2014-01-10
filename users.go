package anaconda

import (
	"net/url"
	"strconv"
)

func (a TwitterApi) GetUsersLookup(usernames string, v url.Values) (u []User, err error) {
	v = cleanValues(v)
	v.Set("screen_name", usernames)
	response_ch := make(chan response)
	a.queryQueue <- query{"http://api.twitter.com/1.1/users/lookup.json", v, &u, _GET, response_ch}
	return u, (<-response_ch).err
}

func (a TwitterApi) GetUsersLookupByIds(ids []int64, v url.Values) (u []User, err error) {
	var pids string
	for w, i := range ids {
		//pids += strconv.Itoa(i)
		pids += strconv.FormatInt(i, 10)
		if w != len(ids)-1 {
			pids += ","
		}
	}
	v = cleanValues(v)
	v.Set("user_id", pids)
	err = a.apiGet("http://api.twitter.com/1.1/users/lookup.json", v, &u)
	return
}

func (a TwitterApi) GetUserSearch(searchTerm string, v url.Values) (u []TwitterUser, err error) {
  v = cleanValues(v)
  v.Set("q", searchTerm)
  // Set other values before calling this method:
  // page, count, include_entities
  err = a.apiGet("http://api.twitter.com/1.1/users/search.json", v, &u)
  return
}

func (a TwitterApi) GetUsersShow(username string, v url.Values) (u User, err error) {
	v = cleanValues(v)
	v.Set("screen_name", username)
	response_ch := make(chan response)
	a.queryQueue <- query{"http://api.twitter.com/1.1/users/show.json", v, &u, _GET, response_ch}
	return u, (<-response_ch).err
}

func (a TwitterApi) GetUsersShowById(id int64, v url.Values) (u User, err error) {
	v = cleanValues(v)
	v.Set("user_id", strconv.FormatInt(id, 10))
	response_ch := make(chan response)
	a.queryQueue <- query{"http://api.twitter.com/1.1/users/show.json", v, &u, _GET, response_ch}
	return u, (<-response_ch).err

}
