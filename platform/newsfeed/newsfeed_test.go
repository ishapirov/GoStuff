package newsfeed

import "testing"

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{"An item!", "Body of item"})

	if len(feed.Items) != 1 {
		t.Errorf("We got an error, item not added")
	}

}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{"An item!", "Body of item"})

	results := feed.GetAll()

	if len(results) != 1 {
		t.Errorf("We got an error, item not added")
	}
}
